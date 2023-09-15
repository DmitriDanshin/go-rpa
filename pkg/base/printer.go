package base

import "log"

type Printer struct{}

func (p Printer) Print(value string) {
	log.Println("Message:", value)
}
