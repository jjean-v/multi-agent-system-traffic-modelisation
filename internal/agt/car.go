package agt

import (
	"fmt"
	"log"
	"project/internal/com"
	"project/internal/graph"
)

type CarState int

const (
	Driving CarState = iota
	Breaking
	CStop
	Crash
	CThinking
	CarInitializing
)

var carStateName = map[CarState]string{
	Driving:         "Driving",
	Breaking:        "Breaking",
	CStop:           "Stop",
	Crash:           "Crash",
	CThinking:       "Thinking",
	CarInitializing: "Initializing",
}

func (cs CarState) String() string {
	return carStateName[cs]
}

type Car struct {
	BaseAgent
	state CarState
}

func NewCar(id int, pos, speed graph.Vec2f, fov FieldOfVision, inbox chan (com.Message)) *Car {
	return &Car{BaseAgent: BaseAgent{id: id, pos: pos, speed: speed, fieldOfVision: fov, inBox: inbox}, state: CarInitializing}
}

func (c *Car) Percept() {
	log.Println("Car :", c.id, "Is Percepting")
}

func (c *Car) Deliberate() {
	log.Println("Car :", c.id, "Is Deliberating")
}

func (c *Car) Act() {
	log.Println("Car :", c.id, "Is Acting")
}

func (c *Car) String() string {
	return fmt.Sprintf("Car[%s, state: %s]", c.BaseAgent.String(), c.state)
}
