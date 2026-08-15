package orchadyn

type Goal struct {
	ID        string   `json:"id"`
	Owner     string   `json:"owner"`
	Outcome   string   `json:"outcome"`
	Deadline  string   `json:"deadline"`
	DependsOn []string `json:"dependsOn,omitempty"`
}

type CapabilityRequirement struct {
	ID             string `json:"id"`
	GoalID         string `json:"goalId"`
	CapabilityType string `json:"capabilityType"`
	RequiredEffect string `json:"requiredEffect"`
}

type Capability struct {
	ID           string   `json:"id"`
	Type         string   `json:"type"`
	Version      string   `json:"version"`
	Effects      []string `json:"effects"`
	Region       string   `json:"region"`
	Cost         int64    `json:"cost"`
	DeliveredAt  string   `json:"deliveredAt,omitempty"`
	Availability int64    `json:"availability,omitempty"`
	SLA          string   `json:"sla,omitempty"`
}

type Constraint struct {
	Kind   string `json:"kind"`
	Hard   bool   `json:"hard"`
	Value  string `json:"value,omitempty"`
	Amount int64  `json:"amount,omitempty"`
}

type AuthorityGrant struct {
	Principal     string   `json:"principal"`
	Scopes        []string `json:"scopes"`
	BudgetCeiling int64    `json:"budgetCeiling"`
	Revoked       bool     `json:"revoked"`
}

type DelegationContract struct {
	Delegator       string   `json:"delegator"`
	Delegatee       string   `json:"delegatee"`
	ParentScopes    []string `json:"parentScopes"`
	DelegatedScopes []string `json:"delegatedScopes"`
	BudgetCeiling   int64    `json:"budgetCeiling"`
	EvidenceDuty    []string `json:"evidenceDuty"`
}

type PlanningRequest struct {
	Goals            []Goal                        `json:"goals"`
	Requirements     []CapabilityRequirement       `json:"requirements"`
	Catalog          []Capability                  `json:"catalog"`
	Constraints      []Constraint                  `json:"constraints"`
	Delegations      map[string]DelegationContract `json:"delegations,omitempty"`
	AuthorityGrants  []AuthorityGrant              `json:"authorityGrants,omitempty"`
	OperationalState []OperationalState            `json:"operationalState,omitempty"`
}

type OperationalState struct {
	CapabilityID string `json:"capabilityId"`
	Available    bool   `json:"available"`
}

// GovernedPlanningInput omits caller-supplied catalog and grants because the service resolves them from control planes.
type GovernedPlanningInput struct {
	Tenant       string                        `json:"tenant"`
	Goals        []Goal                        `json:"goals"`
	Requirements []CapabilityRequirement       `json:"requirements"`
	Constraints  []Constraint                  `json:"constraints"`
	Delegations  map[string]DelegationContract `json:"delegations,omitempty"`
}

type VerificationRequest struct {
	Plan    any             `json:"plan"`
	Request PlanningRequest `json:"request"`
}

type RevisionRequest struct {
	BasePlan               any             `json:"basePlan"`
	PlanningRequest        PlanningRequest `json:"planningRequest"`
	Trigger                string          `json:"trigger"`
	InvalidatedAssumptions []string        `json:"invalidatedAssumptions,omitempty"`
	CompletedNodeIDs       []string        `json:"completedNodeIds,omitempty"`
	ChangedNodeIDs         []string        `json:"changedNodeIds,omitempty"`
}

type ProjectionRequest struct {
	Target string `json:"target"`
}

// AdaptRequest selects a registered third-party adapter target.
type AdaptRequest struct {
	Target string `json:"target"`
}

// SetStateRequest requests a governed plan lifecycle transition.
type SetStateRequest struct {
	State string `json:"state"`
}

type ImpactRequest struct {
	ChangedNodeIDs []string `json:"changedNodeIds"`
}

type PlanSummary struct {
	PlanID     string `json:"planId"`
	PlanDigest string `json:"planDigest"`
}

type PlanEvent struct {
	EventType     string `json:"eventType"`
	PlanID        string `json:"planId"`
	PlanDigest    string `json:"planDigest"`
	OccurredAt    string `json:"occurredAt"`
	CorrelationID string `json:"correlationId"`
}
