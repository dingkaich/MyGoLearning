package mytemplate

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

func MytemplateMain() {
	username := "dingkai"
	data := map[string]interface{}{"Username": username}
	t3, _ := template.ParseFiles("test1.html") //将一个文件读作模板
	t3.Execute(os.Stdout, data)
	fmt.Println(t3.Name(), "\n") //模板名称

}

//注入模板的函数
func tihuan(str string) string {
	return str + "-------" + time.Now().Format("2006-01-02")
}
