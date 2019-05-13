package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"net"
	"net/http"
	"time"
)

// limiter implements a rate limiter.  We allow 1 request per second on average
// with bursts of up to 5 requests per second.
var limiter = rate.NewLimiter(1, 5)

// Index returns the landing page (which contains the scanning form) to the
// user.
func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, IndexPage)
}

// SendResponse sends the given response to the user.
func SendResponse(w http.ResponseWriter, response string) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, response)
}

// ScanDestination extracts the given IP address and port from the POST
// parameters, triggers TCP scanning of the tuple, and returns the result.
func ScanDestination(w http.ResponseWriter, r *http.Request) {

	// We implement rate limiting to prevent someone from abusing this service
	// as a port scanner.
	if limiter.Allow() == false {
		SendResponse(w, "Rate limit exceeded.  Wait for a bit.")
		return
	}

	// The number of seconds we're willing to wait until we decide that the
	// given destination is offline.
	timeout, _ := time.ParseDuration("3s")

	r.ParseForm()
	// These variables will be "" if they're not set.
	address := r.Form.Get("address")
	port := r.Form.Get("port")

	if address == "" {
		SendResponse(w, "No address given.")
		return
	}
	if port == "" {
		SendResponse(w, "No port given.")
		return
	}

	portReachable, err := IsTCPPortReachable(address, port, timeout)
	if portReachable {
		SendResponse(w, SuccessPage)
	} else {
		SendResponse(w, FailurePage(err))
	}
}

// IsTCPPortReachable returns `true' if it can establish a TCP connection with
// the given IP address and port.  If not, it returns `false' and the
// respective error, as reported by `net.DialTimeout'.
func IsTCPPortReachable(addr, port string, timeout time.Duration) (bool, error) {

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", addr, port), timeout)
	if err != nil {
		return false, err
	}
	conn.Close()
	return true, nil
}
