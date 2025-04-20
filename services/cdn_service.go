package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type CdnService struct{}

type purgeRequest struct {
	Files []string `json:"files"`
}

func (cdn *CdnService) GetCDNBaseUrl() string {
	return os.Getenv("CDN_BASE_URL")
}

func (cdn *CdnService) Purge(purgeUrlList ...string) error {
	zoneID := os.Getenv("CDN_CLOUDFLARE_ZONE_ID")
	apiToken := os.Getenv("CDN_CLOUDFLARE_API_TOKEN")

	if zoneID == "" || apiToken == "" {
		return fmt.Errorf("missing Cloudflare credentials")
	}

	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/purge_cache", zoneID)

	body, err := json.Marshal(purgeRequest{Files: purgeUrlList})
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("purge request failed with status: %s", resp.Status)
	}

	return nil
}

func (cdn *CdnService) GetFullPath(relPath string) string {
	return fmt.Sprintf("%s/%s",cdn.GetCDNBaseUrl(), relPath)
}
