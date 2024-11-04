package media

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"time"
)

type (
	// DownloaderConfig represents the configuration for the Downloader.
	DownloaderConfig struct {
		Timeout time.Duration
		MaxSize int64
	}

	// Downloader is responsible for downloading media via HTTP.
	Downloader struct {
		client *http.Client
		config DownloaderConfig
	}
)

var (
	// DefaultDownloaderConfig is the default configuration for the Downloader.
	DefaultDownloaderConfig = DownloaderConfig{
		Timeout: 15 * time.Second,
		MaxSize: 10 * 1024 * 1024, // 10MB
	}
)

// NewDownloader creates a new Downloader with the given configuration.
func NewDownloader(config DownloaderConfig) *Downloader {
	return &Downloader{
		client: &http.Client{
			Timeout: config.Timeout,
		},
		config: config,
	}
}

// Download downloads a media object via HTTP and returns its content as a slice of bytes.
func (d *Downloader) Download(url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), d.config.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := d.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("media.Downloader.Download: failed to close response body: ", err.Error())
		}
	}(resp.Body)

	// Check for a successful HTTP status code
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to download image: " + resp.Status)
	}

	limitedReader := io.LimitReader(resp.Body, d.config.MaxSize+1)
	data, err := io.ReadAll(limitedReader)
	if err != nil {
		return nil, err
	}

	if int64(len(data)) > d.config.MaxSize {
		return nil, errors.New("image is too large")
	}

	return data, nil
}
