package app

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	myproto "github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/config"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/repository"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/repository/impl"
	"github.com/Kamva/mgm/v3"
	"github.com/go-bongo/bongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/option"
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

func NewAppConfig(firebaseApp firebase.App, bongoClient *bongo.Connection, mgoClient *mgo.Session, sagaMessageClient myproto.MessageServiceClient) *config.AppConfig {
	return &config.AppConfig{FirebaseApp: firebaseApp, BongoClient: bongoClient, MgoClient: mgoClient, MessageServiceClient: sagaMessageClient}
}

func getFirebaseServiceAccount() firebase.App {
	currentPath, _ := os.Getwd()
	opt := option.WithCredentialsFile(fmt.Sprintf("%s/firebase-key/passnet-auth-firebase-admin.json", currentPath))
	firebaseConfig := &firebase.Config{ProjectID: "passnet-auth"}
	app, err := firebase.NewApp(context.Background(), firebaseConfig, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return *app
}
func bongoClient() *bongo.Connection {
	config := &bongo.Config{
		ConnectionString: "localhost",
		Database:         "passnet-auth",
	}
	connection, err := bongo.Connect(config)

	if err != nil {
		log.Fatal(err)
	}

	return connection
}

func mgoClient() *mgo.Session {

	dbHost := os.Getenv("DB_HOST")
	session, err := mgo.Dial(dbHost)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	return session
}

func SetMgmClient() {
	dbURI := fmt.Sprintf("mongodb://%s:%s@%s:27017", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"))
	dbName := os.Getenv("DB_NAME")
	err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatalf("error when connecting to database --> %s", err.Error())
	}
}

func GetUserMgmRepository() repository.IUserRepository {
	return impl.NewUserMgmRepository()
}
var appConfigIntance = NewAppConfig(getFirebaseServiceAccount(), bongoClient(), nil, GetSagaMessageGrpcClient())
func GetAppConfigInstance() *config.AppConfig {
	return appConfigIntance
}

