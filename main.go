package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type MatchResult struct {
	TeamA  string
	ScoreA int
	TeamB  string
	ScoreB int
}

type TeamStats struct {
	Name   string
	Points int
}

func UpdateScores(results []MatchResult) map[string]int {
	scores := make(map[string]int)
	for _, r := range results {
		// Explicitly initialize teams if they don't exist in the map
		// This ensures that teams are initialized with 0 points if they haven't played yet
		if _, ok := scores[r.TeamA]; !ok {
			scores[r.TeamA] = 0
		}
		if _, ok := scores[r.TeamB]; !ok {
			scores[r.TeamB] = 0
		}

		if r.ScoreA > r.ScoreB {
			scores[r.TeamA] += 3
		} else if r.ScoreA < r.ScoreB {
			scores[r.TeamB] += 3
		} else {
			scores[r.TeamA] += 1
			scores[r.TeamB] += 1
		}
	}
	return scores
}

func parseTeamScore(s string) (string, int, error) {
	s = strings.TrimSpace(s)
	lastSpace := strings.LastIndex(s, " ")
	if lastSpace == -1 {
		return "", 0, fmt.Errorf("no space between name and score")
	}
	name := s[:lastSpace]
	score, err := strconv.Atoi(s[lastSpace+1:])
	if err != nil {
		return "", 0, err
	}
	return name, score, nil
}

func parseLine(line string) (MatchResult, error) {
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return MatchResult{}, fmt.Errorf("invalid line format")
	}

	teamA, scoreA, err := parseTeamScore(parts[0])
	if err != nil {
		return MatchResult{}, err
	}
	teamB, scoreB, err := parseTeamScore(parts[1])
	if err != nil {
		return MatchResult{}, err
	}

	return MatchResult{
		TeamA:  teamA,
		ScoreA: scoreA,
		TeamB:  teamB,
		ScoreB: scoreB,
	}, nil
}

func ParseMatches(lines []string) ([]MatchResult, error) {
	var results []MatchResult
	for _, line := range lines {
		if line == "" {
			continue
		}
		match, err := parseLine(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Skipping invalid line: %q (%v)\n", line, err)
			continue
		}
		results = append(results, match)
	}
	return results, nil
}

func FormatLeaderboard(scores map[string]int) []string {
	var stats []TeamStats
	for team, pts := range scores {
		stats = append(stats, TeamStats{team, pts})
	}
	sort.Slice(stats, func(i, j int) bool {
		if stats[i].Points == stats[j].Points {
			return stats[i].Name < stats[j].Name
		}
		return stats[i].Points > stats[j].Points
	})

	displayRank := 0
	var output []string
	for i, stat := range stats {
		displayRank = i + 1
		pointLabel := "pts"
		if stat.Points == 1 {
			pointLabel = "pt"
		}
		output = append(output, fmt.Sprintf("%d. %s, %d %s", displayRank, stat.Name, stat.Points, pointLabel))
	}
	return output
}

func main() {
	fmt.Println("Enter match results (Ctrl+D to end):")

	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return
	}

	matches, err := ParseMatches(lines)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Parse error:", err)
		return
	}

	scores := UpdateScores(matches)
	output := FormatLeaderboard(scores)

	for _, line := range output {
		fmt.Println(line)
	}
}
