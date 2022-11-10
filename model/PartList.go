package model

// PartList represents a user's part list.
type PartList struct {
	ID          uint   `json:"id"`
	IsBuildable bool   `json:"is_buildable"`
	Name        string `json:"name"`
	NumParts    uint   `json:"num_parts"`
}

func (s *PartList) Save(filePath string) {
	Save(s, filePath)
}
