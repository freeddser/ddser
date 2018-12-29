package models

import (
	"fmt"
	"github.com/freeddser/ddser/util"
)

//GetOneTaskType("722e9bacd669a1f7b4d2f48a45d3eaea") 得到此token相关的类型
//通过当前登录用户的user_id得到他的相关角色role_id再通过role_id得到task_type类型
// func GetOneTaskType(user_id int) string {
// 	sql := "SELECT * FROM task_type where role_id in (select id from role where id in(select role_id from user_role where user_id=" + util.IntTostring(user_id) + "))"
//
// 	return SQLDatatoJsonString(sql)
// }

// func AddOneTask(taskname, task_shell, shell_params, task_type_desc string, task_type, token_id int) Db_ResultSets {
// 	// "insert into task(task_type,taskname,task_shell,shell_params,create_time,task_type_desc,token_id,status) value()"
// 	sql := fmt.Sprintf("insert into task(task_type,taskname,task_shell,shell_params,create_time,task_type_desc,token_id,status) value('%s','%s','%s','%s',now(),'%s','%s','%s')", util.IntTostring(task_type), taskname, task_shell, shell_params, task_type_desc, util.IntTostring(token_id), util.IntTostring(501))
//
// 	return Sql_insert(sql)
// }

// //通过user_id获取所有role_id,再通过role_id去task_type获取所有token_id,再用token_id去task表中，查出所有相关的task,状态可用的
// func ShowCurrutTokenTask(user_id int) string {
// 	sql := fmt.Sprintf("SELECT * FROM task left join task_status on task.status=task_status.status_id where token_id in (select token_id from task_type where role_id in(SELECT role_id FROM task_type where role_id in (select id from role where id in(select role_id from user_role where user_id=%s)))) and task_bz='Y' order by task.id desc ", util.IntTostring(user_id))
//
// 	return SQLDatatoJsonString(sql)
// }

//clinet来取任务的时候，只有一个固定的token-id
func ShowClientCurrutTokenTask(token_id int) string {
	sql := fmt.Sprintf(" SELECT  DISTINCT t.*,ts.*,tt.id ttid,tt.task_type_desc FROM task t left join task_status ts on t.status=ts.status_id  left join task_type tt on t.task_type=tt.id where t.token_id=%s and t.status=502  and t.task_bz='Y' order by t.id desc", util.IntTostring(token_id))
	// sql := fmt.Sprintf("SELECT * FROM task left join task_status on task.status=task_status.status_id where token_id=%s and status=502  and task_bz='Y' order by task.id desc ", util.IntTostring(token_id))

	return SQLDatatoJsonString(sql)
}

//根据token得到token_id

func GetTokenId(token string) int {
	sql := fmt.Sprintf("SELECT id FROM token where token='%s'", token)
	rows := Sql_Select_Rows(sql)
	var id int
	for rows.Next() {
		rows.Scan(&id)
	}
	return id
}

//在client取到作务第一次来更新task状态时1001 当client告诉server说收到task时，插入一条task_log
func UpdateOneTaskStatusForClient(task_id, status, token_id int) Db_ResultSets {

	// fmt.Println(task_id)
	sql := fmt.Sprintf("update task set status=%s  where id=%s and token_id=%s", util.IntTostring(status), util.IntTostring(task_id), util.IntTostring(token_id))

	return Sql_Update(sql)
}

//插入一条task_log
func AddOneTaskLogForClient(task_id, token_id int) {
	//disable 旧的task-log记录
	dis_tasklog_sql := fmt.Sprintf(`update task_log set log_bz = 1 where task_id=%s`, util.IntTostring(task_id))
	Sql_Update(dis_tasklog_sql)
	tasklog_sql := fmt.Sprintf(`insert into task_log(task_id,token_id,create_time,update_time,task_status) values(%s,%s,now(),now(),'%s')`, util.IntTostring(task_id), util.IntTostring(token_id), "Client_received")
	Sql_insert(tasklog_sql)
}

//更新tasklog
func UpdateTaskLogForClient(task_id, token_id int, task_logs, client_infos, task_status, pusher string) Db_ResultSets {
	sql := fmt.Sprintf("update  task_log set update_time=now(),task_logs='%s',client_infos='%s',task_status='%s',pusher='%s' where task_id=%s and token_id = %s and log_bz = 0", task_logs, client_infos, task_status, pusher, util.IntTostring(task_id), util.IntTostring(token_id))

	return Sql_Update(sql)
}

// func CancelOneTask(task_id int) Db_ResultSets {
// 	sql := fmt.Sprintf("update task set task_bz= 'N' where id=%s", util.IntTostring(task_id))
//
// 	return Sql_Update(sql)
// }

// func UpdateOneTaskStatus(task_id, status int) Db_ResultSets {
// 	sql := fmt.Sprintf("update task set status=%s  where id=%s", util.IntTostring(status), util.IntTostring(task_id))
//
// 	return Sql_Update(sql)
// }

// func GetRoleID(userid int) string {
// 	sql := "select id from role where id in(select role_id from user_role where user_id=" + util.IntTostring(userid) + ") order by id"
// 	// sql := "select * from role where id in(select role_id from user_role where user_id=" + util.IntTostring(id) + ") order by id"
//
// 	return SQLDatatoJsonString(sql)

// }

// SELECT * FROM task_type where role_id in (select id from role where id in(select role_id from user_role where user_id=1) order by id)

/*
// 增加一个resource
func AddOneResource(resource_name, resource_desc, resource_url string, resource_parent int) Db_ResultSets {
	sql := "INSERT INTO resource (resource_name, resource_desc, resource_parent, resource_url, create_date) VALUES ('" + resource_name + "','" + resource_desc + "'," + util.IntTostring(resource_parent) + ",'" + resource_url + "', now())"

	result := Sql_insert(sql)
	return result
}

//修改一个resource

func UpdateOneResource(resource_name, resource_desc, resource_url string, id int) Db_ResultSets {
	sql := "update resource set resource_name = '" + resource_name + "'," + "resource_desc = '" + resource_desc + "'," + "resource_url = '" + resource_url + "'," + "create_date=now() where id = " + util.IntTostring(id)

	return Sql_Update(sql)

}

//END

func DelOneResource(id int) string {
	//首先查一下是否有子菜间
	checksql := "SELECT * FROM resource where resource_parent=" + util.IntTostring(id)
	fmt.Println(checksql)
	if Sql_Select_Bool(checksql) {
		//有子菜单，不能删除,返回0
		return "Delete faild! Please delete child menu first!"
	} else {
		sql := "delete from resource where id=" + util.IntTostring(id)

		if Sql_Delete(sql).RowsAffected >= 0 {
			return "Delete ok!"
		}
		return "delete faild!"

	}

}

//删除一条默认的default资源
func DelResourceDefaultResource(id int) {
	sql := "DELETE FROM role_resource WHERE resource_id=" + util.IntTostring(id)

	Sql_Delete(sql)
}

//增加一条默认的default资源
func AddResourceDefaultResource(id int) {
	sql := "INSERT INTO role_resource (role_id, resource_id) VALUES (" + util.IntTostring(id) + ", 10000)"

	Sql_insert(sql)
}

// func GetOneResourceInfo(id int) string {
// 	sql := "select * from Resource where id=" + util.IntTostring(id)
//
// 	return SQLDatatoJsonString(sql)
// }
*/
