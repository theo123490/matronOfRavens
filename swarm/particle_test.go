package swarm

import (
	"reflect"
	"testing"
)

func TestSubtractWith(t *testing.T) {
	var expectedResult dimensionValues = dimensionValues{2, 1, 0}

	var velocityWeight VelocityWeight = VelocityWeight{
		weight:                     dimensionValues{1.0, 2.0, 2},
		personalBestVelocityWeight: dimensionValues{1.0, 2.0, 2},
		globalBestVelocityWeight:   dimensionValues{1.0, 2.0, 2},
	}

	var particle Particle = Particle{
		position:       dimensionValues{3.0, 2.0, 1},
		personalBest:   dimensionValues{1.0, 1.0, 1},
		globalBest:     dimensionValues{1.0, 2.0, 1},
		velocity:       dimensionValues{1.0, 2.0, 1},
		velocityWeight: velocityWeight,
	}

	result := particle.position.subtractWith(particle.personalBest)

	reflect.DeepEqual(expectedResult, result)
}
