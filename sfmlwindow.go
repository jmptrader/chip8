package chip8

import (
    sf "bitbucket.org/krepa098/gosfml2"
)

type SFMLWindow struct {
    window *sf.RenderWindow
}

func NewSFMLWindow(width, height uint) *SFMLWindow {
    videoMode := sf.VideoMode{Width: width, Height: height, BitsPerPixel: 32}
    window := sf.NewRenderWindow(videoMode, "chip8", sf.StyleDefault, sf.DefaultContextSettings())
    return &SFMLWindow{window: window}
}

func (w *SFMLWindow) Update() {
    for event := w.window.PollEvent(); event != nil; event = w.window.PollEvent() {
        switch event.(type) {
        case sf.EventClosed:
            w.window.Close()
        }
    }
}

func (w *SFMLWindow) Draw(x, y uint, sprite []byte) {
    // TODO: sprite can be split into 4 separate images due to wrapping
    width, height := 8, len(sprite)
    pixels := make([]byte, 4 * width * height, 4 * width * height)
    for j := 0; j < height; j++ {
        for i := 0; i < width; i++ {
            shift := uint(width - i - 1)
            pixels[4*(width*j + i)] = 255 * (sprite[j] >> shift & 0x1)
            pixels[4*(width*j + i) + 1] = 255 * (sprite[j] >> shift & 0x1)
            pixels[4*(width*j + i) + 2] = 255 * (sprite[j] >> shift & 0x1)
            pixels[4*(width*j + i) + 3] = 255 * (sprite[j] >> shift & 0x1)
        }
    }

    image, _ := sf.NewImageFromPixels(uint(width), uint(height), pixels)
    texture, _ := sf.NewTextureFromImage(image, nil)
    sp, _ := sf.NewSprite(texture)

    // TODO: scale SFML sprite by width / 64, height / 32
    // enforce multiples of (64, 32) in cmd
    w.window.Draw(sp, sf.DefaultRenderStates())
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
