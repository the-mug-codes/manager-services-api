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
			categoriesRoute.POST("", middleware.Protected(nil, nil), controllerProductsCategories.Insert)
			categoriesRoute.GET("", middleware.Protected(nil, nil), controllerProductsCategories.ReadAll)
			categoriesRoute.GET(":id", middleware.Protected(nil, nil), controllerProductsCategories.Read)
			categoriesRoute.PUT(":id", middleware.Protected(nil, nil), controllerProductsCategories.Update)
			categoriesRoute.DELETE(":id", middleware.Protected(nil, nil), controllerProductsCategories.Delete)
		}
		productsRoute.POST("", middleware.Protected(nil, nil), controllerProducts.Insert)
		productsRoute.GET("", middleware.Protected(nil, nil), controllerProducts.ReadAll)
		productsRoute.GET(":id", middleware.Protected(nil, nil), controllerProducts.Read)
		productsRoute.PUT(":id", middleware.Protected(nil, nil), controllerProducts.Update)
		productsRoute.DELETE(":id", middleware.Protected(nil, nil), controllerProducts.Delete)
	}
}
