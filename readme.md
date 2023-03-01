# otel-b3-header-propagation

My attempt to propagate headers.

# how to reproduce?

## first run the "request dumper" service:

```
[otel-b3-header-propagation] on main$ go run cmd/request-dumper/main.go

INFO[0000] listening on http://127.0.0.1:8765/dump-request
```

This will simply dump all headers from an incoming request on endpoint http://127.0.0.1:8765/dump-request

## then run the service itself

```
[otel-b3-header-propagation] on main$ go run cmd/service/main.go

INFO[0000] listening on http://127.0.0.1:8080
```

## then make request

Then try to make a request to the service and provide a header.

```
curl -H "X-B3-Traceid: test" localhost:8080/call-downstream/
```

### expected output

I would expect that the provided B3 header (`X-B3-Traceid`) will be propagated to the downstream service (request-dumper).
Unfortunately this is not the case.
