package main

import (
	"reflect"
	"testing"
)

func TestParseMatches_ValidInput(t *testing.T) {
	lines := []string{
		"Lions 3, Snakes 3",
		"Tarantulas 1, FC Awesome 0",
	}
	expected := []MatchResult{
		{"Lions", 3, "Snakes", 3},
		{"Tarantulas", 1, "FC Awesome", 0},
	}
	results, err := ParseMatches(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("ParseMatches = %+v, expected %+v", results, expected)
	}
}

func TestParseMatches_InvalidLines(t *testing.T) {
	lines := []string{
		"Lions 3, Snakes 3",
		"BadFormatLine",
		"AnotherBadInput 2-1",
		"Tarantulas 2, Grouches 0",
	}
	expected := []MatchResult{
		{"Lions", 3, "Snakes", 3},
		{"Tarantulas", 2, "Grouches", 0},
	}
	var parsed []MatchResult
	for _, line := range lines {
		match, err := parseLine(line)
		if err != nil {
			continue
		}
		parsed = append(parsed, match)
	}
	if !reflect.DeepEqual(parsed, expected) {
		t.Errorf("parseLine loop = %+v, expected %+v", parsed, expected)
	}
}

func TestUpdateScores(t *testing.T) {
	matches := []MatchResult{
		{"Lions", 3, "Snakes", 3},
		{"Tarantulas", 1, "FC Awesome", 0},
		{"Lions", 1, "FC Awesome", 1},
		{"Tarantulas", 3, "Snakes", 1},
		{"Lions", 4, "Grouches", 0},
	}
	expected := map[string]int{
		"Tarantulas": 6,
		"Lions":      5,
		"FC Awesome": 1,
		"Snakes":     1,
		"Grouches":   0,
	}
	scores := UpdateScores(matches)

	if !reflect.DeepEqual(scores, expected) {
		t.Errorf("ComputeScores = %+v, expected %+v", scores, expected)
	}
}

func TestFormatLeaderboard(t *testing.T) {
	scores := map[string]int{
		"Tarantulas": 6,
		"Lions":      5,
		"Snakes":     1,
		"FC Awesome": 1,
		"Grouches":   0,
	}
	expected := []string{
		"1. Tarantulas, 6 pts",
		"2. Lions, 5 pts",
		"3. FC Awesome, 1 pt",
		"4. Snakes, 1 pt",
		"5. Grouches, 0 pts",
	}

	output := FormatLeaderboard(scores)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("FormatLeaderboard = %+v, expected %+v", output, expected)
	}
}
