package chip8

type CPU struct {
    pc     uint16   // Program counter
    dt, st byte     // Delay timer, sound timer
    sp     byte     // Stack pointer
    i      uint16   // Address register
    v      [16]byte // General purpose registers
}

func newCPU() *CPU {
    return &CPU{pc: 0x200}
}
