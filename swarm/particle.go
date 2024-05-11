package swarm

import (
	"gonum.org/v1/gonum/mat"
)

type Particle struct {
	dimension         int
	position          *mat.Dense
	personalBest      *mat.Dense
	personalBestValue float64
	velocity          *mat.Dense
	velocityWeight    VelocityWeight
}

type VelocityWeight struct {
	weight                     *mat.Dense
	personalBestVelocityWeight *mat.Dense
	globalBestVelocityWeight   *mat.Dense
}

type FitnessFunction func(*mat.Dense) float64

func newDimensionDense(dimension int, data []float64) *mat.Dense {
	return mat.NewDense(dimension, 1, data)
}

func (p Particle) calculateBestVelocity(bestMatrix *mat.Dense, weightMatrix *mat.Dense) *mat.Dense {
	BestDistance := newDimensionDense(p.dimension, nil)
	BestDistance.Sub(p.position, bestMatrix)
	BestVelocity := newDimensionDense(p.dimension, nil)
	BestVelocity.MulElem(BestDistance, weightMatrix)
	return BestVelocity
}

func (p Particle) calculatePersonalBestVelocity() *mat.Dense {
	return p.calculateBestVelocity(p.personalBest, p.velocityWeight.personalBestVelocityWeight)
}

func (p Particle) calculateGlobalBestVelocity(globalBest *mat.Dense) *mat.Dense {
	return p.calculateBestVelocity(globalBest, p.velocityWeight.globalBestVelocityWeight)
}

func (p Particle) calculateTotalVelocity(globalBest *mat.Dense) *mat.Dense {
	totalVelocity := newDimensionDense(p.dimension, nil)
	totalVelocity.Add(p.calculatePersonalBestVelocity(), p.calculateGlobalBestVelocity(globalBest))
	totalVelocity.Add(totalVelocity, p.velocity)
	return totalVelocity
}

func (p Particle) calculatePositionValue(fitnessFn FitnessFunction) float64 {
	return fitnessFn(p.position)
}
