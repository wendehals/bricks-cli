package model

type CollectionType interface {
	Save(filePath string)
	Print()
}
