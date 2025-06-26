package util

import (
	"strings"
	"time"
)

// TimeFormatFromJSFormat converts a format string to a time format string
func TimeFormatFromJSFormat(jsFormat string) string {
	replacer := strings.NewReplacer(
		"YYYY", "2006", // Full year
		"YY", "06", // Two-digit year
		"MM", "01", // Month with leading zero (01-12)
		"M", "1", // Month without leading zero (1-12)
		"DD", "02", // Day of month with leading zero (01-31)
		"D", "2", // Day of month without leading zero (1-31)
		"HH", "15", // Hour (24-hour) with leading zero (00-23)
		"H", "3", // Hour (12-hour) without leading zero (1-12) - Note: Go's 12-hour uses "3" for 3 PM
		"hh", "03", // Hour (12-hour) with leading zero (01-12) - "03" for 03 PM
		"mm", "04", // Minute with leading zero (00-59)
		"ss", "05", // Second with leading zero (00-59)
		"SSS", "000", // Milliseconds (not directly supported as format, but common in JS, using 000 as placeholder)
		"A", "PM", // AM/PM (uppercase)
		"a", "pm", // am/pm (lowercase)
		"E", "MST", // Timezone abbreviation (e.g., MST, PST)
		"ZZ", "-0700", // Timezone offset (e.g., -0700)
		"Z", "-07:00", // Timezone offset with colon (e.g., -07:00)
	)
	return replacer.Replace(jsFormat)
}

// TimeFormatFromPHPFormat converts a PHP date format string to a Go time format string
// Reference: https://www.php.net/manual/en/datetime.format.php
func TimeFormatFromPHPFormat(phpFormat string) string {
	// Handle special format characters that need to be escaped
	specialFormats := map[string]string{
		"c": "2006-01-02T15:04:05-07:00",       // ISO 8601 format
		"r": "Mon, 02 Jan 2006 15:04:05 -0700", // RFC 2822 format
		"U": "1136185445",                     // Unix timestamp
	}

	// Check for special formats first
	if format, ok := specialFormats[phpFormat]; ok {
		return format
	}

	replacer := strings.NewReplacer(
		// Day
		"d", "02",          // Day of the month, 2 digits with leading zeros (01 to 31)
		"D", "Mon",         // A textual representation of a day, three letters (Mon through Sun)
		"j", "2",           // Day of the month without leading zeros (1 to 31)
		"l", "Monday",      // A full textual representation of the day of the week (Sunday through Saturday)
		"N", "1",           // ISO-8601 numeric representation of the day of the week (1 for Monday, 7 for Sunday)
		"S", "2nd",         // English ordinal suffix for the day of the month, 2 characters (st, nd, rd, th)
		"w", "0",           // Numeric representation of the day of the week (0 for Sunday, 6 for Saturday)
		"z", "002",         // The day of the year (starting from 0, 0 through 365)

		// Week
		"W", "42",          // ISO-8601 week number of year, weeks starting on Monday (42 for the 42nd week in the year)


		// Month
		"F", "January",     // A full textual representation of a month (January through December)
		"m", "01",          // Numeric representation of a month, with leading zeros (01 through 12)
		"M", "Jan",         // A short textual representation of a month, three letters (Jan through Dec)
		"n", "1",           // Numeric representation of a month, without leading zeros (1 through 12)
		"t", "31",          // Number of days in the given month (28 through 31)


		// Year
		"L", "0",           // Whether it's a leap year (1 if it is a leap year, 0 otherwise)
		"o", "2006",        // ISO-8601 week-numbering year
		"Y", "2006",        // A full numeric representation of a year, 4 digits (Examples: 1999 or 2003)
		"y", "06",          // A two digit representation of a year (Examples: 99 or 03)


		// Time
		"a", "pm",          // Lowercase Ante meridiem and Post meridiem (am or pm)
		"A", "PM",          // Uppercase Ante meridiem and Post meridiem (AM or PM)
		"B", "999",         // Swatch Internet time (000 through 999)
		"g", "3",           // 12-hour format of an hour without leading zeros (1 through 12)
		"G", "15",          // 24-hour format of an hour without leading zeros (0 through 23)
		"h", "03",          // 12-hour format of an hour with leading zeros (01 through 12)
		"H", "15",          // 24-hour format of an hour with leading zeros (00 through 23)
		"i", "04",          // Minutes with leading zeros (00 to 59)
		"s", "05",          // Seconds with leading zeros (00 to 59)
		"u", "000000",      // Microseconds (up to six digits)
		"v", "000",         // Milliseconds (up to three digits)

		// Timezone - Note: In Go, timezone handling is different from PHP
		"e", "MST",         // Timezone identifier (UTC, GMT, Atlantic/Azores)
		"I", "0",           // Whether or not the date is in daylight saving time (1 if DST, 0 otherwise)
		"O", "-0700",       // Difference to GMT in hours and minutes (Example: +0200)
		"P", "-07:00",      // Same as O but with colon between hours and minutes (Example: +02:00)
		"p", "Z07:00",      // Same as P but returns Z instead of +00:00 (PHP 8.0.0+)
		"T", "MST",         // Timezone abbreviation (Examples: EST, MDT, +05)
		"Z", "-0700",       // Timezone offset in seconds
	)

	// Special handling for ISO 8601 format with timezone offset
	if strings.Contains(phpFormat, "\\T") {
		// Convert PHP's \T to Go's T without timezone abbreviation
		parts := strings.Split(phpFormat, "\\T")
		if len(parts) == 2 {
			datePart := replacer.Replace(parts[0])
			timePart := replacer.Replace(parts[1])
			// Remove any timezone abbreviation that might have been added
			timePart = strings.Replace(timePart, "MST", "", 1)
			return datePart + "T" + timePart
		}
	}

	return replacer.Replace(phpFormat)
}

func Timestamp() int64 {
	return time.Now().Unix()
}

func TimestampMilli() int64 {
	return time.Now().UnixMilli()
}

func TimestampMicro() int64 {
	return time.Now().UnixMicro()
}

func TimestampNano() int64 {
	return time.Now().UnixNano()
}
