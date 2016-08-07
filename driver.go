package chip8

func Run() {
    var memory [4096]byte

    memory[0] = 0x21;
    memory[1] = 0x23;

    cpu := newCPU()
    context := newContext(cpu)
    context.opcode = uint16(memory[cpu.pc]) << 8 | uint16(memory[cpu.pc + 1])

    cpu.pc += 2
    cpu.v[10] = 't'

    runOpcode(context)
}
