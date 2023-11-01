package models

import "attendance/database"

type Major struct {
	Id    string `json:"id"`    //专业ID(主键)
	Major string `json:"major"` //专业名
}

// SaveMajor 保存专业信息
func SaveMajor(major Major) error {
	return database.GetMySQL().Create(&major).Error
}

// FindAllFromMajor 查询所有专业信息
func FindAllFromMajor() (majors []Major, err error) {
	result := database.GetMySQL().Find(&majors)

	return majors, result.Error
}

// FindMajorByMajorId 根据ID查询专业信息
func FindMajorByMajorId(majorId string) (major Major, err error) {
	result := database.GetMySQL().Where("id = ?", majorId).First(&major)

	return major, result.Error
}

// FindMajorByMajor 根据专业名查询专业信息
func FindMajorByMajor(major string) (majors []Major, err error) {
	result := database.GetMySQL().Where("major = ?", major).Find(&majors)

	return majors, result.Error
}
