package model

import "strings"

// Color represents the color of a Lego part.
type Color struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Color) Compare(other *Color) int {
	return strings.Compare(c.Name, other.Name)
}
