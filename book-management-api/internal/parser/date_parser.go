package parser

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// DateTimeParser handles parsing various date/time formats
type DateTimeParser struct {
	formats []string
	mu      sync.RWMutex
}

var (
	instance *DateTimeParser
	once     sync.Once
)

// GetParser returns the singleton instance of DateTimeParser
func GetParser() *DateTimeParser {
	once.Do(func() {
		instance = &DateTimeParser{
			formats: []string{
				// RFC3339 and ISO 8601 variants
				time.RFC3339,
				time.RFC3339Nano,
				"2006-01-02T15:04:05Z07:00",
				"2006-01-02T15:04:05.000Z07:00",
				"2006-01-02T15:04:05Z",
				"2006-01-02T15:04:05.000Z",
				"2006-01-02T15:04:05",
				"2006-01-02T15:04:05.000",

				// Date only formats
				"2006-01-02",
				"01/02/2006",
				"1/2/2006",
				"02/01/2006", // DD/MM/YYYY
				"2/1/2006",   // D/M/YYYY
				"2006/01/02",
				"2006/1/2",
				"01-02-2006",
				"1-2-2006",
				"02-01-2006", // DD-MM-YYYY
				"2-1-2006",   // D-M-YYYY
				"2006-01-02",
				"Jan 2, 2006",
				"January 2, 2006",
				"Jan 02, 2006",
				"January 02, 2006",
				"2 Jan 2006",
				"02 Jan 2006",
				"2 January 2006",
				"02 January 2006",

				// Time with date formats
				"2006-01-02 15:04:05",
				"2006-01-02 15:04:05.000",
				"01/02/2006 15:04:05",
				"1/2/2006 15:04:05",
				"01/02/2006 3:04:05 PM",
				"1/2/2006 3:04:05 PM",
				"2006-01-02 3:04:05 PM",
				"Jan 2, 2006 15:04:05",
				"Jan 2, 2006 3:04:05 PM",
				"January 2, 2006 15:04:05",
				"January 2, 2006 3:04:05 PM",

				// Time only formats (will use current date)
				"15:04:05",
				"15:04:05.000",
				"3:04:05 PM",
				"3:04 PM",
				"15:04",

				// Unix timestamp handling is done separately
			},
		}
	})
	return instance
}

// AddFormat adds a custom format to the parser (thread-safe)
func (p *DateTimeParser) AddFormat(format string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.formats = append(p.formats, format)
}

// Parse attempts to parse a date/time string using various formats
func (p *DateTimeParser) Parse(dateStr string) (time.Time, error) {
	// Trim whitespace
	dateStr = strings.TrimSpace(dateStr)

	if dateStr == "" {
		return time.Time{}, fmt.Errorf("empty date string")
	}

	// Try to parse as Unix timestamp first
	if t, err := p.parseUnixTimestamp(dateStr); err == nil {
		return t, nil
	}

	// Get formats in a thread-safe way
	p.mu.RLock()
	formats := make([]string, len(p.formats))
	copy(formats, p.formats)
	p.mu.RUnlock()

	// Try each format
	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return t, nil
		}
	}

	// If no format worked, try with different locations
	locations := []*time.Location{
		time.UTC,
		time.Local,
	}

	for _, loc := range locations {
		for _, format := range formats {
			if t, err := time.ParseInLocation(format, dateStr, loc); err == nil {
				return t, nil
			}
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse date string: %s", dateStr)
}

// parseUnixTimestamp attempts to parse Unix timestamps
func (p *DateTimeParser) parseUnixTimestamp(dateStr string) (time.Time, error) {
	// Check if it looks like a Unix timestamp (all digits, possibly with decimal)
	if !isNumeric(dateStr) {
		return time.Time{}, fmt.Errorf("not a numeric timestamp")
	}

	// Try parsing as Unix timestamp (seconds)
	if len(dateStr) == 10 {
		var sec int64
		if _, err := fmt.Sscanf(dateStr, "%d", &sec); err == nil {
			return time.Unix(sec, 0), nil
		}
	}

	// Try parsing as Unix timestamp with milliseconds
	if len(dateStr) == 13 {
		var msec int64
		if _, err := fmt.Sscanf(dateStr, "%d", &msec); err == nil {
			return time.Unix(msec/1000, (msec%1000)*1000000), nil
		}
	}

	// Try parsing as Unix timestamp with microseconds
	if len(dateStr) == 16 {
		var usec int64
		if _, err := fmt.Sscanf(dateStr, "%d", &usec); err == nil {
			return time.Unix(usec/1000000, (usec%1000000)*1000), nil
		}
	}

	// Try parsing as Unix timestamp with nanoseconds
	if len(dateStr) == 19 {
		var nsec int64
		if _, err := fmt.Sscanf(dateStr, "%d", &nsec); err == nil {
			return time.Unix(nsec/1000000000, nsec%1000000000), nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid timestamp format")
}

// isNumeric checks if a string contains only digits and optionally one decimal point
func isNumeric(s string) bool {
	if s == "" {
		return false
	}

	decimalCount := 0
	for _, r := range s {
		if r == '.' {
			decimalCount++
			if decimalCount > 1 {
				return false
			}
		} else if r < '0' || r > '9' {
			return false
		}
	}

	return true
}

// ParseWithLocation parses a date/time string with a specific timezone location
func (p *DateTimeParser) ParseWithLocation(dateStr string, loc *time.Location) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)

	if dateStr == "" {
		return time.Time{}, fmt.Errorf("empty date string")
	}

	// Try to parse as Unix timestamp first
	if t, err := p.parseUnixTimestamp(dateStr); err == nil {
		return t.In(loc), nil
	}

	// Get formats in a thread-safe way
	p.mu.RLock()
	formats := make([]string, len(p.formats))
	copy(formats, p.formats)
	p.mu.RUnlock()

	// Try each format with the specified location
	for _, format := range formats {
		if t, err := time.ParseInLocation(format, dateStr, loc); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse date string: %s", dateStr)
}

// Convenience functions for direct usage

// Parse parses a date/time string using the singleton instance
func ParseDate(dateStr string) (time.Time, error) {
	return GetParser().Parse(dateStr)
}

// ParseWithLocation parses with a specific timezone using the singleton instance
func ParseDateWithLocation(dateStr string, loc *time.Location) (time.Time, error) {
	return GetParser().ParseWithLocation(dateStr, loc)
}

// AddFormat adds a custom format to the singleton parser
func AddFormat(format string) {
	GetParser().AddFormat(format)
}
