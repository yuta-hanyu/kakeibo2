package service

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yuta-hanyu/kakeibo-api/src/app/model"
)

type UserService struct{}

func (UserService) GetUserList() ([]model.User, error) {
	// initialize the DbMap
	dbMap := InitDb()
	defer dbMap.Db.Close()

	var users []model.User

	// ユーザーを全取得
	_, err := dbMap.Select(&users, `SELECT * FROM users`)
	if err != nil {
		fmt.Printf("=====%v\n", err)
		return []model.User{}, err
	}

	return users, nil
}

func (UserService) CreateUser(user *model.User) error {
	// initialize the DbMap
	dbMap := InitDb()
	defer dbMap.Db.Close()

	// トラン ザクションを走らせながらinsert
	tx, _ := dbMap.Begin()

	err := tx.Insert(user)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (UserService) UpdateUser(user *model.User) error {
	// initialize the DbMap
	dbMap := InitDb()
	defer dbMap.Db.Close()

	// トランザクションを走らせながらupdate
	tx, _ := dbMap.Begin()

	_, err := tx.Update(user)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (UserService) DeleteUser(id int) error {
	// initialize the DbMap
	dbMap := InitDb()
	defer dbMap.Db.Close()

	// id から削除するユーザーを取得
	var user model.User
	err := dbMap.SelectOne(&user, `SELECT * FROM users WHERE id = :id`,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		fmt.Printf("error! can't find user by id: %v.\n", id)
		return err
	}

	// トランザクションを走らせながらdelete
	tx, _ := dbMap.Begin()

	_, err = tx.Delete(&user)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
