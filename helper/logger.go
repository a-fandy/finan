package helper

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func LogErrorToMongoDB(req *fiber.Ctx, err error, mongoClient *mongo.Client) {
	// Log error details to MongoDB
	if mongoClient != nil {
		collection := mongoClient.Database("finan").Collection("logs")
		logEntry := map[string]interface{}{
			"timestamp": time.Now(),
			"error":     err.Error(),
			"request": map[string]interface{}{
				"method":    req.Method(),
				"path":      req.Path(),
				"headers":   req.GetReqHeaders(),
				"body":      req.Request().Body(),
				"query":     req.Queries(),
				"params":    req.Params("*"),
				"remote_ip": req.IP(),
			},
			"response": map[string]interface{}{
				"status":   req.Response().StatusCode(),
				"headers":  req.GetRespHeaders(),
				"body":     req.Response().Body(),
				"duration": time.Since(req.Locals("start_time").(time.Time)),
			},
		}

		_, err := collection.InsertOne(req.Context(), logEntry)
		if err != nil {
			log.Error("Error disconnecting from MongoDB")
		}
	}
}
