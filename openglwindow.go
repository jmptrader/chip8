package chip8

import "github.com/go-gl/glfw/v3.2/glfw"

type OpenGLWindow struct {
    glWindow *glfw.Window
}

func InitOpenGLWindow() *OpenGLWindow {
    err := glfw.Init()
    if err != nil {
        panic(err)
    }

    glWindow, err := glfw.CreateWindow(64, 32, "chip8", nil, nil)
    if err != nil {
        panic(err)
    }

    glWindow.MakeContextCurrent()
    return &OpenGLWindow{glWindow: glWindow}
}

func (w *OpenGLWindow) Update() {
    w.glWindow.SwapBuffers()
    glfw.PollEvents()
}

func (w *OpenGLWindow) Draw(x byte, y byte, sprite []byte) {

}

func (w *OpenGLWindow) ShouldClose() bool {
    return w.glWindow.ShouldClose()
}

func (w *OpenGLWindow) Release() {
    glfw.Terminate()
}
