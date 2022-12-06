package services

import (
	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/utils"
)

var (
	colors     *model.Colors
	shapes     *model.Shapes
	alternates *model.Alternates
	molds      *model.Molds
	prints     *model.Prints
)

func GetColors(update bool) *model.Colors {
	if colors != nil {
		return colors
	}

	colorsPath := utils.ColorsPath()
	if utils.FileExists(colorsPath) && !update {
		tmpColors := model.Load[model.Colors](colorsPath)
		colors = &tmpColors
		return colors
	}

	colors = DownloadColors()

	return colors
}

func GetShapes(update bool) *model.Shapes {
	if shapes != nil {
		return shapes
	}

	shapesPath := utils.ShapesPath()
	if utils.FileExists(shapesPath) && !update {
		tmpShapes := model.Load[model.Shapes](shapesPath)
		shapes := &tmpShapes
		return shapes
	}

	shapes = DownloadShapes()

	return shapes
}

func GetAlternates(update bool) *model.Alternates {
	if alternates != nil && !update {
		return alternates
	}

	alternatesPath := utils.AlternatesPath()
	if utils.FileExists(alternatesPath) && !update {
		tmpAlternates := model.Load[model.Alternates](alternatesPath)
		alternates = &tmpAlternates
		return alternates
	}

	alternates, molds, prints = DownloadPartRelationships()

	return alternates
}

func GetMolds(update bool) *model.Molds {
	if molds != nil && !update {
		return molds
	}

	moldsPath := utils.MoldsPath()
	if utils.FileExists(moldsPath) && !update {
		tmpMolds := model.Load[model.Molds](moldsPath)
		molds = &tmpMolds
		return molds
	}

	alternates, molds, prints = DownloadPartRelationships()

	return molds
}

func GetPrints(update bool) *model.Prints {
	if prints != nil && !update {
		return prints
	}

	printsPath := utils.PrintsPath()
	if utils.FileExists(printsPath) && !update {
		tmpPrints := model.Load[model.Prints](printsPath)
		prints = &tmpPrints
		return prints
	}

	alternates, molds, prints = DownloadPartRelationships()

	return prints
}
