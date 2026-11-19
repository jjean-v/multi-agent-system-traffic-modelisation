package env

import (
	"math"
	"project/internal/agt"
	"project/internal/com"
	"project/internal/graph"
	"project/internal/obj"
	"project/internal/utils"
)

type Environment struct {
	mapSize     graph.Vec2f
	agentsGrid  utils.Grid
	objectsGrid utils.Grid

	agents  []*agt.Agent //des pointeurs car on renvoie toujours des pointeurs pour les fonctions get et tout
	objects []*obj.Object

	in chan com.Message //A modifier plus tard je pense mais pour les fonctions de l'env on en a pas besoin mnt. C'est un peu un squelette là
}

func NewEnvironment(mapSize graph.Vec2f, agents []*agt.Agent, objects []*obj.Object) (*Environment, error) {
	agentsGrid, err := utils.NewGrid(mapSize, graph.NewVec2i(10, 10))
	if err != nil {
		return nil, err
	}
	objectsGrid, err := utils.NewGrid(mapSize, graph.NewVec2i(10, 10))
	if err != nil {
		return nil, err
	}

	return &Environment{
		mapSize,
		*agentsGrid,
		*objectsGrid,
		agents,
		objects,
		make(chan com.Message),
	}, nil
}

func (e *Environment) Init() {
	e.UpdateAgentsGrid()
	e.updateObjectsGrid()
}

// Put all agents in the grid
func (e *Environment) UpdateAgentsGrid() {
	e.agentsGrid.Clear()

	for i, a := range e.agents {
		e.agentsGrid.AddTo(a.GetPos(), i)
	}
}

// Put all agents in the grid
func (e *Environment) updateObjectsGrid() {
	e.objectsGrid.Clear()

	for i, o := range e.objects {
		e.objectsGrid.AddTo(o.GetPos(), i)
	}
}

func (e *Environment) GetAgentsAroundMe(agent *agt.Agent) []*agt.Agent {

	var agentsAroundMe []*agt.Agent
	cellsAround := e.agentsGrid.GetCellsAroundMe(agent.GetPos())
	for _, cA := range cellsAround {
		for _, i := range cA {
			a := e.agents[i]
			if agent.GetPos().GetDistanceSquaredWith(a.GetPos()) <= 15*15 { //15 est à redéfinir car on a pas défini ce champ de detection
				if agent.GetId() != a.GetId() {
					agentsAroundMe = append(agentsAroundMe, a)
				}
			}
		}
	}

	return agentsAroundMe
}
func (e *Environment) GetObjectsAroundMe(agent *agt.Agent) []*obj.Object {
	var objectsAroundMe []*obj.Object
	pos := agent.GetPos()

	for _, cell := range e.objectsGrid.GetCellsAroundMe(pos) {
		for _, i := range cell {
			o := e.objects[i]
			if agent.GetId() != o.GetId() && pos.GetDistanceSquaredWith(o.GetPos()) <= 15*15 { // 15 à redéfinir
				objectsAroundMe = append(objectsAroundMe, o)
			}
		}
	}
	return objectsAroundMe
}

// dans le rayon et l'angle de vision de l'agent.
func isInFieldOfVision(agent *agt.Agent, target graph.Vec2f) bool {
	fov := agent.GetFieldOfVision()
	radius := fov.GetRadius()

	if agent.GetPos().GetDistanceSquaredWith(target) > radius*radius {
		return false
	}
	toTarget := target.Sub(agent.GetPos()) // vecteur agent -> cible
	return math.Abs(float64(agent.GetSpeed().GetAngleWith(toTarget))) <= float64(fov.GetAngle())
}

func (e *Environment) GetAgentsISee(agent *agt.Agent) []*agt.Agent {
	var agentsISee []*agt.Agent

	for _, cell := range e.agentsGrid.GetCellsAroundMe(agent.GetPos()) {
		for _, i := range cell {
			a := e.agents[i]
			if agent.GetId() != a.GetId() && isInFieldOfVision(agent, a.GetPos()) {
				agentsISee = append(agentsISee, a)
			}
		}
	}
	return agentsISee
}

func (e *Environment) GetObjectsISee(agent *agt.Agent) []*obj.Object {
	var objectsISee []*obj.Object

	for _, cell := range e.objectsGrid.GetCellsAroundMe(agent.GetPos()) {
		for _, i := range cell {
			o := e.objects[i]
			if agent.GetId() != o.GetId() && isInFieldOfVision(agent, o.GetPos()) {
				objectsISee = append(objectsISee, o)
			}
		}
	}
	return objectsISee
}
