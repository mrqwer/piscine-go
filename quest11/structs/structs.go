package structs

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

type NodeI struct {
	Data int
	Next *NodeI
}
