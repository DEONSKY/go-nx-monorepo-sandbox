package handler

import (
	"log"
	"strconv"

	"go-sandbox/libs/dto/request"
	"go-sandbox/libs/helper"
	"go-sandbox/libs/service"
	"go-sandbox/libs/utils"

	"github.com/gofiber/fiber/v2"
)

// InserIssue is a function to create new Issue
// @Summary Create new Issue
// @Description Creates new issue
// @Tags Issues
// @Accept json
// @Produce json
// @Param Issue body request.IssueCreateRequest true "createIssues"
// @Success 200 {object} helper.Response{data=model.Issue}
// @Failure 400 {object} helper.Response{data=helper.EmptyObj}
// @Security ApiKeyAuth
// @Router /api/issue [post]
type IssueHandler interface {
	InsertIssue(context *fiber.Ctx) error
	GetIssues(context *fiber.Ctx) error
	GetIssuesKanban(context *fiber.Ctx) error
	InsertDependentIssueAssociation(context *fiber.Ctx) error
	AssignieIssueToUser(context *fiber.Ctx) error
}

type issueHandler struct {
	issueService service.IssueService
}

//NewBookController create a new instances of BoookController
func NewIssueHandler(issueSer service.IssueService) IssueHandler {
	return &issueHandler{
		issueService: issueSer,
	}
}

func (c *issueHandler) InsertIssue(context *fiber.Ctx) error {
	userID := context.Locals("user_id").(uint64)
	var IssueCreateDTO request.IssueCreateRequest

	errDTO := context.BodyParser(&IssueCreateDTO)
	if errDTO != nil {
		return utils.ReturnErrorResponse(fiber.StatusInternalServerError, "Request DTO Parse Problem", []string{errDTO.Error()})
	}

	errors := utils.ValidateStruct(IssueCreateDTO)
	if errors != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, "Validation error", errors)
	}

	IssueCreateDTO.ReporterID = userID

	result, err := c.issueService.CreateIssue(IssueCreateDTO)
	if err != nil {
		return err
	}
	response := helper.BuildResponse("Issue created succesfully", result)
	return context.Status(fiber.StatusCreated).JSON(response)

}

// GetIssues is a function to get all issues data from database with dynamic query parameters
// @Summary Get all issues with query parameters
// @Description GetIssues is a function to get all issues data from database with dynamic query parameters
// @Tags Issues
// @Accept json
// @Produce json
// @Param Issue query request.IssueGetQuery true "getIssues"
// @Success 200 {object} helper.Response{data=[]response.IssueResponse}
// @Failure 400 {object} helper.Response{data=[]helper.EmptyObj}
// @Security ApiKeyAuth
// @Router /api/issue [get]
func (c *issueHandler) GetIssues(context *fiber.Ctx) error {
	iq := new(request.IssueGetQuery)
	userID := context.Locals("user_id").(uint64)

	if err := context.QueryParser(iq); err != nil {
		return utils.ReturnErrorResponse(fiber.StatusInternalServerError, "Request Query parse problem", []string{err.Error()})
	}
	log.Println("Params", iq)
	if iq.GetOnlyOrphans != nil && iq.ParentIssueID != nil {
		return utils.ReturnErrorResponse(fiber.StatusInternalServerError, "Request Error", []string{"An issue cannot be orphan and has parent at the same time"})
	}
	result, err := c.issueService.GetIssues(iq, userID)
	if err != nil {
		return err
	}
	response := helper.BuildShortResponse(result)
	return context.Status(fiber.StatusOK).JSON(response)
}

// GetIssuesKanban is a function to get all issues data from database with dynamic query parameters as Kanban format
// @Summary Get all issues as Kanban Format with query parameters
// @Description Get all issues as Kanban Format with query parameters
// @Tags Issues
// @Accept json
// @Produce json
// @Param Issue query request.IssueGetQuery true "getIssues"
// @Success 200 {object} helper.Response{data=[]response.IssueKanbanResponse}
// @Failure 400 {object} helper.Response{data=[]helper.EmptyObj}
// @Security ApiKeyAuth
// @Router /api/issue/kanban [get]
func (c *issueHandler) GetIssuesKanban(context *fiber.Ctx) error {
	iq := new(request.IssueGetQuery)

	userID := context.Locals("user_id").(uint64)

	if err := context.QueryParser(iq); err != nil {
		return utils.ReturnErrorResponse(fiber.StatusInternalServerError, "Request Query parse problem", []string{err.Error()})
	}
	log.Println(iq)
	if iq.GetOnlyOrphans != nil && iq.ParentIssueID != nil {
		return utils.ReturnErrorResponse(fiber.StatusInternalServerError, "Request Error", []string{"An issue cannot be orphan and has parent at the same time"})
	}
	result, err := c.issueService.GetIssuesKanban(iq, userID)
	if err != nil {
		return err
	}
	response := helper.BuildShortResponse(result)
	return context.Status(fiber.StatusOK).JSON(response)
}

// InsertDependentIssueAssociation adds assocation between issue and dependent issue
// @Summary Adds assocation with issue and dependent issue
// @Description Adds assocation with issue and dependent issue
// @Tags Issues
// @Accept json
// @Produce json
// @Param issue_id path string true "Issue ID"
// @Param dependent_issue_id path string true "Dependent Issue ID"
// @Success 200 {object} helper.Response{data=model.Issue}
// @Failure 400 {object} helper.Response{data=[]helper.EmptyObj}
// @Security ApiKeyAuth
// @Router /add-issue-dependency/{issue_id}/{dependent_issue_id} [put]
func (c *issueHandler) InsertDependentIssueAssociation(context *fiber.Ctx) error {
	issueID, err := strconv.ParseUint(context.Params("issue_id"), 10, 64)
	userID := context.Locals("user_id").(uint64)

	log.Println(issueID)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, "Wrong issue parameter", []string{err.Error()})
	}
	dependentIssueID, err := strconv.ParseUint(context.Params("dependent_issue_id"), 10, 64)
	log.Println(dependentIssueID)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, "Wrong dependent issue parameter", []string{err.Error()})
	}

	result, err := c.issueService.InsertDependentIssueAssociation(issueID, dependentIssueID, userID)
	if err != nil {
		return err
	}

	response := helper.BuildResponse("OK", result)
	return context.Status(fiber.StatusCreated).JSON(response)
}

// AssignieIssueToUser adds assocation between issue and assigned user
// @Summary Adds assocation between issue and assigned user
// @Description Adds assocation with issue and dependent issue
// @Tags Issues
// @Accept json
// @Produce json
// @Param issue_id path string true "Issue ID"
// @Param user_id path string true "Assignie User ID"
// @Success 200 {object} helper.Response{data=model.Issue}
// @Failure 400 {object} helper.Response{data=[]helper.EmptyObj}
// @Security ApiKeyAuth
// @Router /assignie-user/{issue_id}/{user_id} [put]
func (c *issueHandler) AssignieIssueToUser(context *fiber.Ctx) error {
	issueID, err := strconv.ParseUint(context.Params("issue_id"), 10, 64)
	userID := context.Locals("user_id").(uint64)

	log.Println(issueID)
	if err != nil {
		res := helper.BuildErrorResponse("Wrong Issue Parameter", err.Error(), helper.EmptyObj{})
		return context.Status(fiber.StatusBadRequest).JSON(res)
	}
	assignieID, err := strconv.ParseUint(context.Params("user_id"), 10, 64)
	log.Println(assignieID)
	if err != nil {
		res := helper.BuildErrorResponse("Wrong User Parameter", err.Error(), helper.EmptyObj{})
		return context.Status(fiber.StatusBadRequest).JSON(res)
	}
	log.Println("here")
	result, err := c.issueService.AssignieIssueToUser(issueID, assignieID, userID)
	if err != nil {
		return err
	}

	response := helper.BuildResponse("OK", result)
	return context.Status(fiber.StatusCreated).JSON(response)
}
