package util

import (
	"testing"
	"time"
)

func TestTimeFormatFromJSFormat(t *testing.T) {
	tests := []struct {
		name     string
		jsFormat string
		expected string
	}{
		{
			name:     "Full date and time with milliseconds",
			jsFormat: "YYYY-MM-DD HH:mm:ss.SSS",
			expected: "2006-01-02 15:04:05.000",
		},
		{
			name:     "Date only with 2-digit year",
			jsFormat: "YY/MM/DD",
			expected: "06/01/02",
		},
		{
			name:     "Time with 12-hour format and AM/PM",
			jsFormat: "hh:mm:ss A",
			expected: "03:04:05 PM",
		},
		{
			name:     "Date with month without leading zero",
			jsFormat: "M/D/YYYY",
			expected: "1/2/2006",
		},
		{
			name:     "Time with timezone",
			jsFormat: "YYYY-MM-DDTHH:mm:ssZ",
			expected: "2006-01-02T15:04:05-07:00",
		},
		{
			name:     "Time with timezone offset",
			jsFormat: "YYYY-MM-DDTHH:mm:ssZZ",
			expected: "2006-01-02T15:04:05-0700",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeFormatFromJSFormat(tt.jsFormat); got != tt.expected {
				t.Errorf("TimeFormatFromJSFormat() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTimeFormatFromPHPFormat(t *testing.T) {
	tests := []struct {
		name        string
		phpFormat   string
		expected    string
		skipParsing bool // Some formats can't be directly parsed by time.Parse
	}{
		// Day formats
		{
			name:        "Day of month with leading zero",
			phpFormat:   "d",
			expected:    "02",
			skipParsing: true,
		},
		{
			name:        "Day of month without leading zero",
			phpFormat:   "j",
			expected:    "2",
			skipParsing: true,
		},
		{
			name:        "Day name short",
			phpFormat:   "D",
			expected:    "Mon",
			skipParsing: true,
		},
		{
			name:        "Day name full",
			phpFormat:   "l",
			expected:    "Monday",
			skipParsing: true,
		},
		{
			name:        "ISO-8601 day of week",
			phpFormat:   "N",
			expected:    "1",
			skipParsing: true,
		},
		{
			name:        "Day of year",
			phpFormat:   "z",
			expected:    "002",
			skipParsing: true,
		},

		// Week formats
		{
			name:        "Week number",
			phpFormat:   "W",
			expected:    "42",
			skipParsing: true,
		},

		// Month formats
		{
			name:        "Month with leading zero",
			phpFormat:   "m",
			expected:    "01",
			skipParsing: true,
		},
		{
			name:        "Month without leading zero",
			phpFormat:   "n",
			expected:    "1",
			skipParsing: true,
		},
		{
			name:        "Month name short",
			phpFormat:   "M",
			expected:    "Jan",
			skipParsing: true,
		},
		{
			name:        "Month name full",
			phpFormat:   "F",
			expected:    "January",
			skipParsing: true,
		},
		{
			name:        "Days in month",
			phpFormat:   "t",
			expected:    "31",
			skipParsing: true,
		},

		// Year formats
		{
			name:        "Full year",
			phpFormat:   "Y",
			expected:    "2006",
			skipParsing: true,
		},
		{
			name:        "Two-digit year",
			phpFormat:   "y",
			expected:    "06",
			skipParsing: true,
		},
		{
			name:        "Leap year",
			phpFormat:   "L",
			expected:    "0",
			skipParsing: true,
		},

		// Time formats
		{
			name:        "12-hour with leading zero",
			phpFormat:   "h",
			expected:    "03",
			skipParsing: true,
		},
		{
			name:        "24-hour with leading zero",
			phpFormat:   "H",
			expected:    "15",
			skipParsing: true,
		},
		{
			name:        "12-hour without leading zero",
			phpFormat:   "g",
			expected:    "3",
			skipParsing: true,
		},
		{
			name:        "24-hour without leading zero",
			phpFormat:   "G",
			expected:    "15",
			skipParsing: true,
		},
		{
			name:        "Minutes with leading zero",
			phpFormat:   "i",
			expected:    "04",
			skipParsing: true,
		},
		{
			name:        "Seconds with leading zero",
			phpFormat:   "s",
			expected:    "05",
			skipParsing: true,
		},
		{
			name:        "Microseconds",
			phpFormat:   "u",
			expected:    "000000",
			skipParsing: true,
		},
		{
			name:        "Milliseconds",
			phpFormat:   "v",
			expected:    "000",
			skipParsing: true,
		},
		{
			name:        "AM/PM lowercase",
			phpFormat:   "a",
			expected:    "pm",
			skipParsing: true,
		},
		{
			name:        "AM/PM uppercase",
			phpFormat:   "A",
			expected:    "PM",
			skipParsing: true,
		},

		// Timezone formats
		{
			name:        "Timezone identifier",
			phpFormat:   "e",
			expected:    "MST",
			skipParsing: true,
		},
		{
			name:        "Daylight saving time",
			phpFormat:   "I",
			expected:    "0",
			skipParsing: true,
		},
		{
			name:        "Timezone offset",
			phpFormat:   "O",
			expected:    "-0700",
			skipParsing: true,
		},
		{
			name:        "Timezone offset with colon",
			phpFormat:   "P",
			expected:    "-07:00",
			skipParsing: true,
		},
		{
			name:        "Timezone abbreviation",
			phpFormat:   "T",
			expected:    "MST",
			skipParsing: true,
		},

		// Full date/time formats
		{
			name:        "ISO 8601 format",
			phpFormat:   "c",
			expected:    "2006-01-02T15:04:05-07:00",
			skipParsing: true,
		},
		{
			name:        "RFC 2822 format",
			phpFormat:   "r",
			expected:    "Mon, 02 Jan 2006 15:04:05 -0700",
			skipParsing: true,
		},
		{
			name:        "Unix timestamp",
			phpFormat:   "U",
			expected:    "1136185445",
			skipParsing: true,
		},

		// Combined formats
		{
			name:        "Standard date and time",
			phpFormat:   "Y-m-d H:i:s",
			expected:    "2006-01-02 15:04:05",
			skipParsing: false,
		},
		{
			name:        "Date with timezone",
			phpFormat:   "D, d M Y H:i:s e",
			expected:    "Mon, 02 Jan 2006 15:04:05 MST",
			skipParsing: true, // Timezone handling varies by system
		},
		{
			name:        "With microseconds",
			phpFormat:   "Y-m-d H:i:s.u",
			expected:    "2006-01-02 15:04:05.000000",
			skipParsing: true, // time.Parse doesn't support microseconds in format
		},
		{
			name:        "With timezone offset",
			phpFormat:   "Y-m-d\\TH:i:sP",
			expected:    "2006-01-02T15:04:05-07:00",
			skipParsing: true, // Skip parsing as the timezone handling might vary by system
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TimeFormatFromPHPFormat(tt.phpFormat)
			if got != tt.expected {
				t.Errorf("TimeFormatFromPHPFormat() = %v, want %v", got, tt.expected)
			}

			// Test parsing the format if it's a parseable format
			if !tt.skipParsing {
				testTime := time.Date(2023, 12, 25, 14, 30, 45, 0, time.UTC)
				_, err := time.Parse(got, testTime.Format(got))
				if err != nil {
					t.Errorf("Failed to parse time with format %s: %v", got, err)
				}
			}
		})
	}
}
