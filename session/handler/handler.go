package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("the-key"))
var userData = make(map[string]Signupdata)

type Signupdata struct {
	ConfirmPassword string
	Email           string
	PhoneNumber     string
	Name            string
	Password        string
}

func HomePage(c *gin.Context) {
	session, _ := store.Get(c.Request, "sessio")
	if session.Values["email"] != nil {
		c.HTML(http.StatusOK, "homepage.html", nil)
	} else {
		c.Redirect(http.StatusSeeOther, "/login")
	}
}

func SignupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signupPage.html", nil)
}

func SignupPost(c *gin.Context) {
	firstname := c.PostForm("firstname")
	password := c.PostForm("password")
	phonenumber := c.PostForm("phonenumber")
	confirmpassword := c.PostForm("confirmpassword")
	email := c.PostForm("email")

	if firstname == "" || email == "" || password == "" || phonenumber == "" || confirmpassword != password {
		c.HTML(http.StatusUnauthorized, "signupPage.html", "Invalid entry")
		return
	}

	userData[email] = Signupdata{
		Email:           email,
		Password:        password,
		Name:            firstname,
		PhoneNumber:     phonenumber,
		ConfirmPassword: confirmpassword,
	}

	c.Redirect(http.StatusSeeOther, "/login")
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "loginPage.html", nil)
}

func Postmethod(c *gin.Context) {
	email := c.PostForm("emailLogin")
	password := c.PostForm("passwordLogin")

	SignupData, ok := userData[email]
	if !ok || SignupData.Password != password {
		c.HTML(http.StatusUnauthorized, "loginPage.html", "Invalid credentials")
		return
	}

	session, _ := store.Get(c.Request, "session")
	session.Values["email"] = email
	session.Save(c.Request, c.Writer)

	c.Redirect(http.StatusSeeOther, "/")
}

func Logout(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	session.Values["email"] = nil
	session.Save(c.Request, c.Writer)

	c.Redirect(http.StatusSeeOther, "/")
}
