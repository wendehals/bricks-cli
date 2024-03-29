package model

// SetLists represents all the user's set lists
type SetLists struct {
	User     string    `json:"user"`
	SetLists []SetList `json:"set_lists"`
}

func NewSetLists() *SetLists {
	return &SetLists{
		SetLists: []SetList{},
	}
}

func (s *SetLists) Save(filePath string) {
	Save(s, filePath)
}
