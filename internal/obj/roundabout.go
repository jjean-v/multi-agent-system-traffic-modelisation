package obj

import "project/internal/graph"

type Roundabout struct {
	BaseObject
	inIDs   []uint
	outIDs  []uint
	nodeIDs []uint
	schema  []byte
}

func NewRoundabout(id int, pos graph.Vec2f, inIDs []uint, outIDs []uint, nodeIDs []uint, schema []byte) *Roundabout {
	return &Roundabout{
		BaseObject: NewBaseObject(id, pos, "Road.Roundabout"),
		inIDs:      inIDs,
		outIDs:     outIDs,
		nodeIDs:    nodeIDs,
		schema:     schema,
	}
}

func (rb *Roundabout) GetInIDs() []uint {
	return rb.inIDs
}

func (rb *Roundabout) GetOutIDs() []uint {
	return rb.outIDs
}

func (rb *Roundabout) GetNodeIDs() []uint {
	return rb.nodeIDs
}

func (rb *Roundabout) GetSchema() []byte {
	return rb.schema
}
