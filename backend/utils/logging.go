package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bắt đầu thời gian xử lý
		t := time.Now()

		// Chuyển tiếp yêu cầu đến handler tiếp theo
		c.Next()

		// Tính toán thời gian xử lý
		latency := time.Since(t)
		// Lấy mã trạng thái HTTP
		status := c.Writer.Status()
		// Lấy phương thức HTTP
		method := c.Request.Method
		// Lấy URL của yêu cầu
		path := c.Request.URL.Path
		// In ra log
		// access the status we are sending

		log.Println(status)
		fmt.Printf("[GIN] %d | %3d | %13v | %s | %s \n",
			status, latency, method, path, c.Errors.ByType(gin.ErrorTypePrivate).String())
	}
}
