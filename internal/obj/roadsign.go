package obj

import "project/internal/graph"

type RoadSign struct {
	BaseObject
	arcID uint
}

func NewRoadSign(id int, pos graph.Vec2f, arcID uint, tags ...Tag) *RoadSign {
	allTags := append([]Tag{"RoadSign"}, tags...)
	return &RoadSign{
		BaseObject: NewBaseObject(id, pos, allTags...),
		arcID:      arcID,
	}
}

func (rs *RoadSign) GetArcID() uint {
	return rs.arcID
}
