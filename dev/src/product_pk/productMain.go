package product

import (
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (p AllProduct) GetNewProductLast10(c *gin.Context) (RetNow []AllProduct) {

	ctx, client, dbName := DBConn()
	defer client.Disconnect(ctx)

	coll := client.Database(dbName).Collection("tbl_Product")
	filter := bson.M{}
	opts := options.Find().SetSort(bson.D{{"DateIn", -1}}).SetLimit(10)
	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		panic(err)
	}

	defer cursor.Close(ctx)
	var result AllProduct
	var results []AllProduct
	var countnow int = 1
	for cursor.Next(ctx) {
		//var result Supplyier
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		/*
			if result.SumTotal == "" {
				result.SumTotal = "0"

			}
		*/
		result.CountNow = countnow
		countnow++
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return results

} // ebd func
