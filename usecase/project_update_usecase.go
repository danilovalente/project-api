package usecase

import (
	"github.com/danilovalente/project-api/appcontext"
	"github.com/danilovalente/project-api/config"
	"github.com/danilovalente/project-api/domain"
)

//ProjectUpdate represents the Usecase which orchestrates the Project update in the database
type ProjectUpdate struct {
	projectRepository domain.ProjectRepository
}

//Execute updates the project
func (u *ProjectUpdate) Execute(project *domain.Project) error {
	logger := config.GetLogger
	defer logger().Sync()
	logger().Debugf("Project %+v \n", project)

	valid, err := project.Valid()
	if !valid {
		logger().Error(err.Error())
		return err
	}
	_, err = u.projectRepository.Update(project)
	if err != nil {
		logger().Errorf("Could not update project into repository. Error %s", err.Error())
		return err
	}
	return nil
}

func buildProjectUpdateUsecase() appcontext.Component {
	return &ProjectUpdate{
		projectRepository: domain.GetProjectRepository(),
	}
}

func init() {
	if config.Values.TestRun {
		return
	}
	appcontext.Current.Add(appcontext.ProjectUpdateUsecase, buildProjectUpdateUsecase)
}
