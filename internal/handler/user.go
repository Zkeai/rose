package handler

import (
	"net/http"

	"github.com/Zkeai/rose/internal/dto"
	"github.com/gin-gonic/gin"
)

func userDetail(c *gin.Context) {
	r := new(dto.UserDetailReq)
	if err := c.Bind(r); err != nil {
		return
	}
	ret, err := svc.UserDetail(c.Request.Context(), r.Uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, ret)
}
