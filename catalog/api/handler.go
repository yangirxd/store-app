package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yangirxd/store-app/catalog/api/dto"
	_ "github.com/yangirxd/store-app/catalog/docs"
	"github.com/yangirxd/store-app/catalog/service"
	"net/http"
)

// @Summary Create a new product
// @Description Create a new product in the catalog (requires authentication)
// @Tags products
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param input body dto.CreateProductRequest true "Product data"
// @Success 201 {object} domain.Product
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/products [post]
func createProductHandler(catalogService *service.CatalogService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product, err := catalogService.CreateProduct(req.Name, req.Description, req.Price, req.Stock)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, product)
	}
}

// @Summary Get product by ID
// @Description Get details of a product by its UUID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} domain.Product
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Product not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/products/{id} [get]
func getProductHandler(catalogService *service.CatalogService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		product, err := catalogService.GetProductByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

// @Summary Get all products
// @Description Get a list of all products (public endpoint)
// @Tags products
// @Produce json
// @Success 200 {array} domain.Product
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/products [get]
func getAllProductsHandler(catalogService *service.CatalogService) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := catalogService.GetAllProducts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

// @Summary Update a product
// @Description Update details of an existing product (requires authentication)
// @Tags products
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Product ID"
// @Param input body dto.UpdateProductRequest true "Updated product data"
// @Success 200 {object} domain.Product
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Product not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/products/{id} [put]
func updateProductHandler(catalogService *service.CatalogService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var req dto.UpdateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product, err := catalogService.UpdateProduct(id, req.Name, req.Description, req.Price, req.Stock)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

// @Summary Delete a product
// @Description Delete a product by its UUID (requires authentication)
// @Tags products
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Product ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Product not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/products/{id} [delete]
func deleteProductHandler(catalogService *service.CatalogService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := catalogService.DeleteProduct(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{})
	}
}
