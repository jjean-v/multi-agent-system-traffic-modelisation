package graph

import "testing"

func TestAddAndDeleteArc(t *testing.T) {
	g := NewGraph()
	n1 := g.CreateNode(0, 0, 1)
	n2 := g.CreateNode(10, 0, 1)

	if !g.AddConnection(n1.id, n2.id) {
		t.Fatal("AddConnection should succeed")
	}
	if g.AddConnection(n1.id, 99) {
		t.Fatal("AddConnection with unknown node should fail")
	}
	if len(g.arcs) != 1 || len(n1.arcs) != 1 {
		t.Fatalf("expected 1 arc in map and in node, got %d / %d", len(g.arcs), len(n1.arcs))
	}

	arc := n1.arcs[0]
	if arc.from != n1 || arc.nextNode != n2 {
		t.Fatal("arc from/nextNode wrong")
	}
	if g.GetArcByID(arc.id) != arc {
		t.Fatal("arc in map and in node should be the same pointer")
	}

	if !g.DeleteArc(arc.id) {
		t.Fatal("DeleteArc should succeed")
	}
	if len(g.arcs) != 0 || len(n1.arcs) != 0 {
		t.Fatalf("arc not removed everywhere: %d / %d", len(g.arcs), len(n1.arcs))
	}
	if g.DeleteArc(arc.id) {
		t.Fatal("deleting twice should fail")
	}
}
