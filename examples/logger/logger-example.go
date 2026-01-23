package main

import (
	"fmt"
	"time"

	"github.com/xelasov/golib/pkg/logger"
)

var log = logger.NewLogger("My App Logger")

func main() {
	fmt.Printf("===== see all 4 levels of messages ====\n")
	// all 4 log messages should show up
	logAllLevels()

	// now just 3 messages should show up
	fmt.Printf("===== Disable Debug Output     ====\n")
	fmt.Printf("===== see 3 levels of messages ====\n")
	logger.SetDebug(false)
	logAllLevels()

	// now just 2 messages should show up
	fmt.Printf("===== Disable Info Output     ====\n")
	fmt.Printf("===== see 2 levels of messages ====\n")
	logger.SetInfo(false)
	logAllLevels()

	// now just 1 messages should show up
	fmt.Printf("===== Disable Warning Output   ====\n")
	fmt.Printf("===== see 1 levels of messages ====\n")
	logger.SetWarning(false)
	logAllLevels()

	// now 0 messages should show up
	fmt.Printf("===== Disable Error Output   ====\n")
	fmt.Printf("===== see 0 levels of messages ====\n")
	logger.SetError(false)
	logAllLevels()

	fmt.Printf("===== The End =====\n")
}

func logAllLevels() {
	log.Debug("--- Simple Debug Message ---")
	log.Info("--- Simple Info Message ---")
	log.Warning("--- Simple Warning Message ---")
	log.Error("--- Simple Error Message ---")

	log.Debug("--- My Debug message with parameters -> number %d, at time: %v ---", 1, time.Now())
	log.Info("--- My Info message with parameters -> number %d, at time: %v ---", 2, time.Now())
	log.Warning("--- My Warning message with parameters -> number %d, at time: %v ---", 3, time.Now())
	log.Error("--- My Error message with parameters -> number %d, at time: %v ---", 4, time.Now())
}
