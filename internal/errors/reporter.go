package errors

import (
	"fmt"
	"os"
)

// ReportError reports an error to the centralized system.
func ReportError(err error) {
	if err == nil {
		return
	}
	// In a real system, this would go to Sentry or a log file.
	// For now, we log to Stderr to avoid circular dependency with logger.
	fmt.Fprintf(os.Stderr, "Internal Error: %v\n", err)
}
