package main

import (
	//"database/sql"

	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	//"log"
	//"net/http"

	_ "github.com/lib/pq"

	"./psql"
	//"./router"
)

func main() {
	psql.Db = psql.ConnectDB()

	psql.GetUserTopic(1)

	// http.HandleFunc("/querytime", router.Query_test_time) //显示所有用户信息
	// http.HandleFunc("/showuserinfo", router.ShowUserInfo)       // 个人信息
	// http.HandleFunc("/showuserquelist", router.ShowUserQueList) // 个人提问列表
	// http.HandleFunc("/showuseranslist", router.ShowUserAnsList) // 个人回答列表
	// http.HandleFunc("/addtopics", router.AddTopics)             // 添加帖子

	// err := http.ListenAndServe(":4000", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	//}
}