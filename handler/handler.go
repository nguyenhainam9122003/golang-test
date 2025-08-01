package handler

import (
	"net/http"
	"strconv"
	"test/domain"
	"test/model"
	"test/utils"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service domain.ProductService
}

func NewProductHandler(r *gin.Engine, service domain.ProductService) {
	h := &ProductHandler{Service: service}

	r.GET("/products", h.GetAll)
	r.POST("/products", h.Create)
	r.GET("products/paginate", h.GetPaginatedProducts)
	r.GET("/products/:id", h.GetByID)
	r.PUT("/products/:id", h.UpdateProduct)
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	products, err := h.Service.GetAll(c)

	if utils.HandleError(c, err, http.StatusInternalServerError, "Failed to fetch products") {
		return
	}

	utils.Respond(c, http.StatusOK, true, "Lấy sản phẩm thành công", products)

}

func (h *ProductHandler) GetPaginatedProducts(c *gin.Context) {
	// Lấy page và limit từ query string, mặc định nếu không truyền
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	query := c.DefaultQuery("q", "")

	// Chuyển thành số nguyên
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var filter model.ProductFilter
	if utils.HandleError(c, c.ShouldBindQuery(&filter), http.StatusBadRequest, "Invalid filter params") {
		return
	}

	if utils.HandleError(c, filter.Validate(), http.StatusBadRequest, "") {
		return
	}

	products, err := h.Service.GetPaginated(c.Request.Context(), limit, offset, query, filter)
	if utils.HandleError(c, c.ShouldBindQuery(&filter), http.StatusInternalServerError, "Unable to fetch products") {
		return
	}

	paged := utils.PaginatedData{
		Page:  page,
		Limit: limit,
		Items: products,
	}
	utils.Respond(c, http.StatusOK, true, "Lấy sản phẩm thành công", paged)
}

func (h *ProductHandler) Create(c *gin.Context) {
	var product model.Product
	if utils.HandleError(c, c.ShouldBindJSON(&product), http.StatusBadRequest, "Invalid input") {
		return
	}
	if utils.HandleError(c, product.Validate(), http.StatusBadRequest, "") {
		return
	}

	err := h.Service.Create(c, &product)

	if utils.HandleError(c, err, http.StatusInternalServerError, "") {
		return
	}
	utils.Respond(c, http.StatusCreated, true, "Tạo phẩm thành công", product)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if utils.HandleError(c, err, http.StatusBadRequest, "Invalid ID\"") {
		return
	}

	product, err := h.Service.GetByID(c, uint(id))
	if utils.HandleError(c, err, http.StatusNotFound, "Product not found\"") {
		return
	}
	utils.Respond(c, http.StatusOK, true, "Lấy sản phẩm thành công", product)

}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if utils.HandleError(c, err, http.StatusBadRequest, "Invalid ID\"") {
		return
	}

	var updatedProduct model.Product
	if utils.HandleError(c, c.ShouldBindJSON(&updatedProduct), http.StatusBadRequest, "Invalid input") {
		return
	}

	if utils.HandleError(c, updatedProduct.Validate(), http.StatusBadRequest, "") {
		return
	}

	err = h.Service.Update(c.Request.Context(), uint64(id), &updatedProduct)
	if utils.HandleError(c, err, http.StatusBadRequest, "") {
		return
	}

	utils.Respond(c, http.StatusOK, true, "Product updated successfully", nil)
}
