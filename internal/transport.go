/*
Package internal unimportable
*/
package internal // import "github.com/1Volk/vksdk/internal"

import (
	"context"
	"net/http"

	"github.com/1Volk/vksdk"
)

// UserAgent module.
const UserAgent = "vksdk/" + vksdk.Version + " (+https://github.com/SevereCloud/vksdk)"

// ContextKey is just an empty struct. It exists so HTTPClient can be
// an immutable public variable with a unique type. It's immutable
// because nobody else can create a ContextKey, being unexported.
type ContextKey int

// ContextKey list.
const (
	HTTPClientKey ContextKey = iota
	UserAgentKey
)

// ContextClient return *http.Client.
func ContextClient(ctx context.Context) *http.Client {
	if ctx != nil {
		if hc, ok := ctx.Value(HTTPClientKey).(*http.Client); ok {
			return hc
		}
	}

	return http.DefaultClient
}

// ContextUserAgent return User-Agent from context.
func ContextUserAgent(ctx context.Context) string {
	if ctx != nil {
		if hc, ok := ctx.Value(UserAgentKey).(string); ok {
			return hc
		}
	}

	return UserAgent
}

// DoRequest sends an HTTP request and returns an HTTP response.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func DoRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	client := ContextClient(ctx)

	if req.UserAgent() == "" {
		req.Header.Set("User-Agent", ContextUserAgent(ctx))
	}

	resp, err := client.Do(req.WithContext(ctx))
	// If we got an error, and the context has been canceled,
	// the context's error is probably more useful.
	if err != nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}
	}

	return resp, err
}
