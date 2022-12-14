package service

import (
	"go-sandbox/libs/dto/request"
	"go-sandbox/libs/model"
	"go-sandbox/libs/repository"
	"go-sandbox/libs/utils"

	"github.com/mashingan/smapping"
)

func AddIssueComment(issueCommentDto request.IssueCommentCreateRequest) (*model.IssueComment, error) {
	issueCommentToCreate := model.IssueComment{}
	err := smapping.FillStruct(&issueCommentToCreate, smapping.MapFields(&issueCommentDto))
	if err != nil {
		return nil, utils.ReturnErrorResponse(400, "Request DTO Parse Problem", []string{err.Error()})
	}
	res, err := repository.InsertIssueComment(issueCommentToCreate)
	if err != nil {
		return nil, utils.ReturnErrorResponse(422, "Issue comment could not be inserted", []string{err.Error()})
	}
	return res, err
}
