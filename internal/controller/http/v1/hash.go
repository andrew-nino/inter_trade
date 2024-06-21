package v1

import (
	"international_trade/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MessageOK = "success "
)

// @Summary		Adding new value to DB
// @Security		ApiKeyAuth
// @Tags			hash
// @Description	Create a new hash if not found.
// @ID				adding-new-hash-value
// @Accept			json
// @Produces		json
// @Param			input	body		entity.StringToHash	true	"key"
// @Success		200		{string}	string				1		"hash"
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/v1/hash/sha256 [post]
func (h *Handler) addingHash(c *gin.Context) {

	var input entity.StringToHash
	typeHash := c.Param("type")

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	hash, err := h.services.ServingString.AddingHash(input.String, typeHash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, hash)
}

// @Summary		Delete hash from service
// @Security		ApiKeyAuth
// @Tags			hash
// @Description	Delete hash from both db
// @ID				delete-hash
// @Accept			json
// @Produces		json
// @Param			input	body		entity.StringToHash	true	"key"
// @Success		200
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/v1/hash/sha256 [delete]
func (h *Handler) deleteHash(c *gin.Context) {

	var input entity.StringToHash
	typeHash := c.Param("type")

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.ServingString.DeleteHash(input.String, typeHash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, MessageOK)
}

// @Summary		Get value
// @Security		ApiKeyAuth
// @Tags			hash
// @Description	Get value from service
// @ID				get-hash
// @Accept			json
// @Produce		json
// @Param		key		query		string	true	"key"
// @Success		200		{object}	string
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/v1/hash/sha256  [get]
func (h *Handler) getValue(c *gin.Context) {

	typeHash := c.Param("type")

	str, _ := c.GetQuery("key")
	hash, err := h.services.ServingString.GetHash(str, typeHash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, hash)
}
