# Quickstart

Run an ORCHADYN Planning API service, then submit the included procurement
planning request:

```sh
curl --request POST http://localhost:1816/plans:generate \
  --header "Content-Type: application/json" \
  --data @examples/procurement-plan.json
```

The response contains a candidate plan, deterministic validation results,
evidence tied to the planning inputs, and a PRAXOVELA reference projection.

The [`examples/`](../examples/) directory also includes
[`scheduling-plan.json`](../examples/scheduling-plan.json), which demonstrates
deadline, availability, and SLA constraints together with goal decomposition:

```sh
curl --request POST http://localhost:1816/plans:generate \
  --header "Content-Type: application/json" \
  --data @examples/scheduling-plan.json
```

## Go SDK

The Go SDK has no runtime dependencies. From `sdk/go`:

```sh
go test ./...
```

Create a client against the API endpoint and submit a typed `PlanningRequest`:

```go
client, err := orchadyn.NewClient("http://localhost:1816", nil)
if err != nil {
    return err
}

var result map[string]any
err = client.Generate(ctx, request, &result)
```

The SDK returns `*orchadyn.APIError` for non-2xx API responses, retaining the
HTTP status code and response body for policy or verification failures.

## Plan lifecycle

Generated plans are governed through `draft`, `approved`, and `executing`
states. Transition a stored plan with `SetPlanState`, and read the recorded
`plan.state_changed` events with `ListPlanEvents`:

```go
var updated map[string]any
err = client.SetPlanState(ctx, planID, orchadyn.SetStateRequest{State: "approved"}, &updated)

var events []orchadyn.PlanEvent
err = client.ListPlanEvents(ctx, planID, &events)
```

## Evidence

Every verification result carries an `attestations` chain that binds the plan
digest to its immutable planning-input digest. Fetch the stored plan package to
inspect the evidence:

```go
var packageJSON map[string]any
err = client.GetPlan(ctx, planID, &packageJSON)
```

Use `ExplainPlan` to read only the recorded candidates, validations, and
revision evidence for a plan.
