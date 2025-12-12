package api

import (
	"net/http"
	"strconv"

	"github.com/Penguin-71630/meme-bot-frontend-dc/tracing"
	"go.uber.org/zap"
)

func (c *Client) GetImages(
	search string,
	limit, page int,
) (ImagesResponse, error) {
	var res ImagesResponse
	resp, err := c.client.R().
		SetResult(&res).
		SetQueryParam("search", search).
		SetQueryParam("limit", strconv.Itoa(limit)).
		SetQueryParam("page", strconv.Itoa(page)).
		Get("/api/images")

	if err != nil || resp.StatusCode() != http.StatusOK {
		tracing.Logger.Ctx(resp.Request.Context()).
			Error("failed to get api images",
				zap.String("search", search),
				zap.Int("limit", limit),
				zap.Int("page", page),
				zap.Error(err))
		return ImagesResponse{}, err
	}
	return res, nil
}
