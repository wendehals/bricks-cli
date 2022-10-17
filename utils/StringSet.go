package utils

import "golang.org/x/exp/slices"

type StringSet struct {
	Values map[string]struct{} `json:"values"`
}

func NewStringSet() StringSet {
	return StringSet{
		Values: make(map[string]struct{}),
	}
}

func (s StringSet) Add(value string) {
	s.Values[value] = struct{}{}
}

func (s StringSet) Contains(value string) bool {
	_, found := s.Values[value]
	return found
}

func (s StringSet) Representative() string {
	if len(s.Values) > 0 {
		keys := make([]string, 0, len(s.Values))
		for k := range s.Values {
			keys = append(keys, k)
		}
		slices.Sort(keys)

		return keys[0]
	}
	return ""
}
