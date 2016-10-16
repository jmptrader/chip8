package chip8

import (
    "math/rand"
    "fmt"
)

type Opcode func(*Context)

var opcodes = [17]Opcode { ops0, jp, call, seb, sneb, se, ldb, addb, ops8,
    sne, ldn, jpn, rnd, drw, opse, opsf }

func runOpcode(context *Context) {
    opcodes[(context.opcode & 0xF000) >> 12](context)
}

func nop(context *Context) {
    // Do nothing
}

// 0xxx opcodes
func ops0(context *Context) {
    switch context.opcode & 0x00FF {
    case 0xE0:
        cls(context)
    case 0xEE:
        ret(context)
    default:
        panic(fmt.Sprintf("Unrecognized opcode %v!", context.opcode))
    }
}

// 00E0 - CLS
// Clear the display
func cls(context *Context) {
    // Clear display
}

// 00EE - RET
// Return from a subroutine
func ret(context *Context) {
    context.cpu.pc = context.stack[context.cpu.sp]
    context.cpu.sp--
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
    b := byte(context.opcode & 0x00FF)
    context.cpu.v[x] = b
    context.cpu.pc++
}

// 7xkk - ADD Vx, byte
// Set Vx = Vx + kk
func addb(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    b := byte(context.opcode & 0x00FF)
    context.cpu.v[x] += b
    context.cpu.pc++
}

// 8xxx opcodes
func ops8(context *Context) {
    switch context.opcode & 0x000F {
    case 0x0:
        mv(context)
    case 0x1:
        or(context)
    case 0x2:
        and(context)
    case 0x3:
        xor(context)
    case 0x4:
        add(context)
    case 0x5:
        sub(context)
    case 0x6:
        shr(context)
    case 0x7:
        subn(context)
    case 0xE:
        shl(context)
    default:
        panic(fmt.Sprintf("Unrecognized opcode %v!", context.opcode))
    }
    context.cpu.pc++
}

// 8xy0 - MV Vx, Vy
// Set Vx = Vy
func mv(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    y := context.opcode & 0x00F0 >> 4
    context.cpu.v[x] = context.cpu.v[y]
}

// 8xy1 - OR Vx, Vy
// Set Vx = Vx OR Vy
func or(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    y := context.opcode & 0x00F0 >> 4
    context.cpu.v[x] = context.cpu.v[x] | context.cpu.v[y]
}

// 8xy2 - AND Vx, Vy
// Set Vx = Vx AND Vy
func and(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    y := context.opcode & 0x00F0 >> 4
    context.cpu.v[x] = context.cpu.v[x] & context.cpu.v[y]
}

// 8xy3 - XOR Vx, Vy
// Set Vx = Vx XOR Vy.
func xor(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    y := context.opcode & 0x00F0 >> 4
    context.cpu.v[x] = context.cpu.v[x] ^ context.cpu.v[y]
}

// 8xy4 - ADD Vx, Vy
// Set Vx = Vx + Vy, set VF = carry
func add(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    y := context.opcode & 0x00F0 >> 4
    sum := uint16(context.cpu.v[x]) + uint16(context.cpu.v[y])
    if sum > 255 {
        context.cpu.v[0xF] = 1
    } else {
        context.cpu.v[0xF] = 0
    }
    context.cpu.v[x] = byte(sum & 0xFF)
}

// 8xy5 - SUB Vx, Vy
// Set Vx = Vx - Vy, set VF = NOT borrow
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

// 8xy6 - SHR Vx {, Vy}
// Set Vx = Vx SHR 1
func shr(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    if context.cpu.v[x] & 0x1 == 0x1 {
        context.cpu.v[0xF] = 1
    } else {
        context.cpu.v[0xF] = 0
    }
    context.cpu.v[x] /= 2
}

// 8xy7 - SUBN Vx, Vy
// Set Vx = Vy - Vx, set VF = NOT borrow
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
    if context.cpu.v[x] & 0x8 == 0x8 {
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
    context.cpu.pc = uint16(context.cpu.v[0]) + n
}

// Cxkk - RND Vx, byte
// Set Vx = random byte AND kk
func rnd(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    b := byte(context.opcode & 0xFF)
    context.cpu.v[x] = b & byte(rand.Intn(256))
    context.cpu.pc++
}

// Dxyn - DRW Vx, Vy, nibble
// Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision
func drw(context *Context) {
    // Use Window/Draw interface
    // TODO
}

func opse(context *Context) {
    switch context.opcode & 0x00FF {
    case 0x9E:
        skp(context)
    case 0xA1:
        sknp(context)
    default:
        panic(fmt.Sprintf("Unrecognized opcode %v!", context.opcode))
    }
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

func opsf(context *Context) {
    switch context.opcode & 0x00FF {
    case 0x07:
        mvfd(context)
    case 0x0A:
        ldk(context)
    case 0x15:
        mvtd(context)
    case 0x18:
        mvts(context)
    case 0x1E:
        addi(context)
    case 0x29:
        ldf(context)
    case 0x33:
        stbcd(context)
    case 0x55:
        st(context)
    case 0x65:
        ld(context)
    default:
        panic(fmt.Sprintf("Unrecognized opcode %v!", context.opcode))
    }
    context.cpu.pc++
}

// Fx07 - MV Vx, DT
// Set Vx = delay timer value
func mvfd(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    context.cpu.v[x] = context.cpu.dt
    context.cpu.pc++
}

// Fx15 - MV DT, Vx
// Set delay timer = Vx.
func mvtd(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    context.cpu.dt = context.cpu.v[x]
    context.cpu.pc++
}

// Fx0A - LD Vx, K
// Wait for a key press, store the value of the key in Vx
func ldk(context *Context) {
    // TODO
}

// Fx18 - MV ST, Vx
// Set sound timer = Vx
func mvts(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    context.cpu.st = context.cpu.v[x]
    context.cpu.pc++
}

// Fx1E - ADD I, Vx
// Set I = I + Vx
func addi(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    context.cpu.i += uint16(context.cpu.v[x])
    context.cpu.pc++
}

// Fx29 - LD F, Vx
// Set I = location of sprite for digit Vx
func ldf(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    context.cpu.i += uint16(context.cpu.v[x])
    context.cpu.pc++
}

// Fx33 - LD B, Vx
// Store BCD representation of Vx in memory locations I, I+1, and I+2
func stbcd(context *Context) {
    x := context.opcode & 0x0F00 >> 8
    num := context.cpu.v[x]
    context.memory[context.cpu.i] = byte((num / 100) % 10)
    context.memory[context.cpu.i + 1] = byte((num / 10) % 10)
    context.memory[context.cpu.i + 2] = byte(num % 10)
}

// Fx55 - ST [I], Vx
// Store registers V0 through Vx in memory starting at location I
func st(context *Context) {
    // TODO
}

// Fx65 - LD Vx, [I]
// Read registers V0 through Vx from memory starting at location I
func ld(context *Context) {
    // TODO
}

