package media

import (
	"github.com/qcserestipy/instagram-api-go-client/pkg/client"
	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/media/client/comments"
)

func GetCommentsByMediaID(mediaId string) (*comments.GetCommentsByMediaIDOK, error) {
	ctx, instagramClient, err := client.ContextWithClient()
	if err != nil {
		return nil, err
	}
	response, err := instagramClient.Media.Comments.GetCommentsByMediaID(ctx, &comments.GetCommentsByMediaIDParams{
		InstagramMediaID: mediaId,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func CreateCommentOnMedia(mediaId string, message string) (*comments.CreateCommentOnMediaOK, error) {
	ctx, instagramClient, err := client.ContextWithClient()
	if err != nil {
		return nil, err
	}
	response, err := instagramClient.Media.Comments.CreateCommentOnMedia(ctx, &comments.CreateCommentOnMediaParams{
		InstagramMediaID: mediaId,
		Message:          message,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
