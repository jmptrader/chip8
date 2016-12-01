package chip8

import (
    sf "bitbucket.org/krepa098/gosfml2"
)

type SFMLWindow struct {
    window *sf.RenderWindow
    scale  sf.Vector2f
    bitmap [4 * 64 * 32]byte
    keys   map[HexKey]sf.KeyCode
}

func NewSFMLWindow(width, height uint) *SFMLWindow {
    videoMode := sf.VideoMode{Width: width, Height: height, BitsPerPixel: 32}
    windowStyle := sf.StyleTitlebar | sf.StyleClose
    window := sf.NewRenderWindow(videoMode, "chip8", windowStyle, sf.DefaultContextSettings())
    bitmap := [4 * 64 * 32]byte{}
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

    return &SFMLWindow{window: window, scale: sf.Vector2f{float32(width / 64), float32(height / 32)},
                       bitmap: bitmap, keys: keys}
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
    for !w.ShouldClose() {
        event := w.window.WaitEvent()
        switch ev := event.(type) {
        case sf.EventKeyPressed:
            // Inefficient lookup. Something more elegant than storing reverse map?
            for k, v := range w.keys {
                if ev.Code == v {
                    return k
                }
            }
        case sf.EventClosed:
            w.window.Close()
        }
    }
    // Dummy return value. Program will exit.
    return 0xFF
}

func (w *SFMLWindow) Draw(screen *[64][32]byte) {
    for j := uint(0); j < 32; j++ {
        for i := uint(0); i < 64; i++ {
            for k := uint(0); k < 4; k++ {
                w.bitmap[4 * (64 * j + i) + k] = 255 * screen[i][j]
            }
        }
    }

    image, _ := sf.NewImageFromPixels(64, 32, w.bitmap[:])
    texture, _ := sf.NewTextureFromImage(image, nil)
    sprite, _ := sf.NewSprite(texture)
    sprite.Scale(w.scale)

    w.window.Draw(sprite, sf.DefaultRenderStates())
    w.window.Display()
}

func (w *SFMLWindow) Clear() {
    w.window.Clear(sf.ColorBlack())
}

func (w *SFMLWindow) ShouldClose() bool {
    return !w.window.IsOpen()
}

func (w *SFMLWindow) Release() {
    // Noop
}
