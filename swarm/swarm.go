package swarm

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Swarm struct {
	dimension       int
	Particles       []Particle
	GlobalBest      mat.Dense
	GlobalBestValue float64
	VelocityWeight  VelocityWeight
}

func CreateSwarm(v VelocityWeight, d int) (Swarm, error) {
	var s Swarm = Swarm{
		dimension:       d,
		Particles:       []Particle{},
		GlobalBest:      *mat.NewDense(d, 1, nil),
		GlobalBestValue: math.SmallestNonzeroFloat64,
		VelocityWeight:  v,
	}
	return s, nil
}

func (s *Swarm) UpdateGlobalBest() error {
	if s.GlobalBestValue == math.SmallestNonzeroFloat64 {
		s.GlobalBestValue = s.Particles[0].personalBestValue
	}
	for _, particle := range s.Particles {
		if particle.personalBestValue > s.GlobalBestValue {
			s.GlobalBestValue = particle.personalBestValue
		}
	}
	return nil
}
