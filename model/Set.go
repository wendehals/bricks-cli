package model

import (
	"math"
	"regexp"
	"strconv"
)

// Set represents a Lego set.
type Set struct {
	Number   string `json:"set_num"`
	Name     string `json:"name"`
	Year     uint   `json:"year"`
	NumParts uint   `json:"num_parts"`
	URL      string `json:"set_url"`
	ImageURL string `json:"set_img_url"`
}

func (s *Set) Compare(other *Set) int {
	number_s, variant_s := s.splitNumber()
	number_other, variant_other := other.splitNumber()

	if number_s < number_other {
		return -1
	} else if number_s > number_other {
		return 1
	} else {
		if variant_s < variant_other {
			return -1
		} else if variant_s > variant_other {
			return 1
		}
	}

	return 0
}

func (s *Set) splitNumber() (int, int) {
	regex := regexp.MustCompile(`(\d+)-(\d+)`)
	match := regex.FindStringSubmatch(s.Number)
	if match == nil {
		return math.MaxInt, 0
	}

	number, _ := strconv.Atoi(match[1])
	variant, _ := strconv.Atoi(match[2])

	return number, variant
}
