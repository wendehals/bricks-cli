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
		log.Printf("Serializing to JSON failed: %s\n", err)
	}

	ioutil.WriteFile(fileName, data, os.ModePerm)
	if err != nil {
		log.Printf("Exporting collection to JSON file '%s' failed: %s\n", fileName, err)
	}

	log.Printf("Exported result to '%s'\n", fileName)
}
