/*
 * Project
 *
 * This is the representation of the domain entity Project
 *
 */
package domain

import (
	"strings"
	"time"

	"github.com/danilovalente/project-api/appcontext"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Project represents the domain entity
type Project struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`

	Name string `bson:"name" json:"name"`

	UnitPrice Money `bson:"unitPrice" json:"unitPrice"`

	TimeUnit string `bson:"timeUnit" json:"timeUnit"`

	DateCreated time.Time `bson:"dateCreated,omitempty" json:"dateCreated,omitempty"`

	DateUpdated time.Time `bson:"dateUpdated,omitempty" json:"dateUpdated,omitempty"`
}

//Valid checks if the instance is in a valid state.
//If the state is not valid, returns an domain.IdentifiableError (warning: it is important to return a domain.IdentifiableError following the example) availabe in the model_error.go file
func (project *Project) Valid() (bool, error) {
	if project == nil {
		return false, ConstraintViolation("The Project is not instantiated")
	}
	if strings.TrimSpace(project.Name) == "" {
		return false, ConstraintViolation("The Project is invalid. The required attribute 'Name' is missing")
	}
	if project.UnitPrice.IsZero() || project.UnitPrice.IsNegative() || strings.TrimSpace(project.UnitPrice.Currency) == "" {
		return false, ConstraintViolation("The Project is invalid. The 'Unit Price' must be in a valid Currency and must be greater than zero")
	}
	if project.TimeUnit != "Hour" &&
		project.TimeUnit != "Day" &&
		project.TimeUnit != "Week" &&
		project.TimeUnit != "Month" {
		return false, ConstraintViolation("The Project is invalid. The 'TimeUnit' must be any of [Hour, Day, Week, Month]")
	}
	return true, nil
}

//ProjectRepository is the specification of the features delivered by a Repository for a Project
type ProjectRepository interface {
	appcontext.Component
	GetAll(lastProjectID string, pageSize int64) ([]*Project, error)
	Get(id string) (*Project, error)
	Save(project *Project) (*Project, error)
	Update(project *Project) (*Project, error)
	Delete(id string) error
}

type ProjectCreateUsecase interface {
	Execute(project *Project) (*Project, error)
}

type ProjectGetAllUsecase interface {
	Execute(lastProjectID string, pageSize int64) ([]*Project, error)
}

type ProjectGetByIDUsecase interface {
	Execute(ID string) (*Project, error)
}

type ProjectUpdateUsecase interface {
	Execute(project *Project) error
}

type ProjectDeleteUsecase interface {
	Execute(ID string) error
}

//GetProjectRepository gets the ProjectRepository current implementation
func GetProjectRepository() ProjectRepository {
	return appcontext.Current.Get(appcontext.ProjectRepository).(ProjectRepository)
}

//GetProjectCreateUsecase gets the ProjectCreateUsecase current implementation
func GetProjectCreateUsecase() ProjectCreateUsecase {
	return appcontext.Current.Get(appcontext.ProjectCreateUsecase).(ProjectCreateUsecase)
}

//GetProjectGetAllUsecase gets the ProjectGetAllUsecase current implementation
func GetProjectGetAllUsecase() ProjectGetAllUsecase {
	return appcontext.Current.Get(appcontext.ProjectGetAllUsecase).(ProjectGetAllUsecase)
}

//GetProjectGetByIDUsecase gets the ProjectGetByIDUsecase current implementation
func GetProjectGetByIDUsecase() ProjectGetByIDUsecase {
	return appcontext.Current.Get(appcontext.ProjectGetByIDUsecase).(ProjectGetByIDUsecase)
}

//GetProjectUpdateUsecase gets the ProjectUpdateUsecase current implementation
func GetProjectUpdateUsecase() ProjectUpdateUsecase {
	return appcontext.Current.Get(appcontext.ProjectUpdateUsecase).(ProjectUpdateUsecase)
}

//GetProjectDeleteUsecase gets the ProjectDeleteUsecase current implementation
func GetProjectDeleteUsecase() ProjectDeleteUsecase {
	return appcontext.Current.Get(appcontext.ProjectDeleteUsecase).(ProjectDeleteUsecase)
}
