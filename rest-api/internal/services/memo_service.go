package services

import (
	"errors"

	"memo-point-com/internal/database"
	"memo-point-com/internal/models"
)

type MemoService struct{}

func (service *MemoService) CreateMemo(userID int, title, content string) error {
	memo := models.Memo{
		UserID:  userID,
		Title:   title,
		Content: content,
	}

	err := database.DbInstance.Create(&memo).Error
	if err != nil {
		return err
	}

	return nil
}

func (service *MemoService) GetMemoByID(memoID int) (*models.Memo, error) {
	var memo models.Memo

	err := database.DbInstance.Where("id = ?", memoID).First(&memo).Error
	if err != nil {
		return nil, err
	}

	return &memo, nil
}

func (service *MemoService) GetMemosByUserID(userID int) ([]models.Memo, error) {
	var memos []models.Memo

	err := database.DbInstance.Where("user_id = ?", userID).Find(&memos).Error
	if err != nil {
		return nil, err
	}

	return memos, nil
}

func (service *MemoService) UpdateMemoByID(userID, memoID int, title, content string) error {
	memo, err := service.GetMemoByID(memoID)
	if err != nil {
		return err
	}

	if memo.UserID != userID {
		return errors.New("unauthorized")
	}

	memo.Title = title
	memo.Content = content

	err = database.DbInstance.Save(&memo).Error
	if err != nil {
		return err
	}

	return nil
}

func (service *MemoService) DeleteMemoByID(userID, memoID int) error {
	memo, err := service.GetMemoByID(memoID)
	if err != nil {
		return err
	}

	if memo.UserID != userID {
		return errors.New("unauthorized")
	}

	err = database.DbInstance.Delete(&memo).Error
	if err != nil {
		return err
	}

	return nil
}
