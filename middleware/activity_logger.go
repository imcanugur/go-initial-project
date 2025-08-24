package middleware

import (
	"bytes"
	"fmt"
	"go-initial-project/entity"
	"go-initial-project/service"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

func ActivityLogger(activityService *service.ActivityService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		c.Next()

		var userID *string
		if val, exists := c.Get("user_id"); exists {
			str := fmt.Sprintf("%v", val)
			userID = &str
		}

		path := c.Request.URL.Path

		activity := &entity.Activity{
			UserID:    userID,
			Action:    "request",
			Path:      path,
			Method:    c.Request.Method,
			IP:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Request:   string(bodyBytes),
			Status:    c.Writer.Status(),
			CreatedAt: time.Now(),
		}

		if err := activityService.Log(activity); err != nil {
			fmt.Println("❌ Activity log DB err:", err)
		}
	}
}
