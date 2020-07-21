package mongodb

import (
	"context"
	"fmt"
	"time"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/danilovalente/project-api/appcontext"
	"github.com/danilovalente/project-api/config"
	"github.com/danilovalente/project-api/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CollectionName in MongoDB
const projectCollectionName = "project"

//ProjectRepository is the specification of the features delivered by a Repository for a Project
type ProjectRepository struct {
	Conn *mongo.Client
}

//Get a Project by ID
func (repo *ProjectRepository) Get(id string) (*domain.Project, error) {
	collection := repo.Conn.Database(DatabaseName).Collection(projectCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	projectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.ConstraintViolation(fmt.Sprintf("Invalid Project ID format: %s . Message: %s", id, err.Error()))
	}
	filter := bson.M{"_id": projectID}
	var project = domain.Project{}
	err = collection.FindOne(ctx, filter).Decode(&project)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return nil, domain.NotFound(fmt.Sprintf("Could not find Project with the ID: %s . Message: %s", id, err.Error()))
	}
	if err != nil {
		return nil, domain.InternalError(fmt.Sprintf("Database fetch error while Getting the Project for ID: %s - Message: %s", id, err.Error()))
	}
	return &project, nil
}

//Save a new project in the collection
func (repo *ProjectRepository) Save(project *domain.Project) (*domain.Project, error) {
	collection := repo.Conn.Database(DatabaseName).Collection(projectCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger := config.GetLogger
	defer logger().Sync()

	logger().Debugf("\n\n\n Before saving to respository %+v \n", project)

	if primitive.NilObjectID != project.ID {
		return nil, domain.InternalError("The Save method should not be used for updating. Please use Update instead")
	}
	project.ID = primitive.NewObjectID()
	project.DateCreated = time.Now()

	_, err := collection.InsertOne(ctx, project)
	if err != nil {
		return nil, domain.InternalError(fmt.Sprintf("Could not create the project. project: %+v - Message: %s", project, err.Error()))
	}
	return project, nil

}

//Update a project in the collection
func (repo *ProjectRepository) Update(project *domain.Project) (*domain.Project, error) {
	collection := repo.Conn.Database(DatabaseName).Collection(projectCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"_id": project.ID}
	existentProject, err := repo.Get(project.ID.Hex())
	if err != nil {
		return nil, err
	}

	project.DateCreated = existentProject.DateCreated
	project.DateUpdated = time.Now()
	_, err = collection.ReplaceOne(ctx, filter, project)
	if err != nil {
		return nil, domain.InternalError(fmt.Sprintf("Could not update the project with ID = %s - Message: %s", project.ID.Hex(), err.Error()))
	}
	return project, nil

}

//GetAll Project
func (repo *ProjectRepository) GetAll(lastProjectID string, pageSize int64) ([]*domain.Project, error) {
	projectList := make([]*domain.Project, 0)
	collection := repo.Conn.Database(DatabaseName).Collection(projectCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var dbfilter interface{}
	if strings.TrimSpace(lastProjectID) == "" {
		dbfilter = bson.D{}
	} else {
		lastProject, err := primitive.ObjectIDFromHex(lastProjectID)
		if err != nil {
			return nil, domain.ConstraintViolation(fmt.Sprintf("Invalid project Id: %s. Message: %s", lastProjectID, err.Error()))
		}
		dbfilter = bson.M{"_id": bson.M{"$gt": lastProject}}
	}
	opts := &options.FindOptions{}
	opts.SetSort(bson.M{"_id": 1})
	opts.SetLimit(pageSize)
	cur, err := collection.Find(ctx, dbfilter, opts)
	defer func() { _ = cur.Close(ctx) }()
	if err != nil {
		return nil, domain.InternalError(fmt.Sprintf("An error occurred while trying to find the project List. Message: %s", err.Error()))
	}
	for cur.Next(ctx) {
		var result domain.Project
		err := cur.Decode(&result)
		if err != nil {
			return nil, domain.InternalError(fmt.Sprintf("An error occured while trying to convert the Project from the database. Message: %s", err.Error()))
		}
		projectList = append(projectList, &result)
	}
	if err := cur.Err(); err != nil {
		return nil, domain.InternalError(fmt.Sprintf("An error occured while trying to convert the list of Project from the database. Message: %s", err.Error()))
	}
	return projectList, nil
}

//Delete a ProjectRepository by ID
func (repo *ProjectRepository) Delete(id string) error {
	collection := repo.Conn.Database(DatabaseName).Collection(projectCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	projectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ConstraintViolation(fmt.Sprintf("Invalid Project ID format: %s . Message: %s", id, err.Error()))
	}
	filter := bson.M{"_id": projectID}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return domain.InternalError(fmt.Sprintf("Database error while deleting the Project with ID: %s - Message: %s", id, err.Error()))
	}
	if result.DeletedCount != 1 {
		return domain.NotFound(fmt.Sprintf("Could not find Project with the ID: %s", id))
	}
	return nil
}

func buildProjectRepository() appcontext.Component {
	dbClient := appcontext.Current.Get(appcontext.DBClient).(*MongoClient)
	return &ProjectRepository{Conn: dbClient.Conn}
}

func init() {
	if config.Values.TestRun {
		return
	}

	appcontext.Current.Add(appcontext.ProjectRepository, buildProjectRepository)
}
