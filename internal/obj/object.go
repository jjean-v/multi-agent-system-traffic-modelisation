package obj

import "project/internal/graph"

type Object struct {
	id  string
	pos graph.Vec2f
}

func NewObject(id string, pos graph.Vec2f) Object {
	return Object{id, pos}
}

func (o *Object) getId() string {
	return o.id
}

func (o *Object) getPos() graph.Vec2f {
	return o.pos
}
