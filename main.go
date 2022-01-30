package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       string `json:"id"`
	Name     string `json:"title"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = []user{
	{
		Name:     "John",
		Email:    "john@asd.com",
		ID:       "1",
		Password: "123",
	},
	{
		Name:     "Brandon",
		Email:    "brandon@asd.com",
		ID:       "2",
		Password: "321",
	}}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func userById(c *gin.Context) {
	id := c.Param("id")
	user, err := getUserById(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func getUserById(id string) (*user, error) {
	for i, u := range users {
		if u.ID == id {
			return &users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func createUser(c *gin.Context) {
	var newUser user
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)

}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", userById)
	router.POST("/users", createUser)
	router.Run("localhost:3333")
}
