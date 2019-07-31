package api

import (
	"errors"
	"net/http"

	"github.com/cenkalti/backoff"
)

// maxRetries defines the number of times an operation should be retried in the
// worst case.
const maxRetries = 5

// operation is a func that performs an API call.
type operation func() (map[string]interface{}, *http.Response, error)

// retry retries an operation using constant backoff. Since non-2xx status
// codes do not cause errors, we handle them separately to force retries for
// 5xx responses.
func retry(fn operation) (payload map[string]interface{}, resp *http.Response, err error) {
	err = backoff.Retry(func() error {
		payload, resp, err = fn()

		// Check the response code. We retry on 500-range responses to allow
		// the server time to recover, as 500's are typically not permanent
		// errors and may relate to outages on the server side. This will catch
		// invalid response codes as well, like 0 and 999.
		if resp != nil && (resp.StatusCode == 0 || resp.StatusCode >= http.StatusInternalServerError) {
			return errors.New(http.StatusText(resp.StatusCode))
		}

		return err
	}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), maxRetries))

	return
}
