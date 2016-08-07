package chip8

import "math/rand"

type Opcode func(*Context)

func runOpcode(context *Context) {
    opcodes[(context.opcode & 0xF000) >> 12](context)
}

func nop(context *Context) {
    // Do nothing
}

// 0xxx opcodes
func ops0(context *Context) {
    opcodes0[context.opcode & 0xF]
}

// 00E0 - CLS
// Clear the display
func cls(context *Context) {
    // Clear display
}

// 00EE - RET
// Return from a subroutine
func ret(context *Context) {

}

// 1nnn - JP nibble
// Jump to location nnn
func jp(context *Context) {
    context.cpu.pc = context.opcode & 0x0FFF
}

// 2nnn - CALL nibble
// Call subroutine at nnn
func call(context *Context) {
    context.cpu.sp++
    context.stack[context.cpu.sp] = context.cpu.pc
    context.cpu.pc = context.opcode & 0x0FFF
}

// 3xkk - SE Vx, byte
// Skip next instruction if Vx = kk
func seb(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    b := byte(context.opcode & 0xFF)
    if context.cpu.v[x] == b {
        context.cpu.pc++
    }
    context.cpu.pc++
}

// 4xkk - SNE Vx, byte
// Skip next instruction if Vx != kk
func sneb(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    b := byte(context.opcode & 0xFF)
    if context.cpu.v[x] != b {
        context.cpu.pc++
    }
    context.cpu.pc++
}

// 5xy0 - SE Vx, Vy
// Skip next instruction if Vx = Vy
func se(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    y := context.opcode & 0x00F0 >> 4
    if context.cpu.v[x] == context.cpu.v[y] {
        context.cpu.pc++
    }
    context.cpu.pc++
}

// 6xkk - LD Vx, byte
// Set Vx = kk
func ldb(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    b := context.opcode & 0x00FF
    context.cpu.v[x] = b
    context.cpu.pc++
}

// 7xkk - ADD Vx, byte
// Set Vx = Vx + kk
func addb(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    b := context.opcode & 0x00FF
    context.cpu.v[x] += b
    context.cpu.pc++
}

// 8xxx opcodes
func ops8(context *Context) {
    opcodes8[context.opcode & 0xF](context);
    context.cpu.pc++
}

func mv(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    y := context.opcode & 0x00F0 >> 4
    context.cpu.v[x] = context.cpu.v[y]
}

func or(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    y := context.opcode & 0x00F0 >> 4
    context.cpu.v[x] = context.cpu.v[x] | context.cpu.v[y]
}

func and(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    y := context.opcode & 0x00F0 >> 4
    context.cpu.v[x] = context.cpu.v[x] & context.cpu.v[y]
}

func xor(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    y := context.opcode & 0x00F0 >> 4
    context.cpu.v[x] = context.cpu.v[x] ^ context.cpu.v[y]
}

func add(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    y := context.opcode & 0x00F0 >> 4
    sum := uint16(context.cpu.v[x]) + uint16(context.cpu.v[y])
    if sum > 255 {
        context.cpu.v[0xF] = 1
    }
    context.cpu.v[x] = byte(sum & 0xFF)
}

func sub(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    y := context.opcode & 0x00F0 >> 4
    if context.cpu.v[x] > context.cpu.v[y] {
        context.cpu.v[0xF] = 1
    } else {
        context.cpu.v[0xF] = 0
    }
    context.cpu.v[x] = context.cpu.v[x] - context.cpu.v[y]
}

func shr(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    if context.cpu.v[x] & 0x1 == 0x1 {
        context.cpu.v[0xF] = 1
    } else {
        context.cpu.v[0xF] = 0
    }
    context.cpu.v[x] /= 2
}

func subn(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    y := context.opcode & 0x00F0 >> 4
    if context.cpu.v[y] > context.cpu.v[x] {
        context.cpu.v[0xF] = 1
    } else {
        context.cpu.v[0xF] = 0
    }
    context.cpu.v[x] = context.cpu.v[x] - context.cpu.v[y]
}

// 8xyE - SHL Vx {, Vy}
// Set Vx = Vx SHL 1
func shl(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    if context.cpu.v[x] & 0x8000 >> 12 == 0x1 {
        context.cpu.v[0xF] = 1
    } else {
        context.cpu.v[0xF] = 0
    }
    context.cpu.v[x] *= 2
}

// 9xy0 - SNE Vx, Vy
// Skip next instruction if Vx != Vy
func sne(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    y := context.opcode & 0x00F0 >> 4
    if context.cpu.v[x] != context.cpu.v[y] {
        context.cpu.pc++
    }
    context.cpu.pc++
}

// Annn - LD I, nibble
// Set I = nnn
func ldn(context *Context) {
    n := context.opcode & 0x0FFF
    context.cpu.i = n
    context.cpu.pc++
}

// Bnnn - JP V0, nibble
// Jump to location V0 + nnn
func jpn(context *Context) {
    n := context.opcode & 0x0FFF
    context.cpu.pc = context.cpu.v[0] + n
}

// Cxkk - RND Vx, byte
// Set Vx = random byte AND kk
func rnd(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    b := context.opcode & 0xFF
    context.cpu.v[x] = b & byte(rand.Intn(256))
    context.cpu.pc++
}

// Dxyn - DRW Vx, Vy, nibble
// Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision
func drw(context *Context) {
    // Use Window/Draw interface
}

// Ex9E - SKP Vx
// Skip next instruction if key with the value of Vx is pressed
func skp(context *Context) {
    // TODO
}

// ExA1 - SKNP Vx
// Skip next instruction if key with the value of Vx is not pressed
func sknp(context *Context) {
    // TODO
}

// Fx07 - MV Vx, DT
// Set Vx = delay timer value
func mvfd(context *Context) {

}

// Fx15 - MV DT, Vx
// Set delay timer = Vx.
func mvtd(context *Context) {

}

// Fx0A - LD Vx, K
// Wait for a key press, store the value of the key in Vx

// Fx18 - LD ST, Vx
// Set sound timer = Vx
func mvts(context *Context) {

}

var opcodes = [17]Opcode { ops0, jp, call, seb, sneb, se, ldb, addb, ops8,
                           sne, nop, nop, nop, nop, nop, nop, nop }

var opcodes0 = [15]Opcode { cls, nop, nop, nop, nop, nop, nop, nop, nop, nop,
                            nop, nop, nop, ret }

var opcodes8 = [16]Opcode { mv, or, and, xor, add, sub, shr,
                            subn, shl, nop, nop, nop, nop, shl }