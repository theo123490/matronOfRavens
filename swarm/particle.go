package swarm

import "gonum.org/v1/gonum/mat"

type Particle struct {
	position       *mat.Dense
	personalBest   *mat.Dense
	globalBest     *mat.Dense
	velocity       *mat.Dense
	velocityWeight VelocityWeight
}

type VelocityWeight struct {
	weight                     *mat.Dense
	personalBestVelocityWeight *mat.Dense
	globalBestVelocityWeight   *mat.Dense
}

func newDimensionDense(dimension int, data []float64) *mat.Dense {
	return mat.NewDense(dimension, 1, data)
}
