package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	r := gin.New()
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/"),
		gin.Recovery(),
	)
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Use token to fetch accounts",
		})
	})

	r.POST("/api/received-accounts", func(c *gin.Context) {
		accountsUrl := "https://sandbox-products.ddp.akoya.com/accounts-info/v2/mikomo"
		token := c.PostForm("token")

		req, err := http.NewRequest("GET", accountsUrl, nil)
		if err != nil {
			panic(err)
		}
		req.Header.Set("accept", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("client: error making http request: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("client: got response!\n")
		fmt.Printf("client: status code: %d\n", res.StatusCode)

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("client: could not read response body: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("client: response body: %s\n", resBody)

		c.Redirect(302, "/api/received-accounts?accounts="+string(resBody))
	})

	r.GET("/api/received-accounts", func(c *gin.Context) {
		accounts := c.Query("accounts")
		c.HTML(http.StatusOK, "accounts.tmpl", gin.H{
			"title":    "Received accounts",
			"accounts": accounts,
		})
	})
	r.Run(":8123") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
