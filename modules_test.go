package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModuleIsSavedToDataFile(t *testing.T) {
	module := Module{
		ID:              "id1",
		Owner:           "owner1",
		Namespace:       "namespace1",
		Name:            "package1",
		Version:         "1.0.0",
		Provider:        "provider1",
		ProviderLogoURL: "http://unknown.com/img1.jpg",
		Description:     "test module",
		Source:          "http://github.com/test_module",
		Tag:             "tag1",
		PublishedAt:     "12:00:00",
		Downloads:       10,
		Verified:        true,
	}

	file, err := os.CreateTemp("", "tmpfile-")

	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	storeModuleData(&module, file.Name())

	data, err := ioutil.ReadFile(file.Name())

	if err != nil {
		log.Fatal(err)
	}

	expected := `{"id":"id1","owner":"owner1","namespace":"namespace1","name":"package1","version":"1.0.0","provider":"provider1","provider_logo_url":"http://unknown.com/img1.jpg","description":"test module","source":"http://github.com/test_module","tag":"tag1","published_at":"12:00:00","downloads":10,"verified":true}`
	actual := strings.SplitN(string(data), "\n", 2)[0]
	fmt.Println(actual)
	os.Remove(file.Name())

	assert.Equal(t, expected, actual)
}

func TestModulesAreLoadedFromFile(t *testing.T) {
	file, err := os.CreateTemp("", "tmpfile-")

	if err != nil {
		log.Fatal(err)
	}

	expected := []Module{
		{
			ID:              "id1",
			Owner:           "owner1",
			Namespace:       "namespace1",
			Name:            "package1",
			Version:         "1.0.0",
			Provider:        "provider1",
			ProviderLogoURL: "http://unknown.com/img1.jpg",
			Description:     "test module",
			Source:          "http://github.com/test_module",
			Tag:             "tag1",
			PublishedAt:     "12:00:00",
			Downloads:       10,
			Verified:        true,
		},
		{
			ID:              "id2",
			Owner:           "owner2",
			Namespace:       "namespace2",
			Name:            "package2",
			Version:         "2.0.0",
			Provider:        "provider2",
			ProviderLogoURL: "http://unknown.com/img2.jpg",
			Description:     "test module",
			Source:          "http://github.com/test_module",
			Tag:             "tag2",
			PublishedAt:     "13:00:00",
			Downloads:       20,
			Verified:        false,
		},
	}

	writeStringNewLine(`{"id":"id1","owner":"owner1","namespace":"namespace1","name":"package1","version":"1.0.0","provider":"provider1","provider_logo_url":"http://unknown.com/img1.jpg","description":"test module","source":"http://github.com/test_module","tag":"tag1","published_at":"12:00:00","downloads":10,"verified":true}`, file)
	writeStringNewLine(`{"id":"id2","owner":"owner2","namespace":"namespace2","name":"package2","version":"2.0.0","provider":"provider2","provider_logo_url":"http://unknown.com/img2.jpg","description":"test module","source":"http://github.com/test_module","tag":"tag2","published_at":"13:00:00","downloads":20,"verified":false}`, file)
	file.Close()

	modules := loadModules(file.Name())

	assert.Len(t, modules, 2)
	assert.Equal(t, expected, modules)
}

func TestPatternMatchesModuleMetadata(t *testing.T) {
	modules := []Module{
		{
			ID:              "id1",
			Owner:           "owner1",
			Namespace:       "namespace1",
			Name:            "package1",
			Version:         "1.0.0",
			Provider:        "provider1",
			ProviderLogoURL: "http://unknown.com/img1.jpg",
			Description:     "test module",
			Source:          "http://github.com/test_module",
			Tag:             "tag1",
			PublishedAt:     "12:00:00",
			Downloads:       10,
			Verified:        true,
		},
		{
			ID:              "id2",
			Owner:           "owner2",
			Namespace:       "namespace2",
			Name:            "package2",
			Version:         "2.0.0",
			Provider:        "provider2",
			ProviderLogoURL: "http://unknown.com/img2.jpg",
			Description:     "test module",
			Source:          "http://github.com/test_module",
			Tag:             "tag2",
			PublishedAt:     "13:00:00",
			Downloads:       20,
			Verified:        false,
		},
	}

	assert.Equal(t, modules[0:1], filterModulesByPattern("ackage1", modules))
	assert.Equal(t, modules[1:2], filterModulesByPattern("ackage2", modules))

}

func filterModulesByPattern(pattern string, modules []Module) []Module {
	var filtered []Module
	for _, module := range modules {
		if patternMatchesModuleMetadata(pattern, &module) {
			filtered = append(filtered, module)
		}
	}
	return filtered
}

func writeStringNewLine(data string, file *os.File) {
	file.WriteString(fmt.Sprintf("%s\n", data))
}
