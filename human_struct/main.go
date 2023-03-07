package main

import "fmt"

type human struct {
	firstName string
	lastName  string
	info      info
}

type info struct {
	addr  string
	phone string
}

func (h human) print() {
	fmt.Printf("%+v", h)
}

func (h *human) changeName(newName string) {
	// (*h).firstName = newName
	h.firstName = newName
}

func main() {
	mrC := human{
		firstName: "Hai",
		lastName:  "Luong",
		info: info{
			addr:  "abcxyz",
			phone: "+8434198552",
		},
	}

	// Go dùng cơ chế "pass by value"
	// mrC_pointer := &mrC
	// mrC_pointer.changeName("Quoc")
	// Rút gọn
	mrC.changeName("Quoc")
	mrC.print()

}
