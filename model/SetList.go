package model

// SetList represents a user's set list.
type SetList struct {
	ID          uint   `json:"id"`
	IsBuildable bool   `json:"is_buildable"`
	Name        string `json:"name"`
	NumSets     uint   `json:"num_sets"`
}

func (s *SetList) Save(filePath string) {
	Save(s, filePath)
}
