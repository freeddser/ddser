package models

import (
	"github.com/freeddser/ddser/util"
)

func GetAllRoleResourceInfos() string {
	sql := "SELECT distinct  role_name,role.create_date,role_id,resource_id,resource_name FROM  role left join  role_resource on role.id = role_resource.role_id left join resource on role_resource.resource_id=resource.id"
	return SQLDatatoJsonString(sql)

}

func GetAllRoleResource() string {
	sql := "select * from resource order by id;"
	return SQLDatatoJsonString(sql)

}

func GetCurrentRoleResource(id int) string {
	sql := "select * from resource where id in(select resource_id from role_resource where role_id=" + util.IntTostring(id) + ")"
	return SQLDatatoJsonString(sql)
}

func DelOneRoleResource(id int) bool {
	sql := "delete from role_resource where resource_id != 10000 and role_id=" + util.IntTostring(id)
	if Sql_Delete(sql).RowsAffected >= 0 {
		return true
	}
	return false

}

func AddOneRoleResource(role_id, resource_id int) Db_ResultSets {
	sql := "insert into role_resource(role_id,resource_id) values (" + util.IntTostring(role_id) + "," + util.IntTostring(resource_id) + ")"
	return Sql_insert(sql)
}
