package orchadyn

import (
	"encoding/json"
	"testing"
)

func TestPlanNodeSubPlanRoundTrip(t *testing.T) {
	node := PlanNode{
		ID:            "n1",
		GoalID:        "g1",
		RequirementID: "r1",
		SubPlan: &SubPlanRef{
			PlanID: "child",
			Digest: "abc",
			Inputs: map[string]string{"region": "${region}"},
		},
	}

	raw, err := json.Marshal(node)
	if err != nil {
		t.Fatalf("marshal PlanNode: %v", err)
	}
	var encoded map[string]any
	if err := json.Unmarshal(raw, &encoded); err != nil {
		t.Fatalf("unmarshal encoded PlanNode: %v", err)
	}
	if _, ok := encoded["capabilityId"]; ok {
		t.Error("expected omitted capabilityId for a sub-plan node")
	}
	subPlan, ok := encoded["subPlan"].(map[string]any)
	if !ok {
		t.Fatalf("expected subPlan object, got %#v", encoded["subPlan"])
	}
	if subPlan["planId"] != "child" || subPlan["digest"] != "abc" {
		t.Errorf("unexpected subPlan encoding: %#v", subPlan)
	}

	var decoded PlanNode
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal PlanNode: %v", err)
	}
	if decoded.SubPlan == nil || decoded.SubPlan.PlanID != "child" {
		t.Fatalf("subPlan did not round-trip: %#v", decoded.SubPlan)
	}
	if decoded.SubPlan.Inputs["region"] != "${region}" {
		t.Errorf("subPlan inputs did not round-trip: %#v", decoded.SubPlan.Inputs)
	}
}

func TestPlanNodeCapabilityOmitsSubPlan(t *testing.T) {
	node := PlanNode{
		ID:                "n1",
		GoalID:            "g1",
		RequirementID:     "r1",
		CapabilityID:      "cap",
		CapabilityVersion: "1.0.0",
		ExpectedEffect:    "approved",
	}

	raw, err := json.Marshal(node)
	if err != nil {
		t.Fatalf("marshal PlanNode: %v", err)
	}
	var encoded map[string]any
	if err := json.Unmarshal(raw, &encoded); err != nil {
		t.Fatalf("unmarshal encoded PlanNode: %v", err)
	}
	if _, ok := encoded["subPlan"]; ok {
		t.Error("expected omitted subPlan for a capability node")
	}
	if encoded["capabilityId"] != "cap" {
		t.Errorf("unexpected capabilityId: %#v", encoded["capabilityId"])
	}
}

func TestSubPlanRefOmitsEmptyOptionals(t *testing.T) {
	raw, err := json.Marshal(SubPlanRef{PlanID: "child"})
	if err != nil {
		t.Fatalf("marshal SubPlanRef: %v", err)
	}
	var encoded map[string]any
	if err := json.Unmarshal(raw, &encoded); err != nil {
		t.Fatalf("unmarshal encoded SubPlanRef: %v", err)
	}
	for _, key := range []string{"digest", "inputs"} {
		if _, ok := encoded[key]; ok {
			t.Errorf("expected omitted %s for an empty value", key)
		}
	}
}
