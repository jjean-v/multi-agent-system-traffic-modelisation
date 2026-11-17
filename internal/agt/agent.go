package agt

import (
	"project/internal/com"
	"project/internal/graph"
)

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

func (a *Agent) GetPos() graph.Vec2f {
	return a.pos
}
