package main

import (
	"FrontAccess/src/handler"
	"FrontAccess/src/service"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

const (
	a1 = iota
	a2 = 100
	a3 = iota
	a4
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

//八进制
var a = 012

// 16进制
var b = 0xff

func main() {

	runtime.GOMAXPROCS(4)

	r := gin.Default()
	pprof.Register(r)
	r.Use(service.LogMiddleWare())

	gin.SetMode(gin.ReleaseMode)
	r.POST("/front_ddc_service_api/v1.0/handler/user", handler.UserHand)
	r.POST("/front_ddc_service_api/v1.0/handler/:imei", handler.DeviceHand)
	r.POST("/front_ddc_service_api/v1.0/handler/:imei/update", handler.UpdateDeviceHand)

	r.GET("/log", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "请求成功",
		})
	})

	r.Run(":8080")

}
