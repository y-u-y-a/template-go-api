package main

import (
	"fmt"
	"net/http"
	"time"
	"y-u-y-a/template-go/config"
	"y-u-y-a/template-go/utils"

	chi "github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/k0kubun/pp"
	"go.uber.org/zap"
)

var location = "Asia/Tokyo"

func init() {
	// タイムゾーンの設定
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
	utils.Clock = utils.NewRealTimeClock()
}

func main() {
	// 環境変数の出力
	pp.Println("appConfig is", config.AppEnv)

	// ロギングの設定
	var logger *zap.Logger
	if config.AppEnv.IsProd() {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}

	/*****************************
	DB接続
	******************************/
	// dbConnection, err := sql.Open(
	// 	"postgres",
	// 	fmt.Sprintf(
	// 		"postgres://%s:%s@%s:%s/%s",
	// 		config.AppEnv.DbUser, config.AppEnv.DbPassword,
	// 		config.AppEnv.DbHost, config.AppEnv.DbPort,
	// 		config.AppEnv.DbName,
	// 	))
	// if err != nil {
	// 	logger.Fatal("DB接続に失敗しました", zap.Error(err))
	// }
	// defer dbConnection.Close()

	/*****************************
	ミドルウェアの設定
	******************************/
	r := chi.NewRouter()
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.Timeout(120 * time.Second))
	r.Use(chiMiddleware.Logger)
	// CORSの設定
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"Set-Cookie"},
		MaxAge:           3600,
	}))

	/*****************************
	サーバー起動
	******************************/
	if err := http.ListenAndServe(fmt.Sprintf(":%v", config.AppEnv.ServerPort), r); err != nil {
		logger.Fatal("サーバー起動に失敗しました", zap.Error(err))
	}

}
