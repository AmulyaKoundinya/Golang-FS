package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config
var mongoUri string = "mongodb://localhost:27017"
var mongoDbName string = "emp_app_db"
var mongoCollectionEmployee string = "employees"

// Database variables
var mongoclient *mongo.Client
var employeeCollection *mongo.Collection

// Model Employee for Collection "employees"
type Employee struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Dept     string             `json:"dept" bson:"dept"`
	Position string             `json:"position" bson:"position"`
}

// Connect to MongoDB
func connectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var errrorConnection error
	mongoclient, errrorConnection = mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if errrorConnection != nil {
		log.Fatal("MongoDB Connection Error:", errrorConnection)
	}

	employeeCollection = mongoclient.Database(mongoDbName).Collection(mongoCollectionEmployee)
	fmt.Println("Connected to MongoDB!")
}

// POST /employees
func createEmployee(c *gin.Context) {
	var jbodyEmployee Employee

	// Bind JSON body to jbodyEmployee
	if err := c.BindJSON(&jbodyEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert employee into MongoDB
	result, err := employeeCollection.InsertOne(ctx, jbodyEmployee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
		return
	}

	// Extract the inserted ID
	employeeId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse inserted ID"})
		return
	}
	jbodyEmployee.ID = employeeId

	// Read the created employee from MongoDB
	var createdEmployee Employee
	err = employeeCollection.FindOne(ctx, bson.M{"_id": jbodyEmployee.ID}).Decode(&createdEmployee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch created employee"})
		return
	}

	// Return created employee
	c.JSON(http.StatusCreated, gin.H{
		"message": "Employee created successfully",
		"employee": createdEmployee,
	})
}

// GET /employees
func readAllEmployees(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := employeeCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
		return
	}
	defer cursor.Close(ctx)

	// Ensure employees is an empty slice, not nil
	employees := []Employee{}
	if err := cursor.All(ctx, &employees); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse employees"})
		return
	}

	c.JSON(http.StatusOK, employees)
}

// GET /employees/:id
func readEmployeeById(c *gin.Context) {
	id := c.Param("id")

	// Convert string ID to primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var employee Employee
	err = employeeCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&employee)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// PUT /employees/:id
func updateEmployee(c *gin.Context) {
	id := c.Param("id")
	// Convert string ID to primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var jbodyEmployee Employee

	if err := c.BindJSON(&jbodyEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var oldEmployee Employee

	err = employeeCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&oldEmployee)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	if jbodyEmployee.Name != "" {
		oldEmployee.Name = jbodyEmployee.Name
	}
	if jbodyEmployee.Dept != "" {
		oldEmployee.Dept = jbodyEmployee.Dept
	}
	if jbodyEmployee.Position != "" {
		oldEmployee.Position = jbodyEmployee.Position
	}
	result, err := employeeCollection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": oldEmployee})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	// Return updated employee
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee updated successfully",
		"employee": oldEmployee,
	})
}

// DELETE /employees/:id
func deleteEmployee(c *gin.Context) {
	id := c.Param("id")
	// Convert string ID to primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, errDelete := employeeCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

func main() {
	// Connect to MongoDB
	connectDB()

	// Set up Gin router
	r := gin.Default()
	// CORS Configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // React frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Routes
	r.POST("/employees", createEmployee)
	r.GET("/employees", readAllEmployees)
	r.GET("/employees/:id", readEmployeeById)
	r.PUT("/employees/:id", updateEmployee)
	r.DELETE("/employees/:id", deleteEmployee)

	// Start server
	r.Run(":8080")
}
