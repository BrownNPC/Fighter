package c

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Safe modulo that behaves like you would expect in a calculator
func Modulo[T Integer](x, y T) T {
	return (x + y) % y
}
