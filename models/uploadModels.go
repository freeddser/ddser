package models

import (
	"fmt"
	"github.com/freeddser/ddser/util"
)

func AddOneUploadFile(user_id int, file_name, file_path string) Db_ResultSets {
	// "insert into upload_file(user_id,file_name,file_path,upload_date,uf_bz) value()"
	sql := fmt.Sprintf("insert into upload_file(user_id,file_name,file_path,upload_date,uf_bz) value('%s','%s','%s',now(),'Y')", util.IntTostring(user_id), file_name, file_path)

	return Sql_insert(sql)
}

func DeleteOneUploadFile(user_id, file_id int) Db_ResultSets {
	sql := fmt.Sprintf("update upload_file set uf_bz='N' where user_id=%s and id=%s", util.IntTostring(user_id), util.IntTostring(file_id))
	return Sql_Update(sql)
}

func ShowUploadFile(user_id int) string {
	sql := fmt.Sprintf("select * from upload_file where uf_bz='Y' and user_id=%s order by id desc", util.IntTostring(user_id))
	return SQLDatatoJsonString(sql)
}
