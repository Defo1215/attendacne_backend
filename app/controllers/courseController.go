package controllers

import (
	"attendance/app/models"
	"attendance/app/result"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SaveCourse 保存课程信息
func SaveCourse(c *gin.Context) {
	var course models.Course
	err := c.ShouldBindJSON(&course)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	course.Id = uuid.NewString()

	err = models.SaveCourse(course)
	if err != nil {
		c.JSON(200, result.Fail("保存失败"))
		return
	}

	c.JSON(200, result.Success("保存成功"))
}

// FindAllFromCourse 查询所有课程信息
func FindAllFromCourse(c *gin.Context) {
	courses, err := models.FindAllFromCourse()
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(courses))
}

// FindCourseByCourseId 根据课程ID查询课程信息
func FindCourseByCourseId(c *gin.Context) {
	courseId := c.Query("courseId")

	course, err := models.FindCourseByCourseId(courseId)
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(course))
}

// FindCourseByCourse 根据课程名查询课程信息
func FindCourseByCourse(c *gin.Context) {
	course := c.Query("course")

	courses, err := models.FindCourseByCourse(course)
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(courses))
}

// FindCourseByGradeAndMajorId 根据年级和专业ID查询课程信息
func FindCourseByGradeAndMajorId(c *gin.Context) {
	grade := c.Query("grade")
	majorId := c.Query("majorId")

	courses, err := models.FindCourseByGradeAndMajorId(grade, majorId)
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(courses))
}

// FindCourseByTeacherId 根据教师ID查询课程信息
func FindCourseByTeacherId(c *gin.Context) {
	teacherId := c.Query("teacherId")

	courses, err := models.FindCourseByTeacherId(teacherId)
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(courses))
}

// FindCourseByGradeAndTeacherId 根据年级和教师ID查询课程信息
func FindCourseByGradeAndTeacherId(c *gin.Context) {
	grade := c.Query("grade")
	teacherId := c.Query("teacherId")

	courses, err := models.FindCourseByGradeAndTeacherId(grade, teacherId)
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(courses))
}

// CheckCourseByCourseIdAndStudentUsername 根据课程ID和学生用户名查询课程信息
func CheckCourseByCourseIdAndStudentUsername(c *gin.Context) {
	courseId := c.Query("courseId")
	studentUsername := c.Query("studentUsername")

	exist, err := models.CheckCourseByCourseIdAndStudentUsername(courseId, studentUsername)
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	if exist {
		c.JSON(200, result.Success("该学生有该课程"))
		return
	}

	c.JSON(200, result.Fail("该学生没有该课程"))
}
