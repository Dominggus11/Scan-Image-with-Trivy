package trivy

import (
	"os"
	"os/exec"
)

func GoTrivy() {
	// to know where we now
	// cmdPwd := exec.Command("cd")
	// cmdPwd.Stdout = os.Stdout
	// cmdPwd.Run()

	cmd := exec.Command("pwd")
	cmd.Stdout = os.Stdout
	cmd.Run()

	cmdTrivy := exec.Command("trivy", "config", "-f", "json", "-o", "/home/rdam/Project_Trivy/DockerFile/resultsImage.json", "Dockerfile")
	cmdTrivy.Dir = "/home/rdam/build/"
	cmdTrivy.Stdout = os.Stdout
	cmdTrivy.Run()
}
