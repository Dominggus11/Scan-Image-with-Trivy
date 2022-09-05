package trivy

import (
	"fmt"
	"io/ioutil"
	"os"

	// "github.com/Jeffail/gabs/v2"
	"github.com/Jeffail/gabs"
	"github.com/tidwall/gjson"
)

func Gjson() {
	// Open our jsonFile
	jsonFile, err := os.Open("fileJson/misconfig.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened Misconfig.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	value := gjson.Get(string(byteValue), "Results")
	println(value.String())

}

func Gabs() {
	// Open our Gabs
	jsonFile, err := gabs.ParseJSONFile("fileJson/misconfig.json")
	if err != nil {
		fmt.Println("Tidak sukses")
	}
	value := jsonFile.Path("Results.Target")
	fmt.Println(value)
}
