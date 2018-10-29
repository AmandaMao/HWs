package song

import (
	"strings"
	"testing"
)

func TestSong(t *testing.T) {
	animals := []string{"fly", "spider", "bird", "cat", "dog", "cow", "horse"}
	song := song(animals)
	for _, animal := range animals {
		strings.Contains(animal, song)
	}
}
