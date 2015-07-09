package speakers_test

import (
	"github.com/stretchr/testify/assert"
	"speakers"
	"testing"
)

func TestSayHello(t *testing.T) {
	english_speaker := speakers.English{Name: "Dan"}
	assert.Equal(t, "Hello World Dan", english_speaker.SayHello(), "Does not say hello")
}
