package graph

type Graph struct {
	nodes map[uint]*Node
	arcs  map[uint]*Arc

	nextNodeID uint
	nextArcID  uint
}

func NewGraph() Graph {
	return Graph{
		nodes: make(map[uint]*Node),
		arcs:  make(map[uint]*Arc),
	}
}

func (g *Graph) GetNodeByID(id uint) *Node {
	return g.nodes[id]
}

func (g *Graph) GetArcByID(id uint) *Arc {
	return g.arcs[id]
}

func (g *Graph) GetNodes() map[uint]*Node {
	return g.nodes
}

func (g *Graph) GetArcs() map[uint]*Arc {
	return g.arcs
}

func (g *Graph) CreateNode(x int, y int, radius uint) *Node {
	n := NewNode(g.nextNodeID, NewVec2f(float32(x), float32(y)), radius)
	g.nextNodeID++
	g.nodes[n.id] = n
	return n
}

func (g *Graph) CreateArc(from *Node, NextNode *Node) *Arc {
	arc := NewArc(g.nextArcID, nil, from, NextNode)
	g.nextArcID++
	g.arcs[arc.id] = arc
	return arc
}

func (g *Graph) AddConnection(N1 uint, N2 uint) bool {
	if g.GetNodeByID(N1) == nil || g.GetNodeByID(N2) == nil {
		return false
	} else {
		from := g.GetNodeByID(N1)
		from.AddArc(g.CreateArc(from, g.GetNodeByID(N2)))
		return true
	}
}

func (g *Graph) DeleteArc(ArcID uint) bool {
	arc := g.GetArcByID(ArcID)
	if arc == nil {
		return false
	}
	arc.from.RemoveArc(arc)
	delete(g.arcs, ArcID)
	return true
}

func (g *Graph) IsThereNode(x float32, y float32, radius uint) *Node {
	vecPose := Vec2f{x: x, y: y}
	var dist float32
	for _, node := range g.nodes {
		dist = node.pos.GetDistanceWith(vecPose)

		if dist <= (float32(node.radius + radius)) {
			return node
		}

	}
	return nil
}
