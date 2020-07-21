package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/danilovalente/project-api/config"
	"github.com/danilovalente/project-api/domain"
)

const (
	DefaultProjectPageSize = 20
)

//CreateProject creates a new Project
func CreateProject(c echo.Context) error {
	logger := config.GetLogger
	defer logger().Sync()

	project := new(domain.Project)
	if err := c.Bind(project); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ConstraintViolation("An error occurred while trying to read the request body: "+err.Error()))
	}

	project, err := domain.GetProjectCreateUsecase().Execute(project)

	if err != nil {
		logger().Errorf("An error occurred while trying to Create the Project: %s", err.Error())
		return c.JSON(err.(domain.IdentifiableError).GetCode(), err)
	}

	return c.JSON(http.StatusCreated, project)
}

//GetProjectList of the collection
func GetProjectList(c echo.Context) error {
	logger := config.GetLogger
	defer logger().Sync()
	var lastProjectID = c.QueryParam("lastProjectId")
	var pageSizeString = c.QueryParam("pageSize")

	var pageSize int
	if pageSizeString == "" {
		pageSize = DefaultProjectPageSize
	} else {
		var err error
		pageSize, err = strconv.Atoi(pageSizeString)
		if err != nil {
			msg := fmt.Sprintf("Invalid format for pageSize %s. Message: %s", pageSizeString, err.Error())
			logger().Error(msg)
			return c.JSON(http.StatusBadRequest, domain.ConstraintViolation(msg))
		}
	}

	projectList, err := domain.GetProjectGetAllUsecase().Execute(lastProjectID, int64(pageSize))
	if err != nil {
		logger().Errorf("An error occurred while trying to Get the Project List: %s", err.Error())
		return c.JSON(err.(domain.IdentifiableError).GetCode(), err)
	}

	return c.JSON(http.StatusOK, projectList)
}

//GetProject provided the projectId
func GetProject(c echo.Context) error {
	logger := config.GetLogger
	defer logger().Sync()
	projectID := strings.TrimSpace(c.Param("projectId"))

	if projectID == "" {
		return c.JSON(http.StatusBadRequest, domain.ConstraintViolation("Bad request. Missing mandatory request value projectId"))
	}

	project, err := domain.GetProjectGetByIDUsecase().Execute(projectID)
	if err != nil {
		logger().Errorf("An error occurred while trying to Get the Project: %s", err.Error())
		return c.JSON(err.(domain.IdentifiableError).GetCode(), err)
	}

	return c.JSON(http.StatusOK, project)
}

//UpdateProject updates the Project
func UpdateProject(c echo.Context) error {
	logger := config.GetLogger
	defer logger().Sync()
	projectID := strings.TrimSpace(c.Param("projectId"))
	if projectID == "" {
		return c.JSON(http.StatusBadRequest, domain.ConstraintViolation("Bad request. Missing mandatory request value projectId"))
	}

	project := domain.Project{}
	if err := c.Bind(&project); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ConstraintViolation("An error occurred while trying to read the request body: "+err.Error()))
	}

	if projectID != project.ID.Hex() {
		return c.JSON(http.StatusBadRequest, domain.ConstraintViolation("The content of the URL Path Parameter projectId is different of the Body's id"))
	}

	err := domain.GetProjectUpdateUsecase().Execute(&project)
	if err != nil {
		logger().Errorf("An error occurred while trying to Update the Project: %s", err.Error())
		return c.JSON(err.(domain.IdentifiableError).GetCode(), err)
	}

	return c.JSON(http.StatusOK, "")

}

//DeleteProject provided the projectId
func DeleteProject(c echo.Context) error {
	logger := config.GetLogger
	defer logger().Sync()
	projectID := strings.TrimSpace(c.Param("projectId"))

	if projectID == "" {
		return c.JSON(http.StatusBadRequest, domain.ConstraintViolation("Bad request. Missing mandatory request value projectId"))
	}

	err := domain.GetProjectDeleteUsecase().Execute(projectID)
	if err != nil {
		logger().Errorf("An error occurred while trying to Delete the Project: %s", err.Error())
		return c.JSON(err.(domain.IdentifiableError).GetCode(), err)
	}

	return c.JSON(http.StatusOK, "")
}
