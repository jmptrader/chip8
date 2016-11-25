package chip8

const windowRight uint = 64
const windowBottom uint = 32

type HexKey byte

type Window interface {
    Update()
    IsKeyPressed(key HexKey) bool
    WaitForKeyPress() HexKey
    Draw(x, y uint, drawable []byte)
    Clear()
    ShouldClose() bool
    Release()
}
