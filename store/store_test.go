package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://go.dev/doc/tutorial/getting-started"
	shortURL := "dysg5Fas"

	// Persist data mapping
	SaveURLInRedis(shortURL, initialLink)

	// Retrieve initial URL
	retrievedUrl := RetrieveInitialURLFromRedis(shortURL)
	assert.Equal(t, initialLink, retrievedUrl)
}
