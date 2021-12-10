package main

//go:generate stringer -type=Opcode
type Opcode byte

const (
	forward Opcode = iota
	down
	up
	invalid
)

func OpcodeFromString(op string) Opcode {
	switch op {
	case forward.String():
		return forward
	case down.String():
		return down
	case up.String():
		return up
	default:
		return invalid
	}
}
