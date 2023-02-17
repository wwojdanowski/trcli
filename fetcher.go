package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SimpleModuleFetcher struct {
	registryUrl string
	limit       int
}

type ModuleFetcherIterator interface {
	hasNext() bool
	next() *ModulesInfo
}

type SimpleModuleFetcherIterator struct {
	fetcher ModuleFetcher
	offset  int
	current *ModulesInfo
}

func (it *SimpleModuleFetcherIterator) hasNext() bool {
	if it.current == nil {
		return true
	}

	return it.current.Meta.NextOffset > 0
}

func (it *SimpleModuleFetcherIterator) next() *ModulesInfo {
	if it.hasNext() {
		return it.fetcher.fetch(it.offset)
	} else {
		return nil
	}
}

type ModuleFetcher interface {
	fetch(offset int) *ModulesInfo
	iterator() ModuleFetcherIterator
}

func (fetcher *SimpleModuleFetcher) fetch(offset int) *ModulesInfo {
	response, err := http.Get(fmt.Sprintf("%s?limit=%d&offset=%d",
		fetcher.registryUrl,
		fetcher.limit, offset))

	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	responseBody, _ := io.ReadAll(response.Body)
	modulesInfo := ModulesInfo{}
	_ = json.Unmarshal(responseBody, &modulesInfo)
	return &modulesInfo
}

func (fetcher *SimpleModuleFetcher) iterator() ModuleFetcherIterator {
	return &SimpleModuleFetcherIterator{fetcher: fetcher}
}
