# obfs4PortScan
This service lets bridge operators test if their bridge's obfs4 port is
publicly reachable.

## Command line arguments arguments
The tool takes as input two command line arguments: a path to a certificate
file (specified by the argument `-cert-file`) and a path to its key file
(specified by the argument `-key-file`), both in PEM format.  We use these
files to run the HTTPS server. An optional third argument (`-addr`) can be used
to specify the address and port to listen on.

## Scanning method
We try to establish a TCP connection with the given IP address and port using
golang's `net.DialTimeout` function.  If we don't get a response within three
seconds, we deem the port unreachable.  We also deem the port unreachable if we
get a RST segment before the timeout.  In both cases, we display the error
message that we got from `net.DialTimeout`.

We implement a simple rate limiter that limits incoming requests to an average
of one per second with bursts of as many as five requests per second.

## Deployment
First, compile the binary:

    go build

Then, shut down the obfs4PortScan service on BridgeDB which runs under the
bridgescan user:

    systemctl --user stop obfs4portscan.service

Then, copy the binary onto BridgeDB's host.  It belongs into the directory
`/home/bridgescan/bin/`.  Once it's there, restart the service:

    systemctl --user start obfs4portscan.service
