package main

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	ID uint64       `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (app *application) SignIn(c *gin.Context) {

	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "please provide { username, password } in json format")
		return
	}

	// authenticate user
	if user.Email != "admin@gmail.com" || user.Password != "admin" {
		c.JSON(http.StatusUnauthorized, "invalid login details")
		return
	}

	// create auth token
	uid := uuid.NewV4().String()
	authToken, err := NewAuthToken(uid, app.hmacSecret)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// write token details to redis with expiry
	expire := time.Unix(authToken.Exp, 0)
	now := time.Now()

	err = app.redisClient.Set(authToken.Uid, strconv.Itoa(int(user.ID)), expire.Sub(now)).Err()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	resp := make(map[string]interface{})
	resp["token"] = authToken.Token
	c.JSON(http.StatusOK, resp)
}

type Task struct {
	Details string
}

func (app *application) CreateTask(c *gin.Context) {
	var t Task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	tokenStr, err := ExtractJWTFromRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "missing jwt")
		return
	}

	err = VerifyClaims(tokenStr, app.hmacSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "missing jwt")
		return
	}
}
