package main

import (
	"os"
)

type Module struct {
	ID              string `json:"id"`
	Owner           string `json:"owner"`
	Namespace       string `json:"namespace"`
	Name            string `json:"name"`
	Version         string `json:"version"`
	Provider        string `json:"provider"`
	ProviderLogoURL string `json:"provider_logo_url"`
	Description     string `json:"description"`
	Source          string `json:"source"`
	Tag             string `json:"tag"`
	PublishedAt     string `json:"published_at"`
	Downloads       int    `json:"downloads"`
	Verified        bool   `json:"verified"`
}

type ModulesInfo struct {
	Meta struct {
		Limit         int    `json:"limit"`
		CurrentOffset int    `json:"current_offset"`
		NextOffset    int    `json:"next_offset"`
		NextURL       string `json:"next_url"`
	} `json:"meta"`
	Modules []Module `json:"modules"`
}
type ModulesCache struct {
	path string
}

func (mc *ModulesCache) PathExists() bool {
	_, error := os.Stat(mc.path)
	if error != nil {
		return false
	}
	return true
}

func (mc *ModulesCache) CacheDir() string {
	return mc.path
}

func (mc *ModulesCache) CacheFile() string {
	return mc.path + "/modules.json"
}

func (mc *ModulesCache) BuildFullPath() {
	//dir := filepath.Dir(mc.path)
	os.MkdirAll(mc.path, 0755)
}
