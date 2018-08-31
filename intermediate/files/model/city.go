package model

import "fmt"

// City represents a city
type City struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

// Sobrescreve a funcao String como toString do java
func (city City) String() string {
	return fmt.Sprintf("%s", city.Name)
}
