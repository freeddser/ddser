package util

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func ExecCMD(cmdstr string) {
	cmd := exec.Command("/bin/bash", "-c", cmdstr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: ", err.Error())
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		return
	}
	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println("ReadAll stderr: ", err.Error())
		return
	}
	if len(bytesErr) != 0 {
		fmt.Printf("stderr is not nil: %s", bytesErr)
		return
	}
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: ", err.Error())
		return
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait: ", err.Error())
		return
	}
	fmt.Printf("stdout: %s", bytes)
}

func ExecCMDString(cmdstr string) string {
	cmd := exec.Command("/bin/bash", "-c", cmdstr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "StdoutPipe: " + err.Error()
		// return "StdoutPipe: "
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		// fmt.Println("StderrPipe: ", err.Error())
		return "StderrPipe: " + err.Error()
	}
	if err := cmd.Start(); err != nil {
		return "Start: " + err.Error()
		// return "Start: "
	}
	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		return "ReadAll stderr: " + err.Error()
		// return "ReadAll stderr: "
	}
	if len(bytesErr) != 0 {
		return "stderr is not nil: " + string(bytesErr)
		// return fmt.Sprintf("stderr is not nil: %s", bytesErr)
	}
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "ReadAll stdout: " + err.Error()
		// return "ReadAll stdout: "
	}
	if err := cmd.Wait(); err != nil {
		return "Wait: " + err.Error()
		// return "Wait: "
	}
	// fmt.Printf("stdout: %s", bytes)
	// fmt.Println(cmdlog)
	//处理脚本的执行结果，将多行合并为一行，并将换行替换为<br>给浏览器输出用

	// 去除换行符
	str := string(bytes)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return strings.TrimSpace(str)
}
