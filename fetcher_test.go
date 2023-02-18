package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleModuleFetcherIterator_itIterates(t *testing.T) {
	iterator := SimpleModuleFetcherIterator{offsetDelta: 3}
	iterator.fetcher = &ModuleFetcherMock{}

	assert.Equal(t, iterator.nextOffset, 0)
	assert.True(t, iterator.hasNext())

	iterator.next()
	assert.Equal(t, iterator.nextOffset, 3)
	assert.True(t, iterator.hasNext())

	iterator.next()
	assert.Equal(t, iterator.nextOffset, 6)
	assert.True(t, iterator.hasNext())

	iterator.next()
	assert.Equal(t, iterator.nextOffset, 9)
	assert.True(t, iterator.hasNext())

	_ = iterator.next()
	assert.Equal(t, iterator.nextOffset, 0)
	assert.False(t, iterator.hasNext())

}
