package controllers

import (
	"net/http"

	"github.com/Hona-Tahlil/Backend/internal/infrastructure/validation"
	"github.com/gin-gonic/gin"
)

func Receive[T any](ctx *gin.Context) T {
	var params T

	if err := ctx.ShouldBind(&params); err != nil {
		// TODO: Proper Error Handling
		ctx.JSON(http.StatusBadRequest, "wrong input")
	}

	if err := ctx.ShouldBindUri(&params); err != nil {
		// TODO: Proper Error Handling
		ctx.JSON(http.StatusBadRequest, "wrong input")
	}

	if err := ctx.ShouldBindQuery(&params); err != nil {
		// TODO: Proper Error Handling
		ctx.JSON(http.StatusBadRequest, "wrong input")
	}

	validation.ValidateFields(params)

	return params
}
