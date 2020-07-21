package usecase

import (
	"fmt"

	"github.com/danilovalente/project-api/appcontext"
	"github.com/danilovalente/project-api/config"
	"github.com/danilovalente/project-api/domain"
)

//ProjectDelete represents the Usecase which orchestrates the Project deletion from the database
type ProjectDelete struct {
	projectRepository domain.ProjectRepository
}

//Execute deletes the Project with the provided ID
func (u *ProjectDelete) Execute(ID string) error {
	logger := config.GetLogger
	defer logger().Sync()

	projectRepository := u.projectRepository
	err := projectRepository.Delete(ID)
	if err != nil {
		msg := fmt.Sprintf("Could not delete the Project with ID: %s. Message: %s\n", ID, err.Error())
		logger().Error(msg)
		return err
	}
	return nil
}

func buildProjectDeleteUsecase() appcontext.Component {

	return &ProjectDelete{
		projectRepository: domain.GetProjectRepository(),
	}
}

func init() {
	if config.Values.TestRun {
		return
	}
	appcontext.Current.Add(appcontext.ProjectDeleteUsecase, buildProjectDeleteUsecase)
}
