package models

import (
	"attendance/database"
	"time"
)

// Record 考勤记录模型
type Record struct {
	Id         string    `json:"id"`         //主键ID
	MajorId    string    `json:"majorId"`    //专业ID
	Grade      string    `json:"grade"`      //年级
	StudentId  string    `json:"studentId"`  //学号
	CourseId   string    `json:"courseId"`   //课程ID
	Status     string    `json:"status"`     //状态
	CreateTime time.Time `json:"createTime"` //创建时间
	UpdateTime time.Time `json:"updateTime"` //更新时间
}

// RecordsGroupResponse 考勤记录分组返回结构体
type RecordsGroupResponse struct {
	CourseId  string    `json:"courseId"`  //课程ID
	Course    string    `json:"course"`    //课程名
	Grade     string    `json:"grade"`     //年级
	Major     string    `json:"major"`     //专业名
	MajorId   string    `json:"majorId"`   //专业ID
	Class     string    `json:"class"`     //班级
	Date      time.Time `json:"date"`      //日期
	StartTime time.Time `json:"startTime"` //开始时间
}

// RecordsResponse 考勤记录返回结构体
type RecordsResponse struct {
	Id          string    `json:"id"`          //主键ID
	Grade       string    `json:"grade"`       //年级
	StudentId   string    `json:"studentId"`   //学号
	StudentName string    `json:"studentName"` //学生姓名
	CourseId    string    `json:"courseId"`    //课程ID
	Status      string    `json:"status"`      //状态
	CreateTime  time.Time `json:"createTime"`  //创建时间
	UpdateTime  time.Time `json:"updateTime"`  //更新时间
}

// NotSignedResponse 未签到学生返回结构体
type NotSignedResponse struct {
	StudentId   string `json:"studentId"`   //学号
	StudentName string `json:"studentName"` //学生姓名
	Grade       string `json:"grade"`       //年级
}

// SaveRecord 保存考勤记录
func SaveRecord(record Record) error {
	return database.GetMySQL().Create(&record).Error
}

// FindAllFromRecord 查询所有考勤记录
func FindAllFromRecord() (records []Record, err error) {
	result := database.GetMySQL().Find(&records)

	return records, result.Error
}

// FindRecordsByRecordId 根据考勤记录ID查询考勤记录
func FindRecordsByRecordId(recordId string) (record Record, err error) {
	result := database.GetMySQL().Where("id = ?", recordId).First(&record)

	return record, result.Error
}

// FindRecordsByCourseId 根据课程ID查询考勤记录
func FindRecordsByCourseId(courseId string) (records []Record, err error) {
	result := database.GetMySQL().Where("course_id = ?", courseId).Find(&records)

	return records, result.Error
}

// FindRecordsByMajorAndGradeAndStudentId 根据专业ID、年级和学号查询考勤记录
func FindRecordsByMajorAndGradeAndStudentId(majorId, grade, studentId string) (records []Record, err error) {
	result := database.GetMySQL().Where("major_id = ? AND grade = ? AND student_id = ?", majorId, grade, studentId).Find(&records)

	return records, result.Error
}

// FindRecordsGroupByTeacherId 根据教师ID查询考勤记录组
func FindRecordsGroupByTeacherId(teacherId string) (recordsGroup []RecordsGroupResponse, err error) {

	sql := `
		SELECT DISTINCT 
			record.course_id,
			course.course,
			record.grade,
			course.major,
			course.major_id,
			course.class,
			DATE(record.create_time) as date,
			course.start_time as start_time
		FROM
		    record
		JOIN
		    course ON record.course_id = course.id
		WHERE
		    teacher_id = ?
		ORDER BY
		    date DESC`

	result := database.GetMySQL().Raw(sql, teacherId).Scan(&recordsGroup)

	return recordsGroup, result.Error
}

// FindNotSignedByMajorIdAndGradeAndClassAndCourseIdAndDate 根据专业ID、年级、班级、课程ID和日期查询未签到学生
func FindNotSignedByMajorIdAndGradeAndClassAndCourseIdAndDate(majorId, grade, courseId, date string, class []string) (records []NotSignedResponse, err error) {
	sql := `
		SELECT
			student.student_id,
			student.student_name,
			student.grade
		FROM
			student
		WHERE
			student.student_id NOT IN (
				SELECT
					record.student_id
				FROM
					record
				WHERE
				    record.course_id = ? AND DATE(record.create_time) = ?
			) AND student.status = '1' AND student.major_id = ? AND student.grade = ? AND student.class IN ? `

	result := database.GetMySQL().Raw(sql, courseId, date, majorId, grade, class).Scan(&records)

	return records, result.Error
}

// FindSignedByCourseIdAndDate 根据课程ID和日期查询已签到学生
func FindSignedByCourseIdAndDate(courseId, date string) (records []RecordsResponse, err error) {
	sql := `
		SELECT
			record.id,
			record.grade,
			record.student_id,
			student.student_name as student_name,
			record.course_id,
			record.status,
			record.create_time,
			record.update_time
		FROM
			record
		JOIN
			student ON record.student_id = student.student_id AND record.grade = student.grade
		WHERE
			record.course_id = ? AND DATE(record.create_time) = ? AND record.status = '1'`

	result := database.GetMySQL().Raw(sql, courseId, date).Scan(&records)

	return records, result.Error
}

// FindAbsenceByCourseIdAndDate 根据课程ID和日期查询缺勤学生
func FindAbsenceByCourseIdAndDate(courseId, date string) (records []RecordsResponse, err error) {
	sql := `
		SELECT
			record.id,
			record.grade,
			record.student_id,
			student.student_name as student_name,
			record.course_id,
			record.status,
			record.create_time,
			record.update_time
		FROM
			record
		JOIN
			student ON record.student_id = student.student_id AND record.grade = student.grade
		WHERE
			record.course_id = ? AND DATE(record.create_time) = ? AND record.status = '0'`

	result := database.GetMySQL().Raw(sql, courseId, date).Scan(&records)

	return records, result.Error
}

// UpdateRecordStatus 更新考勤记录状态
func UpdateRecordStatus(recordId, status string) error {
	return database.GetMySQL().Model(&Record{}).Where("id = ?", recordId).Updates(map[string]interface{}{"status": status, "update_time": time.Now()}).Error
}

// DeleteRecordByRecordId 删除考勤记录
func DeleteRecordByRecordId(recordId string) error {
	return database.GetMySQL().Where("id = ?", recordId).Delete(&Record{}).Error
}
