package api

import (
	"net/http"
	"strings"

	"github.com/baaj2109/webcam_server/global"
	"github.com/baaj2109/webcam_server/model"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Email    string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func Login(c *gin.Context) {
	// email := c.PostForm("email")
	// passwd := c.PostForm("passwd")
	auth := auth{}
	if err := c.ShouldBind(auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to login, bind auth model error",
		})
		return
	}
	auth.Password = global.Md5(auth.Password)
	// passwd = global.Md5(passwd)

	// auth := auth{Email: email, Password: passwd}
	if !model.CheckAuth(auth.Email, auth.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to login, can't find user in database",
		})
		return
	}

	c.SetCookie("passwd", auth.Password, 3600, "/", "localhost", false, false)
	c.SetCookie("email", auth.Email, 3600, "/", "localhost", false, false)
	c.JSON(http.StatusAccepted, gin.H{
		"message": "login successfully",
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("passwd", "", -1, "/", "localhost", false, false)
	c.SetCookie("email", "", -1, "/", "localhost", false, false)
}

func Register(c *gin.Context) {
	auth := auth{}
	if err := c.ShouldBind(auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to login, bind auth model error",
		})
		return
	}
	auth.Password = global.Md5(auth.Password)
	if model.CheckAuth(auth.Email, auth.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to register, account already exist",
		})
		return
	}
	ip := strings.Split(c.Request.RemoteAddr, ":")[0]
	user := model.User{
		Email:    auth.Email,
		Password: auth.Password,
		LastIP:   ip,
	}
	err := global.Db.Create(user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to register, failed set user to database",
		})
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "create user successfully",
	})

}
