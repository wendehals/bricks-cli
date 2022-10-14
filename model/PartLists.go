package model

// PartLists represents all the user's part lists
type PartLists struct {
	User      string     `json:"user"`
	PartLists []PartList `json:"part_lists"`
}

func NewPartLists() *PartLists {
	return &PartLists{
		PartLists: []PartList{},
	}
}
