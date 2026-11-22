package graph

import "fmt"

type typeTrafficLaw int

const (
	YieldToTheRight typeTrafficLaw = iota
	Yield
	Stop
	PriorityRoad
	PedestrianPriority
)

var trafficName = map[typeTrafficLaw]string{
	YieldToTheRight:    "YieldToTheRight",
	Yield:              "Yield",
	Stop:               "Stop",
	PriorityRoad:       "PriorityRoad",
	PedestrianPriority: "PedestrianPriority",
}

type rule struct {
	speedLimit  uint
	trafficLaw  typeTrafficLaw
	illegalness uint
}

func NewRule(speedL uint, sign typeTrafficLaw, illegalness uint) rule {
	return rule{speedLimit: speedL, trafficLaw: sign, illegalness: illegalness}
}

func (tt typeTrafficLaw) String() string {
	return trafficName[tt]
}

func (r *rule) PrintIndent(indent string) {
	fmt.Printf(indent+"Speed Limit: %d\n", r.speedLimit)
	fmt.Printf(indent+"Traffic Rule: %s\n", r.trafficLaw)
	fmt.Printf(indent+"Illegalness: %d\n", r.illegalness)

}

func (r *rule) GetSpeedLimit() uint {
	return r.speedLimit
}

func (r *rule) GetTrafficLaw() typeTrafficLaw {
	return r.trafficLaw
}

func (r *rule) Getillegalness() uint {
	return r.illegalness
}

func (r *rule) SetSpeedLimit(speed uint) {
	r.speedLimit = speed
}

func (r *rule) SetTrafficLaw(traffic typeTrafficLaw) {
	r.trafficLaw = traffic
}

func (r *rule) Setillegalness(newillegalness uint) {
	r.illegalness = newillegalness
}
