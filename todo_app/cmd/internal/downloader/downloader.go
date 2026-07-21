package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const CacheDuration = 10 * time.Minute

type ImgCache struct {
	ImagePath string
	CacheDir  string

	mu         sync.Mutex
	expiresAt  time.Time
	hasImage   bool
	refreshing bool
}

func (c *ImgCache) ImageHandler(w http.ResponseWriter, r *http.Request) {
	if err := c.ensureImage(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, c.ImagePath)
}

func (c *ImgCache) ensureImage() error {
	if err := os.MkdirAll(c.CacheDir, 0755); err != nil {
		return err
	}

	c.mu.Lock()

	if !c.hasImage {
		c.mu.Unlock()
		if err := c.downloadImage(); err != nil {
			return err
		}
		c.mu.Lock()
		c.hasImage = true
		c.expiresAt = time.Now().Add(CacheDuration)
		c.refreshing = false
		c.mu.Unlock()
		return nil
	}

	if time.Now().Before(c.expiresAt) {
		c.mu.Unlock()
		return nil
	}

	if !c.refreshing {
		c.refreshing = true
		c.mu.Unlock()
		go c.refreshInBackground()
	} else {
		c.mu.Unlock()
	}

	return nil
}

func (c *ImgCache) refreshInBackground() {
	if err := c.downloadImage(); err != nil {
		log.Printf("failed to refresh cached image: %v", err)
		c.mu.Lock()
		c.refreshing = false
		c.mu.Unlock()
		return
	}

	c.mu.Lock()
	c.expiresAt = time.Now().Add(CacheDuration)
	c.refreshing = false
	c.mu.Unlock()
}

func (c *ImgCache) downloadImage() error {
	resp, err := http.Get("https://picsum.photos/1200")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status from picsum: %s", resp.Status)
	}

	tmpPath := c.ImagePath + ".tmp"
	file, err := os.Create(tmpPath)
	if err != nil {
		return err
	}

	if _, err := io.Copy(file, resp.Body); err != nil {
		file.Close()
		os.Remove(tmpPath)
		return err
	}
	if err := file.Close(); err != nil {
		os.Remove(tmpPath)
		return err
	}

	return os.Rename(tmpPath, c.ImagePath)
}
