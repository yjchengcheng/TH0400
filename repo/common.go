package repo

import (
	"TH0400/utils"
	"log"
	"time"

	"TH0400/repo/ent"

	_ "github.com/go-sql-driver/mysql"

	"github.com/facebookincubator/ent/dialect/sql"
)

var edb *ent.Client

// 初始化数据库连接
func init() {
	drv, err := sql.Open("mysql", utils.GetSQLURI())
	if err != nil {
		log.Fatal("connect mysql err:", err.Error())
	}

	db := drv.DB()

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	client := ent.NewClient(ent.Driver(drv))

	ctx, f := utils.GetMTimeoutCtx()
	defer f()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %s", err.Error())
	}

	edb = client
}
