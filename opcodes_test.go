package chip8

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type TestWindow struct {
    x, y       uint
    sprite     []byte
    wasCleared bool
}

func (w *TestWindow) Update() {
    // Noop
}

func (w *TestWindow) Draw(x, y uint, sprite []byte) {
    w.x, w.y = x, y
    w.sprite = sprite
}

func (w *TestWindow) Clear() {
    w.wasCleared = true
}

func (w *TestWindow) ShouldClose() bool {
    return false
}

func (w *TestWindow) Release() {
    // Noop
}

func TestRet(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.opcode = 0x00EE
    context.cpu.sp = 1
    context.stack[context.cpu.sp] = 0x321

    runOpcode(context)
    assert.Equal(uint16(0x321), context.cpu.pc)
    assert.Equal(byte(0), context.cpu.sp)
}

func TestJp(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.opcode = 0x1321

    runOpcode(context)
    assert.Equal(uint16(0x321), context.cpu.pc)
}

func TestCall(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.opcode = 0x2321
    context.cpu.sp = 3
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(uint16(0x321), context.cpu.pc)
    assert.Equal(byte(4), context.cpu.sp)
    assert.Equal(pc, context.stack[context.cpu.sp])
}

func TestSebSkip(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.opcode = 0x3111
    context.cpu.v[1] = 17
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(pc + 2, context.cpu.pc)
}

func TestSebNoSkip(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.opcode = 0x3111
    context.cpu.v[1] = 15
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestSnebSkip(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.opcode = 0x4111
    context.cpu.v[1] = 15
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(pc + 2, context.cpu.pc)
}

func TestSnebNoSkip(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.opcode = 0x4111
    context.cpu.v[1] = 17

    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestSeSkip(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.opcode = 0x5120
    context.cpu.v[1] = 17
    context.cpu.v[2] = 17
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(pc + 2, context.cpu.pc)
}

func TestSeNoSkip(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.opcode = 0x5120
    context.cpu.v[1] = 15
    context.cpu.v[2] = 17
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestMv(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.cpu.v[2] = 0x01
    context.opcode = 0x8120
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(context.cpu.v[2], context.cpu.v[1])
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestOr(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.cpu.v[1] = 0x01
    context.cpu.v[2] = 0x10
    context.opcode = 0x8121
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(0x11, int(context.cpu.v[1]))
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestAnd(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.cpu.v[1] = 0x11
    context.cpu.v[2] = 0x10
    context.opcode = 0x8122
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(0x10, int(context.cpu.v[1]))
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestXor(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.cpu.v[1] = 0x11
    context.cpu.v[2] = 0x10
    context.opcode = 0x8123
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(0x01, int(context.cpu.v[1]))
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestAddNoCarry(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.cpu.v[1] = 0x03
    context.cpu.v[2] = 0x02
    context.cpu.v[0xF] = 1
    context.opcode = 0x8124
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(5, int(context.cpu.v[1]))
    assert.Equal(0, int(context.cpu.v[0xF])) // No carry
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestAddCarry(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.cpu.v[1] = 0xFF
    context.cpu.v[2] = 0x05
    context.cpu.v[0xF] = 0
    context.opcode = 0x8124
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(4, int(context.cpu.v[1]))
    assert.Equal(1, int(context.cpu.v[0xF])) // Carry
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestSubBorrow(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.cpu.v[1] = 0x03
    context.cpu.v[2] = 0x05
    context.cpu.v[0xF] = 1
    context.opcode = 0x8125
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(254, int(context.cpu.v[1]))
    assert.Equal(0, int(context.cpu.v[0xF])) // Borrow
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestSubNoBorrow(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.cpu.v[1] = 0x05
    context.cpu.v[2] = 0x03
    context.cpu.v[0xF] = 0
    context.opcode = 0x8125
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(2, int(context.cpu.v[1]))
    assert.Equal(1, int(context.cpu.v[0xF])) // No borrow
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestShrLSB1(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.cpu.v[1] = 0x05
    context.cpu.v[0xF] = 0
    context.opcode = 0x8106
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(2, int(context.cpu.v[1]))
    assert.Equal(1, int(context.cpu.v[0xF]))
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestShrLSB0(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.cpu.v[1] = 0x06
    context.cpu.v[0xF] = 1
    context.opcode = 0x8106
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(3, int(context.cpu.v[1]))
    assert.Equal(0, int(context.cpu.v[0xF]))
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestStBCD(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.cpu.i = 100
    context.cpu.v[1] = 123
    context.opcode = 0xF133
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(1, int(context.memory[100]))
    assert.Equal(2, int(context.memory[101]))
    assert.Equal(3, int(context.memory[102]))
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestSt(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.cpu.i = 100
    context.cpu.v[0] = 1
    context.cpu.v[1] = 2
    context.cpu.v[2] = 3
    context.opcode = 0xF255
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(1, int(context.memory[100]))
    assert.Equal(2, int(context.memory[101]))
    assert.Equal(3, int(context.memory[102]))
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestLd(t *testing.T) {
    assert := assert.New(t)
    context := newContext(newCPU(), new(TestWindow), [4096]byte{})

    context.cpu.i = 100
    context.memory[100] = 1
    context.memory[101] = 2
    context.memory[102] = 3
    context.opcode = 0xF265
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(1, int(context.cpu.v[0]))
    assert.Equal(2, int(context.cpu.v[1]))
    assert.Equal(3, int(context.cpu.v[2]))
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestDrw(t *testing.T) {
    assert := assert.New(t)

    window := new(TestWindow)
    context := newContext(newCPU(), window, [4096]byte{})

    // Stub screen with values
    x, y := 16, 16
    for j := 0; j < 3; j++ {
        for i := 0; i < 8; i++ {
            // Alternating 1s and 0s in the draw bytes
            context.screen[x + i][y + j] = byte(i % 2)
        }
    }

    context.cpu.i = 100
    context.cpu.v[1] = byte(x)
    context.cpu.v[2] = byte(y)
    context.memory[100] = 1
    context.memory[101] = 2
    context.memory[102] = 3
    context.opcode = 0xD123
    pc := context.cpu.pc

    runOpcode(context)
    assert.Equal(uint(x), window.x)
    assert.Equal(uint(y), window.y)
    // screen   ^ sprite   -> draw
    // 01010101 ^ 00000001 -> 01010100 (84)
    // 01010101 ^ 00000010 -> 01010111 (87)
    // 01010101 ^ 00000011 -> 01010110 (86)
    assert.Equal([]byte{84, 87, 86}, window.sprite)
    // VF = 1 since pixels were flipped
    assert.Equal(byte(1), context.cpu.v[0xF])
    assert.Equal(pc + 1, context.cpu.pc)
}

func TestCls(t *testing.T) {
    assert := assert.New(t)

    window := new(TestWindow)
    context := newContext(newCPU(), window, [4096]byte{})
    context.opcode = 0x00E0
    pc := context.cpu.pc

    runOpcode(context)
    assert.True(window.wasCleared)
    assert.Equal(pc + 1, context.cpu.pc)
}
