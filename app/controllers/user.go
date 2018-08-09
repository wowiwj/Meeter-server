package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Meeter/app/response"
	"github.com/lexkong/log"
	"fmt"
)


// Create creates a new user account.
func UserCreate(c *gin.Context)  {
	var r struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var err error
	if err = c.Bind(&r);err != nil {
		c.JSON(http.StatusOK,gin.H{
			"error": response.ErrBind,
		})
		return
	}

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	if r.Username == "" {
		err = response.New(response.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
		log.Errorf(err, "Get an error")
	}

	if response.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty")
	}
	code, message := response.DecodeErr(err)

	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})

}
