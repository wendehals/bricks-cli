package model

import (
	"regexp"
	"strconv"
	"strings"
)

// Shape represents the shape of a Lego part.
type Shape struct {
	Number   string `json:"part_num"`
	Name     string `json:"name"`
	URL      string `json:"part_url"`
	ImageURL string `json:"part_img_url"`
}

func (s *Shape) Compare(other *Shape) int {
	return strings.Compare(s.Name, other.Name)
}

func (s *Shape) Dimensions() string {
	if found, result := s.match(`^(?:Technic )?Brick (\d+) x (\d+)`); found {
		if i, err := strconv.Atoi(result[1]); err == nil && i > 4 {
			return result[0] + " x " + result[1]
		}
	}

	if found, result := s.match(`^(?:Technic )?Plate (\d+) x (\d+)`); found {
		if i, err := strconv.Atoi(result[1]); err == nil && i > 4 {
			return result[0] + " x " + result[1]
		}
	}

	if found, result := s.match(`^Tile (\d+) x (\d+)`); found {
		if i, err := strconv.Atoi(result[1]); err == nil && i > 4 {
			return result[0] + " x " + result[1]
		}
	}

	if found, result := s.match(`Technic Axle (\d+(?:\.\d)?)`); found {
		return result[0]
	}

	if found, result := s.match(`Technic Axle and Pin Connector Angled (#\d)`); found {
		return result[0]
	}

	if found, result := s.match(`Technic Panel Fairing # ?(\d+)`); found {
		return "#" + result[0]
	}

	if found, result := s.match(`Technic Beam 1 x (\d+)`); found {
		if i, err := strconv.Atoi(result[0]); err == nil && i > 4 {
			r := regexp.MustCompile(`Bent`)
			match := r.FindStringSubmatch(s.Name)
			if match == nil {
				return result[0]
			}
		}
	}

	if found, result := s.match(`Hose.* (\d+(?:\.\d)?(?:cm|mm))`); found {
		return result[0]
	}

	if found, result := s.match(`^Dish (\d+) x \d+`); found {
		return result[0]
	}

	if found, result := s.match(` (\d+)L`); found {
		return result[0]
	}

	if found, result := s.match(` (\d+(?:\.\d)?(?:cm|mm))`); found {
		return result[0]
	}

	return ""
}

func (s *Shape) match(regex string) (bool, []string) {
	r := regexp.MustCompile(regex)
	match := r.FindStringSubmatch(s.Name)
	if match != nil {
		return true, match[1:]
	}

	return false, []string{}
}
