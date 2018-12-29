package models

import (
	"database/sql"
	"fmt"
	"github.com/freeddser/ddser/config"
	"github.com/freeddser/ddser/util"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

type Db_ResultSets struct {
	RowsAffected int64
	LastInsertId int64
}

var dbURL string

func InitDB() {
	dbURL = config.MustGetString("mysql.DBUSER") + ":" + config.MustGetString("mysql.DBPASSWD") + "@tcp(" + config.MustGetString("mysql.DBHOST") + ")/" + config.MustGetString("mysql.DBNAME") + "?charset=utf8"
}

func Sql_Delete(sqlju string) Db_ResultSets {
	db, err := sql.Open("mysql", dbURL)
	util.CheckErr(err)
	stmt, err := db.Prepare(sqlju)
	util.CheckErr(err)
	res, err := stmt.Exec()
	util.CheckErr(err)
	db.Close()

	var deleteresult Db_ResultSets
	deleteresult.LastInsertId, _ = res.LastInsertId()
	deleteresult.RowsAffected, _ = res.RowsAffected()

	util.Loger(sqlju)
	return deleteresult
}

func Sql_Update(sqlju string) Db_ResultSets {
	db, err := sql.Open("mysql", dbURL)
	util.CheckErr(err)
	stmt, err := db.Prepare(sqlju)
	util.CheckErr(err)
	res, err := stmt.Exec()
	util.CheckErr(err)
	db.Close()

	var updateresult Db_ResultSets
	updateresult.LastInsertId, _ = res.LastInsertId()
	updateresult.RowsAffected, _ = res.RowsAffected()
	util.Loger(sqlju)
	return updateresult
}

func Sql_Select_Bool(sqlju string) bool {
	db, err := sql.Open("mysql", dbURL)
	util.CheckErr(err)
	rows, err := db.Query(sqlju)
	util.CheckErr(err)
	db.Close()
	util.Loger(sqlju)
	return rows.Next()
}

func Sql_Select_Rows(sqlju string) *sql.Rows {
	db, err := sql.Open("mysql", dbURL)
	util.CheckErr(err)
	rows, err := db.Query(sqlju)
	util.CheckErr(err)
	db.Close()
	util.Loger(sqlju)
	return rows
}

func Sql_insert(sqlju string) Db_ResultSets {
	db, err := sql.Open("mysql", dbURL)
	util.CheckErr(err)

	stmt, err := db.Prepare(sqlju)
	util.CheckErr(err)
	res, err := stmt.Exec()
	util.CheckErr(err)

	var insertresult Db_ResultSets
	insertresult.LastInsertId, _ = res.LastInsertId()
	insertresult.RowsAffected, _ = res.RowsAffected()
	db.Close()
	util.Loger(sqlju)
	return insertresult
}

func SQLDatatoJsonString(sqlju string) string {
	util.Loger(sqlju)
	conn, err := sql.Open("mysql", dbURL)
	defer conn.Close()
	stmt, err := conn.Prepare(sqlju)
	if err != nil {
		fmt.Println("Query Error", err)
		return "Error"
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("Query Error", err)
		return "Error"
	}
	defer rows.Close()
	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))
	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	var jsonstring string
	jsonstring = "{\"timestamp\": \"" + time.Now().Format("2006-01-02 15:04:05") + "\",\"data\":["
	//allcount := 0
	jsonStringArr := make([]string, 0)
	for rows.Next() {
		//jsonstring += "{"
		jsonItem := "{"
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			//          fmt.Println(columns[i], ": ", value)
			if i == len(values)-1 {
				//jsonstring += "\"" + columns[i] + "\":\"" + value + "\""
				//在json中某个key的value是不能用换行和制表符的所以取出数据时，去掉
				value = strings.Replace(value, "\n", "", -1)
				value = strings.Replace(value, "\t", "    ", -1)
				jsonItem += "\"" + columns[i] + "\":\"" + value + "\""
			} else {
				//jsonstring += "\"" + columns[i] + "\":\"" + value + "\","
				value = strings.Replace(value, "\n", "", -1)
				value = strings.Replace(value, "\t", "    ", -1)
				jsonItem += "\"" + columns[i] + "\":\"" + value + "\","
			}
			//          fmt.Println(" :", i, ": ", col, len(values))
		}
		//fmt.Println("replace before :", jsonstring, ": ", len(jsonstring))
		//jsonstring = strings.Replace(jsonstring, ",", " ", len(jsonstring))
		//fmt.Println("replace after :", jsonstring, ": ", len(jsonstring))
		//      fmt.Println("-----------------------------------", allcount)

		//jsonstring += "},"
		jsonItem += "}"
		jsonStringArr = append(jsonStringArr, jsonItem)
		//allcount++
	}
	jsonStringArrLen := len(jsonStringArr)
	//fmt.Println(jsonStringArrLen)
	for i := range jsonStringArr {
		jsonstring += jsonStringArr[i]
		if i < (jsonStringArrLen - 1) {
			jsonstring += ","
		}
	}

	// if allcount > 0 {
	// 	jsonstring = Substr(jsonstring, 0, len(jsonstring)-1)
	// }

	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	/*fmt.Println(strings.LastIndex(jsonstring, ","))
	fmt.Println(len(jsonstring))
	jsonstring = Substr(jsonstring, 0, strings.LastIndex(jsonstring, ","))
	fmt.Println(len(jsonstring))*/
	jsonstring += "]}"
	//jsonstr:=SubString(jsonstring,0,UnicodeIndex(jsonstring,",]}")) +"]}"
	return jsonstring
}
