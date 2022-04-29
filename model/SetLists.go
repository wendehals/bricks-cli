package model

// SetLists represents all the user's set lists
type SetLists struct {
	*abstractSaveable

	User     string    `json:"user"`
	SetLists []SetList `json:"set_lists"`
}
