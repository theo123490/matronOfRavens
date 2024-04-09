package swarm

type Swarm struct {
	Particles      []Particle
	GlobalBest     dimensionValues
	VelocityWeight VelocityWeight
}

// func (s Swarm) calculateGlobalBest() dimensionValues {

// }
