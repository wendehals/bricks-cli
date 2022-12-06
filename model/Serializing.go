package model

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Load reads a model from a JSON encoded file. In case of an error it terminates with a fatal log entry.
func Load[V any](fileName string) V {
	v, err := LoadE[V](fileName)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

// LoadE reads a model from a JSON encoded file. In case of an error it returns it.
func LoadE[V any](fileName string) (V, error) {
	v := new(V)

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return *v, err
	}
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return *v, err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return *v, err
	}

	return *v, nil
}

// Save writes a model into a JSON encoded file. In case of an error it terminates with a fatal log entry.
func Save(v any, fileName string) {
	err := SaveE(v, fileName)
	if err != nil {
		log.Fatalf("serializing to JSON failed: %s", err.Error())
	}
}

// Save writes a model into a JSON encoded file. In case of an error it returns it.
func SaveE(v any, fileName string) error {
	data, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return fmt.Errorf("serializing to JSON failed: %s", err.Error())
	}

	fileName = filepath.FromSlash(fileName)
	os.WriteFile(fileName, data, os.ModePerm)
	if err != nil {
		return fmt.Errorf("exporting collection to JSON file '%s' failed: %s", fileName, err.Error())
	}

	log.Printf("Saved data to file '%s'", fileName)

	return nil
}

// Create a deep clone of the given value.
func DeepClone[V any](v V) V {
	origJSON, err := json.Marshal(v)
	if err != nil {
		log.Fatalf(CLONING_FAILED_MSG, err.Error())
	}

	clone := new(V)
	err = json.Unmarshal(origJSON, clone)
	if err != nil {
		log.Fatalf(CLONING_FAILED_MSG, err.Error())
	}

	return *clone
}
