# Quickstart

Run an ORCHADYN Planning API service, then submit the included procurement
planning request:

```sh
curl --request POST http://localhost:8080/plans:generate \
  --header "Content-Type: application/json" \
  --data @examples/procurement-plan.json
```

The response contains a candidate plan, deterministic validation results,
evidence tied to the planning inputs, and a PRAXOVELA reference projection.

## Go SDK

The Go SDK has no runtime dependencies. From `sdk/go`:

```sh
go test ./...
```

Create a client against the API endpoint and submit a typed `PlanningRequest`:

```go
client, err := orchadyn.NewClient("http://localhost:8080", nil)
if err != nil {
    return err
}

var result map[string]any
err = client.Generate(ctx, request, &result)
```

The SDK returns `*orchadyn.APIError` for non-2xx API responses, retaining the
HTTP status code and response body for policy or verification failures.
