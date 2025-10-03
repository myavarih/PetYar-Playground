package controllers

import "github.com/gin-gonic/gin"

type singleMessageResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func Respond(ctx *gin.Context, data interface{}, err error) {
	// TODO: check err -> generate messages -> Translate them -> Respond
}
