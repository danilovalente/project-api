package usecase

import (
	"github.com/danilovalente/project-api/appcontext"
	"github.com/danilovalente/project-api/config"
	"github.com/danilovalente/project-api/domain"
)

//ProjectCreate represents the Usecase which orchestrates the Project creation in the database
type ProjectCreate struct {
	projectRepository domain.ProjectRepository
}

//Execute creates/persists the project
func (u *ProjectCreate) Execute(project *domain.Project) (*domain.Project, error) {
	logger := config.GetLogger
	defer logger().Sync()
	logger().Debugf("Project %+v \n", project)

	valid, err := project.Valid()
	if !valid {
		logger().Error(err.Error())
		return nil, err
	}
	project, err = u.projectRepository.Save(project)
	if err != nil {
		logger().Errorf("Could not save project into repository. Error %s", err.Error())
		return nil, err
	}
	return project, nil
}

func buildProjectCreateUsecase() appcontext.Component {
	return &ProjectCreate{
		projectRepository: domain.GetProjectRepository(),
	}
}

func init() {
	if config.Values.TestRun {
		return
	}
	appcontext.Current.Add(appcontext.ProjectCreateUsecase, buildProjectCreateUsecase)
}
