package media

import (
	"context"

	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/media/client/comments"
)

func (s *Service) GetCommentsByMediaID(mediaId string, ctx context.Context) (*comments.GetCommentsByMediaIDOK, error) {
	response, err := s.client.Media.Comments.GetCommentsByMediaID(ctx, &comments.GetCommentsByMediaIDParams{
		InstagramMediaID: mediaId,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *Service) CreateCommentOnMedia(mediaId string, message string, ctx context.Context) (*comments.CreateCommentOnMediaOK, error) {
	response, err := s.client.Media.Comments.CreateCommentOnMedia(ctx, &comments.CreateCommentOnMediaParams{
		InstagramMediaID: mediaId,
		Message:          message,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
