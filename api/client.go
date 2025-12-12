package api

import (
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

var ErrRequestFailed = errors.New("request failed")

type Client struct {
	client *resty.Client
}

type Image struct {
	ID             string    `json:"id"`
	UploadedUserID string    `json:"uploaded_user_id"`
	UploadedAt     time.Time `json:"uploaded_at"`
	Aliases        []string  `json:"aliases"`
	URL            string    `json:"url"`
}

type ImagesResponse struct {
	Images []Image `json:"images"`
}

type AliasesResponse struct {
	Aliases []string `json:"aliases"`
}

func NewClient() *Client {
	client := resty.New()
	client.SetBaseURL(viper.GetString("api-endpoint"))
	client.SetAuthToken(viper.GetString("preshared-key"))

	return &Client{
		client: client,
	}
}
