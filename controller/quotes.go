package controller

import (
	"net/http"
	"surgicalsteel-shop/utils"

	"github.com/gin-gonic/gin"
)

func HandleGetBuyerQuoteById(c *gin.Context) {
	idRaw := c.Param("id")

	id := utils.ParseInt64(idRaw)

	quote, err := getBuyerQuoteById(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, quote)

}
