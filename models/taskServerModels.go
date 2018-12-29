package models

import (
	"fmt"
	"github.com/freeddser/ddser/util"
)

//GetOneTaskType("722e9bacd669a1f7b4d2f48a45d3eaea") 得到此token相关的类型
//通过当前登录用户的user_id得到他的相关角色role_id再通过role_id得到task_type类型
func GetOneTaskType(user_id int) string {
	sql := "SELECT * FROM task_type where type='web' and role_id in (select id from role where id in(select role_id from user_role where  user_id=" + util.IntTostring(user_id) + "))"

	return SQLDatatoJsonString(sql)
}

func AddOneTask(pusher, taskname, task_shell, shell_params, task_type_desc string, task_type, token_id int) Db_ResultSets {
	// "insert into task(task_type,taskname,task_shell,shell_params,create_time,task_type_desc,token_id,status) value()"
	sql := fmt.Sprintf("insert into task(pusher,task_type,taskname,task_shell,shell_params,create_time,task_type_desc,token_id,status) value('%s','%s','%s','%s','%s',now(),'%s','%s','%s')", pusher, util.IntTostring(task_type), taskname, task_shell, util.StripSpaceTabToText(shell_params), task_type_desc, util.IntTostring(token_id), util.IntTostring(501))

	return Sql_insert(sql)
}

//显示日志
func ShowTaskLog(task_id int) string {
	sql := fmt.Sprintf("SELECT tl.*,t.task_type_desc,t.taskname,t.shell_params,t.task_shell FROM task_log tl left join task t on tl.task_id=t.id   where tl.task_id=%s and tl.log_bz=0", util.IntTostring(task_id))
	// sql := fmt.Sprintf("SELECT * FROM task_log where task_id=%s and pusher='%s'")

	return SQLDatatoJsonString(sql)
}

//通过user_id获取所有role_id,再通过role_id去task_type获取所有token_id,再用token_id去task表中，查出所有相关的task,状态可用的
func ShowCurrutTokenTask(user_id int) string {
	sql := fmt.Sprintf("select t.*, ts.* from task t join task_status ts left join task_type tt on tt.token_id = t.token_id left join user_role ur on ur.role_id = tt.role_id where ur.user_id = %s and t.task_bz = 'Y' and t.status = ts.status_id group by t.id order by t.create_time desc", util.IntTostring(user_id))

	return SQLDatatoJsonString(sql)
}

// //clinet来取任务的时候，只有一个固定的token-id
// func ShowClientCurrutTokenTask(token_id int) string {
// 	sql := fmt.Sprintf("SELECT * FROM task left join task_status on task.status=task_status.status_id where token_id=%s and task_bz='Y' order by task.id desc ", util.IntTostring(token_id))
//
// 	return SQLDatatoJsonString(sql)
// }

func RepushOneTask(username string, task_id int) Db_ResultSets {
	//disable 旧的task-log记录
	dis_tasklog_sql := fmt.Sprintf("update task_log set log_bz = 1 where task_id=%s", util.IntTostring(task_id))
	Sql_Update(dis_tasklog_sql)

	//502 task_confirmation
	sql := fmt.Sprintf("update task set pusher='%s',status=502,create_time=now() where id=%s", username, util.IntTostring(task_id))

	return Sql_Update(sql)
}

func CancelOneTask(task_id int) Db_ResultSets {
	sql := fmt.Sprintf("update task set task_bz= 'N',create_time=now() where id=%s", util.IntTostring(task_id))

	return Sql_Update(sql)
}

func UpdateOneTaskStatus(task_id, status int) Db_ResultSets {
	sql := fmt.Sprintf("update task set status=%s,create_time=now()  where id=%s", util.IntTostring(status), util.IntTostring(task_id))

	return Sql_Update(sql)
}

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
