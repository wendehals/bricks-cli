package model

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// Load reads a model from a JSON encoded file.
func Load[V any](v V, fileName string) V {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

// Save writes a model into a JSON encoded file.
func Save(v any, fileName string) {
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
