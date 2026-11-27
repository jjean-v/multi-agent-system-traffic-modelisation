package obj

import "project/internal/graph"

type Road struct {
	BaseObject
	inID  uint
	outID uint
}

func NewRoad(id int, pos graph.Vec2f, inID uint, outID uint) *Road {
	return &Road{
		BaseObject: NewBaseObject(id, pos, "Road"),
		inID:       inID,
		outID:      outID,
	}
}

func (r *Road) GetInID() uint {
	return r.inID
}

func (r *Road) GetOutID() uint {
	return r.outID
}
