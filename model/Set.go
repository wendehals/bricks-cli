package model

// Set represents a Lego set.
type Set struct {
	Number   string `json:"set_num"`
	Name     string `json:"name"`
	Year     uint   `json:"year"`
	NumParts uint   `json:"num_parts"`
	URL      string `json:"set_url"`
	ImageURL string `json:"set_img_url"`
}
