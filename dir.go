package main

type Dir int

const (
	DIR_RIGHT Dir = iota
	DIR_LEFT
	DIR_UP
	DIR_DOWN
)

func (dir Dir) String() string {
	switch dir {
	case DIR_RIGHT:
		return "R"
	case DIR_LEFT:
		return "L"
	case DIR_UP:
		return "U"
	case DIR_DOWN:
		return "D"
	default:
		panic(dir) // unreachable
	}
}
