package product

import (
	"context"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBConn() (context.Context, *mongo.Client, string) {
	connectionURI := "mongodb://OZEadmin:MonGoOZE@127.0.0.1:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	dbName := "OZEProject"
	return ctx, client, dbName

} // end func
