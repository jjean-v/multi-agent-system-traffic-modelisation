package agt

import (
	"fmt"
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

type Agent interface {
	Percept()
	Deliberate()
	Act()
}

type BaseAgent struct {
	id int

	pos   graph.Vec2f
	speed graph.Vec2f
	acc   graph.Vec2f

	phone    bool
	inAHurry bool

	itinary []*graph.Node

	hitbox      int
	confortZone int

	fieldOfVision FieldOfVision
	goal          int //c'est un point of interest normalement, j'en ai pas besoin pour faire env

	inBox chan com.Message
}

// Méthodes communes
func (b *BaseAgent) GetId() int {
	return b.id
}

func (b *BaseAgent) GetPos() graph.Vec2f {
	return b.pos
}

func (b *BaseAgent) GetSpeed() graph.Vec2f {
	return b.speed
}

func (b *BaseAgent) GetFieldOfVision() FieldOfVision {
	return b.fieldOfVision
}

func (b *BaseAgent) GetItinary() []*graph.Node {
	return b.itinary
}

func (b *BaseAgent) SetItinary(travel []*graph.Node) {
	b.itinary = travel
}

// Appelé quand on utilie fmt.println(b)
func (b *BaseAgent) String() string {
	return fmt.Sprintf("ID: %d, Pos: (%.2f,%.2f), Speed: (%.2f,%.2f), On a phone: %t, In a Hurry: %t, goal: %d", b.id, b.pos.GetX(), b.pos.GetY(), b.speed.GetX(), b.speed.GetY(), b.phone, b.inAHurry, b.goal)
}
