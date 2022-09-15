package model

import (
	"embed"
	"encoding/json"
	"io"
	"log"
	"os"
)

//go:embed resources/parts_html.gotpl
var embeddedFS embed.FS

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

// LoadE reads a model from a JSON encoded file. In case of an error it returns it.
func LoadE[V any](v V, fileName string) (V, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return v, err
	}
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return v, err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return v, err
	}

	return v, nil
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

func DeepClone[V any](v V, clone V) V {
	origJSON, err := json.Marshal(v)
	if err != nil {
		log.Fatalf(CLONING_FAILED_MSG, err.Error())
	}

	err = json.Unmarshal(origJSON, clone)
	if err != nil {
		log.Fatalf(CLONING_FAILED_MSG, err.Error())
	}

	return clone
}
