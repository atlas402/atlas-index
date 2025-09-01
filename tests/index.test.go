package core

import (
	"testing"
)

func TestAtlasIndex(t *testing.T) {
	index := New("https://facilitator.payai.network")
	if index == nil {
		t.Error("Expected AtlasIndex instance")
	}
}

