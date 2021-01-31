package main

import (
	"blog_server/global"
	"blog_server/model"
	"blog_server/pkg/logger"
	"blog_server/pkg/setting"
	"blog_server/routers"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init setupLogger err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init setupDBEngine err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// description Go project
// @termsOfService https://github.com/iyolee/blog_server
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	// global.Logger.Infof("%s: logger", "blog_server")

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" +
		global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.Newlogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
