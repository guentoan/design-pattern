package prototype

import (
	"fmt"
	"math/rand"
	"time"
)

var colors = []string{
	"red",
	"blue",
	"green",
	"yellow",
	"black",
	"white",
}

type ICar interface {
	Clone() ICar
	GetName() string
	GetColor() string
	CreatedAt() time.Time
	GetModel() string
	Description()
}

type Car struct {
	name, model, color string
	createdAt          time.Time
}

func (t *Car) GetModel() string {
	return t.model
}

func (t *Car) Description() {
	fmt.Printf("%s - %s  %s  build at: %s\n",
		t.name, t.model, t.color, t.createdAt.Format(time.RFC3339Nano))
}

func (t *Car) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Car) GetName() string {
	return t.name
}

func (t *Car) GetColor() string {
	return t.color
}

func (t *Car) Clone() ICar {
	return &Car{
		name:      t.name,
		model:     t.model,
		color:     colors[rand.Intn(len(colors))],
		createdAt: time.Now(),
	}
}
