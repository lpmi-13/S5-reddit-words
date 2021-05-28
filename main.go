package main

import (
	"fmt"
	"github.com/turnage/graw/reddit"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	permalink string
	comments  []string
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		Println("had mongo problems:", err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		Println("mas problemas con Mongo")
	}

	fmt.Println("connected to mongo!")

	bot, err := reddit.NewBotFromAgentFile("bot.agent", 0)
	if err != nil {
		fmt.Println("Failed to do stuff with handle:", err)
		return
	}

	harvest, err := bot.Listing("/r/programming", "")
	if err != nil {
		fmt.Println("failed to fetch /r/programming:", err)
		return
	}

	for _, post := range harvest.Posts[:5] {
		post_info, err := bot.Thread(post.Permalink)
		if err != nil {
			fmt.Println("problem grabbing the comments:", err)
			return
		}
		fmt.Println("number of comments:", len(post_info.Replies))
		for _, comment := range post_info.Replies {
			fmt.Printf("%v\n", comment.Body)
		}
	}
}
