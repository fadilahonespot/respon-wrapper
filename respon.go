package respon_wrapper

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponWrapper struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(c *gin.Context, data interface{}) {
	var returnData = ResponWrapper{
		Success: true,
		Message: "Success",
		Data: data,
	}
	c.JSON(http.StatusOK, returnData)
}

// HandleError is function wrap in respon failed
func HandleError(c *gin.Context, status int, message string) {
	var returnData = ResponWrapper{
		Success: false,
		Message: message,
	}
	c.JSON(status, returnData)
}

func HandleSuccessReturn(data interface{}) string {
	var returnData = ResponWrapper{
		Success: true,
		Message: "Success",
		Data: data,
	}
	res, _ := json.Marshal(&returnData)
	return string(res)
}