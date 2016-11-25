package chip8

import (
    sf "bitbucket.org/krepa098/gosfml2"
)

type SFMLWindow struct {
    window                  *sf.RenderWindow
    widthScale, heightScale float32
    keys                    map[HexKey]sf.KeyCode
}

func NewSFMLWindow(width, height uint) *SFMLWindow {
    videoMode := sf.VideoMode{Width: width, Height: height, BitsPerPixel: 32}
    windowStyle := sf.StyleTitlebar | sf.StyleClose
    window := sf.NewRenderWindow(videoMode, "chip8", windowStyle, sf.DefaultContextSettings())
    keys := map[HexKey]sf.KeyCode {
        0x0: sf.KeyX,
        0x1: sf.KeyNum1,
        0x2: sf.KeyNum2,
        0x3: sf.KeyNum3,
        0x4: sf.KeyQ,
        0x5: sf.KeyW,
        0x6: sf.KeyE,
        0x7: sf.KeyA,
        0x8: sf.KeyS,
        0x9: sf.KeyD,
        0xA: sf.KeyZ,
        0xB: sf.KeyC,
        0xC: sf.KeyNum4,
        0xD: sf.KeyR,
        0xE: sf.KeyF,
        0xF: sf.KeyV,
    }

    return &SFMLWindow{window: window, widthScale: float32(width / 64), heightScale: float32(height / 32), keys: keys}
}

func (w *SFMLWindow) Update() {
    for event := w.window.PollEvent(); event != nil; event = w.window.PollEvent() {
        switch event.(type) {
        case sf.EventClosed:
            w.window.Close()
        }
    }
}

func (w *SFMLWindow) IsKeyPressed(key HexKey) bool {
    return sf.KeyboardIsKeyPressed(w.keys[key])
}

func (w *SFMLWindow) WaitForKeyPress() HexKey {
    for {
        event := w.window.WaitEvent()
        switch ev := event.(type) {
        case sf.EventKeyPressed:
            // Inefficient lookup. Something more elegant than storing reverse map?
            for k, v := range w.keys {
                if ev.Code == v {
                    return k
                }
            }
        }
    }
}

func (w *SFMLWindow) Draw(x, y uint, drawable []byte) {
    // Width of the drawable clamped to the right edge of the screen
    var clampedWidth uint
    if windowRight - x < 8 {
        clampedWidth = windowRight - x
    } else {
        clampedWidth = 8
    }

    // Height of the drawable clamped to the bottom edge of the screen
    var clampedHeight uint
    if windowBottom - y < uint(len(drawable)) {
        clampedHeight = windowBottom - y
    } else {
        clampedHeight = uint(len(drawable))
    }

    // The wrapped coordinates of the bottom right corner of the drawable
    xw, yw := (x + 8) % windowRight, (y + uint(len(drawable))) % windowBottom

    // Top left quadrant of sprite. This is always drawn.
    sprite1 := w.createSprite(x, y, clampedWidth, clampedHeight, drawable[:clampedHeight])
    w.window.Draw(sprite1, sf.DefaultRenderStates())

    // Top right quadrant of sprite. If we span the right edge.
    if x + 8 > windowRight {
        sprite2 := w.createSprite(0, y, xw, clampedHeight, drawable[:clampedHeight])
        w.window.Draw(sprite2, sf.DefaultRenderStates())
    }
    // Bottom left quadrant of sprite. If we span the bottom edge.
    if y + uint(len(drawable)) > windowBottom {
        sprite3 := w.createSprite(x, 0, clampedWidth, yw, drawable[clampedHeight:clampedHeight + yw])
        w.window.Draw(sprite3, sf.DefaultRenderStates())
    }
    // Bottom right quadrant of sprite. If we span both the right and bottom edges.
    if x + 8 > windowRight && y + uint(len(drawable)) > windowBottom {
        sprite4 := w.createSprite(0, 0, xw, yw, drawable[clampedHeight:clampedHeight + yw])
        w.window.Draw(sprite4, sf.DefaultRenderStates())
    }

    w.window.Display()
}

func (w *SFMLWindow) Clear() {
    w.window.Clear(sf.Color{0, 0, 0, 255})
}

func (w *SFMLWindow) ShouldClose() bool {
    return !w.window.IsOpen()
}

func (w *SFMLWindow) Release() {
    // Noop
}

func (w *SFMLWindow) createSprite(x, y, width, height uint, drawable []byte) *sf.Sprite {
    bitmap := make([]byte, 4 * width * height)
    for j := uint(0); j < height; j++ {
        for i := uint(0); i < width; i++ {
            shift := (drawable[j] >> uint(width - i - 1)) & 0x1
            bitmap[4*(width*j + i)] = 255 * shift
            bitmap[4*(width*j + i) + 1] = 255 * shift
            bitmap[4*(width*j + i) + 2] = 255 * shift
            bitmap[4*(width*j + i) + 3] = 255 * shift
        }
    }

    image, _ := sf.NewImageFromPixels(width, height, bitmap)
    texture, _ := sf.NewTextureFromImage(image, nil)
    sprite, _ := sf.NewSprite(texture)
    sprite.SetPosition(sf.Vector2f{X: w.widthScale * float32(x), Y: w.heightScale * float32(y)})
    sprite.Scale(sf.Vector2f{X: w.widthScale, Y: w.heightScale})
    return sprite
}
