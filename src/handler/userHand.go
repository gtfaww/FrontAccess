package handler

import (
	"FrontAccess/src/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func UserHand(c *gin.Context) {

	name := c.PostForm("user_name")
	log.Info(name)
	user, err := model.GetUserByName(name)

	if err != nil {
		log.Error(err)
	}

	c.JSON(200, gin.H{"data": user})

}
