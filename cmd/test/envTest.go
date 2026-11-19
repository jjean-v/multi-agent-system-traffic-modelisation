package main

import (
	"fmt"
	"math"

	"project/internal/agt"
	"project/internal/env"
	"project/internal/graph"
	"project/internal/obj"
)

func main() {
	fov := agt.NewFieldOfVision(20, math.Pi/4)
	right := graph.NewVec2f(1, 0) // l'agent 0 regarde vers la droite

	agents := []*agt.Agent{
		agt.NewAgent(0, graph.NewVec2f(50, 50), right, fov), // l'agent testé
		agt.NewAgent(1, graph.NewVec2f(55, 50), right, fov), // devant, proche
		agt.NewAgent(2, graph.NewVec2f(45, 50), right, fov), // derrière, proche
		agt.NewAgent(3, graph.NewVec2f(90, 90), right, fov), // loin
	}
	objects := []*obj.Object{
		obj.NewObject(100, graph.NewVec2f(52, 50)), // devant, proche
		obj.NewObject(101, graph.NewVec2f(48, 50)), // derrière, proche
		obj.NewObject(102, graph.NewVec2f(10, 10)), // loin
	}

	e, err := env.NewEnvironment(graph.NewVec2f(100, 100), agents, objects)
	if err != nil {
		fmt.Println(err)
		return
	}

	e.Init()

	me := agents[0]
	fmt.Println("Agents autour (attendu 1, 2) :")
	for _, a := range e.GetAgentsAroundMe(me) {
		fmt.Println("  ", a.GetId())
	}
	fmt.Println("Objets autour (attendu 100, 101) :")
	for _, o := range e.GetObjectsAroundMe(me) {
		fmt.Println("  ", o.GetId())
	}
	fmt.Println("Agents vus (attendu 1) :")
	for _, a := range e.GetAgentsISee(me) {
		fmt.Println("  ", a.GetId())
	}
	fmt.Println("Objets vus (attendu 100) :")
	for _, o := range e.GetObjectsISee(me) {
		fmt.Println("  ", o.GetId())
	}
}
