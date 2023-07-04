package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"team_project/models"
)

func PostFeedback(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		res := models.NetResponse{}.Build()
		var feedback models.Feedback
		context.BindJSON(&feedback)
		if len(feedback.PackageName) == 0 || len(feedback.Version) < 0 || len(feedback.Content) < 0 || len(feedback.PackageName) > 100 || len(feedback.Version) > 100 || len(feedback.Content) > models.MAX_CONTENT {
			res.SetStatus(http.StatusBadRequest, models.StatusError, "Feedback not valid")
			context.JSON(res.Generate())
			context.Abort()
			return
		}
		feedback.Ip = context.ClientIP()
		models.InsertFeedback(db, &feedback)
		res.SetStatus(http.StatusOK, models.StatusOk, "Feedback submitted")
		context.JSON(res.Generate())
	}
}
