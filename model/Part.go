package model

import "regexp"

// Part represents the type (shape) of a Lego part.
type Part struct {
	Number   string `json:"part_num"`
	Name     string `json:"name"`
	PartURL  string `json:"part_url"`
	ImageURL string `json:"part_img_url"`
}

func (p *Part) Dimensions() string {
	if found, result := p.match(`^(?:Technic )?Brick (\d+ x \d+)`); found {
		return result
	}

	if found, result := p.match(`^(?:Technic )?Plate (\d+ x \d+)`); found {
		return result
	}

	if found, result := p.match(`^Tile (\d+ x \d+)`); found {
		return result
	}

	if found, result := p.match(`Technic Axle (\d+(?:\.\d)?)`); found {
		return result
	}

	if found, result := p.match(`Technic Panel Fairing (# ?\d+)`); found {
		return result
	}

	if found, result := p.match(`Technic Beam 1 x (\d+)`); found {
		r := regexp.MustCompile(`Bent`)
		match := r.FindStringSubmatch(p.Name)
		if match == nil {
			return result
		}
	}

	if found, result := p.match(`Hose.* (\d+(?:\.\d)?(?:cm|mm))`); found {
		return result
	}

	if found, result := p.match(`^Dish (\d+) x \d+`); found {
		return result
	}

	if found, result := p.match(` (\d+)L`); found {
		return result
	}

	if found, result := p.match(` (\d+(?:\.\d)?(?:cm|mm))`); found {
		return result
	}

	return ""
}

func (p *Part) match(regex string) (bool, string) {
	r := regexp.MustCompile(regex)
	match := r.FindStringSubmatch(p.Name)
	if match != nil {
		return true, match[1]
	}

	return false, ""
}
