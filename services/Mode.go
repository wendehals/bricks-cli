package services

import "strings"

const (
	MODE_COLOR      uint8 = 1
	MODE_ALTERNATES uint8 = 2
	MODE_MOLDS      uint8 = 4
	MODE_PRINTS     uint8 = 8
)

func ColorMode(mode uint8) bool {
	return mode&MODE_COLOR != 0
}

func PrintsMode(mode uint8) bool {
	return mode&MODE_PRINTS != 0
}

func MoldsMode(mode uint8) bool {
	return mode&MODE_MOLDS != 0
}

func AlternatesMode(mode uint8) bool {
	return mode&MODE_ALTERNATES != 0
}

func ModeToUInt8(mode string) uint8 {
	var modeUInt8 uint8 = 0
	if strings.Contains(mode, "c") {
		modeUInt8 = modeUInt8 ^ MODE_COLOR
	}
	if strings.Contains(mode, "a") {
		modeUInt8 = modeUInt8 ^ MODE_ALTERNATES
	}
	if strings.Contains(mode, "m") {
		modeUInt8 = modeUInt8 ^ MODE_MOLDS
	}
	if strings.Contains(mode, "p") {
		modeUInt8 = modeUInt8 ^ MODE_PRINTS
	}

	return modeUInt8
}
