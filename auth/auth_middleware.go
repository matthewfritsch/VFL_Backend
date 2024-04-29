package auth

import (
    "context"
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    supa "github.com/nedpals/supabase-go"
)

func AuthMiddleware(supabase *supa.Client, c *gin.Context) {
    authHeader := c.GetHeader("Authorization")

    if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    token := authHeader[7:]

    ctx := context.Background()
    user, err := supabase.Auth.User(ctx, token)

    if err != nil || user == nil {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }
    c.Set("user", user)
    c.Next()
}
