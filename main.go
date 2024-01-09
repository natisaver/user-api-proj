package main
import (
	"fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)
// album represents data of a user
// string, int, float64
type user struct {
    ID		string	`json:"id"`
    Name	string  `json:"name"`
    Age		int		`json:"age"`
}

// data to store in ram
var users = []user{
	// Populate the array with user data
	{ID: "1", Name: "Alice", Age: 25},
	{ID: "2", Name: "Bob", Age: 30},
	{ID: "3", Name: "Charlie", Age: 22},
}

//take in a parameter of type gin.Context
// it contains request details
// Context.IndentedJSON serailises the struct data into JSON
func getUsers(c *gin.Context) {
	// send client http status code of 200
	c.IndentedJSON(http.StatusOK, users)
}

//get by id
func getUserByID(c *gin.Context) {
    id := c.Param("id")

    // loop over data, look for user ID value matches the parameter.
    for _, x := range users {
        if x.ID == id {
            c.IndentedJSON(http.StatusOK, x)
            return
        }
    }
	// status 404
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

//update handler
func updateUserByID(c *gin.Context) {
    id := c.Param("id")

    // loop over data, look for user ID value matches the parameter.
    for k, x := range users {
		// if user found
        if x.ID == id {
			var updatedUser user
			if err := c.BindJSON(&updatedUser); err != nil {
				// check if updated user details from request body has error
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
				return
			}
			// update the user in the data
			users[k] = updatedUser
			c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %s updated successfully", id)})
            return
        }
    }
	// status 404
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

// post handler
func postUsers(c *gin.Context) {
    var newUser user

    // Call BindJSON to bind the received JSON to
    // newUser.
	// := is used to declare and assign
	// SO here we assign an error, then check if it exists then return
    if err := c.BindJSON(&newUser); err != nil {
        return
    }

    // Add the new album to the data slice.
	//status 201
    users = append(users, newUser)
    c.IndentedJSON(http.StatusCreated, newUser)
} 

//initialises the gin router using Default
// then we assign the getUsers handler to the endpoint /users
func main() {
    router := gin.Default()
    router.GET("/users", getUsers)
	router.POST("/users", postUsers)
	// path parameter
	router.GET("/users/:id", getUserByID)
	router.PUT("/users/:id", updateUserByID)
	//attach router to server and run the server
    router.Run("localhost:8080")
}

