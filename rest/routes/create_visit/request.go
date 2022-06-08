package create_visit

import (
	"github.com/gin-gonic/gin"
	"queueing-clean-demo/domain/usecase/clinical_diagnose/contract"
)

type jsonBody struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

func makeRequest(ctx *gin.Context) (contract.CreateVisit, error) {
	var body jsonBody
	if err := ctx.ShouldBindJSON(body); err != nil {
		return contract.CreateVisit{}, err
	}

	return contract.CreateVisit{
		Name:   body.Name,
		Gender: body.Gender,
		Age:    body.Age,
	}, nil
}
