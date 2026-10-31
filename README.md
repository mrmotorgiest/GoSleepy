# GoSleepy
An open source music software , made on golang with an emphasis on customization.
Made by MrMotorgiest in 2026.

# Build instructions
To build you just need the latest version of golang and then run ```go run .``` from the directory.
Im not sure if the code works on windows tho.

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
    opacity = "1 override 1 override",
})
```
