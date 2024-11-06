package models

type Package struct {
	Weight              float64 `json:"weight"`
	Length              float64 `json:"length"`
	Width               float64 `json:"width"`
	Height              float64 `json:"height"`
	Contents            string  `json:"contents"`
	IsFragile           bool    `json:"is_fragile"`
	SpecialRequirements string  `json:"special_requirements"`
}