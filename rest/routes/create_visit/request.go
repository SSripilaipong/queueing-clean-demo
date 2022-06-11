package create_visit

import (
	"github.com/gin-gonic/gin"
	"queueing-clean-demo/domain/clinical_diagnose/contract"
)

type jsonBody struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

func makeRequest(ctx *gin.Context) (clinical_diagnose.CreateVisit, error) {
	var body jsonBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return clinical_diagnose.CreateVisit{}, err
	}

	return clinical_diagnose.CreateVisit{
		Name:   body.Name,
		Gender: body.Gender,
		Age:    body.Age,
	}, nil
}
