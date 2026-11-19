package agt

import (
	"project/internal/com"
	"project/internal/graph"
)

type FieldOfVision struct {
	radius, angle float32
}

func NewFieldOfVision(radius, angle float32) FieldOfVision {
	return FieldOfVision{radius, angle}
}

func (fov FieldOfVision) GetRadius() float32 {
	return fov.radius
}

func (fov FieldOfVision) GetAngle() float32 {
	return fov.angle
}

type Agent struct {
	id int

	pos   graph.Vec2f
	speed graph.Vec2f
	acc   graph.Vec2f

	phone    bool
	inAHurry bool

	itinary graph.Graph

	hitbox      int
	confortZone int

	fieldOfVision FieldOfVision
	goal          int //c'est un point of interest normalement, j'en ai pas besoin pour faire env

	inBox chan com.Message
}

func NewAgent(id int, pos, speed graph.Vec2f, fov FieldOfVision) *Agent {
	return &Agent{
		id:            id,
		pos:           pos,
		speed:         speed,
		fieldOfVision: fov,
		inBox:         make(chan com.Message),
	}
}

func (a *Agent) GetId() int {
	return a.id
}

func (a *Agent) GetPos() graph.Vec2f {
	return a.pos
}

func (a *Agent) GetSpeed() graph.Vec2f {
	return a.speed
}

func (a *Agent) GetFieldOfVision() FieldOfVision {
	return a.fieldOfVision
}
