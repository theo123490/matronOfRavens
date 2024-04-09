package swarm

import (
	"reflect"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestSubtractWith(t *testing.T) {
	var dimensions int = 3
	var expectedResult *mat.Dense = newDimensionDense(dimensions, []float64{2, 1, 0})

	var velocityWeight VelocityWeight = VelocityWeight{
		weight:                     newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		personalBestVelocityWeight: newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		globalBestVelocityWeight:   newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
	}

	var particle Particle = Particle{
		position:       newDimensionDense(dimensions, []float64{3.0, 2.0, 1}),
		personalBest:   newDimensionDense(dimensions, []float64{1.0, 1.0, 1}),
		globalBest:     newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
		velocity:       newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
		velocityWeight: velocityWeight,
	}

	result := newDimensionDense(dimensions, nil)
	result.Sub(particle.position, particle.personalBest)
	if !reflect.DeepEqual(expectedResult.RawMatrix(), result.RawMatrix()) {
		t.Error("result and expectedresult are not the sames")
	}
}
