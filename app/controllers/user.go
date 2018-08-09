package controllers

import (
	"Meeter/app/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"net/http"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

// Create creates a new user account.
func UserCreate(c *gin.Context) {
	var r CreateRequest
	var err error
	if err = c.Bind(&r); err != nil {
		response.Json(c, response.ErrBind, nil)
		c.JSON(http.StatusOK, gin.H{
			"error": response.ErrBind,
		})
		return
	}

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	if r.Username == "" {
		err = response.New(response.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
		response.Json(c, err, nil)
		log.Errorf(err, "Get an error")
		return
	}

	if response.IsErrUserNotFound(err) {
		response.Json(c, err, nil)
		log.Debug("err type is ErrUserNotFound")
		return
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty")
	}
	rsp := CreateResponse{
		Username: r.Username,
	}
	response.Json(c, nil, rsp)

}
