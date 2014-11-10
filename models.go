package main

import (
	"fmt"
)

type Neuron struct {
	id       int
	x        float64
	x2       float64
	group_id int
}

func (n Neuron) String() string {
	// return fmt.Sprintf("Neuron{id: %d}", n.id)
	return fmt.Sprintf("%d", n.id)
}
