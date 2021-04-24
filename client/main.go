package main

import (
	"fmt"
	"github.com/SubashMourougayane/gRPC/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

func main (){

	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err!= nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	g := gin.Default()

	g.GET("/add/:a/:b", func(context *gin.Context) {
		a, err := strconv.ParseUint(context.Param("a"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error" : "Invalid Params",
			})
			return
		}

		b, err := strconv.ParseUint(context.Param("b"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error" : "Invalid Params",
			})
			return
		}

		req := &proto.Request{A:int64(a), B: int64(b)}

		if response, err := client.Add(context, req); err == nil {
			context.JSON(http.StatusOK, gin.H{
				"result" : fmt.Sprint(response.Result),
			})
		}else {
			context.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
		}


	})

	g.GET("/multiply/:a/:b", func(context *gin.Context) {
		a, err := strconv.ParseUint(context.Param("a"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error" : "Invalid Params",
			})
			return
		}

		b, err := strconv.ParseUint(context.Param("b"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error" : "Invalid Params",
			})
			return
		}

		req := &proto.Request{A:int64(a), B: int64(b)}

		if response, err := client.Multiply(context, req); err == nil {
			context.JSON(http.StatusOK, gin.H{
				"result" : fmt.Sprint(response.Result),
			})
		}else {
			context.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
