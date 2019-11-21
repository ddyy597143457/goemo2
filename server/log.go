package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func InitLogger() {
	//访问日志
	requestlogfile := "C:/Users/59714/go/src/ddyy/goemo2/logs/log.txt"
	f, err := os.OpenFile(requestlogfile, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = f

	//错误日志
	errorlogfile := "C:/Users/59714/go/src/ddyy/goemo2/logs/errlog.txt"
	f1, err1 := os.OpenFile(errorlogfile, os.O_CREATE|os.O_APPEND, 0666)
	if err1 != nil {
		panic(err1)
	}
	log.SetOutput(f1)
}
