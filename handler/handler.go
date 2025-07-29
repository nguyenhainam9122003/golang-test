package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/domain"
)

type ProductHandler struct {
	usecase domain.ProductUsecase
}

func NewProductHandler(u domain.ProductUsecase) *ProductHandler {
	return &ProductHandler{u}
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.usecase.FetchAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.usecase.FindByID(c, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}
