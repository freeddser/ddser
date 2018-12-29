package models

import (
	"fmt"
	"github.com/freeddser/ddser/util"
)

func GetAllResourceInfos() string {
	sql := "SELECT * FROM resource order by id;"
	return SQLDatatoJsonString(sql)

}

// 增加一个resource
func AddOneResource(resource_name, resource_desc, resource_url string, resource_parent int) Db_ResultSets {
	sql := "INSERT INTO resource (resource_name, resource_desc, resource_parent, resource_url, create_date) VALUES ('" + resource_name + "','" + resource_desc + "'," + util.IntTostring(resource_parent) + ",'" + resource_url + "', now())"
	result := Sql_insert(sql)
	return result
}

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

func DelResourceDefaultResource(id int) {
	sql := "DELETE FROM role_resource WHERE resource_id=" + util.IntTostring(id)
	Sql_Delete(sql)
}

func AddResourceDefaultResource(id int) {
	sql := "INSERT INTO role_resource (role_id, resource_id) VALUES (" + util.IntTostring(id) + ", 10000)"
	Sql_insert(sql)
}
