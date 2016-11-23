package chip8

type Context struct {
    opcode uint16 // Current opcode
    stack  [16]uint16
    memory [4096]byte
    cpu    *CPU
    window Window // Interface for audio, graphics, and input
    screen [64][32]byte // Internal representation of screen independent of window
}

func newContext(cpu *CPU, window Window, memory [4096]byte) *Context {
    return &Context{opcode: 0, cpu: cpu, window: window, memory: memory}
}
