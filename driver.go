package chip8

type Driver struct {
    screen  Screen
    context *Context
}

func NewDriver(screen Screen) *Driver {
    return &Driver{screen: screen}
}

func (d *Driver) Run() {
    //var memory [4096]byte
    //cpu := newCPU()
    //context := newContext(cpu, memory)

    screen := d.screen
    for !screen.ShouldClose() {
        screen.Update()
        //context.opcode = d.nextOpcode()
        //runOpcode(context)
    }
}

func (d *Driver) nextOpcode() uint16 {
    memory := d.context.memory
    cpu := d.context.cpu
    return uint16(memory[cpu.pc]) << 8 | uint16(memory[cpu.pc + 1])
}
