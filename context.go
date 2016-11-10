package chip8

type Context struct {
    opcode uint16 // Current opcode
    stack  [16]uint16
    memory [4096]byte
    cpu    *CPU
    screen Screen
}

func newContext(cpu *CPU, screen Screen, memory [4096]byte) *Context {
    return &Context{opcode: 0, cpu: cpu, screen: screen, stack: [16]uint16{}, memory: memory}
}
