package main

import (
	"errors"
	"os"

	"github.com/hanzhang418/spark"
	"github.com/hanzhang418/spark/internal/config"
)

func main() {
	// Example 1: Using the default logger
	spark.Info("This is an info message using default logger")
	spark.Warn("This is a warning message", spark.String("component", "example"))
	spark.Error("This is an error message", spark.Err(errors.New("example error")))

	// Example 2: Creating a custom logger with JSON format
	jsonLogger, err := spark.NewLogger(
		config.WithFormat(config.JSONFormat),
		config.WithLevel(config.DebugLevel),
	)
	if err != nil {
		panic(err)
	}

	jsonLogger.Debug("Debug message in JSON format")
	jsonLogger.Info("Info message with fields",
		spark.String("user", "john"),
		spark.Int("age", 30),
		spark.Bool("active", true),
	)

	// Example 3: Using context logger with fields
	contextLogger := spark.With(
		spark.String("service", "user-service"),
		spark.String("version", "1.0.0"),
	)

	contextLogger.Info("Processing user request",
		spark.String("user_id", "12345"),
		spark.String("action", "login"),
	)

	// Example 4: Creating a logger with file output
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileLogger, err := spark.NewLogger(
		config.WithOutput(file),
		config.WithFormat(config.JSONFormat),
		config.WithCaller(true),
	)
	if err != nil {
		panic(err)
	}

	fileLogger.Info("This message will be written to app.log",
		spark.String("destination", "file"),
		spark.Int("line_number", 42),
	)

	// Example 5: Different log levels
	logger := spark.MustNewLogger(
		config.WithLevel(config.InfoLevel),
		config.WithFormat(config.ConsoleFormat),
	)

	logger.Debug("This debug message won't be shown (level is Info)")
	logger.Info("This info message will be shown")
	logger.Warn("This warning message will be shown")
	logger.Error("This error message will be shown")

	// Example 6: Using structured logging for better observability
	userLogger := spark.With(
		spark.String("module", "user"),
		spark.String("function", "CreateUser"),
	)

	userLogger.Info("Starting user creation process",
		spark.String("email", "user@example.com"),
		spark.String("role", "admin"),
	)

	// Simulate some processing
	userLogger.Debug("Validating user input")
	userLogger.Debug("Checking if user exists")

	userLogger.Info("User created successfully",
		spark.String("user_id", "user_123"),
		spark.Int64("created_at", 1640995200),
	)

	// Don't forget to sync the logger to ensure all logs are flushed
	if err := spark.Sync(); err != nil {
		// Handle sync error if needed
		spark.Error("Failed to sync logger", spark.Err(err))
	}
}
