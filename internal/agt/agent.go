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

func (fov *FieldOfVision) GetRadius() float32 {
	return fov.radius
}

func (fov *FieldOfVision) GetAngle() float32 {
	return fov.angle
}

type Agent struct {
	id string

	pos   graph.Vec2f
	speed graph.Vec2f
	acc   graph.Vec2f

	phone    bool
	inAHurry bool

	itinary graph.Graph

	hitbox      int
	confortZone int

	fieldOfVision int //à bien définir c'est juste pour faire env
	goal          int //c'est un point of interest normalement, j'en ai pas besoin pour faire env

	inBox chan com.Message
}

//Constructeur maquant

func (a *Agent) GetID() string {
	return a.id
}

func (a *Agent) GetPos() graph.Vec2f {
	return a.pos
}
