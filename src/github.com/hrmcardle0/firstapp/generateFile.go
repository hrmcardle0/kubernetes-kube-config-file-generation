package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func generateJson(kc* kubeConfig) (retValue string) {

	// Generate final Json File
	file, _ := json.MarshalIndent(kc, "", " ")
 
	// Write to output
	if err := ioutil.WriteFile("kubeConfig.json", file, 0644); err != nil {
		log.Fatal(err)
		return "Failure: See log for details"
	}

	return "Success"
}
