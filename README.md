# go-tools

A comprehensive Go toolkit library providing utility functions for common tasks across multiple domains including datetime handling, number formatting, string operations, HTTP request/response handling, and logging.

## Table of Contents

- [Installation](#installation)
- [Features](#features)
  - [DateTime Tools](#datetime-tools)
  - [String Tools](#string-tools)
  - [Number Tools](#number-tools)
  - [Handler Tools](#handler-tools)
  - [General Tools](#general-tools)
  - [Zerolog Tools](#zerolog-tools)
- [Quick Start](#quick-start)

---

## Installation

Add this module to your Go project:

```bash
go get github.com/bagus-aulia/go-tools
```

---

## Features

### DateTime Tools

Located in `tools/datetime/`, provides functions to handle date and time formatting with Indonesian locale support.

#### Available Functions:

- **`IndonesiaDate(t time.Time) string`** - Converts time to Indonesian date format
- **`GenFullIndonesianDate(t time.Time) string`** - Generates full Indonesian date with day name
- **`GenIndonesianDateTime(t time.Time) string`** - Generates Indonesian date-time format

#### Example Usage:

```go
package main

import (
	"fmt"
	"time"
	"github.com/bagus-aulia/go-tools/tools/datetime"
)

func main() {
	now := time.Now()
	
	// Output: 17 September 2026
	fmt.Println(datetime.IndonesiaDate(now))
	
	// Output: Kamis, 17 September 2026
	fmt.Println(datetime.GenFullIndonesianDate(now))
	
	// Output: 17 September 2026, 14:30
	fmt.Println(datetime.GenIndonesianDateTime(now))
}
```

---

### String Tools

Located in `tools/string/`, provides string manipulation utilities.

#### Available Functions:

- **`RandString(n int) string`** - Generates a random alphanumeric string of length n
- **`Capitalize(s string) string`** - Capitalizes the first character of a string

#### Example Usage:

```go
package main

import (
	"fmt"
	"github.com/bagus-aulia/go-tools/tools/string"
)

func main() {
	// Generate random string of 10 characters
	// Output: aBcDeFgHiJ (random characters)
	fmt.Println(string.RandString(10))
	
	// Capitalize first character
	// Output: Hello world
	fmt.Println(string.Capitalize("hello world"))
	
	// Empty string remains empty
	fmt.Println(string.Capitalize(""))
}
```

---

### Number Tools

Located in `tools/number/`, provides number formatting and manipulation utilities.

#### Available Functions:

**Currency:**
- **`GenerateRupiah(value int, decimalPrecision int) string`** - Formats number as Indonesian Rupiah currency

**Number Operations:**
- **`AddTimestampFraction(value int, unixTimestamp int64, digit int) float64`** - Adds last N digits of timestamp as decimal fraction
- **`CountDigits(n int64) int`** - Counts the number of digits in a number
- **`RoundUp(number, dividedBy int) int`** - Rounds up division result
- **`OrdinalSuffix(n int) string`** - Converts number to ordinal form (1st, 2nd, 3rd, etc.)
- **`GenerateIndoNumFormat(num int64) string`** - Formats number with Indonesian style (dot separator)
- **`RoundFloat(val float64, precision uint) float64`** - Rounds float to specified decimal places

#### Example Usage:

```go
package main

import (
	"fmt"
	"github.com/bagus-aulia/go-tools/tools/number"
)

func main() {
	// Currency formatting
	// Output: Rp100.000,-
	fmt.Println(number.GenerateRupiah(100000, 0))
	// Output: Rp100.000,50
	fmt.Println(number.GenerateRupiah(100000, 2))
	
	// Number operations
	// Output: 4 (digits in 1234)
	fmt.Println(number.CountDigits(1234))
	
	// Round up division
	// Output: 4 (10 / 3 = 3.33 -> rounds up to 4)
	fmt.Println(number.RoundUp(10, 3))
	
	// Ordinal suffix
	// Output: 1st, 2nd, 3rd, 4th
	fmt.Println(number.OrdinalSuffix(1))  // 1st
	fmt.Println(number.OrdinalSuffix(2))  // 2nd
	fmt.Println(number.OrdinalSuffix(3))  // 3rd
	fmt.Println(number.OrdinalSuffix(4))  // 4th
	
	// Indonesian number format
	// Output: 1.000.000
	fmt.Println(number.GenerateIndoNumFormat(1000000))
	
	// Float rounding
	// Output: 3.14
	fmt.Println(number.RoundFloat(3.14159, 2))
}
```

---

### Handler Tools

Located in `tools/handler/`, provides HTTP request/response handling utilities.

#### 1. Response Builder (`tools/handler/response/`)

A fluent builder pattern for constructing HTTP responses with type-safe generics.

**Available Methods:**

- **`Factory[T any]() *Response[T]`** - Creates a new Response instance
- **`WithData(dt T) *Response[T]`** - Sets response data
- **`ErrorMessage(message string) *Response[T]`** - Sets error message
- **`ErrorReason(reason string) *Response[T]`** - Sets error reason

**HTTP Status Methods:**
- `Success()` - HTTP 200
- `Created()` - HTTP 201
- `BadRequest()` - HTTP 400
- `Unauthorized()` - HTTP 401
- `Forbidden()` - HTTP 403
- `NotFound()` - HTTP 404
- `Conflict()` - HTTP 409
- `EntityTooLarge()` - HTTP 413
- `UnprocessableEntity()` - HTTP 422
- `InternalServerError()` - HTTP 500
- `GeneralError(statusCode int)` - Custom status code

#### Example Usage:

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/bagus-aulia/go-tools/tools/handler/response"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// Success response with data
	user := User{ID: 1, Name: "John Doe"}
	res := response.Factory[User]().
		WithData(user).
		Success()
	
	data, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(data))
	// Output:
	// {
	//   "status": 200,
	//   "data": {
	//     "id": 1,
	//     "name": "John Doe"
	//   }
	// }
	
	// Error response
	errRes := response.Factory[User]().
		ErrorMessage("User not found").
		ErrorReason("not found").
		NotFound()
	
	errData, _ := json.MarshalIndent(errRes, "", "  ")
	fmt.Println(string(errData))
	// Output:
	// {
	//   "status": 404,
	//   "error": {
	//     "message": "Sorry, the feature that you're looking for is not found",
	//     "reason": "not found"
	//   }
	// }
}
```

#### 2. Request Parameter Parser (`tools/handler/mux/`)

- **`GetParamValue(r *http.Request, param string, variable string) (string, error)`** - Extracts parameter value from POST form data, URL query string, or pre-populated variable

#### Example Usage:

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/bagus-aulia/go-tools/tools/handler/mux"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Tries to get parameter in this order:
	// 1. From pre-populated variable (if provided)
	// 2. From POST form data
	// 3. From URL query string
	
	// Example: /search?q=golang
	value, err := mux.GetParamValue(r, "q", "")
	if err != nil {
		fmt.Println("Parameter not found:", err)
		return
	}
	
	fmt.Println("Search query:", value) // Output: golang
}
```

---

### General Tools

Located in `tools/general/`, provides general-purpose utility functions.

#### Available Functions:

- **`IndexOf[T comparable](data []T, element T) int`** - Generic function to find index of element in slice (returns -1 if not found)
- **`GenFileHeaderFaker() (*multipart.FileHeader, error)`** - Generates a fake file header for unit testing purposes

#### Example Usage:

```go
package main

import (
	"fmt"
	"github.com/bagus-aulia/go-tools/tools/general"
)

func main() {
	// Find index in slice
	numbers := []int{10, 20, 30, 40}
	
	// Output: 2
	fmt.Println(general.IndexOf(numbers, 30))
	
	// Output: -1 (not found)
	fmt.Println(general.IndexOf(numbers, 99))
	
	// Works with strings too
	names := []string{"Alice", "Bob", "Charlie"}
	
	// Output: 1
	fmt.Println(general.IndexOf(names, "Bob"))
	
	// For unit testing - generate fake file header
	fakeHeader, err := general.GenFileHeaderFaker()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Fake file:", fakeHeader.Filename) // Output: test.png
}
```

---

### Zerolog Tools

Located in `tools/zerolog/`, provides centralized logging configuration with log rotation support using zerolog.

#### Features:

- **Structured logging** with JSON output
- **Console and file output** support
- **Log rotation** with configurable size, backup count, and retention
- **Automatic compression** of rotated logs
- **Customizable log levels** (TRACE, DEBUG, INFO, WARN, ERROR, FATAL)
- **Stack trace support** for error messages
- **Caller information** tracking

#### Available Functions:

- **`DefaultLoggerConfig() LoggerConfig`** - Returns default logger configuration
- **`ConfigureZerolog(level string, enableConsole bool, logFile string)`** - Quick configuration with basic settings
- **`ConfigureZerologWithConfig(config LoggerConfig)`** - Full configuration with all options

#### LoggerConfig Struct:

```go
type LoggerConfig struct {
	Level         string // Log level: TRACE, DEBUG, INFO, WARN, ERROR, FATAL
	EnableConsole bool   // Output to console
	LogFile       string // Log file path
	TimeFormat    string // Time format (e.g., time.RFC822)
	EnableCaller  bool   // Include caller information
	EnableStack   bool   // Include stack trace for errors
	MaxSize       int    // Maximum file size in MB before rotation (default: 100)
	MaxBackups    int    // Number of old log files to keep (default: 3)
	MaxAge        int    // Delete files older than N days (default: 30)
	Compress      bool   // Compress rotated files (default: true)
}
```

#### Example Usage:

```go
package main

import (
	"github.com/rs/zerolog/log"
	"github.com/bagus-aulia/go-tools/tools/zerolog"
)

func main() {
	// Quick setup with defaults
	zerolog.ConfigureZerolog("INFO", true, "app.log")
	
	// Log messages
	log.Info().Msg("Application started")
	log.Warn().Msg("This is a warning")
	log.Error().Err(nil).Msg("An error occurred")
	
	// With custom configuration
	config := zerolog.LoggerConfig{
		Level:         "DEBUG",
		EnableConsole: true,
		LogFile:       "logs/app.log",
		TimeFormat:    "2006-01-02 15:04:05",
		EnableCaller:  true,
		EnableStack:   true,
		MaxSize:       50,      // 50 MB
		MaxBackups:    5,       // Keep 5 old files
		MaxAge:        7,       // Keep for 7 days
		Compress:      true,    // Compress rotated logs
	}
	
	zerolog.ConfigureZerologWithConfig(config)
	
	log.Debug().Str("user", "john").Msg("User logged in")
	log.Error().Err(nil).Stack().Msg("Unrecoverable error")
}
```

**Log Rotation Behavior:**

When log file reaches configured size (e.g., 50 MB):
1. Current log file is renamed with timestamp (e.g., `app.log.2026-09-17T14-30-00-000`)
2. A new `app.log` file is created
3. Old backup files are deleted if count exceeds `MaxBackups`
4. Files older than `MaxAge` days are automatically deleted
5. Rotated files are compressed (`.gz` extension) if `Compress: true`

---

## Quick Start

Here's a minimal example using multiple tools:

```go
package main

import (
	"fmt"
	"time"
	"github.com/rs/zerolog/log"
	"github.com/bagus-aulia/go-tools/tools/datetime"
	"github.com/bagus-aulia/go-tools/tools/number"
	"github.com/bagus-aulia/go-tools/tools/string"
	"github.com/bagus-aulia/go-tools/tools/zerolog"
)

func main() {
	// Setup logging
	zerolog.ConfigureZerolog("INFO", true, "app.log")
	
	// Use datetime
	now := time.Now()
	log.Info().Msg(datetime.GenFullIndonesianDate(now))
	
	// Use number formatting
	amount := 500000
	log.Info().Msg(number.GenerateRupiah(amount, 0))
	
	// Use string utilities
	randomID := string.RandString(8)
	log.Info().Str("id", randomID).Msg("Generated new ID")
}
```

---

## License

This project is open source and available for use in your Go applications.
