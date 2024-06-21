package v1

import (
	"international_trade/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addingHash(c *gin.Context) {

	var input entity.StringToHash

	typeHash := c.Param("type")

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	add_time, err := h.services.ServingString.CreateNewHash(input.String, typeHash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, add_time)
}
