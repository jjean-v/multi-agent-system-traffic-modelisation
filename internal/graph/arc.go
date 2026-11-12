package graph

import "fmt"

type typeTrafficLaw int

const (
	YieldToTheRight typeTrafficLaw = iota
	Yield
	Stop
	PriorityRoad
	PedestrianPriority
)

type rule struct {
	speedLimit uint
	trafficLaw typeTrafficLaw
}

type Arc struct {
	Id       uint
	rule     rule
	nextNode *Node
}

func NewArc(id uint, r rule, nxtNode *Node) *Arc {
	return &Arc{Id: id, rule: r, nextNode: nxtNode}
}

func (a *Arc) Print() {
	fmt.Println("Id: %d", a.Id)
	fmt.Println("Rule: %d", a.rule)
	fmt.Println("NextNode: %d", a.nextNode.Id)
}
