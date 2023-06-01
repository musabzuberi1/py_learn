// myapp/config/config.go
package config

type MongoDBConfig struct {
	ConnectionString string
	Database         string
}

func GetMongoDBConfig() *MongoDBConfig {
	return &MongoDBConfig{
		ConnectionString: "mongodb://localhost:27017",
		Database:         "myapp",
	}
}

func ConnectMongoDB(cfg *config.MongoDBConfig) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(cfg.ConnectionString)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}