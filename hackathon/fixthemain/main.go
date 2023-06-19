package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = true
	return true
}

func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?")
	return !Door.state
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	ptrDoor.state = false
	return false
}

func main() {
	door := &Door{}

	OpenDoor(door)
	PrintStr("\n")
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	PrintStr("\n")
	if IsDoorOpen(door) {
		PrintStr("\n")
		CloseDoor(door)
	}
	PrintStr("\n")
	if !door.state {
		PrintStr("\n")
		CloseDoor(door)

	}
}
