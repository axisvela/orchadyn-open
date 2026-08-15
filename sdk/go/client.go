// Package orchadyn is the Go client for the public ORCHADYN Planning API.
package orchadyn

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
}

type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("ORCHADYN API returned %d: %s", e.StatusCode, e.Body)
}

// NewClient creates a client for an ORCHADYN Planning API base URL.
func NewClient(baseURL string, httpClient *http.Client) (*Client, error) {
	parsed, err := url.Parse(strings.TrimRight(baseURL, "/"))
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return nil, fmt.Errorf("invalid ORCHADYN API URL %q", baseURL)
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{baseURL: parsed, httpClient: httpClient}, nil
}

func (c *Client) Health(ctx context.Context) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.endpoint("/healthz"), nil)
	if err != nil {
		return err
	}
	return c.do(request, nil)
}

// Generate sends a PlanningRequest and decodes the complete API response into response.
func (c *Client) Generate(ctx context.Context, planningRequest PlanningRequest, response any) error {
	return c.post(ctx, "/plans:generate", planningRequest, response)
}

// GenerateGoverned resolves capability and authority inputs through configured control planes.
func (c *Client) GenerateGoverned(ctx context.Context, input GovernedPlanningInput, response any) error {
	return c.post(ctx, "/governed-plans:generate", input, response)
}

// Verify verifies a plan against the immutable planning inputs that created it.
func (c *Client) Verify(ctx context.Context, verificationRequest VerificationRequest, response any) error {
	return c.post(ctx, "/plans:verify", verificationRequest, response)
}

// Revise creates a new plan package without changing completed effects.
func (c *Client) Revise(ctx context.Context, revisionRequest RevisionRequest, response any) error {
	return c.post(ctx, "/plans:revise", revisionRequest, response)
}

// Project projects a persisted plan to a supported execution runtime.
func (c *Client) Project(ctx context.Context, planID string, projectionRequest ProjectionRequest, response any) error {
	return c.post(ctx, "/plans/"+url.PathEscape(planID)+":project", projectionRequest, response)
}

// AnalyzeImpact identifies the smallest plan subgraph affected by changed nodes.
func (c *Client) AnalyzeImpact(ctx context.Context, planID string, request ImpactRequest, response any) error {
	return c.post(ctx, "/plans/"+url.PathEscape(planID)+":impact", request, response)
}

// Adapt adapts a persisted plan to a registered third-party runtime.
func (c *Client) Adapt(ctx context.Context, planID string, request AdaptRequest, response any) error {
	return c.post(ctx, "/plans/"+url.PathEscape(planID)+":adapt", request, response)
}

// SetPlanState transitions a persisted plan through the governed lifecycle.
func (c *Client) SetPlanState(ctx context.Context, planID string, request SetStateRequest, response any) error {
	return c.post(ctx, "/plans/"+url.PathEscape(planID)+":state", request, response)
}

func (c *Client) GetPlan(ctx context.Context, planID string, response any) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.endpoint("/plans/"+url.PathEscape(planID)), nil)
	if err != nil {
		return err
	}
	return c.do(request, response)
}

// ListPlans returns plan summaries available to the authenticated tenant.
func (c *Client) ListPlans(ctx context.Context, response any) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.endpoint("/plans"), nil)
	if err != nil {
		return err
	}
	return c.do(request, response)
}

// ExplainPlan returns only candidate, validation, and revision evidence recorded for a plan.
func (c *Client) ExplainPlan(ctx context.Context, planID string, response any) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.endpoint("/plans/"+url.PathEscape(planID)+"/explain"), nil)
	if err != nil {
		return err
	}
	return c.do(request, response)
}

// ListPlanEvents returns immutable verification and revision events recorded for a plan.
func (c *Client) ListPlanEvents(ctx context.Context, planID string, response any) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.endpoint("/plans/"+url.PathEscape(planID)+"/events"), nil)
	if err != nil {
		return err
	}
	return c.do(request, response)
}

func (c *Client) post(ctx context.Context, path string, payload, response any) error {
	encoded, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("encode request: %w", err)
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint(path), bytes.NewReader(encoded))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	return c.do(request, response)
}

func (c *Client) do(request *http.Request, response any) error {
	httpResponse, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()
	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}
	if httpResponse.StatusCode < http.StatusOK || httpResponse.StatusCode >= http.StatusMultipleChoices {
		return &APIError{StatusCode: httpResponse.StatusCode, Body: string(body)}
	}
	if response == nil || len(body) == 0 {
		return nil
	}
	if err := json.Unmarshal(body, response); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	return nil
}

func (c *Client) endpoint(path string) string {
	endpoint := *c.baseURL
	endpoint.Path = strings.TrimRight(endpoint.Path, "/") + path
	return endpoint.String()
}
