package main
import (
   "fmt"
	4     "reflect" // 这里引入reflect模块
	5 )
	6 type User struct {
		7     Name   string "user name" //这引号里面的就是tag
		8     Passwd string "user passsword"
		9 }
		10 func main() {
			11     user := &User{"chronos", "pass"}
			12     s := reflect.TypeOf(user).Elem() //通过反射获取type定义
			13     for i := 0; i < s.NumField(); i++ {
				14         fmt.Println(s.Field(i).Tag) //将tag输出出来
				15     }
				16 }
