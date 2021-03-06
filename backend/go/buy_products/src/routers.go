package buy_products

import (
	// "fmt"
	"errors"
	"net/http"
	"gobuys_products/common"
	"github.com/gin-gonic/gin"
)

func Buy_ProductsRegister(router *gin.RouterGroup) {
	router.POST("/:slug", Buy_ProductCreate)
	router.GET("/", Buy_ProductList)
	router.GET("/:slug", Buy_ProductRetrieve)
}

func Buy_ProductCreate(c *gin.Context) {

	slug := c.Param("slug")

	err := InsertBuyProduct(&ProductModel{Slug: slug})


	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data":"okey"})
}

func Buy_ProductList(c *gin.Context) {
	buy_productModels, modelCount, err := FindManyBuy_Products()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("buy_products", errors.New("Invalid param")))
		return
	}
	serializer := Buy_ProductsSerializer{c, buy_productModels}
	c.JSON(http.StatusOK, gin.H{"buy_products": serializer.Response(), "buy_productsCount": modelCount})
}
func Buy_ProductFeed(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	myUsers := c.MustGet("my_user_model").(Users)
	if myUsers.ID == 0 {
		c.AbortWithError(http.StatusUnauthorized, errors.New("{error : \"Require auth!\"}"))
		return
	}
	buy_productUsers := GetBuy_ProductUsers(myUsers)
	buy_productModels, modelCount, err := buy_productUsers.GetBuy_ProductFeed(limit, offset)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("buy_products", errors.New("Invalid param")))
		return
	}
	serializer := Buy_ProductsSerializer{c, buy_productModels}
	c.JSON(http.StatusOK, gin.H{"buy_products": serializer.Response(), "buy_productsCount": modelCount})
}

func Buy_ProductRetrieve(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "feed" {
		Buy_ProductFeed(c)
		return
	}
	buy_productModel, err := FindOneBuy_Product(&Buy_ProductModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("buy_products", errors.New("Invalid slug")))
		return
	}
	serializer := Buy_ProductSerializer{c, buy_productModel}
	c.JSON(http.StatusOK, gin.H{"buy_product": serializer.Response()})
}