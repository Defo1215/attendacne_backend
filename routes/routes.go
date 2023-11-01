package routes

import (
	"attendance/app/controllers"
	"attendance/app/middleware/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/public", "./public") // 静态文件服务

	r.Use(cors.Cors()) //启用跨域中间件

	//路由组件
	root := r.Group("/api")
	{
		//学生路由组
		student := root.Group("/student")
		{
			student.POST("/save", controllers.SaveStudent)
			student.GET("/findAll", controllers.FindAllFromStudent)
			student.GET("/findStudentByMajorIdAndGradeAndStudentId", controllers.FindStudentByMajorIdAndGradeAndStudentId)
			student.POST("/login", controllers.LoginStudent)

		}
		teacher := root.Group("/teacher")
		{
			teacher.POST("/login", controllers.LoginTeacher)
		}
		// 专业路由组
		major := root.Group("/major")
		{
			major.POST("/save", controllers.SaveMajor)
			major.GET("/findAll", controllers.FindAllFromMajor)
			major.GET("/findByMajorId", controllers.FindMajorByMajorId)
			major.GET("/findByMajor", controllers.FindMajorByMajor)
		}
		//年级路由组
		grade := root.Group("/grade")
		{
			grade.GET("/hello", func(c *gin.Context) {
				c.JSON(200, "这是一个测试接口")
			})
		}
		//课程路由组
		course := root.Group("/course")
		{
			course.POST("/save", controllers.SaveCourse)
			course.GET("/findAll", controllers.FindAllFromCourse)
			course.GET("/findByCourseId", controllers.FindCourseByCourseId)
			course.GET("/findByCourse", controllers.FindCourseByCourse)
			course.GET("/findByGradeAndMajorId", controllers.FindCourseByGradeAndMajorId)
			course.GET("/findByTeacherId", controllers.FindCourseByTeacherId)
			course.GET("/findByGradeAndTeacherId", controllers.FindCourseByGradeAndTeacherId)
			course.GET("/checkByCourseIdAndStudentUsername", controllers.CheckCourseByCourseIdAndStudentUsername)
		}
		//记录路由组
		record := root.Group("/record")
		{
			record.POST("/save", controllers.SaveRecord)
			record.GET("/findGroupByTeacherId", controllers.FindRecordsGroupByTeacherId)
			record.GET("/findNotSignedByMajorIdAndGradeAndClassAndCourseIdAndDate", controllers.FindNotSignedByMajorIdAndGradeAndClassAndCourseIdAndDate)
			record.GET("/findSignedByCourseIdAndDate", controllers.FindSignedByCourseIdAndDate)
			record.GET("/findAbsenceByCourseIdAndDate", controllers.FindAbsenceByCourseIdAndDate)
			record.POST("/updateStatus", controllers.UpdateRecordStatus)
			record.POST("/deleteByRecordId", controllers.DeleteRecordByRecordId)
		}
		qrcode := root.Group("/qrcode")
		{
			qrcode.POST("/init", controllers.InitQrcode)
			qrcode.GET("/findById", controllers.FindQrcodeById)
			qrcode.GET("/findAll", controllers.FindQrcodeList)
		}
	}

	return r
}
