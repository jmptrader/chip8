package chip8

type Driver struct {
    window  Window
    context *Context
}

func NewDriver(window Window) *Driver {
    return &Driver{window: window}
}

func (d *Driver) Run() {
    sprite := []byte{255, 255, 255, 255, 255, 255, 255, 255}
    window := d.window
    for !window.ShouldClose() {
        window.Update()
        window.Draw(0, 0, sprite)
        //context.opcode = d.nextOpcode()
        //runOpcode(context)
    }
}

func (d *Driver) nextOpcode() uint16 {
    memory := d.context.memory
    cpu := d.context.cpu
    return uint16(memory[cpu.pc]) << 8 | uint16(memory[cpu.pc + 1])
}
