package router

import (
	"2022/ginseckill/backend/backendAuthController"
	"2022/ginseckill/middleware"
	"github.com/gin-gonic/gin"
)

func BackendAuthGroup(backendGroup  *gin.RouterGroup) {
	backendAuthGroup := backendGroup.Group("/auth")
	backendAuthGroup.Use(middleware.BackendAuth())
	{
		backendAuthGroup.GET("/ping",backendAuthController.Ping)
		// 所有商品
		backendAuthGroup.POST("/goods/list",backendAuthController.GoodsList)
		// 更新商品
		backendAuthGroup.POST("/goods/update",backendAuthController.GoodsUpdate)
		// 添加商品
		backendAuthGroup.POST("/goods/add",backendAuthController.InsertGoods)
		// 删除一件商品
		backendAuthGroup.POST("/goods/remove",backendAuthController.DeleteGoodsOne)
		// 单条件查询
		backendAuthGroup.POST("/goods/select",backendAuthController.SelectOneGoods)

		// 所有订单
		backendAuthGroup.POST("/order/list",backendAuthController.OrderList)
		// 更新订单
		backendAuthGroup.POST("/order/update",backendAuthController.OrderUpdate)
		// 添加订单
		backendAuthGroup.POST("/order/add",backendAuthController.InsertOrder)
		// 删除一件订单
		backendAuthGroup.POST("/order/remove",backendAuthController.DeleteOrderOne)
		// 单条件查询
		backendAuthGroup.POST("/order/select",backendAuthController.SelectOneOrder)
	}
}