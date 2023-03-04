package service

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/yuta-hanyu/kakeibo-api/src/app/model"

	"github.com/go-gorp/gorp"
	"github.com/joho/godotenv"
	"github.com/yuta-hanyu/kakeibo-api/src/app/database/migration"
)

// gorp初期化処理
func InitDb() *gorp.DbMap {
	err := godotenv.Load("src/app/environment/local.env")
	if err != nil {
		log.Fatalln("envを読み込み出来ませんでした", err)
	}

	// mysql接続
	db, err := sql.Open(os.Getenv("DB_DRIVER"),
		fmt.Sprintf("%s:%s@tcp(%s)/",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
		))
	if err != nil {
		log.Fatalln("mysqlの接続に失敗しました", err)
	}
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// データベース作成
	cmdU := fmt.Sprintf(`CREATE DATABASE %s;`, os.Getenv("DB_NAME"))
	dbMap.Exec(cmdU)

	// kakeiboデータベース接続
	Db, err := sql.Open(
		os.Getenv("DB_DRIVER"),
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
		),
	)
	if err != nil {
		log.Fatalln("kakeiboデータベースの接続に失敗しました", err)
	}

	// テーブルマッピング
	dbMap = &gorp.DbMap{Db: Db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// テーブル作成
	err = migration.InitMigration(Db)
	if err != nil {
		log.Fatal("テーブルの初期化に失敗しました", err)
	}

	dbMap.AddTableWithName(model.User{}, "users").SetKeys(true, "Id")

	err = dbMap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal("usersテーブルの作成に失敗しました", err)
	}

	// ログの取得
	dbMap.TraceOn("[gorp]", log.New(os.Stdout, "go-iris-sample:", log.Lmicroseconds))

	return dbMap
}
