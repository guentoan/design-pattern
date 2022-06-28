package singleton

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	s := GetInstance()
	s.SetValue(0)
	s.Increase()
	assert.Equal(t, 1, s.GetValue())
	s.Decrease()
	assert.Equal(t, 0, s.GetValue())
	s.Decrease()
	assert.Equal(t, -1, s.GetValue())

	s2 := GetInstance()
	assert.Equal(t, -1, s2.GetValue())
	assert.Equal(t, s, s2)
}

func TestParallel(t *testing.T) {
	const count = 100
	ch := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(count)
	instances := [count]ISingleton{}
	for i := 0; i < count; i++ {
		go func(index int) {
			<-ch
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	close(ch)
	wg.Wait()
	for i := 1; i < count; i++ {
		assert.Equal(t, instances[i], instances[i-1])
	}
}
