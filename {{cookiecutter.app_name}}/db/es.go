package db

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	es "github.com/olivere/elastic/v7"

	"log"
	"os"
	"time"
)

func ESConnect() *es.Client {
	err := godotenv.Load()
	options := []es.ClientOptionFunc{
		es.SetHealthcheck(true),
		es.SetSniff(false),
		es.SetURL(os.Getenv("ES_HOST")),
		es.SetRetrier(es.NewBackoffRetrier(es.NewConstantBackoff(5 * time.Second))),
	}
	c, err := es.NewClient(options...)
	if err != nil {
		log.Fatalf("Got error when connect elasticsearch, the error is '%v'", err)
	}
	return c
}

func ESInit(c *gin.Context) *es.Client {
	elastic := c.MustGet("ES").(*es.Client)
	return elastic
}
