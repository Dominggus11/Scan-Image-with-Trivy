package main

import (
	postfile "ProjectTrivy/postFile"
	"ProjectTrivy/trivy"
)

func main() {

	// for Scan Docker File
	// trivy.GoTrivy()

	// For Get Something JSON File with Module Gabs
	trivy.Gabs()

	// For Get Something JSON File with Module Gjson
	trivy.Gjson()

	//Upload File
	postfile.UploadFile()
}
