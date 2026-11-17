package env

import (
	"math"
	"project/internal/agt"
	"project/internal/com"
	"project/internal/obj"
)

type Environment struct {
	Agents  []*agt.Agent //des pointeurs car on renvoie toujours des pointeurs pour les fonctions get et tout
	Objects []*obj.Object
	in      chan com.Message //A modifier plus tard je pense mais pour les fonctions de l'env on en a pas besoin mnt. C'est un peu un squelette là
}

func NewEnvironment(agents []*agt.Agent, objects []*obj.Object) Environment {
	return Environment{agents, objects, make(chan com.Message)}
}

func (e *Environment) GetAgentsAroundMe(agent *agt.Agent) []*agt.Agent {

	var agentsAroundMe []*agt.Agent

	//C'est bien trop gourmand en perf, faudra plus tard bien opti
	for _, a := range e.Agents {
		if agent.GetPos().GetDistanceSquaredWith(a.GetPos()) <= 15*15 { //15 est à redéfinir car on a pas défini ce champ de detection
			if agent.GetId() != a.GetId() {
				agentsAroundMe = append(agentsAroundMe, a)
			}
		}
	}

	return agentsAroundMe
}

func (e *Environment) GetObjectsAroundMe(agent *agt.Agent) []*obj.Object {

	var objectsAroundMe []*obj.Object

	//C'est bien trop gourmand en perf, faudra plus tard bien opti
	for _, o := range e.Objects {
		if agent.GetPos().GetDistanceSquaredWith(o.GetPos()) <= 15*15 { //15 est à redéfinir car on a pas défini ce champ de detection
			if agent.GetId() != o.GetId() {
				objectsAroundMe = append(objectsAroundMe, o)
			}
		}
	}

	return objectsAroundMe
}

func (e *Environment) GetAgentsISee(agent *agt.Agent) []*agt.Agent {

	var agentsISee []*agt.Agent

	radius := agent.GetFieldOfVision().GetRadius()
	angle := float64(agent.GetFieldOfVision().GetAngle())

	//C'est bien trop gourmand en perf, faudra plus tard bien opti
	for _, a := range e.Agents {
		if agent.GetPos().GetDistanceSquaredWith(a.GetPos()) <= radius*radius { //15 est à redéfinir car on a pas défini ce champ de detection
			if agent.GetId() != a.GetId() {
				btw := agent.GetPos().Sub(a.GetPos())

				if math.Abs(float64(agent.GetSpeed().GetAngleWith(btw))) <= angle {
					agentsISee = append(agentsISee, a)
				}
			}
		}
	}

	return agentsISee
}

func (e *Environment) GetObjectsISee(agent *agt.Agent) []*obj.Object {

	var objectsISee []*obj.Object

	radius := agent.GetFieldOfVision().GetRadius()
	angle := float64(agent.GetFieldOfVision().GetAngle())

	//C'est bien trop gourmand en perf, faudra plus tard bien opti
	for _, o := range e.Objects {
		if agent.GetPos().GetDistanceSquaredWith(o.GetPos()) <= radius*radius { //15 est à redéfinir car on a pas défini ce champ de detection
			if agent.GetId() != o.GetId() {
				btw := agent.GetPos().Sub(o.GetPos())

				if math.Abs(float64(agent.GetSpeed().GetAngleWith(btw))) <= angle {
					objectsISee = append(objectsISee, o)
				}
			}
		}
	}

	return objectsISee
}
