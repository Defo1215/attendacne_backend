package controllers

import (
	"attendance/app/models"
	"attendance/app/result"
	"attendance/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

// SaveStudent 保存学生信息
func SaveStudent(c *gin.Context) {

	var student models.Student
	err := c.ShouldBindJSON(&student) //绑定JSON数据
	if err != nil {
		fmt.Println(err)
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	student.Id = uuid.NewString()   //生成主键ID
	student.Password = "123456"     //设置默认密码
	student.Status = "1"            //设置默认状态
	student.CreateTime = time.Now() //设置创建时间
	student.UpdateTime = time.Now() //设置更新时间

	err = models.SaveStudent(student)
	if err != nil {
		c.JSON(200, result.Fail("保存失败"))
		return
	}

	c.JSON(200, result.Success("保存成功"))

}

// FindAllFromStudent 查询所有学生信息
func FindAllFromStudent(c *gin.Context) {

	students, err := models.FindAllFromStudent()
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(students))
}

// FindStudentByMajorIdAndGradeAndStudentId 根据专业Id、年级和学号查询学生信息
func FindStudentByMajorIdAndGradeAndStudentId(c *gin.Context) {

	majorId := c.Query("majorId")
	grade := c.Query("grade")
	studentId := c.Query("studentId")

	student, err := models.FindStudentByMajorIdAndGradeAndStudentId(majorId, grade, studentId)
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(student))
}

// FindStudentByUsername 根据用户名查询学生信息
func FindStudentByUsername(c *gin.Context) {

	username := c.Query("username")

	student, err := models.FindStudentByUsername(username)
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(student))
}

// LoginStudent 学生登录
func LoginStudent(c *gin.Context) {

	var student models.StudentLogin
	err := c.ShouldBindJSON(&student) //绑定JSON数据
	if err != nil {
		fmt.Println(err)
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	studentInfo, err := models.LoginStudent(student.Username, student.Password)
	if err != nil {
		c.JSON(200, result.Fail("登录失败"))
		return
	}

	token, err := utils.GenerateToken(studentInfo.Id, studentInfo.Username)
	if err != nil {
		return
	}

	studentInfo.Token = token //设置token

	c.JSON(200, result.Success(studentInfo))
}
