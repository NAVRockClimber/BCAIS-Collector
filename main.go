package main

import (
	"context"
	"fmt"
	"log"
	"os"

	arghelper "github.com/NAVRockClimber/BCAIS-Collector/helper"
	"github.com/NAVRockClimber/BCAIS-Collector/helper/feedhelper"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/feed"
)

func main() {
	arguments := os.Args[1:]
	var personalAccessToken, _ = arghelper.GetArgValue(arguments, "PAT")
	var organizationUrl, _ = arghelper.GetArgValue(arguments, "organizationUrl")
	var feedName, _ = arghelper.GetArgValue(arguments, "Feed")
	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)
	ctx := context.Background()

	feedClient, err := feed.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	Feeds, _ := feedClient.GetFeeds(ctx, feed.GetFeedsArgs{})
	myFeed, feedId, err := feedhelper.GetFeedByNameFromFeeds(*Feeds, feedName)
	if err != nil {
		df := log.Default()
		df.Fatal(err.Error())
	}
	fmt.Println(*myFeed.FullyQualifiedName)
	fmt.Println(feedId)
}
