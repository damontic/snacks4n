package main

import (
    "time"
    "net/http"

    "github.com/gin-gonic/gin"
)

func listUsers(c *gin.Context) {
    time.Sleep(5 * time.Second)
    c.String(http.StatusOK, "list of users")
}

func listProducts(c *gin.Context) {
    time.Sleep(5 * time.Second)
    c.String(http.StatusOK, "list of products")
}