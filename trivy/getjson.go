package trivy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Json function
func OpenJSON() {
	// Open our jsonFile
	jsonFile, err := os.Open("DockerFile/misconfig.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened Misconfig.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["ArtifactName"])

}
