package main

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Add(a, b float64) float64 {
	return a + b
}

func Decrease(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) float64 {
	return a / b
}

const defaultPort = "8080"

func main() {
	// args := os.Args[1:]
	// a, _ := strconv.ParseFloat(args[0], 64)
	// b, _ := strconv.ParseFloat(args[1], 64)

	// fmt.Println(Add(a, b))
	// fmt.Println(Decrease(a, b))
	// fmt.Println(Multiply(a, b))
	// fmt.Println(Divide(a, b))

	// fmt.Println("Hello Go")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	f, _ := os.Create("gin.log")
	defer f.Close()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()

	router.GET("/add", func(c *gin.Context) {
		params := c.Request.URL.Query()
		a, _ := strconv.ParseFloat(params.Get("a"), 64)
		b, _ := strconv.ParseFloat(params.Get("b"), 64)

		c.JSON(http.StatusOK, gin.H{
			"result": Add(a, b),
		})
	})
	router.GET("/substract", func(c *gin.Context) {
		params := c.Request.URL.Query()
		a, _ := strconv.ParseFloat(params.Get("a"), 64)
		b, _ := strconv.ParseFloat(params.Get("b"), 64)

		c.JSON(http.StatusOK, gin.H{
			"result": Decrease(a, b),
		})
	})
	router.GET("/multiply", func(c *gin.Context) {
		params := c.Request.URL.Query()
		a, _ := strconv.ParseFloat(params.Get("a"), 64)
		b, _ := strconv.ParseFloat(params.Get("b"), 64)

		c.JSON(http.StatusOK, gin.H{
			"result": Multiply(a, b),
		})
	})
	router.GET("/divide", func(c *gin.Context) {
		params := c.Request.URL.Query()
		a, _ := strconv.ParseFloat(params.Get("a"), 64)
		b, _ := strconv.ParseFloat(params.Get("b"), 64)

		c.JSON(http.StatusOK, gin.H{
			"result": Divide(a, b),
		})
	})

	router.Run(":" + port)
}
