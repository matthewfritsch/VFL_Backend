package auth

import (
    "context"
    "log"
    "net/http"
    supa "github.com/nedpals/supabase-go"
    "github.com/gin-gonic/gin"
)

func Register(supabase *supa.Client, c *gin.Context) {
    email := c.PostForm("email")
    password := c.PostForm("password")

    if email == "" || password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
        return
    }

    ctx := context.Background()
    user, err := supabase.Auth.SignUp(ctx, supa.UserCredentials{
        Email:    email,
        Password: password,
    })

    if err != nil {
        log.Printf("Error registering user: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}


func SignIn(supabase *supa.Client, c *gin.Context) {
    email := c.PostForm("email")
    password := c.PostForm("password")

    if email == "" || password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
        return
    }

    ctx := context.Background()
    user, err := supabase.Auth.SignIn(ctx, supa.UserCredentials{
        Email:    email,
        Password: password,
    })

    if err != nil {
        log.Printf("Error logging in user: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error logging in user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}
