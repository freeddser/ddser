package models

import (
	"fmt"
	"github.com/freeddser/ddser/util"
)

func GetAllRoleInfos() string {
	sql := "SELECT * FROM role order by id;"
	return SQLDatatoJsonString(sql)

}

func DelOneRole(id int) bool {
	sql := "delete from role where id=" + util.IntTostring(id)
	if Sql_Delete(sql).RowsAffected >= 0 {
		//delete default resource from role_resource
		DelResourceDefaultResource(id)
		//删除默认的token
		DelRoleDefaultToken(id)
		return true
	}
	return false

}
func DelRoleDefaultToken(id int) {
	sql := "DELETE FROM token WHERE role_id=" + util.IntTostring(id)
	Sql_Delete(sql)
}

func DelRoleDefaultResource(id int) {
	sql := "DELETE FROM role_resource WHERE role_id=" + util.IntTostring(id)
	Sql_Delete(sql)
}

func AddOneRole(Rolename string) Db_ResultSets {
	sql := "insert into role(role_name,create_date) values ('" + Rolename + "', now())"
	result := Sql_insert(sql)
	AddRoleDefaultResource(int(result.LastInsertId))
	AddRoleDefaultToken(int(result.LastInsertId), Rolename)
	return result
}

func AddRoleDefaultToken(roleid int, rolename string) {
	sql := fmt.Sprintf("INSERT INTO token (token, role_name,role_id) VALUES (md5('%s'),'%s',%s)", rolename+"gavincool2016", rolename, util.IntTostring(roleid))
	Sql_insert(sql)
}

func AddRoleDefaultResource(id int) {
	sql := "INSERT INTO role_resource (role_id, resource_id) VALUES (" + util.IntTostring(id) + ", 10000)"
	Sql_insert(sql)
}

func GetOneRoleInfo(id int) string {
	sql := "select * from role where id=" + util.IntTostring(id)
	return SQLDatatoJsonString(sql)
}

func UpdateOneRole(Rolename string, id int) Db_ResultSets {
	sql := "update role set role_name = '" + Rolename + "',create_date=now() where id = " + util.IntTostring(id)
	UpdateOneRoleToken(Rolename, id)
	return Sql_Update(sql)

}

func UpdateOneRoleToken(Rolename string, id int) Db_ResultSets {
	sql := "update token set role_name = '" + Rolename + "' where role_id = " + util.IntTostring(id)
	//把token中的角色名更新一下,这里我并没有更新token的md5值，因为我考虑到可能client已经配置完token,服务端一修改，client会无法正常用。如果你喜欢这里可以跟着修改
	return Sql_Update(sql)
}

/*
\


def MUpdateGetRole(id):
    sql="select * from Role where id=%s" % (id)
    return config.dbw.query(sql).list()

def MUpdateRole(Rolename,password,id):
    sql="update Role set Rolename = '%s', password= '%s',create_date=now() where id = %s" % (Rolename,password,id)
    return config.dbw.query(sql)
*/
