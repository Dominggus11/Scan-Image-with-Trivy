package trivy

import (
	"os"
	"os/exec"
)

func GoTrivy() {
	//Code to Run Trivy In Golang
	cmdTrivy := exec.Command("trivy", "config", "-f", "json", "-o", "/home/rdam/Project_Trivy/DockerFile/resultsImage.json", "Dockerfile")
	cmdTrivy.Dir = "/home/rdam/build/"
	cmdTrivy.Stdout = os.Stdout
	cmdTrivy.Run()
}

func TrivyUpload() {
	//Code to Run Trivy In Golang
	cmdTrivy := exec.Command("trivy", "config", "-f", "json", "-o", "fileUpload/resultsImage.json", "Dockerfile")
	cmdTrivy.Dir = "/home/rdam/build/"
	cmdTrivy.Stdout = os.Stdout
	cmdTrivy.Run()
}

func TrivyWrite() {
	//Code to Run Trivy In Golang
	cmdTrivy := exec.Command("trivy", "config", "-f", "json", "-o", "/home/rdam/Project_Trivy/DockerFile/resultsImage.json", "Dockerfile")
	cmdTrivy.Dir = "/home/rdam/build/"
	cmdTrivy.Stdout = os.Stdout
	cmdTrivy.Run()
}
