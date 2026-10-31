package layout

import (
	"sort"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	fyneLayout "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"main/config"
)

type UItem = config.Component

func BuildContent(components map[string]config.Component, objects ...fyne.CanvasObject) fyne.CanvasObject {
	type positionedComponent struct {
		name      string
		component config.Component
	}

	registered := make([]positionedComponent, 0, len(components))
	for name, component := range components {
		registered = append(registered, positionedComponent{name: name, component: component})
	}
	sort.SliceStable(registered, func(left, right int) bool {
		if registered[left].component.Priority == registered[right].component.Priority {
			return registered[left].name < registered[right].name
		}
		return registered[left].component.Priority < registered[right].component.Priority
	})

	var top, bottom, left, right, center []fyne.CanvasObject
	for _, entry := range registered {
		object := buildComponent(entry.component)
		if object == nil {
			continue
		}
		switch strings.ToLower(entry.component.Position) {
		case "top":
			top = append(top, object)
		case "bottom":
			bottom = append(bottom, object)
		case "left":
			left = append(left, object)
		case "right":
			right = append(right, object)
		default:
			center = append(center, object)
		}
	}
	center = append(center, objects...)
	return container.NewBorder(
		stack(top), stack(bottom), stack(left), stack(right), stack(center),
	)
}

func buildComponent(component config.Component) fyne.CanvasObject {
	spacing := parseDimension(component.Spacing)
	items := make([]fyne.CanvasObject, 0, len(component.Items))
	selectionGroups := make(map[string][]*widget.Button)
	for _, item := range component.Items {
		if !strings.EqualFold(item.Type, "button") {
			continue
		}
		button := widget.NewButton(item.Label, nil)
		group := actionGroup(item.Action)
		if group != "" {
			selectionGroups[group] = append(selectionGroups[group], button)
			selectedButton := button
			button.OnTapped = func() {
				for _, groupButton := range selectionGroups[group] {
					groupButton.Importance = widget.MediumImportance
					groupButton.Refresh()
				}
				selectedButton.Importance = widget.HighImportance
				selectedButton.Refresh()
			}
		}
		items = append(items, button)
	}
	if len(items) == 0 {
		return nil
	}

	var horizontal bool
	switch strings.ToLower(component.Type) {
	case "hbox":
		horizontal = true
	case "vbox":
	default:
		return nil
	}
	content := container.New(boxLayout{horizontal: horizontal, spacing: spacing}, items...)
	padding := parseDimension(component.Padding)
	return container.New(fyneLayout.NewCustomPaddedLayout(padding, padding, padding, padding), content)
}

func actionGroup(action string) string {
	if strings.HasPrefix(action, "tool_group:") {
		return "tool_group"
	}
	return ""
}

func parseDimension(value string) float32 {
	dimension, err := strconv.ParseFloat(strings.TrimSpace(value), 32)
	if err != nil || dimension < 0 {
		return 0
	}
	return float32(dimension)
}

func stack(objects []fyne.CanvasObject) fyne.CanvasObject {
	if len(objects) == 0 {
		return nil
	}
	return container.New(boxLayout{}, objects...)
}
