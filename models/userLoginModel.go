package models

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/freeddser/ddser/util"
)

func CheckUser(username, password string) bool {

	SQL_statement := "SELECT * FROM user where username='" + username + "' and password='" + util.GetMd5Str(password) + "'"
	return Sql_Select_Bool(SQL_statement)
}

type RoleInfos struct {
	id              int
	resource_name   string
	resource_desc   string
	resource_parent int
	resource_url    string
	create_data     string
}

func GetUserRoleInfos(username string) string {
	sql := "select distinct r.* from resource r  left join role_resource rr on r.id = rr.resource_id left join role ro on ro.id = rr.role_id left join user_role ur on ro.id = ur.role_id left join user u on ur.user_id = u.id where u.username ='" + username + "'"
	return SQLDatatoJsonString(sql)

}

func GetUserId(username string) int {
	sql := fmt.Sprintf("SELECT id FROM user where username='%s'", username)
	rows := Sql_Select_Rows(sql)
	var id int
	for rows.Next() {
		rows.Scan(&id)
	}
	return id
}

func GetUrlAccess(userprivileges string) map[string]interface{} {

	js, err := simplejson.NewJson([]byte(userprivileges))
	if err != nil {
		panic(err.Error())
	}
	arrnodes, _ := js.Get("data").Array()
	dat := make(map[string]interface{})
	for _, v := range arrnodes {
		for l, lv := range util.GetTypeMap(v) {
			if l == "resource_url" {
				// fmt.Println(lv)
				dat[util.GetTypeString(lv)] = "gavincool"
			}
		}
	}
	return dat
}
