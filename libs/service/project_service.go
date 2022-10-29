package service

import (
	"go-sandbox/libs/dto/request"
	"go-sandbox/libs/dto/response"
	"go-sandbox/libs/model"
	"go-sandbox/libs/repository"
	"go-sandbox/libs/utils"

	"github.com/mashingan/smapping"
)

func CreateProject(projectCreateDTO request.ProjectCreateRequest) (*model.Project, error) {
	projectToCreate := model.Project{}
	err := smapping.FillStruct(&projectToCreate, smapping.MapFields(&projectCreateDTO))
	if err != nil {
		return nil, utils.ReturnErrorResponse(400, "Request DTO Parse Problem", []string{err.Error()})
	}
	res, err := repository.InsertProject(projectToCreate)
	if err != nil {
		return nil, utils.ReturnErrorResponse(422, "Project could not be inserted", []string{err.Error()})
	}
	return res, err
}

func GetProjectsByUserId(userID uint64) ([]response.ProjectNavTreeResponse, error) {

	subjectRes, subjectErr := repository.GetSubjectsByUserId(userID)
	if subjectErr != nil {
		return nil, utils.ReturnErrorResponse(400, "Cannot get subjects by user id", []string{subjectErr.Error()})
	}

	subjectNavTreeMap := make(map[uint64][]response.SubjectNavTreeResponse)
	var projectIdSlice []uint64
	for _, subject := range subjectRes {
		if subjectNavTreeMap[subject.ProjectID] == nil {
			projectIdSlice = append(projectIdSlice, subject.ProjectID)
		}
		subjectNavTreeMap[subject.ProjectID] = append(subjectNavTreeMap[subject.ProjectID], subject)
	}

	res, err := repository.GetProjectsBySubjectIds(projectIdSlice)
	if err != nil {
		return nil, utils.ReturnErrorResponse(400, "Cannot get projects by project ids", []string{subjectErr.Error()})
	}

	for i := range res {
		res[i].Subjects = subjectNavTreeMap[res[i].ID]
	}
	return res, nil
}
