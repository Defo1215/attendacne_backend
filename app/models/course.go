package models

import (
	"attendance/database"
	"time"
)

// Course 课程模型
type Course struct {
	Id          string    `json:"id"`          //课程ID(主键)
	Course      string    `json:"course"`      //课程名
	Grade       string    `json:"grade"`       //年级
	Class       string    `json:"class"`       //班级
	MajorId     string    `json:"majorId"`     //专业ID
	Major       string    `json:"major"`       //专业名
	TeacherId   string    `json:"teacherId"`   //教师ID
	TeacherName string    `json:"teacherName"` //教师名
	Classroom   string    `json:"classroom"`   //教室
	StartTime   time.Time `json:"startTime"`   //开始时间
	EndTime     time.Time `json:"endTime"`     //结束时间
}

// SaveCourse 保存课程信息
func SaveCourse(course Course) error {
	return database.GetMySQL().Create(&course).Error
}

// FindAllFromCourse 查询所有课程信息
func FindAllFromCourse() (courses []Course, err error) {
	result := database.GetMySQL().Find(&courses)

	return courses, result.Error
}

// FindCourseByCourseId 根据课程ID查询课程信息
func FindCourseByCourseId(courseId string) (course Course, err error) {
	result := database.GetMySQL().Where("id = ?", courseId).First(&course)

	return course, result.Error
}

// FindCourseByCourse 根据课程名查询课程信息
func FindCourseByCourse(course string) (courses []Course, err error) {
	result := database.GetMySQL().Where("course = ?", course).Find(&courses)

	return courses, result.Error
}

// FindCourseByGradeAndMajorId 根据年级和专业ID查询课程信息
func FindCourseByGradeAndMajorId(grade string, majorId string) (courses []Course, err error) {
	result := database.GetMySQL().Where("grade = ? AND major_id = ?", grade, majorId).Find(&courses)

	return courses, result.Error
}

// FindCourseByTeacherId 根据教师ID查询课程信息
func FindCourseByTeacherId(teacherId string) (courses []Course, err error) {
	result := database.GetMySQL().Where("teacher_id = ?", teacherId).Find(&courses)

	return courses, result.Error
}

// FindCourseByGradeAndTeacherId 根据年级和教师ID查询课程信息
func FindCourseByGradeAndTeacherId(grade string, teacherId string) (courses []Course, err error) {
	result := database.GetMySQL().Where("grade = ? AND teacher_id = ?", grade, teacherId).Find(&courses)

	return courses, result.Error
}

// CheckCourseByCourseIdAndStudentUsername 根据课程ID和学生用户名查询课程信息
func CheckCourseByCourseIdAndStudentUsername(courseId, username string) (bool, error) {
	sql := `
		SELECT student.student_id, course.id AS course_id
		FROM student
		INNER JOIN course ON student.major_id = course.major_id
                 	AND student.grade = course.grade
                 	AND FIND_IN_SET(student.class, course.class) > 0
		WHERE  course.id = ? AND student.username = ? 
`
	result := database.GetMySQL().Raw(sql, courseId, username).Scan(nil)

	return result.RowsAffected > 0, result.Error
}
