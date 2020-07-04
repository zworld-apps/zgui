# zgui

[![Build Status](https://travis-ci.com/zworld-apps/zgui.svg?branch=master)](https://travis-ci.com/zworld-apps/zgui)
[![GoDoc](https://godoc.org/github.com/zworld-apps/zgui?status.svg)](https://godoc.org/github.com/zworld-apps/zgui)
[![Go Report Card](https://goreportcard.com/badge/github.com/zworld-apps/zgui)](https://goreportcard.com/report/github.com/zworld-apps/zgui)

**zgui** is a GUI system based on Constraints, which uses `raylib`.

[![Example image](https://i.imgur.com/3IA243l.png)]

## Example

In zgui, a component definition is always preceded by a constraints definition
(`constraints`) and the struct with the options of the component (`BoxOptions`).

``` go
display := zgui.GetDisplay()

constraints := zgui.DefaultConstraints()
constraints.SetY(zgui.NewCenterConstraint())
constraints.SetWidth(zgui.NewRelativeConstraint(func(x float32) float32 {
    return x * 0.3
}))
```

After describing the component, you can add it to a parent element. The root
component is `display`.

``` go
greenBox := zgui.NewBoxComponent(&zgui.BoxOptions{
    Color:     rl.Green,
    Roundness: 0.1,
    Segments:  20,
})

display.Add(greenBox, constraints)
```

Finally, the `raylib` main loop should look like this:
``` go
for !rl.WindowShouldClose() {
    rl.BeginDrawing()
    rl.ClearBackground(rl.Black)

    display.Update(rl.GetFrameTime())
    display.Draw()

    rl.EndDrawing()
}
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. Remember to update tests according to the changes.

## License
[MIT](https://github.com/zworld-apps/zgui/blob/master/LICENSE)
