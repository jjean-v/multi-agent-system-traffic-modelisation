package graph

import "fmt"

type Node struct {
	id     uint
	pos    Vec2f
	arcs   []*Arc
	radius uint
}

func NewNode(id uint, pos Vec2f, r uint) Node {
	return Node{id: id, pos: pos, radius: r}
}

// To be deleted after, I think
func (n *Node) AddArc(a *Arc) {
	n.arcs = append(n.arcs, a)
}

func (n *Node) Print() {
	fmt.Printf("Node:\n")
	fmt.Printf("\tId: %d\n", n.id)
	fmt.Printf("\tRadius: %d\n", n.radius)
	n.pos.Print()
	fmt.Printf("\tArc: \n")
	for arc := range n.arcs {
		n.arcs[arc].PrintIndent("\t\t")
	}
	fmt.Println()

}
