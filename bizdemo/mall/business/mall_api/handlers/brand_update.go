// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package handlers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/dal/client"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/kitex_gen/cmp/ecom/product"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/kitex_gen/cmp/ecom/shop"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/conf"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
	"github.com/gin-gonic/gin"
)

// BrandUpdate godoc
// @Summary 商家更新品牌信息
// @Description 商家更新品牌信息
// @Tags 商品模块-品牌子模块
// @Accept json
// @Produce json
// @Param shopSettleParam body handlers.BrandEditParam true "品牌信息"
// @Param Authorization header string true "Bearer $token"
// @Success 200 {object} handlers.Response
// @Router /product/brand_update [post]
func BrandUpdate(c *gin.Context) {
	var brandEditParam BrandEditParam
	if err := c.ShouldBind(&brandEditParam); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	claims := jwt.ExtractClaims(c)
	userID := int64(claims[conf.IdentityKey].(float64))
	shopId, err := client.GetShopIdByUserId(c, &shop.GetShopIdByUserIdReq{UserId: userID})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	req := &product.UpdateBrandReq{
		BrandId:    brandEditParam.BrandId,
		ShopId:     shopId,
		BrandName:  brandEditParam.BrandName,
		Logo:       brandEditParam.Logo,
		BrandStory: brandEditParam.BrandStory,
	}
	err = client.UpdateBrand(c, req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
