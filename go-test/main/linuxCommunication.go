package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("cmd", "C:/Users/THUNDER/Downloads/kkFileView-2.2.0-SNAPSHOT/kkFileView-2.2.0-SNAPSHOT/bin/startup.bat")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

//fmt.Println(runtime.GOOS)
//cmd := exec.Command("/bin/bash", "-c", `cat /etc/issue`)
//out, err := cmd.Output()
//LOG.Debug(string(out))
//if err != nil {
//LOG.Debug(err)
//LOG.Error("System Information Acquisition Failure", err)
//}
//if strings.Contains(string(out), "SUSE") {
//LOG.Debug("System Information Acquisition Success")
//cmdString = FormatCmd(CONVERT_COMMAND_ENV, names, values)
//} else {
//LOG.Debug("System Information Acquisition Success And No SUSE")
//cmdString = FormatCmd(CONVERT_COMMAND_ENV_FS, names, values)
//LOG.Debug("图片转换测试", names, values)
//}
//}
