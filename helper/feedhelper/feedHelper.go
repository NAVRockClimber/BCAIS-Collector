package feedhelper

import (
	"errors"

	"github.com/microsoft/azure-devops-go-api/azuredevops/feed"
)

func GetFeedByNameFromFeeds(FeedArray []feed.Feed, feedName string) (*feed.Feed, string, error) {
	var devopsFeed *feed.Feed
	var feedId string
	for _, curfeed := range FeedArray {
		if *curfeed.FullyQualifiedName == feedName {
			devopsFeed = &curfeed
			feedId = *curfeed.FullyQualifiedId
			break
		}
	}
	if devopsFeed == nil {
		var errMsg = "No feed with name '" + feedName + "' found."
		return devopsFeed, "", errors.New(errMsg)
	}
	return devopsFeed, feedId, nil
}
