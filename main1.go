package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//_ = r.Run() // listen and serve on 0.0.0.0:8080

	var tmpTime string
	tmpTime = "1556561820"
	timestamp, _ := strconv.Atoi(tmpTime)

	fmt.Println(timestamp)

	tm := time.Unix(1556561217, 0)
	fmt.Println(tm)
}
