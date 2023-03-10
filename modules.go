package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

func (mc *ModulesCache) DeleteCacheFile() {
	os.Remove(mc.path + "/modules.json")
}

func (mc *ModulesCache) BuildFullPath() {
	os.MkdirAll(mc.path, 0755)
}

func loadModules(fileName string) []Module {
	var modules []Module

	f, err := os.OpenFile(fileName, os.O_RDONLY, 0644)

	if err != nil {
		var pathError *os.PathError
		if errors.As(pathError, &pathError) {
			fmt.Println("modules cache doesn't exists, run `trcli update` to get latest content from registry")
			return nil
		} else {
			log.Panicf("Unexpected error %s\n", err)
		}
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var module Module
		data := fileScanner.Text()
		json.Unmarshal([]byte(data), &module)
		modules = append(modules, module)
	}

	return modules

}

// what about concurrent access?
func (mc *ModulesCache) Find(packageName string) *Module {
	modules := loadModules(mc.CacheFile())

	makeSearchableName := func(module *Module) string {
		return fmt.Sprintf("%s/%s/%s", module.Namespace, module.Name, module.Provider)
	}
	for i := 0; i < len(modules); i++ {
		if packageName == makeSearchableName(&modules[i]) {
			return &modules[i]
		}
	}
	return nil
}
