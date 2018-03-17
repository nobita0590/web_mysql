package config

import (
	"flag"
	"fmt"
)

var (
	/*FilePath  = "/usr/go/src/github.com/nobita0590/web_mysql"
	Host = "localhost"
	Port = "8080"
	MysqlConnect = "root:@/test?charset=utf8&parseTime=True"*/
	Host = "2study.edu.vn"
	Port = "80"
	FilePath = "/opt/web"
	MysqlConnect = "root:@Va123456@tcp(45.77.170.201)/study_db?charset=utf8&parseTime=True"
)

func Init()  {
	prod := flag.Bool("prod", false, "a string")
	flag.Parse()
	if *prod {
		Host = "2study.edu.vn"
		Port = "80"
		FilePath = "/opt/web"
		MysqlConnect = "root:@Va123456@tcp(45.77.170.201)/study_db?charset=utf8&parseTime=True"
		fmt.Println(FilePath)
		fmt.Println("product")
	}else{
		fmt.Println("develop")
	}
}
// /public/dist/backend/
// /public/dist/frondtend/