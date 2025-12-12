package api

import (
	"net/http"

	"github.com/Penguin-71630/meme-bot-frontend-dc/tracing"
	"go.uber.org/zap"
)

type GenLoginURLRequest struct {
	UserId string `json:"userId"`
}

type GenLoginURLResponse struct {
	LoginURL string `json:"loginUrl"`
}

func (c *Client) PostGenLoginURL(
	userId string,
) (string, error) {
	var res GenLoginURLResponse
	resp, err := c.client.R().
		SetBody(GenLoginURLRequest{UserId: userId}).
		Post("/auth/gen-login-url")
	if err != nil || resp.StatusCode() != http.StatusOK {
		tracing.Logger.Ctx(resp.Request.Context()).
			Error("failed to post gen-login-url",
				zap.Error(err))
		return "", ErrRequestFailed
	}
	return res.LoginURL, nil
}
