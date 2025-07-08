package libvirt

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ConfigureLogging sets up the standard log package to avoid interfering
// with Terraform's JSON output mode. All logging should go through tflog
// instead of the standard log package.
func ConfigureLogging() {
	// When TF_LOG is not set, disable standard logging entirely
	// to prevent corruption of JSON output
	if os.Getenv("TF_LOG") == "" {
		log.SetOutput(io.Discard)
	} else {
		// Even with TF_LOG set, we should use tflog for proper formatting
		// but keep standard logging for backwards compatibility
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		log.SetPrefix("[terraform-provider-libvirt] ")
	}
}

// LogDebug is a wrapper that uses tflog when available, falls back to log.Printf
func LogDebug(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	if ctx == nil {
		// Fallback to standard logging if no context
		if os.Getenv("TF_LOG") != "" {
			log.Printf("[DEBUG] %s", msg)
		}
		return
	}
	
	tflog.Debug(ctx, msg, additionalFields...)
}

// LogInfo is a wrapper that uses tflog when available, falls back to log.Printf
func LogInfo(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	if ctx == nil {
		// Fallback to standard logging if no context
		if os.Getenv("TF_LOG") != "" {
			log.Printf("[INFO] %s", msg)
		}
		return
	}
	
	tflog.Info(ctx, msg, additionalFields...)
}

// LogWarn is a wrapper that uses tflog when available, falls back to log.Printf
func LogWarn(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	if ctx == nil {
		// Fallback to standard logging if no context
		if os.Getenv("TF_LOG") != "" {
			log.Printf("[WARN] %s", msg)
		}
		return
	}
	
	tflog.Warn(ctx, msg, additionalFields...)
}

// LogError is a wrapper that uses tflog when available, falls back to log.Printf
func LogError(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	if ctx == nil {
		// Fallback to standard logging if no context
		if os.Getenv("TF_LOG") != "" {
			log.Printf("[ERROR] %s", msg)
		}
		return
	}
	
	tflog.Error(ctx, msg, additionalFields...)
}