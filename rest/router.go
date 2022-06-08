package rest

import (
	"github.com/gin-gonic/gin"
	d "queueing-clean-demo/rest/deps"
	"queueing-clean-demo/rest/ext"
	"queueing-clean-demo/rest/routes/check_visits"
	"queueing-clean-demo/rest/routes/create_doctor_queue"
	"queueing-clean-demo/rest/routes/create_visit"
	"queueing-clean-demo/rest/routes/push_visit_to_doctor_queue"
	"queueing-clean-demo/rest/routes/submit_assessment"
)

func getApiRouter(deps *d.RestDeps) *gin.Engine {
	router := gin.Default()

	router.GET("/doctor-queues/:doctorId", ext.R(check_visits.Route, deps))
	router.POST("/doctor-queues/:doctorId", ext.R(create_doctor_queue.Route, deps))
	router.POST("/doctor-queues/:doctorId/visits", ext.R(push_visit_to_doctor_queue.Route, deps))

	router.POST("/his/visits", ext.R(create_visit.Route, deps))
	router.PUT("/his/visits/:visitId/assessment", ext.R(submit_assessment.Route, deps))

	return router
}
