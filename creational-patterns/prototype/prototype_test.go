package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPrototype(t *testing.T) {
	toyota := &Car{
		name:      "Toyota",
		color:     "Red",
		model:     "Roomy",
		createdAt: time.Now(),
	}
	const count = 10
	var cars []ICar
	for i := 0; i < count; i++ {
		car := toyota.Clone()
		assert.Equal(t, toyota.name, car.GetName())
		assert.Equal(t, toyota.model, car.GetModel())
		assert.NotEqual(t, toyota.createdAt, car.CreatedAt())
		cars = append(cars, car)
		car.Description()
	}
	assert.Equal(t, count, len(cars))
}
