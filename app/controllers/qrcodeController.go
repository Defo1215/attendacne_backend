package controllers

import (
	"attendance/app/result"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"sync"
	"time"
)

// QrcodeInfo 二维码信息
type QrcodeInfo struct {
	Major          string `json:"major"`          //专业名
	MajorId        string `json:"majorId"`        //专业ID
	Grade          string `json:"grade"`          //年级
	Course         string `json:"course"`         //课程名
	CourseId       string `json:"courseId"`       //课程ID
	Class          string `json:"class"`          //班级
	QrcodeDuration int    `json:"qrcodeDuration"` //二维码有效时长
	QrcodeRefresh  int    `json:"qrcodeRefresh"`  //二维码刷新频率
	ExpiresAt      int64  `json:"expiresAt"`      //二维码过期时间
}

// QRCodeTask 二维码任务
type QRCodeTask struct {
	ID         string
	ExpiresAt  time.Time
	UpdateRate time.Duration
	ClientData QrcodeInfo
	Data       []byte
}

var qrCodeTasks map[string]*QRCodeTask
var qrCodeMutex sync.RWMutex

func InitQrCodeTasks() {
	qrCodeTasks = make(map[string]*QRCodeTask)
}

// InitQrcode 初始化二维码
func InitQrcode(c *gin.Context) {

	var qrcodeInfo QrcodeInfo

	err := c.ShouldBindJSON(&qrcodeInfo)

	if err != nil {
		c.JSON(200, result.Fail("参数错误"))
		return
	}

	id := qrcodeInfo.MajorId + "-" + qrcodeInfo.Grade + "-" + qrcodeInfo.CourseId + "-" + qrcodeInfo.Class

	// 检查是否已存在相同ID的认为,如果存在则返回错误或处理方式
	if _, exists := qrCodeTasks[id]; exists {
		// 删除已存在的二维码任务
		delete(qrCodeTasks, id)
	}

	// 获取二维码的有效期
	expiresAt := time.Now().Add(time.Duration(qrcodeInfo.QrcodeDuration) * time.Minute)
	// 获取二维码的刷新频率
	updateRate := time.Duration(qrcodeInfo.QrcodeRefresh) * time.Second

	qrcodeInfo.ExpiresAt = time.Now().UnixMilli() + int64(qrcodeInfo.QrcodeRefresh)*1000

	// 使用客户端数据生成新的二维码内容
	qrcodeContent := generateNewQRCodeContent(qrcodeInfo)
	// 生成新的二维码
	qrcodeData, err := qrcode.Encode(qrcodeContent, qrcode.Medium, 256)

	// 创建二维码任务
	task := &QRCodeTask{
		ID:         id,
		ExpiresAt:  expiresAt,
		UpdateRate: updateRate,
		ClientData: qrcodeInfo,
		Data:       qrcodeData,
	}

	// 将二维码任务添加到任务列表
	qrCodeTasks[id] = task

	// 启动一个 go routine 来定时更新二维码
	go func() {
		for {
			select {
			case <-time.After(task.UpdateRate):
				// 使用客户端数据生成新的二维码内容
				task.ClientData.ExpiresAt = time.Now().UnixMilli() + int64(task.ClientData.QrcodeRefresh)*1000
				qrcodeContent := generateNewQRCodeContent(task.ClientData)
				// 生成新的二维码
				qrcodeData, err := qrcode.Encode(qrcodeContent, qrcode.Medium, 256)
				if err != nil {
					fmt.Println("生成二维码失败:", err)
					continue
				}
				// 更新二维码数据
				task.Data = qrcodeData
			case <-time.After(time.Until(task.ExpiresAt)):
				// 二维码过期后删除任务
				qrCodeMutex.Lock()
				delete(qrCodeTasks, task.ID)
				qrCodeMutex.Unlock()
				return
			}
		}
	}()

	// 返回二维码任务ID
	c.JSON(200, result.Success(task))
}

// FindQrcodeById 根据ID查找二维码
func FindQrcodeById(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(200, result.Fail("二维码ID不能为空"))
		return
	}

	//qrCodeMutex.RLock()
	task, exists := qrCodeTasks[id]
	//qrCodeMutex.RUnlock()

	if !exists {
		c.JSON(200, result.Fail("二维码不存在"))
		return
	}

	// 检查二维码任务是否过期
	if time.Now().After(task.ExpiresAt) {
		c.JSON(200, result.Fail("二维码已过期"))
		return
	}

	c.Data(200, "image/png", task.Data)
}

// FindQrcodeList 查找二维码列表
func FindQrcodeList(c *gin.Context) {
	c.JSON(200, result.Success(qrCodeTasks))
}

// 生成新的二维码内容
func generateNewQRCodeContent(qrcodeInfo QrcodeInfo) string {
	qrcodeContent := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v", qrcodeInfo.Major, qrcodeInfo.MajorId, qrcodeInfo.Grade, qrcodeInfo.Course, qrcodeInfo.CourseId, qrcodeInfo.Class, qrcodeInfo.QrcodeDuration, qrcodeInfo.QrcodeRefresh, qrcodeInfo.ExpiresAt, time.Now())
	qrcodeContentBase64 := base64.StdEncoding.EncodeToString([]byte(qrcodeContent))

	return qrcodeContentBase64
}
