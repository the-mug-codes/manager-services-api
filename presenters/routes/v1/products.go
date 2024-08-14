package v1

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/the-mug-codes/adapters-service-api/server/middlewares"
	controllerProducts "github.com/the-mug-codes/service-manager-api/presenters/controllers/products"
	controllerProductsCategories "github.com/the-mug-codes/service-manager-api/presenters/controllers/products/categories"
)

func Products(router *gin.RouterGroup) {
	productsRoute := router.Group("products")
	{
		categoriesRoute := productsRoute.Group("categories")
		{
			categoriesRoute.POST("", middleware.Protected(&[]string{"admin:full"}), controllerProductsCategories.Insert)
			categoriesRoute.GET("", middleware.Protected(&[]string{"admin:full"}), controllerProductsCategories.ReadAll)
			categoriesRoute.GET(":id", middleware.Protected(&[]string{"admin:full"}), controllerProductsCategories.Read)
			categoriesRoute.PUT(":id", middleware.Protected(&[]string{"admin:full"}), controllerProductsCategories.Update)
			categoriesRoute.DELETE(":id", middleware.Protected(&[]string{"admin:full"}), controllerProductsCategories.Delete)
		}
		productsRoute.POST("", middleware.Protected(&[]string{"admin:full"}), controllerProducts.Insert)
		productsRoute.GET("", middleware.Protected(&[]string{"admin:full"}), controllerProducts.ReadAll)
		productsRoute.GET(":id", middleware.Protected(&[]string{"admin:full"}), controllerProducts.Read)
		productsRoute.PUT(":id", middleware.Protected(&[]string{"admin:full"}), controllerProducts.Update)
		productsRoute.DELETE(":id", middleware.Protected(&[]string{"admin:full"}), controllerProducts.Delete)
	}
}
