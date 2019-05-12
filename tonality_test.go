package tonality_test

import (
	"testing"

	"github.com/tombell/tonality"
)

func TestConvertCamelotToOpenKey(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1A", "6M"},
		{"1B", "6D"},
		{"2A", "7M"},
		{"2B", "7D"},
		{"3A", "8M"},
		{"3B", "8D"},
		{"4A", "9M"},
		{"4B", "9D"},
		{"5A", "10M"},
		{"5B", "10D"},
		{"6A", "11M"},
		{"6B", "11D"},
		{"7A", "12M"},
		{"7B", "12D"},
		{"8A", "1M"},
		{"8B", "1D"},
		{"9A", "2M"},
		{"9B", "2D"},
		{"10A", "3M"},
		{"10B", "3D"},
		{"11A", "4M"},
		{"11B", "4D"},
		{"12A", "5M"},
		{"12B", "5D"},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			key, err := tonality.ConvertKeyToNotation(tc.input, tonality.OpenKey)
			if err != nil {
				t.Fatalf("expected ConvertKeyToNotation to not return an error: got %v", err)
			}

			if key != tc.expected {
				t.Fatalf("expected ConvertKeyToNotation to return key %v: got %v", tc.expected, key)
			}
		})
	}
}

func TestConvertCamelotToMusical(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1A", "Abm"},
		{"1B", "B"},
		{"2A", "Ebm"},
		{"2B", "Gb"},
		{"3A", "Bbm"},
		{"3B", "Db"},
		{"4A", "Fm"},
		{"4B", "Ab"},
		{"5A", "Cm"},
		{"5B", "Eb"},
		{"6A", "Gm"},
		{"6B", "Bb"},
		{"7A", "Dm"},
		{"7B", "F"},
		{"8A", "Am"},
		{"8B", "C"},
		{"9A", "Em"},
		{"9B", "G"},
		{"10A", "Bm"},
		{"10B", "D"},
		{"11A", "Gbm"},
		{"11B", "A"},
		{"12A", "Dbm"},
		{"12B", "E"},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			key, err := tonality.ConvertKeyToNotation(tc.input, tonality.Musical)
			if err != nil {
				t.Fatalf("expected ConvertKeyToNotation to not return an error: got %v", err)
			}

			if key != tc.expected {
				t.Fatalf("expected ConvertKeyToNotation to return key %v: got %v", tc.expected, key)
			}
		})
	}
}

func TestConvertCamelotToMusicalAlt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1A", "G#m"},
		{"1B", "B"},
		{"2A", "Ebm"},
		{"2B", "F#"},
		{"3A", "A#m"},
		{"3B", "Db"},
		{"4A", "Fm"},
		{"4B", "G#"},
		{"5A", "Cm"},
		{"5B", "D#"},
		{"6A", "Gm"},
		{"6B", "Bb"},
		{"7A", "Dm"},
		{"7B", "F"},
		{"8A", "Am"},
		{"8B", "C"},
		{"9A", "Em"},
		{"9B", "G"},
		{"10A", "Bm"},
		{"10B", "D"},
		{"11A", "F#m"},
		{"11B", "A"},
		{"12A", "C#m"},
		{"12B", "E"},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			key, err := tonality.ConvertKeyToNotation(tc.input, tonality.MusicalAlt)
			if err != nil {
				t.Fatalf("expected ConvertKeyToNotation to not return an error: got %v", err)
			}

			if key != tc.expected {
				t.Fatalf("expected ConvertKeyToNotation to return key %v: got %v", tc.expected, key)
			}
		})
	}
}

func TestConvertCamelotToBeatport(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1A", "G#m"},
		{"1B", "Bmaj"},
		{"2A", "Ebm"},
		{"2B", "Gb"},
		{"3A", "Bbm"},
		{"3B", "Db"},
		{"4A", "Fmin"},
		{"4B", "Ab"},
		{"5A", "Cmin"},
		{"5B", "Eb"},
		{"6A", "Gmin"},
		{"6B", "Bb"},
		{"7A", "Dmin"},
		{"7B", "Fmaj"},
		{"8A", "Amin"},
		{"8B", "Cmaj"},
		{"9A", "Emin"},
		{"9B", "Gmaj"},
		{"10A", "Bmin"},
		{"10B", "Dmaj"},
		{"11A", "F#m"},
		{"11B", "Amaj"},
		{"12A", "C#m"},
		{"12B", "Emaj"},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			key, err := tonality.ConvertKeyToNotation(tc.input, tonality.Beatport)
			if err != nil {
				t.Fatalf("expected ConvertKeyToNotation to not return an error: got %v", err)
			}

			if key != tc.expected {
				t.Fatalf("expected ConvertKeyToNotation to return key %v: got %v", tc.expected, key)
			}
		})
	}
}

func TestConvertInvalidKey(t *testing.T) {
	_, err := tonality.ConvertKeyToNotation("invalid_key", tonality.Beatport)
	if err == nil {
		t.Fatal("expected ConvertKeyToNotation to return an error: got nil", err)
	}
}
