package models

import "attendance/database"

// Grade 年级模型
type Grade struct {
	Id    string `json:"id"`    //主键ID
	Grade string `json:"grade"` //年级名
}

// AddGrade 添加年级
func AddGrade(grade Grade) (id string, err error) {

	result := database.GetMySQL().Create(&grade)

	if result.Error != nil {
		err = result.Error
	}

	return grade.Id, err
}

// FindAllGrade 查询所有年级
func FindAllGrade() (grade []Grade, err error) {

	result := database.GetMySQL().Find(&grade)

	if result.Error != nil {
		err = result.Error
	}

	return grade, err
}
