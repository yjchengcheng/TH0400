package main

import (
	th0400 "TH0400"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	th0400.RunRestServer()
	// todo 启动grpc
}
