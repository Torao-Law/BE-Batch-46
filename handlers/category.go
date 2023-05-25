package handlers

import (
	categorydto "dumbmerch/dto/category"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerCategory struct {
	CategoryRepository repositories.CategoryRepository
}

func HandlerCategory(CategoryRepository repositories.CategoryRepository) *handlerCategory {
	return &handlerCategory{CategoryRepository}
}

func (h *handlerCategory) FindCategories(c echo.Context) error {
	category, err := h.CategoryRepository.FindCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: category})
}

func (h *handlerCategory) GetCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	category, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCategory(category)})
}

func (h *handlerCategory) CreateCategory(c echo.Context) error {
	request := new(categorydto.CategoryRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	category := models.Category{
		Name: request.Name,
	}

	category, err = h.CategoryRepository.CreateCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	category, _ = h.CategoryRepository.GetCategory(category.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCategory(category)})
}

func convertResponseCategory(u models.Category) models.CategoryResponse {
	return models.CategoryResponse{
		Name:      u.Name,
		ProductID: u.ID,
		Product:   u.Product,
	}
}
