package util

func Condition[T any](v bool, t, f T) T {
	if v {
		return t
	} else {
		return f
	}
}

func If(v bool, t func()) {
	if v {
		t()
	}
}

func IfElse(v bool, t, f func()) {
	if v {
		t()
	} else {
		f()
	}
}
