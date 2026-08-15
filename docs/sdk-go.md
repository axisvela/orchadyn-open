# Go SDK

The dependency-free Go SDK in `sdk/go` covers every published ORCHADYN API
operation.

## Setup

```go
import (
	"context"
	"fmt"

	orchadyn "github.com/axisrobo/orchadyn-open/sdk/go"
)
```

Create a client and run the SDK tests from `sdk/go`:

```sh
go test ./...
```

## Planning

```go
client, err := orchadyn.NewClient("http://localhost:8080", nil)
if err != nil {
	return err
}

request := orchadyn.PlanningRequest{
	Goals: []orchadyn.Goal{{ID: "g1", Owner: "owner", Outcome: "approved", Deadline: "2026-09-30T00:00:00Z"}},
	Requirements: []orchadyn.CapabilityRequirement{{
		ID: "r1", GoalID: "g1", CapabilityType: "approval", RequiredEffect: "approved",
	}},
	Catalog: []orchadyn.Capability{{
		ID: "c1", Type: "approval", Version: "1.0.0", Effects: []string{"approved"},
		Region: "eu", Cost: 100, Availability: 99, SLA: "gold",
	}},
	Constraints: []orchadyn.Constraint{
		{Kind: "budget", Hard: true, Amount: 100},
		{Kind: "region", Hard: true, Value: "eu"},
		{Kind: "deadline", Hard: true, Value: "2026-09-30T00:00:00Z"},
		{Kind: "availability", Hard: true, Amount: 98},
		{Kind: "sla", Hard: true, Value: "gold"},
	},
}

var generated map[string]any
if err := client.Generate(ctx, request, &generated); err != nil {
	return err
}
```

## Verification and evidence

```go
var evidence map[string]any
if err := client.Verify(ctx, orchadyn.VerificationRequest{Plan: generated, Request: request}, &evidence); err != nil {
	return err
}
// evidence["attestations"] binds the plan and input digests.
```

## Lifecycle and execution tracking

```go
// Transition draft -> approved -> executing.
var updated map[string]any
if err := client.SetPlanState(ctx, planID, orchadyn.SetStateRequest{State: "approved"}, &updated); err != nil {
	return err
}

// Record per-node progress for execution traceability.
var recorded map[string]any
if err := client.RecordNodeProgress(ctx, planID, nodeID, orchadyn.NodeProgressRequest{Status: "started"}, &recorded); err != nil {
	return err
}

// Read the recorded events.
var events []orchadyn.PlanEvent
if err := client.ListPlanEvents(ctx, planID, &events); err != nil {
	return err
}
```

The SDK returns `*orchadyn.APIError` for non-2xx responses with the HTTP status
and body, so policy and verification failures are inspectable.
