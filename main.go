package main

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	TvAuthenticationService "github.com/nayanmakasare/TvAuthenticationService/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main()  {
	service := grpc.NewService(
		micro.Name("TvAuthenticationService"),
		micro.Address(":50058"),
		micro.Version("1.0"),
	)
	service.Init()
	authRedisCleint := GetRedisClient()
	go syncMongoAndRedis(authRedisCleint)
	handler := TvAuthenticationServiceHandler{AuthRedisConnection:authRedisCleint}
	err := TvAuthenticationService.RegisterTvAuthenticationServiceHandler(service.Server(), &handler)
	if err != nil {
		log.Fatal(err)
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func syncMongoAndRedis(authRedisClient *redis.Client){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI( "mongodb://192.168.1.9:27017"))
	if err != nil {
		log.Debug(err)
	}
	emacMongoCollection := mongoClient.Database("cwtx2devel").Collection("authwalls")
	myPipes := mongo.Pipeline{
		bson.D{{"$project", bson.D{{"_id", 0}, {"emac", 1}}}},
	}
	cur, err := emacMongoCollection.Aggregate(ctx, myPipes)
	if err != nil{
		log.Fatal("Failed to fetch mongo emac ", err)
		return
	}
	for cur.Next(context.TODO()){
		authRedisClient.SAdd("cloudwalkerEmac", cur.Current.Lookup("emac").StringValue())
	}
	cur.Close(context.TODO())
	return
}

func GetRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6378",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to redis %v", err)
	}
	return client
}