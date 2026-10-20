# GoSleepy
An open source music software , made on golang with an emphasis on customization.

# Build instructions

## Hyprland

For the Splash screen to work properly update ur hyprland config with the following
```
hl.window_rule({
    name = "gosleepy-splash",
    match = {
        title = "^GoSleepy Splash$",
    },

    float = true,
    size = { 400, 400 },
    center = true,
    rounding = 0,
    opacity = "0.9 override 0.9 override",
})
```
