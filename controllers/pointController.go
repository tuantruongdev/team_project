package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"team_project/models"
)

func GetPoints(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		res := models.NetResponse{}.Build()
		var update models.PointUpdate
		var currentDevice models.Device
		context.BindJSON(&update)
		if update.DeviceID == "" {
			res.SetStatus(http.StatusBadRequest, models.StatusError, "Missing deviceId")
			context.JSON(res.Generate())
			return
		}
		currentDevice, err := models.GetDevice(db, update.DeviceID)
		res.Set("firstTime", false)
		if err != nil {
			currentDevice = models.Device{
				DeviceId: update.DeviceID,
				Point:    models.DEFAULT_POINTS,
			}
			ok := models.InsertDevice(db, &currentDevice)
			if !ok {
				res.SetStatus(http.StatusBadRequest, models.StatusError, "Something went wrong")
				context.JSON(res.Generate())
				return
			}
			res.Set("firstTime", true)
		}
		res.SetStatus(http.StatusOK, models.StatusOk, "get point successfully")
		res.Set("user_id", currentDevice.Id)
		res.Set("points", currentDevice.Point)
		context.JSON(res.Generate())
		return

	}
}
func UpdatePoints(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		res := models.NetResponse{}.Build()
		var update models.PointUpdate
		var currentDevice models.Device
		context.BindJSON(&update)
		if update.DeviceID == "" {
			res.SetStatus(http.StatusBadRequest, models.StatusError, "Missing deviceId")
			context.JSON(res.Generate())
			return
		}
		if update.AdjustPoint > 500 || update.AdjustPoint < 1 {
			res.SetStatus(http.StatusBadRequest, models.StatusError, "Point not valid")
			context.JSON(res.Generate())
			return
		}

		currentDevice, err := models.GetDevice(db, update.DeviceID)
		res.Set("firstTime", false)
		if err != nil {
			currentDevice = models.Device{
				DeviceId: update.DeviceID,
				Point:    models.DEFAULT_POINTS,
			}
			ok := models.InsertDevice(db, &currentDevice)
			if !ok {
				res.SetStatus(http.StatusBadRequest, models.StatusError, "Something went wrong")
				context.JSON(res.Generate())
				return
			}
			res.Set("firstTime", true)
		}
		if update.UpdateType == 1 {
			currentDevice.Point += update.AdjustPoint
			ok := models.UpdatePoint(db, &currentDevice)
			if !ok {
				res.SetStatus(http.StatusBadRequest, models.StatusError, "Something went wrong 2")
				context.JSON(res.Generate())
				return
			}
			res.SetStatus(http.StatusOK, models.StatusOk, "update successfully")
			res.Set("points", currentDevice.Point)
			context.JSON(res.Generate())
			return
		} else if update.UpdateType == 2 {
			currentDevice.Point -= update.AdjustPoint
			if currentDevice.Point < 0 {
				res.SetStatus(http.StatusBadRequest, models.StatusError, "Negative point")
				context.JSON(res.Generate())
				return
			}
			ok := models.UpdatePoint(db, &currentDevice)
			if !ok {
				res.SetStatus(http.StatusBadRequest, models.StatusError, "Something went wrong 3")
				context.JSON(res.Generate())
				return
			}
			res.SetStatus(http.StatusOK, models.StatusOk, "update successfully")
			res.Set("points", currentDevice.Point)
			context.JSON(res.Generate())
			return
		} else {
			res.SetStatus(http.StatusBadRequest, models.StatusError, "update type not valid")
			context.JSON(res.Generate())
			return
		}

	}
}
func GetLeaderboard(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		res := models.NetResponse{}.Build()
		var leaderboard []models.Device
		leaderboard, err := models.GetLeaderboard(db)
		if err != nil {
			res.SetStatus(http.StatusBadRequest, models.StatusError, "Something went wrong")
			context.JSON(res.Generate())
			return
		}
		res.Set("data", leaderboard)
		res.SetStatus(http.StatusOK, models.StatusOk, "get leaderboard successfully")
		context.JSON(res.Generate())
	}
}
