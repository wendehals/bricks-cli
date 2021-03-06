package model

// UserSet represents a set owned by the user.
type UserSet struct {
	Quantity       uint `json:"quantity"`
	IncludesSpares bool `json:"include_spares"`
	Set            Set  `json:"set"`
}
