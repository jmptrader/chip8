package chip8

import (
    "time"
)

// Update approximately 60 times a second
const msPerTick = 16

type Driver struct {
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
    return &Driver{context: context}
}

func (d *Driver) Run() {
    window := d.context.window
    prev := time.Now().Nanosecond() / 1000000
    delay := 0

    for !window.ShouldClose() {
        now := time.Now().Nanosecond() / 1000000
        delay += now - prev
        prev = now

        window.Update()

        // Processing opcodes and updating timers happen each tick
        for delay >= msPerTick {
            d.updateTimers()
            //runNextOpcode(context)
            delay -= msPerTick
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

func (d *Driver) runNextOpcode() {
    memory := d.context.memory
    cpu := d.context.cpu
    d.context.opcode = uint16(memory[cpu.pc]) << 8 | uint16(memory[cpu.pc + 1])
    runOpcode(d.context)
}
