package main

func main() {
	r := gin.Default()

	// setup routing
	route.SetUp(r, client)

	// todo release
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
