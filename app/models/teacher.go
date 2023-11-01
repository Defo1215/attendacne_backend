package models

import "attendance/database"

// Teacher 教师模型
type Teacher struct {
	Id          string `json:"id"`           //主键ID
	Username    string `json:"username"`     //用户名
	Password    string `json:"password"`     //密码
	TeacherName string `json:"teacher_name"` //姓名
	CreateTime  int64  `json:"createTime"`   //创建时间
	UpdateTime  int64  `json:"updateTime"`   //更新时间
}

// TeacherLogin 教师登录模型
type TeacherLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// TeacherLoginResponse 教师登录返回模型
type TeacherLoginResponse struct {
	Id          string `json:"id"`          //主键ID
	Username    string `json:"username"`    //用户名
	TeacherName string `json:"teacherName"` //姓名
	Token       string `json:"token"`       //token
}

func SaveTeacher(teacher Teacher) error {
	return database.GetMySQL().Create(&teacher).Error
}

// FindAllFromTeacher 查询所有教师信息
func FindAllFromTeacher() (teachers []Teacher, err error) {
	result := database.GetMySQL().Find(&teachers)

	return teachers, result.Error
}

// FindTeacherByTeacherId 根据教师ID查询教师信息
func FindTeacherByTeacherId(teacherId string) (teacher Teacher, err error) {
	result := database.GetMySQL().Where("id = ?", teacherId).First(&teacher)

	return teacher, result.Error
}

// LoginTeacher 教师登录
func LoginTeacher(username, password string) (teacher TeacherLoginResponse, err error) {
	sql := "SELECT id, username, teacher_name FROM teacher WHERE username = ? AND password = ?"

	result := database.GetMySQL().Raw(sql, username, password).Scan(&teacher)

	return teacher, result.Error
}
