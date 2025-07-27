package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// User represents a user in our system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Code    int         `json:"code,omitempty"`
}

// In-memory storage
var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: 25},
	{ID: 3, Name: "Bob Wilson", Email: "bob@example.com", Age: 35},
}
var nextID = 4

func main() {
	// TODO: Create Gin router
	router := gin.Default()
	router.GET("/pong", getPing)
	// TODO: Setup routes
	// GET /users - Get all users
	router.GET("/users", getAllUsers)
	// GET /users/:id - Get user by ID
	router.GET("/users/:id", getUserByID)
	// POST /users - Create new user
	router.POST("/users", createUser)
	// PUT /users/:id - Update user
	router.PUT("/users/:id", updateUser)
	// DELETE /users/:id - Delete user
	router.DELETE("/users/:id", deleteUser)
	// GET /users/search - Search users by name
	router.GET("/users/search", searchUsers)
	// TODO: Start server on port 8080
	router.Run("localhost:8080")
}

// TODO: Implement handler functions

// getAllUsers handles GET /users
func getAllUsers(c *gin.Context) {
	// TODO: Return all users
	resp := Response{
		Success: true,
		Data:    users,
		Message: "users retrieved successfully",
		Code:    200,
	}

	c.IndentedJSON(http.StatusOK, resp)
}

// getUserByID handles GET /users/:id
func getUserByID(c *gin.Context) {
	// TODO: Get user by ID
	// Handle invalid ID format
	// Return 404 if user not found
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	for _, usr := range users {
		if usr.ID == id {
			resp := Response{
				Success: true,
				Data:    usr,
				Message: "user retrieved successfully",
				Code:    200,
			}
			c.IndentedJSON(http.StatusOK, resp)
			return
		}
	}

	resp := Response{
		Success: false,
		Error:   "Not found",
		Code:    404,
	}

	c.IndentedJSON(http.StatusNotFound, resp)
}

func getPing(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
}

// createUser handles POST /users
func createUser(c *gin.Context) {
	// TODO: Parse JSON request body
	// Validate required fields
	// Add user to storage
	// Return created user

	var usr User

	if err := c.BindJSON(&usr); err != nil {
		resp := Response{
			Success: false,
			Error:   "invalid JSON",
			Code:    http.StatusBadRequest,
		}
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	if usr.Name == "" || usr.Email == "" {
		resp := Response{
			Success: false,
			Error:   "name and email are required",
			Code:    http.StatusBadRequest,
		}
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	usr.ID = nextID
	nextID++
	users = append(users, usr)

	resp := Response{
		Success: true,
		Data:    usr,
		Message: "user created successfully",
		Code:    http.StatusCreated,
	}

	// Add to storage
	c.IndentedJSON(http.StatusCreated, resp)
}

// updateUser handles PUT /users/:id
func updateUser(c *gin.Context) {
	// TODO: Get user ID from path
	// Parse JSON request body
	// Find and update user
	// Return updated user
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp := Response{
			Success: false,
			Error:   "invalid user ID",
			Code:    http.StatusBadRequest,
		}
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	var updatedUser User
	if err := c.BindJSON(&updatedUser); err != nil {
		resp := Response{
			Success: false,
			Error:   "invalid JSON",
			Code:    http.StatusBadRequest,
		}
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	if updatedUser.Name == "" || updatedUser.Email == "" {
		resp := Response{
			Success: false,
			Error:   "name and email are required",
			Code:    http.StatusBadRequest,
		}
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	for _, usr := range users {
		if usr.ID == id {
			usr.Name = updatedUser.Name
			usr.Email = updatedUser.Email
			usr.Age = updatedUser.Age

			resp := Response{
				Success: true,
				Data:    usr,
				Message: "user updated successfully",
				Code:    http.StatusOK,
			}
			c.IndentedJSON(http.StatusOK, resp)
			return
		}
	}

	resp := Response{
		Success: false,
		Error:   "user not found",
		Code:    http.StatusNotFound,
	}

	c.IndentedJSON(http.StatusNotFound, resp)
}

// deleteUser handles DELETE /users/:id
func deleteUser(c *gin.Context) {
	// TODO: Get user ID from path
	// Find and remove user
	// Return success message

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp := Response{
			Success: false,
			Error:   "invalid user ID",
			Code:    http.StatusBadRequest,
		}
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}
	for i, usr := range users {
		if usr.ID == id {
			users = append(users[:i], users[i+1:]...)
			resp := Response{
				Success: true,
				Message: "user deleted successfully",
				Code:    http.StatusOK,
			}
			c.IndentedJSON(http.StatusOK, resp)
			return
		}
	}
	resp := Response{
		Success: false,
		Error:   "user not found",
		Code:    http.StatusNotFound,
	}
	c.IndentedJSON(http.StatusNotFound, resp)
}

// searchUsers handles GET /users/search?name=value
func searchUsers(c *gin.Context) {
	// TODO: Get name query parameter
	// Filter users by name (case-insensitive)
	// Return matching users

	name := c.Query("name")
	if name == "" {
		resp := Response{
			Success: false,
			Error:   "name query parameter is required",
			Code:    http.StatusBadRequest,
		}
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}
	results := make([]User, 0)
	lowerName := strings.ToLower(name)

	for _, usr := range users {
		if usr.Name != "" && strings.Contains(strings.ToLower(usr.Name), lowerName) {
			results = append(results, usr)
		}
	}

	resp := Response{
		Success: true,
		Data:    results,
		Message: "users found",
		Code:    http.StatusOK,
	}
	c.IndentedJSON(http.StatusOK, resp)
}

// Helper function to find user by ID
func findUserByID(id int) (*User, int) {
	// TODO: Implement user lookup
	// Return user pointer and index, or nil and -1 if not found
	return nil, -1
}

// Helper function to validate user data
func validateUser(user User) error {
	// TODO: Implement validation
	// Check required fields: Name, Email
	// Validate email format (basic check)
	return nil
}
