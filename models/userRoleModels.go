package models

import (
	"github.com/freeddser/ddser/util"
)

func GetAllUserRoleInfos() string {
	sql := "SELECT distinct  username,user.create_date,user_id,role_id,role_name FROM  user left join  user_role on user.id = user_role.user_id left join role on user_role.role_id=role.id"

	return SQLDatatoJsonString(sql)

}

func GetAllRole() string {
	sql := "select * from role order by id;"

	return SQLDatatoJsonString(sql)

}

func GetCurrentUserRole(id int) string {
	sql := "select * from role where id in(select role_id from user_role where user_id=" + util.IntTostring(id) + ") order by id"

	return SQLDatatoJsonString(sql)
}

func DelOneUserRole(id int) bool {
	sql := "delete from user_role where user_id=" + util.IntTostring(id)

	if Sql_Delete(sql).RowsAffected >= 0 {
		return true
	}
	return false

}

func AddOneUserRole(user_id, role_id int) Db_ResultSets {
	sql := "insert into user_role(user_id,role_id) values (" + util.IntTostring(user_id) + "," + util.IntTostring(role_id) + ")"

	return Sql_insert(sql)
}
