package remote

import (
	"encoding/json"
	"fmt"
	"github.com/coding-wepack/carctl/pkg/settings"
	"github.com/coding-wepack/carctl/pkg/util/httputil"
	"github.com/pkg/errors"
	"io"
	"net/url"
)

type NexusArtListResp struct {
	Items []Item `json:"items"`
}

type Item struct {
	ID         string  `json:"id"`
	Repository string  `json:"repository"`
	Format     string  `json:"format"`
	Group      string  `json:"group"`
	Name       string  `json:"name"`
	Version    string  `json:"version"`
	Assets     []Asset `json:"assets"`
}

type Asset struct {
	DownloadUrl string `json:"downloadUrl"`
	Path        string `json:"path"`
	ID          string `json:"id"`
	Repository  string `json:"repository"`
	Format      string `json:"format"`
	ContentType string `json:"contentType"`
}

func FindFileListFromNexus(nexusUrl *url.URL) ([]Asset, error) {

	resp, err := httputil.DefaultClient.GetWithAuth(nexusUrl.String(), settings.SrcUsername, settings.SrcPassword)
	if err != nil {
		return nil, errors.Wrap(err, "Send http request failed.")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Close resp failed:", err)
		}
	}(resp.Body)

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read from resp failed.:", err)
		return nil, err
	}

	var nar NexusArtListResp

	err = json.Unmarshal(bytes, &nar)
	if err != nil {
		return nil, err
	}

	assets := make([]Asset, 0)
	for _, item := range nar.Items {
		for _, asset := range item.Assets {
			assets = append(assets, asset)
		}
	}

	return assets, nil
}
