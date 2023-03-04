package migration

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var err error

const (
	tableNameUser = "users"
)

// usersテーブル作成
func InitMigration(Db *sql.DB) error {
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id int(11) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		name VARCHAR(15) NOT NULL,
		mail VARCHAR(255) NOT NULL,
		updated_at datetime,
		created_at datetime default current_timestamp);`, tableNameUser)
	_, err := Db.Exec(cmdU)
	return err
}
