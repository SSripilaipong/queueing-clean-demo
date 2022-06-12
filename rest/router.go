package rest

import (
	"github.com/gin-gonic/gin"
	d "queueing-clean-demo/rest/deps"
	"queueing-clean-demo/rest/routes/check_visits"
	"queueing-clean-demo/rest/routes/create_doctor_queue"
	"queueing-clean-demo/rest/routes/create_visit"
	"queueing-clean-demo/rest/routes/push_visit_to_doctor_queue"
	"queueing-clean-demo/rest/routes/submit_assessment"
	"queueing-clean-demo/toolbox/endpoint"
)

func getApiRouter(deps *d.RestDeps) *gin.Engine {
	router := gin.Default()

	router.GET("/doctor-queues/:doctorId", endpoint.R(check_visits.Route, deps))
	router.POST("/doctor-queues/:doctorId", endpoint.R(create_doctor_queue.Route, deps))
	router.POST("/doctor-queues/:doctorId/visits", endpoint.R(push_visit_to_doctor_queue.Route, deps))

	router.POST("/his/visits", endpoint.R(create_visit.Route, deps))
	router.PUT("/his/visits/:visitId/assessment", endpoint.R(submit_assessment.Route, deps))

	return router
}
