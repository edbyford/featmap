package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func claudeAPI(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Use(RequireAPIKey())

		// List all projects in workspace
		r.Get("/projects", listProjectsForClaude)

		// Get all features for a project with full context
		r.Get("/projects/{PROJECT_ID}/features", getProjectFeaturesForClaude)

		// Update feature status
		r.Post("/features/{FEATURE_ID}/status", updateFeatureStatusForClaude)

		// Update feature annotations
		r.Post("/features/{FEATURE_ID}/annotations", updateFeatureAnnotationsForClaude)

		// API key management (authenticated via existing API key)
		r.Get("/api-keys", listAPIKeys)
	})
}

// listProjectsForClaude returns all projects in the workspace
func listProjectsForClaude(w http.ResponseWriter, r *http.Request) {
	s := GetEnv(r).Service

	projects := s.GetProjects()

	type projectSummary struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	type response struct {
		WorkspaceID string           `json:"workspaceId"`
		Projects    []projectSummary `json:"projects"`
	}

	summaries := make([]projectSummary, len(projects))
	for i, p := range projects {
		summaries[i] = projectSummary{
			ID:          p.ID,
			Title:       p.Title,
			Description: p.Description,
		}
	}

	render.JSON(w, r, response{
		WorkspaceID: s.GetWorkspaceObject().ID,
		Projects:    summaries,
	})
}

// getProjectFeaturesForClaude returns all features with full context and instructions
func getProjectFeaturesForClaude(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "PROJECT_ID")

	s := GetEnv(r).Service

	resp, err := s.GetProjectForClaude(projectID)
	if err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, resp)
}

type updateStatusRequest struct {
	Status string `json:"status"`
}

func (p *updateStatusRequest) Bind(r *http.Request) error {
	return nil
}

// updateFeatureStatusForClaude updates a feature's status
func updateFeatureStatusForClaude(w http.ResponseWriter, r *http.Request) {
	featureID := chi.URLParam(r, "FEATURE_ID")

	data := &updateStatusRequest{}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	s := GetEnv(r).Service

	feature, err := s.UpdateFeatureStatusForClaude(featureID, data.Status)
	if err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, feature)
}

type updateAnnotationsRequest struct {
	Annotations string `json:"annotations"`
}

func (p *updateAnnotationsRequest) Bind(r *http.Request) error {
	return nil
}

// updateFeatureAnnotationsForClaude updates a feature's annotations
func updateFeatureAnnotationsForClaude(w http.ResponseWriter, r *http.Request) {
	featureID := chi.URLParam(r, "FEATURE_ID")

	data := &updateAnnotationsRequest{}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	s := GetEnv(r).Service

	feature, err := s.UpdateFeatureAnnotationsForClaude(featureID, data.Annotations)
	if err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, feature)
}

// listAPIKeys returns all API keys for the workspace
func listAPIKeys(w http.ResponseWriter, r *http.Request) {
	s := GetEnv(r).Service

	keys := s.GetAPIKeysByWorkspace()

	render.JSON(w, r, keys)
}
