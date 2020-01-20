package main

import (
	"fmt"
)

// IndexPage contains the landing page of the service.  It implements a simple
// form that asks for an IP address and TCP port.
var IndexPage = `<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width">
  <title>Test your obfs4 bridge&rsquo;s TCP port</title>
</head>

<body>
  <form method="GET" action="scan">
    <h2>TCP reachability test</h2>
    <p>This service allows you to test if your obfs4 bridge port is reachable
    to the rest of the world.</p>
    <p>Enter your bridge&rsquo;s IP address (enclose IPv6 addresses in square
    brackets) and obfs4 port, and click
    &ldquo;Scan&rdquo;.  The service will then try to establish a TCP
    connection with your bridge and tell you if it succeeded.</p>
    <input type="text" required name="address" placeholder="IP address">
    <input type="text" required name="port" placeholder="Obfs4 port">
    <label ></label>
    <button type="submit">Scan</button>
  </form>
</body>

</html>
`

// SuccessPage is shown when the given address and port are reachable.
var SuccessPage = `<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width">
  <title>Success!</title>
</head>

<body>
  <div align='center'>
    <h2 style='color:green'>TCP port is reachable!</h2>
  </div>
</body>

</html>
`

// FailurePage2 is shown when the given address and port are unreachable.
func FailurePage(reason error) string {

	var failurePage1 = `<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width">
  <title>Failure!</title>
</head>

<body>
  <div align='center'>
    <h2 style='color:red'>TCP port is <i>not</i> reachable!</h2>
    <p>Here&rsquo;s the error message we are getting:</p>
    <tt>
`

	var failurePage2 = `</tt>
  </div>
</body>

</html>
`

	return fmt.Sprintf("%s%s%s", failurePage1, reason, failurePage2)
}
