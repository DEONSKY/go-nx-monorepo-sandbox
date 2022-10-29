package repository

import (
	"go-sandbox/libs/config"
	"go-sandbox/libs/model"
)

func InsertIssueComment(issueComment model.IssueComment) (*model.IssueComment, error) {
	if result := config.DB.Save(&issueComment); result.Error != nil {
		return nil, result.Error
	}
	//config.DB.Preload("User").Find(&issue)
	return &issueComment, nil
}
