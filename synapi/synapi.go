package synapi

type Client struct {
	client   *http.Client // HTTP client used to communicate with the API

	// Services used for talking to different parts of the Synology API
	API 			*APIService
	//...
}

// Client returns the http.Client used by this client
func (c *Client) Client() *http.Client {

}

// NewClient returns a new Synology API client.
func NewClient() *Client {
	c := &Client{client: httpClient}
	c.API = (*APIService)
	return c
}

