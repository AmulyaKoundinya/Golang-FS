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
var EmployeeCollection *mongo.Collection

// dept Employee for Collection "Employees"
type Employee struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `json:"name" bson:"name"`
	Dept  string             `json:"dept" bson:"dept"`
	Position   string             `json:"position" bson:"position"`
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

	EmployeeCollection = mongoclient.Database(mongoDbName).Collection(mongoCollectionEmployee)
	if EmployeeCollection==nil{
		log.Fatal("Record Error:", "big problem")
		return
	}
	fmt.Println("Connected to MongoDB!")
}

// POST /Employees
func createEmployee(c *gin.Context) {
	var jbodyEmployee Employee

	// Bind JSON body to jbodyEmployee
	if err := c.BindJSON(&jbodyEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert Employee into MongoDB
	result, err := EmployeeCollection.InsertOne(ctx, jbodyEmployee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Employee"})
		return
	}

	// Extract the inserted ID
	EmployeeId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse inserted ID"})
		return
	}
	jbodyEmployee.ID = EmployeeId

	// Read the created Employee from MongoDB
	var createdEmployee Employee
	err = EmployeeCollection.FindOne(ctx, bson.M{"_id": jbodyEmployee.ID}).Decode(&createdEmployee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch created Employee"})
		return
	}

	// Return created Employee
	c.JSON(http.StatusCreated, gin.H{
		"message": "Employee created successfully",
		"Employee":     createdEmployee,
	})
}

// GET /Employees
func readAllEmployees(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := EmployeeCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Employees"})
		return
	}
	defer cursor.Close(ctx)

	// Ensure Employees is an empty slice, not nil
	Employees := []Employee{}
	if err := cursor.All(ctx, &Employees); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse Employees"})
		return
	}

	c.JSON(http.StatusOK, Employees)
}

// GET /Employees/:id
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

	var Employee Employee
	err = EmployeeCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&Employee)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, Employee)
}

// PUT /Employees/:id
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

	err = EmployeeCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&oldEmployee)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	if  jbodyEmployee.Name!= ""{
		oldEmployee.Name=jbodyEmployee.Name
	}
	if  jbodyEmployee.Dept!= ""{
		oldEmployee.Dept=jbodyEmployee.Dept
	}
	if  jbodyEmployee.Position!= ""{
		oldEmployee.Position=jbodyEmployee.Position
	}
	result, err := EmployeeCollection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": oldEmployee})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Employee"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	// Return updated Employee
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee updated successfully",
		"Employee":     oldEmployee,
	})
}

// DELETE /Employees/:id
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

	result, errDelete := EmployeeCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Employee"})
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
		AllowHeaders:     []string{"Origin", "Content-position", "Authorization"},
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