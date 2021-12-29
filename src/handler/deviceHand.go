package handler

import (
	"FrontAccess/src/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func DeviceHand(c *gin.Context) {

	imei := c.Param("imei")
	log.Info(imei)
	DeviceID, err := model.GetDeviceId(imei)

	if err != nil {
		log.Error(err)
	}

	c.JSON(200, gin.H{"data": DeviceID})

}

func UpdateDeviceHand(c *gin.Context) {

	imei := c.Param("imei")
	log.Info(imei)

	DeviceName := c.PostForm("DeviceName")
	ch := make(chan error, 3000)

	go func() {
		err := model.UpdateDeviceId(imei, DeviceName)
		ch <- err
	}()

	err := <-ch
	if err != nil {
		log.Error(err)
		c.JSON(500, gin.H{"data": err})
	} else {
		c.JSON(200, gin.H{"data": "OK"})
	}

}
