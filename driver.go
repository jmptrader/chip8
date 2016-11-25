package chip8

import (
    "time"
)

// Tick at approximately 60 fps
const tick = 16

type Driver struct {
    window  Window
    context *Context
}

func NewDriver(window Window) *Driver {
    memory := [4096]byte {
        0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
        0x20, 0x60, 0x20, 0x20, 0x70, // 1
        0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
        0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
        0x90, 0x90, 0xF0, 0x10, 0x10, // 4
        0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
        0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
        0xF0, 0x10, 0x20, 0x40, 0x40, // 7
        0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
        0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
        0xF0, 0x90, 0xF0, 0x90, 0x90, // A
        0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
        0xF0, 0x80, 0x80, 0x80, 0xF0, // C
        0xE0, 0x90, 0x90, 0x90, 0xE0, // D
        0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
        0xF0, 0x80, 0xF0, 0x80, 0x80, // F
    }
    context := newContext(newCPU(), window, memory)
    return &Driver{window: window, context: context}
}

func (d *Driver) Run() {
    drawable := []byte{255, 255, 255, 255, 255, 255, 255, 255}
    window := d.window

    prev := time.Now().Nanosecond() / 1000000
    accum := 0

    for !window.ShouldClose() {
        now := time.Now().Nanosecond() / 1000000
        accum += now - prev
        prev = now

        window.Update()

        // Processing opcodes and updating timers happens each tick
        for accum >= tick {
            window.Draw(60, 28, drawable)
            d.updateTimers()
            //runNextOpcode(context)
            accum -= tick
        }
    }
}

func (d *Driver) updateTimers() {
    if d.context.cpu.dt > 0 {
        d.context.cpu.dt--
    }
    if d.context.cpu.st > 0 {
        d.context.cpu.st--
    }
}

func (d *Driver) runNextOpcode() uint16 {
    memory := d.context.memory
    cpu := d.context.cpu
    opcode := uint16(memory[cpu.pc]) << 8 | uint16(memory[cpu.pc + 1])
    runOpcode(opcode)
}
