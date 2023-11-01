package controllers

import (
	"attendance/app/models"
	"attendance/app/result"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
	"time"
)

// SaveRecord 保存考勤记录
func SaveRecord(c *gin.Context) {
	var record models.Record
	err := c.ShouldBindJSON(&record)
	if err != nil {
		c.JSON(200, result.Fail("参数错误"))
	}

	record.Id = uuid.NewString()

	fmt.Println(record.CreateTime)
	// 判断是否有创建时间，没有则设置创建时间
	if record.CreateTime.IsZero() {
		record.CreateTime = time.Now()
	}
	record.UpdateTime = time.Now()

	err = models.SaveRecord(record)
	if err != nil {
		c.JSON(200, result.Fail("保存失败"))
		return
	}

	c.JSON(200, result.Success("保存成功"))
}

// FindRecordsGroupByTeacherId 根据教师ID查询考勤记录组
func FindRecordsGroupByTeacherId(c *gin.Context) {
	teacherId := c.Query("teacherId")

	recordsGroup, err := models.FindRecordsGroupByTeacherId(teacherId)
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(recordsGroup))
}

// FindNotSignedByMajorIdAndGradeAndClassAndCourseIdAndDate 根据专业ID、年级、班级、课程ID和日期查询未签到学生
func FindNotSignedByMajorIdAndGradeAndClassAndCourseIdAndDate(c *gin.Context) {
	majorId := c.Query("majorId")
	grade := c.Query("grade")
	class := c.Query("class")
	courseId := c.Query("courseId")
	date := c.Query("date")

	classSlice := strings.Split(class, ",")

	records, err := models.FindNotSignedByMajorIdAndGradeAndClassAndCourseIdAndDate(majorId, grade, courseId, date, classSlice)
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(records))
}

// FindSignedByCourseIdAndDate 根据课程ID和日期查询已签到学生
func FindSignedByCourseIdAndDate(c *gin.Context) {
	courseId := c.Query("courseId")
	date := c.Query("date")

	records, err := models.FindSignedByCourseIdAndDate(courseId, date)
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(records))
}

// FindAbsenceByCourseIdAndDate 根据课程ID和日期查询缺勤学生
func FindAbsenceByCourseIdAndDate(c *gin.Context) {
	courseId := c.Query("courseId")
	date := c.Query("date")

	records, err := models.FindAbsenceByCourseIdAndDate(courseId, date)
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(records))
}

// UpdateRecordStatus 更新考勤记录状态
func UpdateRecordStatus(c *gin.Context) {
	type recordIdAndStatus struct {
		Id     string `json:"id"`
		Status string `json:"status"`
	}

	var record recordIdAndStatus

	err := c.ShouldBindJSON(&record)
	if err != nil {
		c.JSON(200, result.Fail("参数错误"))
	}

	err = models.UpdateRecordStatus(record.Id, record.Status)
	if err != nil {
		c.JSON(200, result.Fail("更新失败"))
		return
	}

	c.JSON(200, result.Success("更新成功"))
}

// DeleteRecordByRecordId 删除考勤记录
func DeleteRecordByRecordId(c *gin.Context) {
	type recordId struct {
		Id string `json:"id"`
	}

	var record recordId
	err := c.ShouldBindJSON(&record)
	if err != nil {
		c.JSON(200, result.Fail("参数错误"))
	}

	err = models.DeleteRecordByRecordId(record.Id)
	if err != nil {
		c.JSON(200, result.Fail("删除失败"))
		return
	}

	c.JSON(200, result.Success("删除成功"))
}
