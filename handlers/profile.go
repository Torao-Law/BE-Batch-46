package handlers

import (
	profiledto "dumbmerch/dto/profile"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type profileHandler struct {
	ProfileRepository repositories.ProfileRepository
}

func ProfileHandler(ProfileRepository repositories.ProfileRepository) *profileHandler {
	return &profileHandler{ProfileRepository}
}

func (h *profileHandler) GetProfile(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	profile, err := h.ProfileRepository.GetProfile(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProfile(profile)})
}

func convertResponseProfile(u models.Profile) profiledto.ProfileResponse {
	return profiledto.ProfileResponse{
		ID:      u.ID,
		Phone:   u.Phone,
		Gender:  u.Gender,
		Address: u.Address,
		UserID:  u.UserID,
		User:    u.User,
	}
}
