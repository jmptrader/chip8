package chip8

type HexKey byte

type Window interface {
    Update()
    IsKeyPressed(key HexKey) bool
    WaitForKeyPress() HexKey
    Draw(screen *[64][32]byte)
    Clear()
    ShouldClose() bool
    Release()
}
