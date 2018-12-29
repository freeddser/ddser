package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/freeddser/ddser/config"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func StripSpaceTabToText(str string) string {

	str = strings.TrimFunc(str, func(r rune) bool {
		if r == ' ' || r == '	' || r == '\n' {
			// fmt.Println(1)
			return true
		}
		// fmt.Println(2)
		return false
	})
	return str
}

func StripHtmlToText(str string) string {

	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	str = re.ReplaceAllStringFunc(str, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	str = re.ReplaceAllString(str, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	str = re.ReplaceAllString(str, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	str = re.ReplaceAllString(str, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	str = re.ReplaceAllString(str, "\n")

	//去除空格

	str = strings.Replace(str, "&nbsp;", "", -1)
	str = strings.Replace(str, "&nbsp;&nbsp;&nbsp;&nbsp;", "\t", -1)
	str = strings.Replace(str, "<br>", "\n", -1)

	str = strings.Replace(str, "&#39;", "'", -1)
	str = strings.Replace(str, "&#92;", "\\", -1)
	str = strings.Replace(str, "<span>$</strong>", "$", -1)
	str = strings.Replace(str, "&amp;", "&", -1)

	str = strings.Replace(str, "&nbsp;&nbsp;&nbsp;&nbsp;", "\t", -1)
	str = strings.Replace(str, "&nbsp;", " ", -1)
	str = strings.Replace(str, "<br>", "\n", -1)

	str = strings.Replace(str, "&#39;", "'", -1)
	str = strings.Replace(str, "&#92;", "\\", -1)
	str = strings.Replace(str, "<span>$</strong>", "$", -1)
	str = strings.Replace(str, "&lt;", "<", -1)
	str = strings.Replace(str, "&gt;", ">", -1)

	str = strings.Replace(str, "&amp;", "&", -1)

	return str

	// return strings.TrimSpace(str)

}
func GetRandMd5Key() string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nowkey := r.Intn(100000000)
	key := strconv.Itoa(nowkey) + config.MustGetString("server.AppName") + strconv.Itoa(r.Intn(10000000))
	h := md5.New()
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum(nil))

}

func GetMd5Str(str string) string {
	h := md5.New()
	h.Write([]byte(str))                  // 需要加密的字符串为 str
	return hex.EncodeToString(h.Sum(nil)) // 输出加密结果
}

/*
def string_to_mysqlstr(s):
    #s=str(s).decode('UTF-8')
    s = s.replace("&", "&amp;")
    s = s.replace("<", "&lt;")
    s = s.replace(">", "&gt;")
    s = s.replace("\t", "&nbsp;&nbsp;&nbsp;&nbsp;")
    s = s.replace("\r\n", "\n")
    s = s.replace("\n", "<br>")

    s = s.replace("'", "&#39;")
    s = s.replace("\\", "&#92;")
    s = s.replace("$", "<span>$</strong>")
    return s

def mysql_to_str(s):

    s = s.replace("&amp;","&")
    s = s.replace("&lt;","<")
    s = s.replace("<<","&lt;&lt;")
    s = s.replace("&gt;",">")
    s = s.replace("&nbsp;&nbsp;&nbsp;&nbsp;","\t" )
    s = s.replace( "\n","\r\n")
    s = s.replace( "<br>","\n")
    s = s.replace("&#39;","'")
    s = s.replace( "&#92;","\\")
    return s

def html_to_null(s):
    s = s.replace("&amp;","")
    s = s.replace("&lt;","")
    s = s.replace("<<","")
    s = s.replace("&gt;","")
    s = s.replace("&nbsp;&nbsp;&nbsp;&nbsp;","" )
    s = s.replace( "\n","")
    s = s.replace( "<br>","")
    s = s.replace("&#39;","")
    s = s.replace( "&#92;","")
    s = s.replace("div","")
    return s
*/

func HtmlToMysql(str string) string {
	//&
	str = strings.Replace(str, "Y7Q", "&", -1)
	//;
	str = strings.Replace(str, "YQF", ";", -1)
	//#
	str = strings.Replace(str, "YQ3", "#", -1)
	return str
}

func MysqlHtmlToString(str string) string {
	//http://www.w3school.com.cn/html/html_entities.asp
	//str = strings.Replace(str, ".", "&sdot;", -1)
	//str = strings.Replace(str, "...", "&hellip;", -1)
	//str = strings.Replace(str, " ", "&nbsp;", -1)

	//str = strings.Replace(str, "<", "&lt;", -1)
	//str = strings.Replace(str, ">", "&gt;", -1)
	//str = strings.Replace(str, "\t", "&nbsp;&nbsp;&nbsp;&nbsp;", -1)
	//str = strings.Replace(str, "\r\n", "\n", -1)
	//str = strings.Replace(str, "\n", "<br>", -1)
	//str = strings.Replace(str, "<p>", "", -1)
	//str = strings.Replace(str, "</p>", "</br>", -1)

	str = strings.Replace(str, "'", "&#39;", -1)
	str = strings.Replace(str, "\\", "&#92;", -1)
	//发现html中的得直接替换成空，要不然显示样式会多一个""，无法正常使用
	str = strings.Replace(str, "\"", "", -1)
	//str = strings.Replace(str, "$", "<span>$</strong>", -1)
	//str = strings.Replace(str, "&", "&amp;", -1)
	// //&
	// str = strings.Replace(str, "&", "Y7Q", -1)
	// //;
	// str = strings.Replace(str, ";", "YQF", -1)
	// //#
	// str = strings.Replace(str, "#", "YQ3", -1)
	return str
}

func StringFileContentToMysql(str string) string {
	// str = strings.Replace(str, ".", "&sdot;", -1)
	// str = strings.Replace(str, "...", "&hellip;", -1)
	// str = strings.Replace(str, " ", "&nbsp;", -1)
	// str = strings.Replace(str, "<", "&lt;", -1)
	// str = strings.Replace(str, ">", "&gt;", -1)
	str = strings.Replace(str, "\t", "&nbsp;&nbsp;&nbsp;&nbsp;", -1)
	str = strings.Replace(str, "\r\n", "\n", -1)
	str = strings.Replace(str, "\n", "<br>", -1)
	str = strings.Replace(str, "<p>", "", -1)
	str = strings.Replace(str, "</p>", "</br>", -1)

	str = strings.Replace(str, "'", "&#39;", -1)
	str = strings.Replace(str, "\\", "&#92;", -1)
	str = strings.Replace(str, "$", "<span>$</strong>", -1)
	str = strings.Replace(str, "&", "&amp;", -1)
	// //&
	// str = strings.Replace(str, "&", "Y7Q", -1)
	// //;
	// str = strings.Replace(str, ";", "YQF", -1)
	// //#
	// str = strings.Replace(str, "#", "YQ3", -1)
	return str
}

func MysqlToString(str string) string {
	str = strings.Replace(str, "&nbsp;&nbsp;&nbsp;&nbsp;", "\t", -1)
	str = strings.Replace(str, "&nbsp;", " ", -1)
	str = strings.Replace(str, "<br>", "\n", -1)

	str = strings.Replace(str, "&#39;", "'", -1)
	str = strings.Replace(str, "&#92;", "\\", -1)
	str = strings.Replace(str, "<span>$</strong>", "$", -1)
	str = strings.Replace(str, "&lt;", "<", -1)
	str = strings.Replace(str, "&gt;", ">", -1)

	str = strings.Replace(str, "&amp;", "&", -1)

	return str
}

func StringToMysql(str string) string {
	str = strings.Replace(str, ".", "&sdot;", -1)
	str = strings.Replace(str, "...", "&hellip;", -1)
	str = strings.Replace(str, " ", "&nbsp;", -1)
	str = strings.Replace(str, "<", "&lt;", -1)
	str = strings.Replace(str, ">", "&gt;", -1)
	str = strings.Replace(str, "\t", "&nbsp;&nbsp;&nbsp;&nbsp;", -1)
	str = strings.Replace(str, "\r\n", "\n", -1)
	str = strings.Replace(str, "\n", "<br>", -1)
	str = strings.Replace(str, "'", "&#39;", -1)
	str = strings.Replace(str, "\\", "&#92;", -1)
	str = strings.Replace(str, "$", "<span>$</strong>", -1)
	str = strings.Replace(str, "&", "&amp;", -1)
	//&
	str = strings.Replace(str, "&", "Y7Q", -1)
	//;
	str = strings.Replace(str, ";", "YQF", -1)
	//#
	str = strings.Replace(str, "#", "YQ3", -1)
	return str
}
