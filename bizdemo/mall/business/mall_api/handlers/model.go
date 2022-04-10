package handlers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

var AuthMiddleware *jwt.GinJWTMiddleware

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type LoginResponse struct {
	Code   int64  `json:"code"`
	Expire string `json:"expire"`
	Token  string `json:"token"`
}

type ShopSettleParam struct {
	ShopName string `json:"shop_name"`
}

type BrandAddParam struct {
	BrandName  string `json:"brand_name"`
	Logo       string `json:"logo"`
	BrandStory string `json:"brand_story"`
}

type BrandEditParam struct {
	BrandId    int64   `json:"brand_id" binding:"required"`
	BrandName  *string `json:"brand_name"`
	Logo       *string `json:"logo"`
	BrandStory *string `json:"brand_story"`
}

type BrandDelParam struct {
	BrandId int64 `json:"brand_id" binding:"required"`
}
