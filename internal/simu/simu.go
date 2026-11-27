package simu

import (
	"encoding/json"
	"project/internal/graph"
	"project/internal/obj"
)

type Simulation struct {
	graph        graph.Graph
	objects      []obj.Object
	nextObjectID int
}

func NewSimulation() *Simulation {
	return &Simulation{
		graph:   graph.NewGraph(),
		objects: make([]obj.Object, 0),
	}
}

func (s *Simulation) GetGraph() *graph.Graph {
	return &s.graph
}

func (s *Simulation) GetObjects() []obj.Object {
	return s.objects
}

func (s *Simulation) GetObjectByID(id int) obj.Object {
	for _, o := range s.objects {
		if o.GetId() == id {
			return o
		}
	}
	return nil
}

func (s *Simulation) GetObjectsByTag(tag obj.Tag) []obj.Object {
	res := make([]obj.Object, 0)
	for _, o := range s.objects {
		if o.GetTags().HasTag(tag) {
			res = append(res, o)
		}
	}
	return res
}

func (s *Simulation) AddObject(o obj.Object) {
	s.objects = append(s.objects, o)
}

func (s *Simulation) AddRoad(inID uint, outID uint) *obj.Road {
	if !s.graph.AddConnection(inID, outID) {
		return nil
	}
	n1 := s.graph.GetNodeByID(inID)
	n2 := s.graph.GetNodeByID(outID)
	mid := n1.GetPose().Add(n2.GetPose()).Scale(0.5)

	road := obj.NewRoad(s.nextObjectID, mid, inID, outID)
	s.nextObjectID++
	s.objects = append(s.objects, road)
	return road
}

func (s *Simulation) AddRoundabout(schema []byte, pos graph.Vec2f) (*obj.Roundabout, error) {
	var data graph.GraphData
	if err := json.Unmarshal(schema, &data); err != nil {
		return nil, err
	}

	inIDs, outIDs, nodeIDs := s.graph.AddSubGraph(data, pos)
	rb := obj.NewRoundabout(s.nextObjectID, pos, inIDs, outIDs, nodeIDs, schema)
	s.nextObjectID++
	s.objects = append(s.objects, rb)
	return rb, nil
}

func (s *Simulation) AddRoadSign(pos graph.Vec2f, arcID uint, tags ...obj.Tag) *obj.RoadSign {
	if s.graph.GetArcByID(arcID) == nil {
		return nil
	}
	rs := obj.NewRoadSign(s.nextObjectID, pos, arcID, tags...)
	s.nextObjectID++
	s.objects = append(s.objects, rs)
	return rs
}
