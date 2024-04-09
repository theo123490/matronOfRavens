package swarm

import "gonum.org/v1/gonum/mat"

type Swarm struct {
	Particles      []Particle
	GlobalBest     mat.Dense
	VelocityWeight VelocityWeight
}

// func (s Swarm) calculateGlobalBest() dimensionValues {

// }
