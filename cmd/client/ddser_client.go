package main

import (
	"flag"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/freeddser/ddser/config"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var ExecuteTime int
var DdserURL string
var Token string
var SlackWebHookUrl string
var SlackChannel string
var SlackBotName string
var SlackSwitch string

func main() {
	// httpGet()
	//每30s执行一次

	configFile := flag.String("c", "", "Configuration File")
	flag.Parse()

	if *configFile == "" {
		fmt.Println("\n\nUse -h to get more information on command line options\n")
		fmt.Println("You must specify a configuration file")
		os.Exit(1)
	}

	err := config.Initialize(*configFile)
	if err != nil {
		fmt.Printf("Error reading configuration: %s\n", err.Error())
		os.Exit(1)
	}
	//init variable
	DdserURL = config.MustGetString("server.ddser_server_url")
	Token = config.MustGetString("server.ddser_server_token")
	ExecuteTime = config.MustGetInt("server.execute_time")

	SlackSwitch = config.MustGetString("switch.slack_message")
	SlackWebHookUrl = config.MustGetString("slack.web_hook_url")
	SlackChannel = config.MustGetString("slack.slack_channel")
	SlackBotName = config.MustGetString("slack.bot_name")

DOC:
	start_time := time.Now()
	fmt.Println(time.Now())
	time.Sleep(time.Duration(int64(ExecuteTime)) * time.Second)

	ParseTask(GetClientTaskList())
	//输出执行时间
	end_time := time.Now()
	fmt.Println(end_time.Sub(start_time))
	goto DOC

}

func UpdateClientTaskLogs(task_id, task_logs, client_infos, task_status, pusher string) string {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)

		}
	}()

	serverUrl := fmt.Sprintf("%s/taskclient/show?type=update_task_log&token=%s&task_id=%s&client_infos=%s&task_status=%s&pusher=%s&task_logs=%s", DdserURL, Token, task_id, client_infos, task_status, pusher, task_logs)
	fmt.Println(serverUrl)
	resp, err := http.Get(serverUrl)
	if err != nil {
		CheckErr(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		CheckErr(err)
	}
	fmt.Println(string(body))
	return string(body)
}

func ParseTask(jsondata string) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)

		}
	}()
	js, err := simplejson.NewJson([]byte(jsondata))
	if err != nil {
		CheckErr(err)
	}
	arrnodes, err := js.Get("data").Array()
	if err != nil {
		fmt.Println(jsondata)
		CheckErr(err)
	}

	fmt.Println("-------------------------")
	//once time one task
	if len(arrnodes) == 0 {
		fmt.Println("no new task!")
	} else {

		task_type := GetTypeString(GetTypeMap(arrnodes[0])["task_type_desc"])

		//1.此任务类型是否是允许我执行的类型，如果是我再做第2步，如果不是就直接跳过
		if config.MustGetString("task."+task_type) == "allow" {
			fmt.Println("I Recivied a New task!")
			task_id := GetTypeString(GetTypeMap(arrnodes[0])["id"])
			// task_id := GetTypeString(GetTypeMap(arrnodes[0])["id"])
			// task_shell:=GetTypeMap(arrnodes[0])["task_shell"])
			task_shell := GetTypeString(GetTypeMap(arrnodes[0])["task_shell"])
			shell_params := GetTypeString(GetTypeMap(arrnodes[0])["shell_params"])
			pusher := GetTypeString(GetTypeMap(arrnodes[0])["pusher"])
			taskname := GetTypeString(GetTypeMap(arrnodes[0])["taskname"])
			create_time := GetTypeString(GetTypeMap(arrnodes[0])["create_time"])

			client_infos := config.MustGetString("server.hostname") + "<->" + config.MustGetString("server.ip")

			taskStr := fmt.Sprintf("UserName:%s At %s Push a Task\n Task_id: %s \n Task_Name:%s \n  Task_Type:%s \n Task_Content:%s\t %s  ", pusher, create_time, task_id, taskname, task_type, task_shell, shell_params)
			//taskStrHtml := fmt.Sprintf("UserName:%s At %s Push a Task</br> Task_id: %s </br> Task_Name:%s </br> Task_Type:%s </br> Task_Content:%s\t %s  </br>", pusher, create_time, task_id, taskname, task_type, task_shell, shell_params)

			// fmt.Println(taskStr)
			//发送slack消息
			PostSlackMessage(taskStr)
			//客户端已经接收到，通知server
			// fmt.Println(StringToInt(UpdateClientTaskStatus(IntTostring(19), IntTostring(1002))) + 2)
			if StringToInt(UpdateClientTaskStatus(task_id, IntTostring(1001))) >= 0 {
				fmt.Println("Send Client_received to Server.")
				//检查脚本是否存在

				fmt.Println("Check " + task_shell + " On Local.")
				if IsExists(task_shell) {
					// #1002	Client_executing
					UpdateClientTaskStatus(task_id, IntTostring(1002))

					UpdateClientTaskLogs(task_id, "Client-start-Exec", client_infos, "Client_executing", pusher)
					CmdStr := task_shell + " " + shell_params
					fmt.Println("Now go to exec: " + CmdStr)
					task_logs := ExecCMD(CmdStr)
					fmt.Println("-------Exec--Logs----------")
					fmt.Println(task_logs)
					if strings.Contains(task_logs, "OUT") {
						// #1003	Client_execting _done
						UpdateClientTaskStatus(task_id, IntTostring(1003))
						UpdateClientTaskLogs(task_id, task_logs, client_infos, "Client_execting_done", pusher)
						fmt.Println("Exec Done")

					} else {
						UpdateClientTaskStatus(task_id, IntTostring(1004))
						UpdateClientTaskLogs(task_id, task_logs, client_infos, "Client_exec_Error", pusher)
						fmt.Println("Exec Done")

					}

				} else {
					//脚本不存在
					UpdateClientTaskStatus(task_id, IntTostring(102))
					//发送邮件或者slack报警，接收到任务了，但是脚本不存在
					errinfos := "client Infos:" + client_infos + "\n" + taskStr + "\nTask Faild!!!" + task_shell + "\tStatus:Client_script_does_not_exist"
					//errinfos_html := "client Infos:HostName:" + iniconf.String("hostname") + " ip:" + iniconf.String("ip") + "</br>" + taskStrHtml + "</br>Task Faild!!!" + task_shell + "\tStatus:Client_script_does_not_exist</br>"
					UpdateClientTaskLogs(task_id, "Client_script_does_not_exist", client_infos, "Client_script_does_not_exist", pusher)
					//slack
					PostSlackMessage(errinfos)
					//email
					//SendToMail("Client_script_does_not_exist", errinfos_html)

				}

			}

		} else {
			fmt.Println("not my task!")
		}

		// cmdstr := GetTypeString(GetTypeMap(arrnodes[0])["task_shell"]) + " " + GetTypeString(GetTypeMap(arrnodes[0])["shell_params"])
		// fmt.Println(cmdstr)
		// ExecCMD(cmdstr)
	}

}

//comm func
//type=update_task_status&task_id=1&status=1001  已经接收到
func UpdateClientTaskStatus(task_id, status string) string {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)

		}
	}()

	serverUrl := fmt.Sprintf(`%s/taskclient/show?type=update_task_status&token=%s&task_id=%s&status=%s`, DdserURL, Token, task_id, status)
	resp, err := http.Get(serverUrl)
	if err != nil {
		CheckErr(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		CheckErr(err)
	}
	return string(body)
}

func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func GetClientTaskList() string {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)

		}
	}()

	serverUrl := fmt.Sprintf("%s/taskclient/show?type=auth_token&token=%s", DdserURL, Token)
	// fmt.Println(serverUrl)
	resp, err := http.Get(serverUrl)
	if err != nil {
		CheckErr(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		CheckErr(err)
	}
	return string(body)
}

func PostSlackMessage(content string) {
	if SlackSwitch == "on" {
		jsonstr := fmt.Sprintf(`{"channel": "%s", "username": "`+SlackBotName+`", "text": "%s", "icon_emoji": ":imp:"}`, SlackChannel, content)
		resp, err := http.Post(SlackWebHookUrl,
			"application/x-www-form-urlencoded",
			strings.NewReader("payload="+jsonstr))
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
			fmt.Println(err)
		}
		fmt.Println("Send slack Message:", string(body))

	} else {
		fmt.Println("Send slack not enable!")
	}

}

func IntTostring(i int) string {
	str1 := strconv.Itoa(i)
	return str1

}

func GetTypeString(value interface{}) string {

	switch inst := value.(type) {
	case string:
		return inst
	default:
		return "unknow"

	}

}

func GetTypeMap(value interface{}) map[string]interface{} {

	switch inst := value.(type) {
	case map[string]interface{}:
		return inst
	default:
		return nil

	}

}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
func ExecCMD(cmdstr string) string {
	cmd := exec.Command("/bin/bash", "-c", cmdstr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return StringToMysql("StdoutPipe: " + err.Error())
		// return "StdoutPipe: "
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		// fmt.Println("StderrPipe: ", err.Error())
		return StringToMysql("StderrPipe: " + err.Error())
	}
	if err := cmd.Start(); err != nil {
		return StringToMysql("Start: " + err.Error())
		// return "Start: "
	}
	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		return StringToMysql("ReadAll stderr: " + err.Error())
		// return "ReadAll stderr: "
	}
	if len(bytesErr) != 0 {
		return StringToMysql("stderr is not nil: " + string(bytesErr))
		// return fmt.Sprintf("stderr is not nil: %s", bytesErr)
	}
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return StringToMysql("ReadAll stdout: " + err.Error())
		// return "ReadAll stdout: "
	}
	if err := cmd.Wait(); err != nil {
		return StringToMysql("Wait: " + err.Error())
		// return "Wait: "
	}
	// fmt.Printf("stdout: %s", bytes)
	// fmt.Println(cmdlog)
	//处理脚本的执行结果，将多行合并为一行，并将换行替换为<br>给浏览器输出用
	str := string(bytes)
	// fmt.Println("----->---")
	// fmt.Println(str)

	// 去除换行符
	return StringToMysql("OUT:" + str)
}

func StringToMysql(str string) string {
	// str = strings.Replace(str, ".", "&sdot;", -1)
	str = strings.Replace(str, "...", "&hellip;", -1)
	str = strings.Replace(str, " ", "&nbsp;", -1)
	str = strings.Replace(str, "<", "&lt;", -1)
	str = strings.Replace(str, ">", "&gt;", -1)
	str = strings.Replace(str, "\t", "&nbsp;&nbsp;&nbsp;&nbsp;", -1)
	str = strings.Replace(str, "\r\n", "\n", -1)
	str = strings.Replace(str, "\n", "<br>", -1)
	str = strings.Replace(str, "'", "&#39;", -1)
	str = strings.Replace(str, "\\", "&#92;", -1)
	str = strings.Replace(str, "$", "<span>$</strong>", -1)
	// str = strings.Replace(str, "&", "&amp;", -1)
	//&
	str = strings.Replace(str, "&", "Y7Q", -1)
	//;
	str = strings.Replace(str, ";", "YQF", -1)
	//#
	str = strings.Replace(str, "#", "YQ3", -1)
	return str
}
