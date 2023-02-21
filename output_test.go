package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModuleOutputIsProperlyFormatted(t *testing.T) {
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
	outputString := formatModuleOutput(&module)

	assert.Equal(t, "namespace1/package1/provider1 - 1.0.0 - test module", outputString)
}

func TestModuleInfoOutputIsProperlyFormatted(t *testing.T) {
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

	outputString := formatModuleInfo(&module)
	expectedOutput := `namespace1/package1/provider1
Namespace: namespace1
Name: package1
Provider: provider1
Version: 1.0.0
Description: test module
`
	assert.Equal(t, expectedOutput, outputString)
}
