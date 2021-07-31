package main

import (
	"context"
	"fmt"
	"log"
	"os"

	arghelper "github.com/NAVRockClimber/BCAIS-Collector/helper"
	confighelper "github.com/NAVRockClimber/BCAIS-Collector/helper/configHelper"
	"github.com/NAVRockClimber/BCAIS-Collector/helper/feedhelper"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/feed"
)

func main() {
	arguments := os.Args[1:]
	cfg := initialize(arguments)
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

func initialize(args []string) confighelper.Config {
	var configFile, _ = arghelper.GetArgValue(args, "Config")
	var cfg confighelper.Config
	confighelper.ReadFile(configFile, &cfg)
	confighelper.ReadEnv(&cfg)
	OrgCount := len(cfg.Organizations)
	if OrgCount == 0 {
		var org confighelper.Organizations
	}
	var personalAccessToken, _ = arghelper.GetArgValue(args, "PAT")
	var organizationUrl, _ = arghelper.GetArgValue(args, "organizationUrl")
	var feedName, _ = arghelper.GetArgValue(args, "Feed")
}
