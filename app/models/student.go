package models

import (
	"attendance/database"
	"time"
)

// Student 学生模型
type Student struct {
	Id          string    `json:"id"`          //主键ID
	StudentId   string    `json:"studentId"`   //学号
	StudentName string    `json:"studentName"` //姓名
	MajorId     string    `json:"majorId"`     //专业ID
	Grade       string    `json:"grade"`       //年级
	Class       string    `json:"class"`       //班级
	Gender      string    `json:"gender"`      //性别
	Username    string    `json:"username"`    //用户名
	Password    string    `json:"password"`    //密码
	Status      string    `json:"status"`      //状态
	CreateTime  time.Time `json:"createTime"`  //创建时间
	UpdateTime  time.Time `json:"updateTime"`  //更新时间
}

// StudentLogin 学生登录结构体
type StudentLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// StudentLoginResponse 学生登录返回结构体
type StudentLoginResponse struct {
	Id          string `json:"id"`          //主键ID
	StudentId   string `json:"studentId"`   //学号
	StudentName string `json:"studentName"` //姓名
	MajorId     string `json:"majorId"`     //专业ID
	Major       string `json:"major"`       //专业名
	Grade       string `json:"grade"`       //年级
	Class       string `json:"class"`       //班级
	Gender      string `json:"gender"`      //性别
	Username    string `json:"username"`    //用户名
	Token       string `json:"token"`       //token
}

// SaveStudent 保存学生信息
func SaveStudent(student Student) error {
	return database.GetMySQL().Create(&student).Error
}

// FindAllFromStudent 查询所有学生信息
func FindAllFromStudent() (students []Student, err error) {
	result := database.GetMySQL().Find(&students)

	return students, result.Error
}

// FindStudentByMajorIdAndGradeAndStudentId 根据专业ID、年级和学号查询学生信息
func FindStudentByMajorIdAndGradeAndStudentId(majorId, grade, studentId string) (student Student, err error) {
	result := database.GetMySQL().Where("major_id = ? AND grade = ? AND student_id = ?", majorId, grade, studentId).First(&student)

	return student, result.Error
}

// FindStudentByMajorIdAndGradeAndClass 根据专业Id、年级和班级查询学生信息
func FindStudentByMajorIdAndGradeAndClass(majorId, grade, class string) (students []Student, err error) {
	result := database.GetMySQL().Where("major_id = ? AND grade = ? AND class = ?", majorId, grade, class).Find(&students)

	return students, result.Error
}

// FindStudentByUsername 根据用户名查询学生信息
func FindStudentByUsername(username string) (student Student, err error) {
	result := database.GetMySQL().Where("username = ?", username).First(&student)

	return student, result.Error
}

// LoginStudent 学生登录
func LoginStudent(username, password string) (student StudentLoginResponse, err error) {
	sql := `
		SELECT
			student.id,
			student.student_id,
			student.student_name,
			student.major_id,
			major.major,
			student.grade,
			student.class,
			student.gender,
			student.username
		FROM
		    student
		JOIN major  on major.id = student.major_id
		WHERE
		    username = ? AND password = ?
`
	result := database.GetMySQL().Raw(sql, username, password).Scan(&student)

	return student, result.Error
}
