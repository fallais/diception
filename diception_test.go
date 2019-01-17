package diception

import (
	"testing"
)

// TestOMG tests the OMG function.
func TestOMG(t *testing.T) {
	solutions := map[string]string{
		"Martine dirige le pays !": "rige",
		"J'suis chaud pour une bonne dichotomie !": "chotomie",
		"Il a une légère dyslexie le bonhomme !": "slexie",
		"dithyrambique": "thyrambique",
	}

	for key, value := range solutions {
		if OMG(key) != value {
			t.Fatal("Nop")
		}
	}
}
