package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/xlab/suplog"
	"google.golang.org/grpc"
)

// readEnv is a special utility that reads `.env` file into actual environment variables
// of the current app, similar to `dotenv` Node package.
func readEnv() {
	if envdata, _ := os.ReadFile(".env"); len(envdata) > 0 {
		s := bufio.NewScanner(bytes.NewReader(envdata))
		for s.Scan() {
			parts := strings.Split(s.Text(), "=")
			if len(parts) != 2 {
				continue
			}
			strValue := strings.Trim(parts[1], `"`)
			if err := os.Setenv(parts[0], strValue); err != nil {
				log.WithField("name", parts[0]).WithError(err).Warningln("failed to override ENV variable")
			}
		}
	}
}

// logLevel converts vague log level name into typed level.
func logLevel(s string) log.Level {
	switch s {
	case "1", "error":
		return log.ErrorLevel
	case "2", "warn":
		return log.WarnLevel
	case "3", "info":
		return log.InfoLevel
	case "4", "debug":
		return log.DebugLevel
	default:
		return log.FatalLevel
	}
}

// toBool is used to parse vague bool definition into typed bool.
func toBool(s string) bool {
	switch strings.ToLower(s) {
	case "true", "1", "t", "yes":
		return true
	default:
		return false
	}
}

// duration parses duration from string with a provided default fallback.
func duration(s string, defaults time.Duration) time.Duration {
	dur, err := time.ParseDuration(s)
	if err != nil {
		dur = defaults
	}
	return dur
}

// checkStatsdPrefix ensures that the statsd prefix really
// have "." at end.
func checkStatsdPrefix(s string) string {
	if !strings.HasSuffix(s, ".") {
		return s + "."
	}
	return s
}

// orShutdown fatals the app if there was an error.
func orShutdown(err error) {
	if err != nil && err != grpc.ErrServerStopped {
		log.WithError(err).Fatalln("unable to start hyperion")
	}
}

func formatRPCs(input string) map[string][]string {
	pairs := strings.Split(input, ",")
	result := make(map[string][]string)

	for _, pair := range pairs {
		parts := strings.SplitN(pair, ":", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			result[key] = append(result[key], value)
		}
	}

	return result
}

func formatHyperionIds(input string) map[int]int {
	pairs := strings.Split(input, ",")
	result := make(map[int]int)

	for _, pair := range pairs {
		parts := strings.SplitN(pair, ":", 2)
		if len(parts) == 2 {
			key, err := strconv.Atoi(parts[0])
			if err != nil {
				continue
			}
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				continue
			}
			result[key] = value
		}
	}

	return result
}
