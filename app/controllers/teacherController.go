package controllers

import (
	"attendance/app/models"
	"attendance/app/result"
	"attendance/utils"
	"github.com/gin-gonic/gin"
)

func LoginTeacher(c *gin.Context) {
	var teacher models.TeacherLogin

	err := c.ShouldBindJSON(&teacher)
	if err != nil {
		c.JSON(200, result.Fail("参数错误"))
	}

	teacherInfo, err := models.LoginTeacher(teacher.Username, teacher.Password)

	if err != nil {
		c.JSON(200, result.Fail("登录失败"))
		return
	}

	token, err := utils.GenerateToken(teacherInfo.Id, teacherInfo.Username)

	if err != nil {
		c.JSON(200, result.Fail("登录失败"))
		return
	}

	teacherInfo.Token = token

	c.JSON(200, result.Success(teacherInfo))
}
