package models

import (
	"fmt"
	"github.com/freeddser/ddser/util"
)

func GetAllUserInfos() string {
	sql := "SELECT * FROM user;"
	return SQLDatatoJsonString(sql)

}

func DelOneUser(id int) bool {
	sql := "delete from user where id=" + util.IntTostring(id)
	if Sql_Delete(sql).RowsAffected >= 0 {
		DelUserDefaultRole(id)
		return true
	}
	return false

}

func AddOneUser(username, password string) Db_ResultSets {
	sql := fmt.Sprintf("insert into user(username,password,api_key,create_date) values ('%s','%s','%s',now())", username, util.GetMd5Str(password), util.GetRandMd5Key())
	//fmt.Println(sql)
	// return Sql_insert(sql)
	result := Sql_insert(sql)
	AddUserDefaultRole(int(result.LastInsertId))
	// Sql_Select_Rows("select id from role where role_name=you")
	return result
}

func GetOneUserInfo(id int) string {
	sql := "select * from user where id=" + util.IntTostring(id)
	return SQLDatatoJsonString(sql)
}

func UpdateOneUser(username, password string, id int) Db_ResultSets {
	sql := "update user set username = '" + username + "', password= '" + util.GetMd5Str(password) + "',create_date=now() where id = " + util.IntTostring(id)
	return Sql_Update(sql)

}

//删除一条默认的default资源
func DelUserDefaultRole(id int) {
	sql := "DELETE FROM user_role WHERE user_id=" + util.IntTostring(id)
	Sql_Delete(sql)
}

//增加一条默认的default资源
func AddUserDefaultRole(id int) {
	sql := "INSERT INTO user_role (user_id, role_id) VALUES (" + util.IntTostring(id) + ", 100)"
	Sql_insert(sql)
}
