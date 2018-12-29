package models

import (
	"fmt"
	"github.com/freeddser/ddser/util"
)

func GetAllTokenInfos() string {
	sql := "SELECT * FROM token;"
	return SQLDatatoJsonString(sql)

}

func GetOneTokenTaskTypelist(id string) string {
	sql := fmt.Sprintf("SELECT * FROM task_type where token_id=%s;", id)
	return SQLDatatoJsonString(sql)

}

func AddOneTokenTask(tk_type, task_shell, shell_params, task_type_desc string, role_id, token_id int) Db_ResultSets {
	// "insert into task(task_type,taskname,task_shell,shell_params,create_time,task_type_desc,token_id,status) value()"
	sql := fmt.Sprintf("insert into task_type(type,task_shell,task_shell_params,create_time,task_type_desc,token_id,role_id) value('%s','%s','%s',now(),'%s','%s','%s')", tk_type, task_shell, shell_params, task_type_desc, util.IntTostring(token_id), util.IntTostring(role_id))
	util.Loger(sql)
	return Sql_insert(sql)
}

func DelOneTask(id int) Db_ResultSets {
	sql := "update task_type set role_id=0,token_id=0  where id=" + util.IntTostring(id)
	return Sql_Update(sql)

}
