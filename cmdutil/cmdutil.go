package cmdutil

import (
	"bytes"
	"os/exec"
	"strings"
)

// ExecCommand 执行cmd命令
// use shell /bin/bash -c to execute command
func ExecCommand(command string) (stdout, stderr string, err error) {
	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Stdout = &out
	cmd.Stderr = &errOut
	err = cmd.Run()
	if err != nil {
		stderr = errOut.String()
	}
	stdout = out.String()
	stdout = strings.Trim(stdout, "\n")
	return
}
