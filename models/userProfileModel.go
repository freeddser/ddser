package models

import (
	"fmt"
	"github.com/freeddser/ddser/util"
)

func MChangePassword(user_id int, newpass string) Db_ResultSets {
	sql := fmt.Sprintf("update user set password= '%s',create_date=now() where id=%s ", util.GetMd5Str(newpass), util.IntTostring(user_id))
	return Sql_Update(sql)
}
