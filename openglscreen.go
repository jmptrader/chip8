package chip8

import "github.com/go-gl/glfw/v3.2/glfw"

type OpenGLScreen struct {
    window *glfw.Window
}

func InitOpenGLScreen() *OpenGLScreen {
    err := glfw.Init()
    if err != nil {
        panic(err)
    }

    window, err := glfw.CreateWindow(64, 32, "chip8", nil, nil)
    if err != nil {
        panic(err)
    }

    window.MakeContextCurrent()
    return &OpenGLScreen{window: window}
}

func (s *OpenGLScreen) Update() {
    s.window.SwapBuffers()
    glfw.PollEvents()
}

func (s *OpenGLScreen) Draw(x uint, y uint, sprite []byte) {
    // TODO
}

func (s *OpenGLScreen) Clear() {
    // TODO
}

func (s *OpenGLScreen) ShouldClose() bool {
    return s.window.ShouldClose()
}

func (s *OpenGLScreen) Release() {
    glfw.Terminate()
}
