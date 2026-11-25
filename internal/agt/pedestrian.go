package agt

import (
	"fmt"
	"log"
	"project/internal/com"
	"project/internal/graph"
)

type PedestrianState int

const (
	Walking PedestrianState = iota
	Thinking
	Stop
	Dead
	Initializing
)

var pedestrianStateName = map[PedestrianState]string{
	Walking:      "Walking",
	Thinking:     "Thinking",
	Stop:         "Stop",
	Dead:         "Dead",
	Initializing: "Initializing",
}

func (ps PedestrianState) String() string {
	return pedestrianStateName[ps]
}

type Pedestrian struct {
	BaseAgent
	state PedestrianState
}

func NewPedestrian(id int, pos, speed graph.Vec2f, fov FieldOfVision, inbox chan (com.Message)) *Pedestrian {
	return &Pedestrian{BaseAgent: BaseAgent{id: id, pos: pos, speed: speed, fieldOfVision: fov, inBox: inbox}, state: Initializing}
}

func (p *Pedestrian) Percept() {
	log.Println("Pedestrian :", p.id, "Is Percepting")
}

func (p *Pedestrian) Deliberate() {
	log.Println("Pedestrian :", p.id, "Is Deliberating")
}

func (p *Pedestrian) Act() {
	log.Println("Pedestrian :", p.id, "Is Acting")
}

func (p *Pedestrian) String() string {
	return fmt.Sprintf("Pieton[%s, state: %s]", p.BaseAgent.String(), p.state)
}
