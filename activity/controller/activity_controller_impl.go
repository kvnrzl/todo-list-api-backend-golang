package controller

import (
	"strconv"
	"technical_test_skyshi/activity/service"
	"technical_test_skyshi/helper"
	"technical_test_skyshi/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ActivityControllerImpl struct {
	activityService service.ActivityService
}

func NewActivityControllerImpl(activityService service.ActivityService) ActivityController {
	return &ActivityControllerImpl{
		activityService: activityService,
	}
}

func (a *ActivityControllerImpl) Create(c *gin.Context) {
	activity := &model.Activity{}
	c.ShouldBindJSON(&activity)

	activity, err := a.activityService.Create(c, activity)
	if err == model.ErrTitleRequired {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	if err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseStatusCreated(c, &model.ActivityCreate{
		CreatedAt: activity.CreatedAt,
		UpdatedAt: activity.UpdatedAt,
		ID:        activity.ID,
		Title:     activity.Title,
		Email:     activity.Email,
	})
}

func (a *ActivityControllerImpl) GetAll(c *gin.Context) {
	activities, err := a.activityService.GetAll(c)
	if err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseStatusOK(c, activities)

}

func (a *ActivityControllerImpl) GetByID(c *gin.Context) {
	id := c.Param("id")
	activityID, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	parseID := int64(activityID)
	activity, err := a.activityService.GetByID(c, parseID)
	// if err == model.ErrActivityNotFound {
	if err == gorm.ErrRecordNotFound {
		helper.ResponseActivityNotFound(c, parseID)
		return
	}

	if err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseStatusOK(c, activity)
}

func (a *ActivityControllerImpl) Update(c *gin.Context) {
	id := c.Param("id")
	activityID, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	activity := &model.Activity{}
	c.ShouldBindJSON(&activity)
	activity.ID = int64(activityID)

	activity, err = a.activityService.Update(c, activity)
	if err == gorm.ErrRecordNotFound {
		helper.ResponseActivityNotFound(c, int64(activityID))
		return
	}

	if err == model.ErrTitleRequired {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	if err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseStatusOK(c, activity)
}

func (a *ActivityControllerImpl) Delete(c *gin.Context) {
	id := c.Param("id")
	activityID, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	parseID := int64(activityID)
	err = a.activityService.Delete(c, parseID)
	// if err == model.ErrActivityNotFound {
	if err == gorm.ErrRecordNotFound {
		helper.ResponseActivityNotFound(c, parseID)
		return
	}

	if err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseStatusOK(c, struct{}{})
}
