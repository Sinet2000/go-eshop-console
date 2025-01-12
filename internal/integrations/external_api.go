// internal/integrations/external_api.go
package integrations

import (
	"context"
	"encoding/json"
	"net/http"
)

type ExternalAPIClient struct {
    client *http.Client
}

func (c *ExternalAPIClient) FetchExternalData(ctx context.Context, url string) (interface{}, error) {
    resp, err := c.client.Get(url)
    if err != nil {
        return nil, err
    }
    var result interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }
    return result, nil
}
