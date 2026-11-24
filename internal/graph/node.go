package graph

import (
	"fmt"
	"slices"
)

type Node struct {
	id     uint
	pos    Vec2f
	arcs   []*Arc
	radius uint
}

func NewNode(id uint, pos Vec2f, r uint) *Node {
	return &Node{id: id, pos: pos, radius: r}
}

func (n *Node) GetId() uint {
	return n.id
}

func (n *Node) GetPose() Vec2f {
	return n.pos
}

func (n *Node) GetArcs() []*Arc {
	return n.arcs
}

func (n *Node) GetRadius() uint {
	return n.radius
}

func (n *Node) SetId(id uint) {
	n.id = id
}

func (n *Node) SetPos(v Vec2f) {
	n.pos = v
}

func (n *Node) SetRadius(r uint) {
	n.radius = r
}

func (n *Node) AddArc(a *Arc) {
	n.arcs = append(n.arcs, a)
}

func (n *Node) RemoveArc(a *Arc) {
	i := slices.Index(n.arcs, a)
	if i != -1 {
		n.arcs = slices.Delete(n.arcs, i, i+1) // retire de i a i+1
	}
}

func (n *Node) Print() {
	fmt.Printf("Node:\n")
	fmt.Printf("\tId: %d\n", n.id)
	fmt.Printf("\tRadius: %d\n", n.radius)
	n.pos.Print()
	if len(n.arcs) != 0 {
		fmt.Printf("\tArc: \n")
		for arc := range n.arcs {
			n.arcs[arc].PrintIndent("\t\t")
		}
	}
	fmt.Println()

}
