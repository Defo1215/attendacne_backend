package controllers

import (
	"attendance/app/models"
	"attendance/app/result"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SaveMajor 保存专业信息
func SaveMajor(c *gin.Context) {
	var major models.Major
	err := c.ShouldBindJSON(&major)
	if err != nil {
		return
	}

	major.Id = uuid.NewString()

	err = models.SaveMajor(major)
	if err != nil {
		c.JSON(200, result.Fail("保存失败"))
		return
	}

	c.JSON(200, result.Success("保存成功"))
}

// FindAllFromMajor 查询所有专业信息
func FindAllFromMajor(c *gin.Context) {
	majors, err := models.FindAllFromMajor()
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(majors))
}

// FindMajorByMajorId 根据专业ID查询专业信息
func FindMajorByMajorId(c *gin.Context) {
	majorId := c.Query("majorId")

	major, err := models.FindMajorByMajorId(majorId)
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(major))
}

// FindMajorByMajor 根据专业名查询专业信息
func FindMajorByMajor(c *gin.Context) {
	major := c.Query("major")

	majors, err := models.FindMajorByMajor(major)
	if err != nil {
		c.JSON(200, result.Fail("查询失败"))
		return
	}

	c.JSON(200, result.Success(majors))
}
