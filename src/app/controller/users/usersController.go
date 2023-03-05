package users

import (
	"fmt"
	"kakeibo2/src/app/model"
	"kakeibo2/src/app/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UsersController struct {
	UserService service.UserService
	Ctx         iris.Context
}

// メソッド名でパスの違い・リクエストメソッド・パラメータを受け付けます（超便利）
// [例]
// GetList()                GET: "http://localhost:8080/users/list"
// Post()                   POST: "http://localhost:8080/users"
// PutDetails()             PUT: "http://localhost:8080/users/details"
// PutDetailsBy(id uint)    PUT: "http://localhost:8080/users/details/{id:int}"
// DeleteDetails()          DELETE: "http://localhost:8080/users/details"
// DeleteDetailsBy(id uint) DELETE: "http://localhost:8080/users/details/{id:int}"
// [その他の例]
// POST: http://localhost:8080/users/details/example -> PostDetailsExample()
// PUT: http://localhost:8080/users -> Put()

func (c *UsersController) GetList() mvc.Response {
	// 一覧取得
	users, err := c.UserService.GetUserList()
	fmt.Printf("==== =%v\n", users)
	if err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError, // エラーハンドリング
		}
	}
	// Iris に備え付きのレスポンス用構造体（struct）
	return mvc.Response{
		Code:   iris.StatusOK,
		Object: users,
	}
}

func (c *UsersController) Post() mvc.Response {
	// リクエストボディのjsonデータを構造体（struct）に格納する
	var user model.User
	err := c.Ctx.ReadJSON(&user)

	// エラーハンドリング（Iris 備え付きのもので作れます）
	if err != nil {
		c.Ctx.StopWithError(iris.StatusBadRequest, err)
		return mvc.Response{
			Code: iris.StatusBadRequest,
		}
	}

	// 新規作成
	err = c.UserService.CreateUser(&user)
	if err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError, // エラーハンドリング
		}
	}

	// Iris 備え付きのレスポンス用構造体（struct）
	return mvc.Response{Code: iris.StatusCreated}
}

// func (c *UsersController) PutDetailsBy(id int) mvc.Response {
// 	// リクエストボディのjsonデータを構造体（struct）に格納する
// 	var user model.User
// 	err := c.Ctx.ReadJSON(&user)

// 	// エラーハンドリング（Iris 備え付きのもので作れます）
// 	if err != nil {
// 		c.Ctx.StopWithError(iris.StatusBadRequest, err)
// 		return mvc.Response{
// 			Code: iris.StatusBadRequest,
// 		}
// 	}

// 	user.Id = uint32(id)

// 	// 更新
// 	err = userService.UpdateUser(&user)
// 	if err != nil {
// 		return mvc.Response{
// 			Code: iris.StatusInternalServerError, // エラーハンドリング
// 		}
// 	}

// 	// Iris 備え付きのレスポンス用構造体（struct）
// 	return mvc.Response{Code: iris.StatusOK}
// }

func (c *UsersController) DeleteDetailsBy(id int) mvc.Response {
	// 削除
	err := c.UserService.DeleteUser(id)
	if err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError, // エラーハンドリング
		}
	}

	// Iris 備え付きのレスポンス用構造体（struct）
	return mvc.Response{Code: iris.StatusOK}
}
