package module_b

import (
	c "github.com/KeiichiHirobe/gomodule-example-c"
)

func GetThroughB() string {
	return c.GetVersionConfirm()
}

func SetThroughB(s string) {
	c.SetVersionConfirm(s)
}

func GetDecolatedThroughB() string {
	return c.GetDecolatedVersionConfirm()
}
