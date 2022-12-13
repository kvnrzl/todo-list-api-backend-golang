package service

import (
	"context"
	"technical_test_skyshi/activity/repository"
	"technical_test_skyshi/model"

	"gorm.io/gorm"
)

type ActivityServiceImpl struct {
	activityRepository repository.ActivityRepository
	db                 *gorm.DB
}

func NewActivityServiceImpl(activityRepository repository.ActivityRepository, db *gorm.DB) ActivityService {
	return &ActivityServiceImpl{
		activityRepository: activityRepository,
		db:                 db,
	}
}

func (a *ActivityServiceImpl) Create(ctx context.Context, activity *model.Activity) (*model.Activity, error) {
	if activity.Title == "" {
		return nil, model.ErrTitleRequired
	}

	if activity.Email == "" {
		activity.Email = "no email"
	}

	return a.activityRepository.Create(ctx, a.db, activity)
}

func (a *ActivityServiceImpl) GetAll(ctx context.Context) ([]*model.Activity, error) {
	return a.activityRepository.GetAll(ctx, a.db)
}

func (a *ActivityServiceImpl) GetByID(ctx context.Context, id int64) (*model.Activity, error) {
	return a.activityRepository.GetByID(ctx, a.db, id)
}

func (a *ActivityServiceImpl) Update(ctx context.Context, activity *model.Activity) (*model.Activity, error) {
	if activity.Title == "" {
		return nil, model.ErrTitleRequired
	}

	_, err := a.GetByID(ctx, activity.ID)
	if err != nil {
		return nil, err
	}

	activity, err = a.activityRepository.Update(ctx, a.db, activity)
	if err != nil {
		return nil, err
	}

	return a.GetByID(ctx, activity.ID)
}

func (a *ActivityServiceImpl) Delete(ctx context.Context, id int64) error {
	_, err := a.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return a.activityRepository.Delete(ctx, a.db, id)
}
