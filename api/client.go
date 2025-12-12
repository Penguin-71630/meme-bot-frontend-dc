package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	token      string
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

type GenLoginURLRequest struct {
	UserID string `json:"userId"`
}

type GenLoginURLResponse struct {
	LoginURL string `json:"loginUrl"`
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) SetToken(token string) {
	c.token = token
}

func (c *Client) GetImages(search string, limit, page int) (*ImagesResponse, error) {
	url := fmt.Sprintf("%s/api/images?search=%s&limit=%d&page=%d", c.baseURL, search, limit, page)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var result ImagesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetImage(id string) (*Image, error) {
	url := fmt.Sprintf("%s/api/images/%s", c.baseURL, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var image Image
	if err := json.NewDecoder(resp.Body).Decode(&image); err != nil {
		return nil, err
	}

	return &image, nil
}

func (c *Client) GetAliases() (*AliasesResponse, error) {
	url := fmt.Sprintf("%s/api/aliases", c.baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var result AliasesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UploadImage(imageData []byte, aliases []string) (*Image, error) {
	// TODO: Implement multipart form upload
	return nil, fmt.Errorf("not implemented")
}

func (c *Client) DeleteImage(id string) error {
	url := fmt.Sprintf("%s/api/images/%s", c.baseURL, id)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) AddAlias(imageID, alias string) (*Image, error) {
	url := fmt.Sprintf("%s/api/images/%s/aliases", c.baseURL, imageID)

	body := map[string]string{"alias": alias}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var image Image
	if err := json.NewDecoder(resp.Body).Decode(&image); err != nil {
		return nil, err
	}

	return &image, nil
}

func (c *Client) GetImageFile(id string) ([]byte, error) {
	url := fmt.Sprintf("%s/api/images/%s/file", c.baseURL, id)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func (c *Client) GenerateLoginURL(userID string) (string, error) {
	url := fmt.Sprintf("%s/auth/gen-login-url", c.baseURL)

	reqBody := GenLoginURLRequest{UserID: userID}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	// TODO: Set preshared key authorization
	req.Header.Set("Authorization", "Bearer poop")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	var result GenLoginURLResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.LoginURL, nil
}
