package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	// set defaut user's account
	u := user{ID: 1, Email: "admin@admin.com", Password: "admin"}

	members.Members = append(members.Members, u)
}

func login(c *gin.Context) {
	isLogin, _ := c.Get("logined")

	fmt.Println(isLogin)
	if isLogin == true {
		fmt.Println("Logined")
		c.Redirect(http.StatusFound, "/result")
	}

	reade(c, "login.html", gin.H{
		"title":     "login",
		"actionUrl": "/login",
		// "msg":       msg,
	})
}

func attempt(c *gin.Context) {
	email, _ := c.GetPostForm("email")
	password, _ := c.GetPostForm("password")

	if email == "" || password == "" {
		c.Redirect(http.StatusFound, "/login")
	} else if members.verify(email, password) {
		log.Println("Login Scuuess")

		// rand => generate a rand number
		// strconv.FormatInt => int convert to string
		c.SetCookie("login", strconv.FormatInt(rand.Int63(), 20), 600, "/", "", false, true)

		c.Redirect(http.StatusFound, "/result")

	} else {
		c.Redirect(http.StatusFound, "/login")
	}

}

func logout(c *gin.Context) {
	log.Println("logout")

	isLogin, _ := c.Get("logined")

	if isLogin != true {
		c.Redirect(http.StatusFound, "/login")
	}

	c.SetCookie("login", "", -1, "", "", false, true)

	// 重新定向到 /login
	c.Redirect(http.StatusFound, "/login")
}

func (m account) verify(email string, pwd string) bool {
	isVerify := false
	for _, u := range m.Members {
		if email == u.Email && pwd == u.Password {
			isVerify = true
			break
		}
	}
	return isVerify
}

func checkLogined() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("login"); err == nil && token != "" {
			c.Set("logined", true)
		} else {
			c.Set("logined", false)
		}
	}
}
