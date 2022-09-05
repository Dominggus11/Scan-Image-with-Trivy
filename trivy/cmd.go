package trivy

import (
	"os"
	"os/exec"
)

func GoTrivy() {
	// to know where we now
	cmdPwd := exec.Command("cd")
	cmdPwd.Stdout = os.Stdout
	cmdPwd.Run()

	cmd := exec.Command("pwd")
	cmd.Stdout = os.Stdout
	cmd.Run()

	cmdTrivy := exec.Command("trivy", "image", "-f", "json", "-o", "resultsImage.json", "golang:1.12-alpine")
	cmdTrivy.Stdout = os.Stdout
	cmdTrivy.Run()
}
