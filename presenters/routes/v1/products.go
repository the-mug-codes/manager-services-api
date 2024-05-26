package v1

import (
	"github.com/gin-gonic/gin"
	controllerProducts "github.com/kodit-tecnologia/service-manager/presenters/controllers/products"
	controllerProductsCategories "github.com/kodit-tecnologia/service-manager/presenters/controllers/products/categories"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
)

func Products(router *gin.RouterGroup) {
	productsRoute := router.Group("products")
	{
		categoriesRoute := productsRoute.Group("categories")
		{
			categoriesRoute.POST("", controllerProductsCategories.Insert)
			categoriesRoute.GET("", controllerProductsCategories.ReadAll)
			categoriesRoute.GET(":id", controllerProductsCategories.Read)
			categoriesRoute.PUT(":id", controllerProductsCategories.Update)
			categoriesRoute.DELETE(":id", middleware.Protected(nil, nil), controllerProductsCategories.Delete)
		}
		productsRoute.POST("", controllerProducts.Insert)
		productsRoute.GET("", controllerProducts.ReadAll)
		productsRoute.GET(":id", controllerProducts.Read)
		productsRoute.PUT(":id", controllerProducts.Update)
		productsRoute.DELETE(":id", middleware.Protected(nil, nil), controllerProducts.Delete)
	}
}
