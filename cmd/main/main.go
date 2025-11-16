package main

import (
	"fmt"

	"github.com/qcserestipy/instagram-api-go-client/pkg/instagram"

	"github.com/sirupsen/logrus"
)

func main() {
	reels, err := instagram.GetReels("17841464714098258", nil, nil)
	if err != nil {
		logrus.Fatalf("fatal error: %v", err)
	}
	fmt.Println(reels)

	// 	postResp, err := media.CreateCommentOnMedia(mediaItem.ID, "This is a remark from the Instagram API Go Client!")
	// 	if err != nil {
	// 		logrus.Warnf("failed to create comment on media %s: %v", mediaItem.ID, err)
	// 	} else {
	// 		fmt.Printf("Success! Comment created on media %s: %+v\n", mediaItem.ID, postResp)
	// 	}
	// }
}
