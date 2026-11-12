package graph

import "fmt"

type Node struct {
	Id     uint
	Pos    int
	Arc    []Arc
	radius uint
}

func NewNodePointer(id uint, pos int, r uint) *Node {
	return &Node{Id: id, Pos: pos, radius: r}
}

func NewNode(id uint, pos int, r uint) Node {
	return Node{Id: id, Pos: pos, radius: r}
}

func (n *Node) Print() {
	fmt.Println("Node:")
	fmt.Println("\tId:", n.Id)
	fmt.Println("\tPos:", n.Pos)
	fmt.Println("\tRadius:", n.radius)
	fmt.Printf("\tArc: \n")
	for arc := range n.Arc {
		fmt.Println(arc)
	}

}
