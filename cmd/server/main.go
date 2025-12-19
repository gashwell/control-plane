package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"version": "0.1.0",
		})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/models", listModels)
		api.POST("/api-keys", createAPIKey)
		api.GET("/analytics", getAnalytics)
		api.GET("/nginx/status", getNginxStatus)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("🚀 Control plane starting on :%s", port)
	log.Printf("📡 API available at: http://localhost:%s/api/v1", port)
	r.Run(":" + port)
}

func listModels(c *gin.Context) {
	models := []map[string]interface{}{
		{
			"id":               "gpt-4",
			"name":             "GPT-4",
			"provider":         "OpenAI",
			"cost_per_1k":      0.03,
			"max_tokens":       8192,
		},
		{
			"id":               "gpt-3.5-turbo",
			"name":             "GPT-3.5 Turbo",
			"provider":         "OpenAI",
			"cost_per_1k":      0.002,
			"max_tokens":       4096,
		},
		{
			"id":               "claude-3-sonnet",
			"name":             "Claude 3 Sonnet",
			"provider":         "Anthropic",
			"cost_per_1k":      0.015,
			"max_tokens":       200000,
		},
	}

	c.JSON(200, gin.H{"models": models})
}

func createAPIKey(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
		Tier string `json:"tier"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// TODO: Generate real API key and store in database
	c.JSON(201, gin.H{
		"api_key":    "sk-test-" + req.Name,
		"name":       req.Name,
		"tier":       req.Tier,
		"created_at": "2024-01-01T00:00:00Z",
	})
}

func getAnalytics(c *gin.Context) {
	c.JSON(200, gin.H{
		"total_requests": 1234,
		"total_tokens":   567890,
		"total_cost":     45.67,
		"by_model": map[string]interface{}{
			"gpt-4": map[string]int{
				"requests": 234,
				"tokens":   123456,
			},
			"gpt-3.5-turbo": map[string]int{
				"requests": 1000,
				"tokens":   444434,
			},
		},
	})
}

func getNginxStatus(c *gin.Context) {
	// TODO: Query Nginx Plus API
	c.JSON(200, gin.H{
		"status":  "running",
		"version": "nginx-plus-r31",
		"upstreams": []string{"openai", "anthropic", "ollama"},
	})
}
