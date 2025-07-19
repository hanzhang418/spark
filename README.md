# Spark

A convenient and easy-to-use Go logging library based on [Uber's Zap](https://github.com/uber-go/zap). Spark provides a simple, structured logging interface while maintaining the high performance of Zap.

## Features

- 🚀 **High Performance**: Built on top of Uber's Zap logger
- 📝 **Structured Logging**: Support for structured fields and JSON output
- 🎨 **Multiple Formats**: Console (colored) and JSON output formats
- 📊 **Multiple Log Levels**: Debug, Info, Warn, Error, Fatal
- 🔧 **Configurable**: Flexible configuration options
- 🌍 **Global Logger**: Convenient global logger for simple use cases
- 🧵 **Thread Safe**: Safe for concurrent use
- 📁 **File Output**: Support for logging to files

## Installation

```bash
go get github.com/hanzhang418/spark
```

## Quick Start

### Basic Usage

```go
package main

import (
    "github.com/hanzhang418/spark"
)

func main() {
    // Simple logging
    spark.Info("Hello, World!")
    spark.Warn("This is a warning", spark.String("component", "main"))
    spark.Error("An error occurred", spark.Err(err))

    // Don't forget to sync at the end
    defer spark.Sync()
}
```

### Custom Logger

```go
package main

import (
    "github.com/hanzhang418/spark"
    "github.com/hanzhang418/spark/internal/config"
)

func main() {
    // Create a custom logger with JSON format
    logger, err := spark.NewLogger(
        config.WithFormat(config.JSONFormat),
        config.WithLevel(config.DebugLevel),
    )
    if err != nil {
        panic(err)
    }

    logger.Info("Custom logger message",
        spark.String("user", "john"),
        spark.Int("age", 30),
    )
}
```

### Context Logger

```go
// Create a logger with context fields
contextLogger := spark.With(
    spark.String("service", "user-service"),
    spark.String("version", "1.0.0"),
)

contextLogger.Info("Processing request",
    spark.String("user_id", "12345"),
    spark.String("action", "login"),
)
```

## Configuration Options

Spark provides several configuration options:

- `WithLevel(level)` - Set the minimum log level
- `WithFormat(format)` - Set output format (JSON or Console)
- `WithOutput(writer)` - Set output destination
- `WithTimeFormat(format)` - Set time format
- `WithCaller(enabled)` - Enable/disable caller information
- `WithStacktrace(enabled)` - Enable/disable stacktrace

### Log Levels

- `config.DebugLevel` - Debug messages
- `config.InfoLevel` - Informational messages (default)
- `config.WarnLevel` - Warning messages
- `config.ErrorLevel` - Error messages
- `config.FatalLevel` - Fatal messages (exits program)

### Output Formats

- `config.ConsoleFormat` - Human-readable colored output (default)
- `config.JSONFormat` - Structured JSON output

## Field Types

Spark supports various field types for structured logging:

```go
spark.Info("User action",
    spark.String("name", "John Doe"),
    spark.Int("age", 30),
    spark.Int64("timestamp", 1640995200),
    spark.Float64("score", 95.5),
    spark.Bool("active", true),
    spark.Err(err),
    spark.Any("metadata", map[string]interface{}{"key": "value"}),
)
```

## Examples

See the [examples](./examples) directory for more detailed usage examples.

## Performance

Spark inherits Zap's excellent performance characteristics:
- Zero allocation in most cases
- Structured logging without reflection
- Configurable sampling for high-volume logs

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
