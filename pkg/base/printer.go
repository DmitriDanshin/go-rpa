package base

import "fmt"

type Printer struct{}

func (p Printer) Print(value string) {
	fmt.Println(value)
}
