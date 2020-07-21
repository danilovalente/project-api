package usecase

import (
	"fmt"
	"github.com/danilovalente/project-api/appcontext"
	"github.com/danilovalente/project-api/config"
	"github.com/danilovalente/project-api/domain"
)

//ProjectGetAll represents the Usecase which orchestrates the Project creation in the database
type ProjectGetAll struct {
	projectRepository domain.ProjectRepository
}

//Execute with paging
func (u *ProjectGetAll) Execute(lastProjectID string, pageSize int64) ([]*domain.Project, error) {
	logger := config.GetLogger
	defer logger().Sync()

	projectList, err := u.projectRepository.GetAll(lastProjectID, pageSize)
	if err != nil {
		msg := fmt.Sprintf("Could not get the Project list. Message: %s\n", err.Error())
		logger().Error(msg)
		return nil, err
	}
	return projectList, nil
}

func buildProjectGetAllUsecase() appcontext.Component {
	return &ProjectGetAll{
		projectRepository: domain.GetProjectRepository(),
	}
}

func init() {
	if config.Values.TestRun {
		return
	}
	appcontext.Current.Add(appcontext.ProjectGetAllUsecase, buildProjectGetAllUsecase)
}
