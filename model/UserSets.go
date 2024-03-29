package model

// UserSets represents all the user's sets
type UserSets struct {
	User string    `json:"user"`
	ID   uint      `json:"set_list_id"`
	Name string    `json:"set_list_name"`
	Sets []UserSet `json:"sets"`
}

func NewUserSets() *UserSets {
	return &UserSets{
		Sets: []UserSet{},
	}
}

func (s *UserSets) Save(filePath string) {
	Save(s, filePath)
}
