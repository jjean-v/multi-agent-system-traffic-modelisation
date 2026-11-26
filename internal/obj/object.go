package obj

import "project/internal/graph"

type Object interface {
	GetId() int
	GetPos() graph.Vec2f
	GetTags() *TagContainer
}

type BaseObject struct {
	id   int
	pos  graph.Vec2f
	tags TagContainer
}

func NewBaseObject(id int, pos graph.Vec2f, tags ...Tag) BaseObject {
	return BaseObject{
		id:   id,
		pos:  pos,
		tags: NewTagContainer(tags...),
	}
}

func (o *BaseObject) GetId() int {
	return o.id
}

func (o *BaseObject) GetPos() graph.Vec2f {
	return o.pos
}

func (o *BaseObject) SetPos(pos graph.Vec2f) {
	o.pos = pos
}

func (o *BaseObject) GetTags() *TagContainer {
	return &o.tags
}
