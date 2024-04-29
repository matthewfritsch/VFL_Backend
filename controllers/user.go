package controllers

import (
    "fantasy-valorant/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func RegisterUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // hashedPassword, err := auth.HashPassword(user.Password)
    // if err != nil {
    //     c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
    //     return
    // }
    // user.Password = hashedPassword

    // // Save the user to the database
    // // ...
    //
    // c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func LoginUser(c *gin.Context) {
    var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&credentials); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    _ = user
    // Retrieve the user from the database based on the email
    // ...

    // if !auth.CheckPasswordHash(credentials.Password, user.Password) {
    //     c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    //     return
    // }
    //
    // token, err := auth.GenerateToken(user)
    // if err != nil {
    //     c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
    //     return
    // }
    //
    // c.JSON(http.StatusOK, gin.H{"token": token})
}
