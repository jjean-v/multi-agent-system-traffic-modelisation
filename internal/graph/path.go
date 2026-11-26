package graph

import (
	"math"

	"github.com/beefsack/go-astar"
)

type Profile struct {
	Lambda float64 // poids du danger
}

type searchNode struct {
	node    *Node
	profile *Profile
}

// Cost function of the A* Algorithm, Euclidian distance + illegalness * security coefficient of the driver(lambda)
// The higher is lambda, the safer drives the agent
func (p *Profile) cost(node *Node, arc *Arc) float64 {
	return p.Lambda*float64(arc.rule.illegalness) + float64(node.pos.GetDistanceWith(arc.nextNode.pos))
}

// Function of the A* package, modified for our needs
func (s searchNode) PathNeighbors() []astar.Pather {
	res := make([]astar.Pather, 0, len(s.node.arcs))
	for _, arc := range s.node.arcs {
		res = append(res, searchNode{arc.nextNode, s.profile})

	}
	return res
}

func (s searchNode) PathNeighborCost(to astar.Pather) float64 {
	t := to.(searchNode).node
	best := math.Inf(1)
	for _, arc := range s.node.arcs {
		if arc.nextNode == t {
			best = math.Min(best, s.profile.cost(s.node, arc))
		}
	}
	return best
}

func (s searchNode) PathEstimatedCost(to astar.Pather) float64 {
	t := to.(searchNode).node
	return float64(s.node.pos.GetDistanceWith(t.pos))
}

// The function to call
func FindPath(from, to *Node, p *Profile) ([]*Node, float64, bool) {
	invertedPath, cost, found := astar.Path(searchNode{from, p}, searchNode{to, p}) // invertedPath is from last to first
	if !found {
		return nil, 0, false
	}
	// Invert the order, first--last
	path := make([]*Node, len(invertedPath))
	for i, x := range invertedPath {
		path[len(invertedPath)-1-i] = x.(searchNode).node
	}
	return path, cost, true
}
