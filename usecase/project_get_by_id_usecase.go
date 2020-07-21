package usecase

import (
	"fmt"

	"github.com/danilovalente/project-api/appcontext"
	"github.com/danilovalente/project-api/config"
	"github.com/danilovalente/project-api/domain"
)

//ProjectGetByID represents the Usecase which orchestrates the Project get from the database
type ProjectGetByID struct {
	projectRepository domain.ProjectRepository
}

//Execute get the Project with the provided ID
func (u *ProjectGetByID) Execute(ID string) (*domain.Project, error) {
	logger := config.GetLogger
	defer logger().Sync()

	projectRepository := u.projectRepository
	project, err := projectRepository.Get(ID)
	if err != nil {
		msg := fmt.Sprintf("Could not get the Project. Message: %s\n", err.Error())
		logger().Error(msg)
		return nil, err
	}
	return project, nil
}

func buildProjectGetByIDUsecase() appcontext.Component {

	return &ProjectGetByID{
		projectRepository: domain.GetProjectRepository(),
	}
}

func init() {
	if config.Values.TestRun {
		return
	}
	appcontext.Current.Add(appcontext.ProjectGetByIDUsecase, buildProjectGetByIDUsecase)
}
