package graph

import "fmt"

type Arc struct {
	id       uint
	rule     rule
	nextNode *Node
}

func NewArc(id uint, r rule, nxtNode *Node) Arc {
	return Arc{id: id, rule: r, nextNode: nxtNode}
}

func (a *Arc) GetId() uint {
	return a.id
}

func (a *Arc) GetRule() rule {
	return a.rule
}

func (a *Arc) GetNextNode() *Node {
	return a.nextNode
}

func (a *Arc) SetId(id uint) {
	a.id = id
}

func (a *Arc) SetRule(r rule) {
	a.rule = r
}

func (a *Arc) SetNextNode(node *Node) {
	a.nextNode = node
}

func (a *Arc) Print() {
	fmt.Println("Id:", a.id)
	fmt.Println("Rule:", a.rule)
	fmt.Println("NextNode:", a.nextNode.id)
}

func (a *Arc) PrintIndent(indent string) {
	fmt.Printf(indent+"Id: %d\n", a.id)
	fmt.Println(indent + "Rule:")
	a.rule.PrintIndent(indent + "\t")
	fmt.Printf(indent+"NextNode: %d\n", a.nextNode.id)
}
