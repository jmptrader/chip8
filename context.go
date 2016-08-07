package chip8

type Context struct {
    opcode uint16 // Current opcode
    stack  [16]uint16
    memory [4096]byte
    cpu    *CPU
}

func newContext(cpu *CPU) *Context {
    return &Context{opcode: 0, cpu: cpu}
}
