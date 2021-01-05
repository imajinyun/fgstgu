package main

// -> dlv debug
// -> break main.man
// -> disassemble
// -> break main.asmHelloWorld
// -> continue
// -> regs
// -> print *(*[5]byte)(uintptr(0xxxx...))
func main() { asmHelloWorld() }

func output(s string) { println(s) }
func asmHelloWorld()
