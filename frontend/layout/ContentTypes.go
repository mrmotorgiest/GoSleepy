package layout

import "fyne.io/fyne/v2"

type boxLayout struct {
	horizontal bool
	spacing    float32
}

func (layout boxLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	position := float32(0)
	for _, object := range objects {
		objectSize := object.MinSize()
		if layout.horizontal {
			object.Move(fyne.NewPos(position, (size.Height-objectSize.Height)/2))
			position += objectSize.Width + layout.spacing
		} else {
			object.Move(fyne.NewPos((size.Width-objectSize.Width)/2, position))
			position += objectSize.Height + layout.spacing
		}
		object.Resize(objectSize)
	}
}

func (layout boxLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	var width, height float32
	for _, object := range objects {
		objectSize := object.MinSize()
		if layout.horizontal {
			width += objectSize.Width
			if objectSize.Height > height {
				height = objectSize.Height
			}
		} else {
			height += objectSize.Height
			if objectSize.Width > width {
				width = objectSize.Width
			}
		}
	}
	if len(objects) > 1 {
		spacing := float32(len(objects)-1) * layout.spacing
		if layout.horizontal {
			width += spacing
		} else {
			height += spacing
		}
	}
	return fyne.NewSize(width, height)
}
