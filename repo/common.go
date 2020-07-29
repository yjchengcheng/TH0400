package repo

import (
	"TH0400/utils"
	"log"

	"TH0400/repo/ent"
)

var edb *ent.Client

// 初始化数据库连接
func init() {
	client, err := ent.Open("mysql", "root:962464@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal("connect mysql err:", err.Error())
	}

	ctx, f := utils.GetTimeoutCtx()
	defer f()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %s", err.Error())
	}

	edb = client
}
