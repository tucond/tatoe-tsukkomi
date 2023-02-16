package main

import (
	"net/http"
	"os"

	oapi "tatoe-tsukkomi/api/generated"

	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type apiController struct{}

// レスポンス
type TsukkomiResponse struct {
	Tsukkomi *string `json:"tsukkomi"`
}

// OpenAPI で定義された (GET /tsukkomi) の実装
func (apiController) GetTsukkomi(ctx echo.Context, params oapi.GetTsukkomiParams) error {
	// OpenApi で生成された Tsukkomi モデルを使ってレスポンスを返す
	return ctx.JSON(http.StatusOK, TsukkomiResponse{
		Tsukkomi: getGptResponse(*params.Tsukkomi),
	})
}

func main() {

	// Echo のインスタンス作成
	e := echo.New()

	// OpenApi 仕様に沿ったリクエストかバリデーションをするミドルウェアを設定
	swagger, err := oapi.GetSwagger()
	if err != nil {
		panic(err)
	}
	e.Use(oapiMiddleware.OapiRequestValidator(swagger))
	// ロガーのミドルウェアを設定
	e.Use(middleware.Logger())
	// APIがエラーで落ちてもリカバーするミドルウェアを設定
	e.Use(middleware.Recover())

	// OpenAPI の仕様を満たす構造体をハンドラーとして登録する
	api := apiController{}
	oapi.RegisterHandlers(e, api)

	// Cloud Run用設定
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 8080ポートで Echo サーバ起動
	e.Logger.Fatal(e.Start(":" + port))

}
