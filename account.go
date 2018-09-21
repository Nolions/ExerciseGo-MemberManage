package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// data struct of user
type user struct {
	ID       int
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type account struct {
	Members []user
}

var members account

func result(c *gin.Context) {
	if isLogin, _ := c.Get("logined"); isLogin != true {
		c.Redirect(http.StatusFound, "/login")
	}

	reade(c, "result.html", gin.H{
		"msg":               "Welcome back",
		"addAccountAction":  "/create",
		"editAccountAction": "/edit",
		"logoutAcion":       "/logout",
		"accounts":          members.Members,
	})
}

func addAccount(c *gin.Context) {
	if isLogin, _ := c.Get("logined"); isLogin != true {
		c.Redirect(http.StatusFound, "/login")
	}

	reade(c, "addAccount.html", gin.H{
		"createAction": "/create",
		"CancelUrl":    "/result",
	})
}

func editAccount(c *gin.Context) {
	// strconv.Atoi => string cover to int
	// c.Params 讀取route's param value
	id, _ := strconv.Atoi(c.Param("id"))

	members.getUser(id)

	reade(c, "editAccount.html", gin.H{})
}

func sotreAccount(c *gin.Context) {
	var form user

	if isLogin, _ := c.Get("logined"); isLogin != true {
		c.Redirect(http.StatusFound, "/login")
	}

	if c.ShouldBind(&form) == nil {
		log.Println("Create Account is Successed")

		email, _ := c.GetPostForm("email")
		password, _ := c.GetPostForm("password")
		members.addUser(email, password)
		c.Redirect(http.StatusFound, "/result")
	} else {
		log.Println("Create Account is Failed")

		c.Redirect(http.StatusFound, "/create")
	}
}

func (m account) addUser(email, pwd string) bool {
	u := user{ID: len(m.Members) + 1, Email: email, Password: pwd}

	members.Members = append(members.Members, u)

	return true
}

func (m account) getUser(id int) (*user, error) {
	for _, member := range m.Members {
		if member.ID == id {
			return &member, nil
		}
	}

	return nil, errors.New("The account is not found")
}
