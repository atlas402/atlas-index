package core

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ServiceDiscovery struct {
	ID          string
	Name        string
	Description string
	Endpoint    string
	Category    string
	Network     string
	Accepts     []PaymentAccept
	Metadata    map[string]interface{}
}

type PaymentAccept struct {
	Asset            string
	PayTo            string
	Network          string
	MaxAmountRequired string
	Scheme           string
	MimeType         string
}

type AtlasIndex struct {
	facilitatorURL string
	serviceCache   map[string]*ServiceDiscovery
	httpClient     *http.Client
}

func New(facilitatorURL string) *AtlasIndex {
	return &AtlasIndex{
		facilitatorURL: facilitatorURL,
		serviceCache:   make(map[string]*ServiceDiscovery),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (idx *AtlasIndex) Discover(ctx context.Context, options *DiscoveryOptions) ([]*ServiceDiscovery, error) {
	url := fmt.Sprintf("%s/discovery/resources", idx.facilitatorURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if options.Category != "" {
		q.Add("category", options.Category)
	}
	if options.Network != "" {
		q.Add("network", options.Network)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := idx.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("discovery failed with status %d", resp.StatusCode)
	}

	var result struct {
		Resources []*ServiceDiscovery `json:"resources"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	for _, service := range result.Resources {
		idx.serviceCache[service.ID] = service
	}

	return result.Resources, nil
}

type DiscoveryOptions struct {
	Category string
	Network  string
	Scheme   string
	Limit    int
	Offset   int
}

