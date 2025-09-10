package c

type Input uint8

const INeutral Input = 0
const (
	IUp Input = 1 << iota
	IDown
	IForward
	IBack
)

