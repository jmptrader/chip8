package chip8

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestJp(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.opcode = 0x1321

    runOpcode(context)
    assert.Equal(context.cpu.pc, uint16(0x321))
}

func TestCall(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.opcode = 0x2321
    // pc := context.cpu.pc
    context.cpu.pc = 15
    context.cpu.sp = 3

    runOpcode(context)
    assert.Equal(context.cpu.pc, uint16(0x321))
    assert.Equal(context.cpu.sp, byte(4))
    assert.Equal(context.stack[context.cpu.sp], uint16(15))
}

func TestSebSkip(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.opcode = 0x3111
    // Explicitly set or fetch pc
    // context.cpu.pc =
    context.cpu.v[1] = 17

    runOpcode(context)
    assert.Equal(uint16(2), context.cpu.pc)
}

func TestSebNoSkip(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.opcode = 0x3111
    context.cpu.v[1] = 15

    runOpcode(context)
    assert.Equal(uint16(1), context.cpu.pc)
}

func TestSnebSkip(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.opcode = 0x4111
    context.cpu.v[1] = 15

    runOpcode(context)
    assert.Equal(uint16(2), context.cpu.pc)
}

func TestSnebNoSkip(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.opcode = 0x4111
    context.cpu.v[1] = 17

    runOpcode(context)
    assert.Equal(uint16(1), context.cpu.pc)
}

func TestSeSkip(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.opcode = 0x5120
    context.cpu.v[1] = 17
    context.cpu.v[2] = 17

    runOpcode(context)
    assert.Equal(uint16(2), context.cpu.pc)
}

func TestSeNoSkip(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.opcode = 0x5120
    context.cpu.v[1] = 15
    context.cpu.v[2] = 17

    runOpcode(context)
    assert.Equal(uint16(1), context.cpu.pc)
}

func TestMv(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.cpu.v[2] = 0x01
    context.opcode = 0x8120

    runOpcode(context)
    assert.Equal(context.cpu.v[2], context.cpu.v[1])
    assert.Equal(uint16(1), context.cpu.pc)
}

func TestOr(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.cpu.v[1] = 0x01
    context.cpu.v[2] = 0x10
    context.opcode = 0x8121

    runOpcode(context)
    assert.Equal(0x11, int(context.cpu.v[1]))
    assert.Equal(uint16(1), context.cpu.pc)
}

func TestAnd(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.cpu.v[1] = 0x11
    context.cpu.v[2] = 0x10
    context.opcode = 0x8122

    runOpcode(context)
    assert.Equal(0x10, int(context.cpu.v[1]))
    assert.Equal(uint16(1), context.cpu.pc)
}

func TestXor(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.cpu.v[1] = 0x11
    context.cpu.v[2] = 0x10
    context.opcode = 0x8123

    runOpcode(context)
    assert.Equal(0x01, int(context.cpu.v[1]))
    assert.Equal(uint16(1), context.cpu.pc)
}

func TestAddNoCarry(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.cpu.v[1] = 0x03
    context.cpu.v[2] = 0x02
    context.opcode = 0x8124

    runOpcode(context)
    assert.Equal(5, int(context.cpu.v[1]))
    assert.Equal(0, int(context.cpu.v[0xF])) // No carry
    assert.Equal(uint16(1), context.cpu.pc)
}

func TestAddCarry(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.cpu.v[1] = 0xFF
    context.cpu.v[2] = 0x05
    context.opcode = 0x8124

    runOpcode(context)
    assert.Equal(4, int(context.cpu.v[1]))
    assert.Equal(1, int(context.cpu.v[0xF])) // Carry
    assert.Equal(uint16(1), context.cpu.pc)
}

func TestSubBorrow(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.cpu.v[1] = 0x03
    context.cpu.v[2] = 0x05
    context.opcode = 0x8125

    runOpcode(context)
    assert.Equal(254, int(context.cpu.v[1]))
    assert.Equal(0, int(context.cpu.v[0xF])) // Borrow
    assert.Equal(uint16(1), context.cpu.pc)
}

func TestSubNoBorrow(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.cpu.v[1] = 0x05
    context.cpu.v[2] = 0x03
    context.opcode = 0x8125

    runOpcode(context)
    assert.Equal(2, int(context.cpu.v[1]))
    assert.Equal(1, int(context.cpu.v[0xF])) // No borrow
    assert.Equal(uint16(1), context.cpu.pc)
}

func TestShrLSB1(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.cpu.v[1] = 0x05
    context.opcode = 0x8106
    // context.cpu.v[0xF] = 0

    runOpcode(context)
    assert.Equal(2, int(context.cpu.v[1]))
    assert.Equal(1, int(context.cpu.v[0xF]))
    assert.Equal(uint16(1), context.cpu.pc)
}

func TestShrLSB0(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU())

    context.cpu.v[1] = 0x06
    context.opcode = 0x8106
    // context.cpu.v[0xF] = 1

    runOpcode(context)
    assert.Equal(3, int(context.cpu.v[1]))
    assert.Equal(0, int(context.cpu.v[0xF]))
    assert.Equal(uint16(1), context.cpu.pc)
}
