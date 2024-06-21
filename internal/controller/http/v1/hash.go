package v1

import (
	"international_trade/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MessageOK = "success "
)

func (h *Handler) addingHash(c *gin.Context) {

	var input entity.StringToHash

	typeHash := c.Param("type")

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	hash, err := h.services.ServingString.CreateNewHash(input.String, typeHash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, hash)
}

func (h *Handler) deleteHash(c *gin.Context) {

	var input entity.StringToHash
	typeHash := c.Param("type")

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.services.ServingString.DeleteteHash(input.String, typeHash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, MessageOK)
}
