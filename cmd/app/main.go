package main

// import (
// 	"TH0400/service"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	userregister := service.UserService{
// 		UserName:  "yjc",
// 		PassWord:  "123456",
// 		School:    "qinhua",
// 		Level:     3,
// 		Likes:     5,
// 		Views:     6,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 		IsDeleted: false,
// 	}

// 	resUserRegister, err := userregister.Register()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	fmt.Println(resUserRegister)

// }

import (
	th0400 "TH0400"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	th0400.RunRestServer()
	// todo 启动grpc

}
