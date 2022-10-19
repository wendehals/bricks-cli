package services

import (
	"github.com/wendehals/bricks/api"
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
		colors = model.Load(model.NewColors(), colorsPath)
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
		shapes = model.Load(model.NewShapes(), shapesPath)
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
		alternates = model.Load(model.NewAlternates(), alternatesPath)
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
		molds = model.Load(model.NewMolds(), moldsPath)
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
		prints = model.Load(model.NewPrints(), printsPath)
		return prints
	}

	alternates, molds, prints = DownloadPartRelationships()

	return prints
}

func GetPartImageURL(part string, bricksAPI *api.BricksAPI) string {
	shapes := GetShapes(false)

	if shape, found := shapes.Shapes[part]; found {
		if shape.ImageURL != "" {
			return shape.ImageURL
		}
		if bricksAPI != nil {
			details := bricksAPI.GetPart(part)
			shape.ImageURL = details.ImageURL
			shapes.Shapes[part] = shape
		}
	}

	return ""
}
