package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortLinkGenerator(t *testing.T) {
	initialLink := "https://go.dev/doc/tutorial/getting-started"
	userId := "UwQPr3aIf9dM5x7r"
	shortLink := GenerateShortURL(initialLink, userId)

	assert.Equal(t, shortLink, "dysg5Fas")
}
