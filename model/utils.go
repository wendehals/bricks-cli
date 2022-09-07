package model

import (
	"encoding/json"
	"log"
	"os"
)

func Save(v interface{}, fileName string) {
	data, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		log.Fatalf("serializing to JSON failed: %s", err.Error())
	}

	os.WriteFile(fileName, data, os.ModePerm)
	if err != nil {
		log.Fatalf("exporting collection to JSON file '%s' failed: %s", fileName, err.Error())
	}

	log.Printf("Saved data to file '%s'\n", fileName)
}
