package model

// PartType represents the type of a Lego part.
type PartType struct {
	Number   string `json:"part_num"`
	Name     string `json:"name"`
	PartURL  string `json:"part_url"`
	ImageURL string `json:"part_img_url"`
}

// ColorType represents the color of a Lego part.
type ColorType struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// PartEntry represents an entry of a single part in a set with its quantity and color.
type PartEntry struct {
	Quantity int       `json:"quantity"`
	Part     PartType  `json:"part"`
	Color    ColorType `json:"color"`
	IsSpare  bool      `json:"is_spare"`
}

// SetListEntry represents an entry in the user's sets list.
type SetListEntry struct {
	ID          uint   `json:"id"`
	IsBuildable bool   `json:"is_buildable"`
	Name        string `json:"name"`
	NumSets     uint   `json:"num_sets"`
}

// PartListEntry represents an entry in the user's part lists.
type PartListEntry struct {
	ID          uint   `json:"id"`
	IsBuildable bool   `json:"is_buildable"`
	Name        string `json:"name"`
	NumParts    uint   `json:"num_parts"`
}

// PartLists represents a list of user's part lists
type PartLists struct {
	PartLists []PartListEntry `json:"partLists"`
}

// UsersSet represents a set owned by the user.
type UsersSet struct {
	Quantity       uint    `json:"quantity"`
	IncludesSpares bool    `json:"include_spares"`
	Set            SetType `json:"set"`
}

// SetType represents a Lego set.
type SetType struct {
	SetNum      string `json:"set_num"`
	Name        string `json:"name"`
	Year        uint   `json:"year"`
	NumParts    uint   `json:"num_parts"`
	SetURL      string `json:"set_url"`
	SetImageURL string `json:"set_img_url"`
}

type PartColor struct {
	ColorId  uint   `json:"color_id"`
	ImageURL string `json:"part_img_url"`
}
