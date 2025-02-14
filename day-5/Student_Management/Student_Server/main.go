package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors" // Uncomment this
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config
var mongoUri string = "mongodb://localhost:27017"
var mongoDbName string = "student_management_db"
var mongoCollectionStudent string = "students"

// Database variables
var mongoclient *mongo.Client
var studentCollection *mongo.Collection

// Model Student
type Student struct {
	USN     string `bson:"usn" json:"usn"`
	Name    string `bson:"name" json:"name"`
	Section string `bson:"section" json:"section"`
	Type    string `bson:"type" json:"type"`
}

// Connect to MongoDB
func connectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var errConnection error
	mongoclient, errConnection = mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if errConnection != nil {
		log.Fatal("MongoDB Connection Error:", errConnection)
	}

	studentCollection = mongoclient.Database(mongoDbName).Collection(mongoCollectionStudent)
	fmt.Println("Connected to MongoDB!")
}

// POST /students
func createStudent(c *gin.Context) {
	var jbodyStudent Student

	if err := c.BindJSON(&jbodyStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := studentCollection.InsertOne(ctx, jbodyStudent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Student created successfully",
		"student": jbodyStudent,
	})
}

// GET /students
func readAllStudents(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := studentCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}
	defer cursor.Close(ctx)

	students := []Student{}
	if err := cursor.All(ctx, &students); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse students"})
		return
	}

	c.JSON(http.StatusOK, students)
}

// GET /students/:usn
func readStudentByUSN(c *gin.Context) {
	usn := c.Param("usn")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var student Student
	err := studentCollection.FindOne(ctx, bson.M{"usn": usn}).Decode(&student)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// PUT /students/:usn
func updateStudent(c *gin.Context) {
	usn := c.Param("usn")
	var jbodyStudent Student

	if err := c.BindJSON(&jbodyStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := studentCollection.UpdateOne(ctx, bson.M{"usn": usn}, bson.M{"$set": jbodyStudent})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
}

// DELETE /students/:usn
func deleteStudent(c *gin.Context) {
	usn := c.Param("usn")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, errDelete := studentCollection.DeleteOne(ctx, bson.M{"usn": usn})
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}

func main() {
	connectDB()
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // React frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/students", createStudent)
	r.GET("/students", readAllStudents)
	r.GET("/students/:usn", readStudentByUSN)
	r.PUT("/students/:usn", updateStudent)
	r.DELETE("/students/:usn", deleteStudent)

	r.Run(":8080")
}
