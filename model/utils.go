package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Export writes a struct into a JSON encoded file.
func ExportToJSON(fileName string, v interface{}) {
	data, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		log.Fatalf("serializing to JSON failed: %s", err.Error())
	}

	ioutil.WriteFile(fileName, data, os.ModePerm)
	if err != nil {
		log.Fatalf("exporting collection to JSON file '%s' failed: %s", fileName, err.Error())
	}

	log.Printf("Exported JSON file to '%s'\n", fileName)
}
