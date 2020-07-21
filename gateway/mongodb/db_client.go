package mongodb

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/danilovalente/project-api/appcontext"
	"github.com/danilovalente/project-api/config"
)

//DatabaseName in MongoDB
const DatabaseName = "project"

//MongoClient is a struct to keep the DB connection and URI
type MongoClient struct {
	DBURI string
	Conn  *mongo.Client
}

func getCustomTLSConfig(caFile string) (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	certs, err := ioutil.ReadFile(caFile)

	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, errors.New("Failed parsing pem file")
	}

	return tlsConfig, nil
}

//buildMongoClientWithTLS creates the MongoClient instance with TLS encryption
func buildMongoClientWithTLS() appcontext.Component {
	var mongoClient = MongoClient{}

	caFilePath := config.Values.DBConnectionCertificateFileName

	tlsConfig, err := getCustomTLSConfig(caFilePath)
	if err != nil {
		log.Fatalf("Failed getting TLS configuration: %v", err)
	}

	mongoClient.DBURI = config.Values.DBConnectionString
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoClient.DBURI).SetTLSConfig(tlsConfig))
	if err != nil {
		log.Fatal("An error occurred while trying to open a DB connection.  Error message: " + err.Error())
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("An error occurred while trying to open a DB connection. Error message: " + err.Error())
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping cluster: %v", err)
	}
	mongoClient.Conn = client
	return &mongoClient
}

//buildMongoClient creates the MongoClient instance
func buildMongoClient() appcontext.Component {
	var mongoClient = MongoClient{}

	mongoClient.DBURI = config.Values.DBConnectionString
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoClient.DBURI))
	if err != nil {
		log.Fatal("An error occurred while trying to open a DB connection.  Error message: " + err.Error())
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("An error occurred while trying to open a DB connection. Error message: " + err.Error())
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping cluster: %v", err)
	}
	mongoClient.Conn = client
	return &mongoClient
}

//StartTransactionalSession starts a Database Session with an open Transaction
func (client *MongoClient) StartTransactionalSession() (mongo.Session, error) {
	var session mongo.Session
	var err error
	if session, err = client.Conn.StartSession(); err != nil {
		return nil, errors.New("An error occurred while trying to create a DB session. Message: " + err.Error())
	}
	if err = session.StartTransaction(); err != nil {
		return nil, errors.New("An error occurred while trying to start a DB transaction. Message: " + err.Error())
	}

	return session, nil
}

//CollectionExists in the Database?
func CollectionExists(db *mongo.Database, collectionName string) (bool, error) {
	ctx := context.Background()
	options := options.ListCollectionsOptions{}
	options.SetNameOnly(true)
	cur, err := db.ListCollections(ctx, bson.M{}, &options)
	if err != nil {
		return false, errors.New("Failed to get collection names: " + err.Error())
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result string
		err := cur.Decode(&result)
		if err != nil {
			return false, err
		}
		if result == collectionName {
			return true, nil
		}
	}
	return false, nil
}

func init() {
	if config.Values.TestRun {
		return
	}

	if config.Values.DBConnectionCertificateFileName == "" {
		appcontext.Current.Add(appcontext.DBClient, buildMongoClient)
	} else {
		appcontext.Current.Add(appcontext.DBClient, buildMongoClientWithTLS)
	}
}
