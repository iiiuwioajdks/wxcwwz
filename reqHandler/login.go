package reqHandler

import (
	"encoding/json"
	"github.com/asmcos/requests"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"wvCheck/global"
)

func Login(c *gin.Context) {
	var school global.LoginForm
	err := c.ShouldBind(&school)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
			"err": err.Error(),
		})
		return
	}

	req := requests.Requests()

	targetUrl := "https://xxcapp.xidian.edu.cn/uc/wap/login/check"

	payload := requests.Datas{
		"username": school.Username,
		"password": school.Password,
	}

	var msg Check

	rsp, err := req.Post(targetUrl, payload)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(rsp.Content(), &msg)

	if msg.E != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "学号或者密码错误",
		})
		return
	}

	var lm LoginModel
	lm.UserName = school.Username
	lm.Password = school.Password
	var count int64
	global.DB.Model(&LoginModel{}).Where("user_name=?", lm.UserName).Count(&count)
	if count == 0 {
		global.DB.Save(&lm)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}

func GetAll() []*LoginModel {
	var lms []*LoginModel
	global.DB.Find(&lms)
	return lms
}

type LoginModel struct {
	gorm.Model
	UserName string
	Password string
}
