# League Ranking CLI

This command-line application calculates the ranking table for a league based on match results.

---

## Setup

This project is written in [Go](https://golang.org/).

### Requirements

- Go 1.24 installed
- A UNIX-like environment (e.g., macOS or Linux)

---

## Running the App

### Option 1: Run from source

```bash
go run main.go
```

```bash
Lions 3, Snakes 3
Tarantulas 1, FC Awesome 0
Lions 1, FC Awesome 1
Tarantulas 3, Snakes 1
Lions 4, Grouches 0
```

### Option 2: Run from source

For convenience, a precompiled binary (league-ranking) is included in this repository.

```bash
chmod +x league-ranking
./league-ranking
```

### 🧪 Running Tests

```bash
go test
```

## Disclaimer

I included the compiled binary (league-ranking) in the repository for convenience only to simplify execution in macOS environments.

In general, I would not commit compiled binaries to version control — they are platform-specific, large, and not source-controlled. In a production project, this file would be ignored via .gitignore.
