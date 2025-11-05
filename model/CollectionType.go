package model

type CollectionType interface {
	Save(filePath string)
	Print()
	Clone() CollectionType
}
