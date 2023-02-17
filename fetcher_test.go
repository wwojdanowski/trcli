package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleModuleFetcherIterator_hasNext(t *testing.T) {
	type fields struct {
		fetcher ModuleFetcher
		offset  int
		current *ModulesInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := &SimpleModuleFetcherIterator{
				fetcher: tt.fields.fetcher,
				offset:  tt.fields.offset,
				current: tt.fields.current,
			}
			assert.Equalf(t, tt.want, it.hasNext(), "hasNext()")
		})
	}
}

func TestSimpleModuleFetcherIterator_itIterates(t *testing.T) {

	iterator := SimpleModuleFetcherIterator{
		fetcher: &ModuleFetcherMock{3},
	}

	assert.True(t, iterator.hasNext())

	_ = iterator.next()
	assert.True(t, iterator.hasNext())

	_ = iterator.next()
	assert.True(t, iterator.hasNext())

	_ = iterator.next()
	assert.False(t, iterator.hasNext())

}
