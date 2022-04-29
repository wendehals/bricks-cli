package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type saveable interface {
	Save(fileName string)
}

type abstractSaveable struct{}

var _ saveable = &abstractSaveable{}

func (t *abstractSaveable) Save(fileName string) {
	data, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		log.Fatalf("serializing to JSON failed: %s", err.Error())
	}

	ioutil.WriteFile(fileName, data, os.ModePerm)
	if err != nil {
		log.Fatalf("exporting collection to JSON file '%s' failed: %s", fileName, err.Error())
	}

	log.Printf("Exported JSON file to '%s'\n", fileName)

}
