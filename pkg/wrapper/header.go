package wrapper

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func CustomMatcher(key string) (string, bool) {
	switch key {
	case "X-Platform-Hash":
		return key, true
	case "X-Admin-User-Id":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}
