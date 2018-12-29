package controllers

import (
	"fmt"
	"github.com/freeddser/ddser/config"
	UploadM "github.com/freeddser/ddser/models"
	"github.com/freeddser/ddser/util"
	"io"
	"net/http"
	"os"
)

func UploadSQLFileHandler(w http.ResponseWriter, r *http.Request) {
	sess := GlobalSessions.SessionStart(w, r)
	if r.Method == "GET" {
		r.ParseForm()
		switch r.FormValue("type") {
		case "showsqlfile":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, UploadM.ShowUploadFile(util.GetTypeInt(sess.Get("user_id"))))

		case "deletesqlfile":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, UploadM.DeleteOneUploadFile(util.GetTypeInt(sess.Get("user_id")), util.StringToInt(r.FormValue("file_id"))).RowsAffected)

		default:
			util.RenderHtml(w, "upload_sql", nil)
		}
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("uploadFile")

		util.CheckErr(err)
		filename := h.Filename
		defer f.Close()
		// 1/sql/
		uploadDIR := config.MustGetString("server.UPLOAD_DIR")
		baseDIR := config.MustGetString("server.Base_Path")
		uploadFileDIR := config.MustGetString("server.Upload_file_Path")
		shellPathDIR := config.MustGetString("server.Shell_Path")
		toolLibDIR := config.MustGetString("server.ToolLib")

		userpath := "user/" + util.IntTostring(util.GetTypeInt(sess.Get("user_id"))) + "/sql/"
		// ./static/uploads/user/1/sql/
		// ./static/uploads/user/
		createPath(uploadDIR + "/" + "user/")
		// ./static/uploads/user/1/
		createPath(uploadDIR + "/" + "user/" + util.IntTostring(util.GetTypeInt(sess.Get("user_id"))))
		//检./static/uploads/user/1/sql/
		createPath(uploadDIR + "/" + userpath)
		// /static/uploads/1/sql/xxx.sql
		t, err := os.Create(uploadDIR + "/" + userpath + filename)
		util.CheckErr(err)
		defer t.Close()
		_, err = io.Copy(t, f)
		util.CheckErr(err)

		mysqlSyntaxecheck := baseDIR + shellPathDIR + "mysqlSyntaxCheck.sh"
		toolspath := baseDIR + toolLibDIR + "tmysqlparse/"
		sqlFileName := baseDIR + uploadFileDIR + userpath + filename
		logfile := baseDIR + uploadFileDIR + userpath + filename

		checksqlcmd := mysqlSyntaxecheck + " " + toolspath + " " + sqlFileName + " " + logfile
		cmdResult := util.StringToInt(util.ExecCMDString(checksqlcmd))

		w.Header().Set("Content-Type", "application/json")
		if cmdResult == 0 {
			UploadM.AddOneUploadFile(util.GetTypeInt(sess.Get("user_id")), filename, uploadFileDIR+userpath+filename)
			resjson := fmt.Sprintf(`{"status":0,"filename":"%s"}`, uploadFileDIR+userpath+filename)
			fmt.Fprintln(w, resjson)
		} else if cmdResult == 1 {
			resjson := fmt.Sprintf(`{"status":1,"xmllog":"%s"}`, uploadFileDIR+userpath+filename+".xml")
			os.Remove(uploadDIR + "/" + userpath + filename)
			fmt.Fprintln(w, resjson)
		} else {
			fmt.Fprintln(w, `{"status":"unknow"`)
		}

	}
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func createPath(path string) {
	if !isExists(path) {
		err := os.Mkdir(path, 0777)
		CheckErr(err)
	}

}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
