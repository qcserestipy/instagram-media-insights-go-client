package main

import (
	"context"
	"fmt"
	"log"

	"github.com/qcserestipy/instagram-api-go-client/pkg/account"
	"github.com/qcserestipy/instagram-api-go-client/pkg/client"
	"github.com/qcserestipy/instagram-api-go-client/pkg/instagram"
	"github.com/qcserestipy/instagram-api-go-client/pkg/media"
	"github.com/sirupsen/logrus"
)

func main() {

	igClient, err := client.NewDefault()
	if err != nil {
		log.Fatalf("failed to create instagram client: %v", err)
	}

	accountSvc := account.NewService(igClient)
	mediaSvc := media.NewService(igClient)

	ctx := context.Background()

	dyn, err := instagram.GetFollowerDynamics(ctx, accountSvc, "17841464714098258", "last_21_days")
	if err != nil {
		logrus.Fatalf("fatal error: %v", err)
	}
	logrus.Infof("Follower Dynamics: %+v", dyn)

	demographics, err := instagram.GetAccountDemographics(ctx, accountSvc, "17841464714098258")
	if err != nil {
		logrus.Fatalf("fatal error: %v", err)
	}
	logrus.Infof("Follower Demographics: %+v", demographics.Follower.Countries["DE"])
	logrus.Infof("Engaged Audience Demographics: %+v", demographics.Engaged.Countries["DE"])

	info, err := instagram.GetAccountInfo(ctx, accountSvc, "17841464714098258")
	if err != nil {
		logrus.Fatalf("fatal error: %v", err)
	}
	logrus.Infof("Info: %s, Followers: %d, Following: %d, Media: %d", info.Username, info.FollowersCount, info.FollowingCount, info.MediaCount)

	// err := utils.RefreshAccessToken("627978530406248")
	// if err != nil {
	// 	logrus.Fatalf("fatal error: %v", utils.ParseAPIError(err, "refresh access token"))
	// }

	reels, err := instagram.GetReels(ctx, accountSvc, mediaSvc, "17841464714098258", nil, nil)
	if err != nil {
		logrus.Fatalf("fatal error: %v", err)
	}
	for _, reel := range reels {
		fmt.Printf("Reel ID: %s, Time: %s, Views: %d, Reach: %d, Likes: %d, Comments: %d, Shares: %d, Saves: %d, Engagement Views: %.2f%%\n",
			reel.ID, reel.DateTime, reel.Views, reel.Reach, reel.Likes, reel.Comments, reel.Shares, reel.Saves, reel.EngagementViews)
	}

	// stories, err := instagram.GetStories("17841464714098258")
	// if err != nil {
	// 	logrus.Fatalf("fatal error: %v", err)
	// }
	// for _, story := range stories {
	// 	fmt.Printf("Story ID: %s, Views: %d, Reach: %d, Replies: %d, Shares: %d, Navigation: %d, Profile Activity: %d, Profile Visits: %d, Follows: %d, Total Interactions: %d, Engagement Views: %.2f%%\n",
	// 		story.ID, story.Views, story.Reach, story.Replies, story.Shares, story.Navigation, story.ProfileActivity, story.ProfileVisits, story.Follows, story.TotalInteractions, story.EngagementViews)
	// }

	// 	postResp, err := media.CreateCommentOnMedia(mediaItem.ID, "This is a remark from the Instagram API Go Client!")
	// 	if err != nil {
	// 		logrus.Warnf("failed to create comment on media %s: %v", mediaItem.ID, err)
	// 	} else {
	// 		fmt.Printf("Success! Comment created on media %s: %+v\n", mediaItem.ID, postResp)
	// 	}
	// }
}
