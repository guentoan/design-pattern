package singleton

import "sync"

type ISingleton interface {
	Increase()
	Decrease()
	SetValue(val int)
	GetValue() int
}

type singleton struct {
	value int
}

func (s *singleton) Increase() {
	s.value++
}

func (s *singleton) Decrease() {
	s.value--
}

func (s *singleton) SetValue(val int) {
	s.value = val
}

func (s *singleton) GetValue() int {
	return s.value
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() ISingleton {
	once.Do(func() {
		instance = &singleton{
			value: 1,
		}
	})

	return instance
}
