package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

type GinResponse struct {
	C       *gin.Context
	ErrCode int
}

type ResponseSuccess struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseError struct {
	Status string `json:"status"`
	//Code       int    `json:"code"`
	ErrMessage string `json:"err_message"`
}

func (ginres *GinResponse) ResSuccess(data interface{}) {
	var ressuccess ResponseSuccess
	ressuccess.Status = "ok"
	ressuccess.Data = data
	ginres.C.JSON(200, &ressuccess)
}

func (ginres *GinResponse)  ResError(errcode int,err error) {
	var reserror ResponseError
	reserror.Status = "error"
	reserror.ErrMessage = GetErrMessage(errcode)
	if err != nil {
		log.Print(err)
	}
	ginres.C.JSON(200, reserror)
}
