# Claude Code integration for Featmap

This document describes how to use Featmap with Claude Code for AI-driven feature implementation.

## Overview

This integration allows Claude Code to:
- Fetch features from your story map via API
- Receive instructions to clarify requirements before implementing
- Update feature status as work progresses

## Quick start

### 1. Start Featmap locally

```bash
# Create environment file
cat > .env << 'EOF'
FEATMAP_DB=featmap
FEATMAP_DB_USER=featmap
FEATMAP_DB_PASSWORD=changeme
FEATMAP_HTTP_PORT=5000
EOF

# Create conf.json (copy from config/conf.json and adjust as needed)
cp config/conf.json conf.json

# Start with Docker
docker-compose up --build
```

Open http://localhost:5000 in your browser.

### 2. Create a workspace and project

1. Register an account (email confirmation not required in self-hosted mode)
2. Create a workspace
3. Create a project with your features/user stories

### 3. Generate an API key

Using curl (after logging in via the browser to get a session):

```bash
# You'll need to be logged in and have the JWT cookie
# The workspace ID is visible in the URL when viewing your workspace

curl -X POST "http://localhost:5000/v1/api-keys" \
  -H "Content-Type: application/json" \
  -H "Workspace: YOUR_WORKSPACE_ID" \
  -H "Cookie: jwt=YOUR_JWT_TOKEN" \
  -d '{"name": "claude-code"}'
```

The response includes the raw API key (shown only once):

```json
{
  "apiKey": { "id": "...", "name": "claude-code", ... },
  "rawKey": "fm_abc123...",
  "warning": "Store this key securely. It will only be shown once."
}
```

### 4. Configure environment variables

```bash
export FEATMAP_API_KEY="fm_your_key_here"
export FEATMAP_PROJECT_ID="your_project_id"
```

### 5. Use with Claude Code

Start Claude Code in your project directory and ask it to fetch features:

```
Fetch features from featmap at http://localhost:5000 and help me implement them
```

Claude Code will:
1. Call the API to get features with context
2. Receive instructions to use AskUserQuestion for clarifications
3. Ask you clarifying questions before implementing
4. Update feature status as work progresses

## API reference

All endpoints require the `X-API-Key` header.

### List projects

```bash
curl "http://localhost:5000/v1/claude/projects" \
  -H "X-API-Key: $FEATMAP_API_KEY"
```

### Get features for a project

```bash
curl "http://localhost:5000/v1/claude/projects/$FEATMAP_PROJECT_ID/features" \
  -H "X-API-Key: $FEATMAP_API_KEY"
```

Response includes:
- `instructions` - Guidance for Claude on how to process features
- `project` - Project metadata
- `features` - Array of features with full context (milestone, workflow, subworkflow, comments)

### Update feature status

```bash
curl -X POST "http://localhost:5000/v1/claude/features/FEATURE_ID/status" \
  -H "X-API-Key: $FEATMAP_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"status": "IN_PROGRESS"}'
```

Valid status values: `OPEN`, `IN_PROGRESS`, `CLOSED`

### Update feature annotations

```bash
curl -X POST "http://localhost:5000/v1/claude/features/FEATURE_ID/annotations" \
  -H "X-API-Key: $FEATMAP_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"annotations": "DISCUSSION"}'
```

## Workflow

```
┌─────────────────────────────────────────────────────────────────┐
│  1. Create requirements in featmap UI                           │
│     (titles, brief descriptions - don't over-specify)           │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  2. Generate API key and set environment variables              │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  3. Start Claude Code                                           │
│     "Fetch features from featmap and help me implement them"    │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  4. Claude fetches features via API                             │
│     Response includes instructions to clarify requirements      │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  5. Claude uses AskUserQuestion to clarify requirements         │
│     "For user login, should I support OAuth or just email?"     │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  6. Claude updates status to IN_PROGRESS, implements feature    │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│  7. Claude updates status to CLOSED                             │
│     → You see the update in featmap UI                          │
└─────────────────────────────────────────────────────────────────┘
```

## Next steps

### Frontend UI for API key management

The backend supports API key CRUD operations, but the React frontend doesn't yet have a UI for managing API keys. To add this:

1. Add an "API Keys" section to the workspace settings page
2. Call `GET /v1/api-keys` to list keys
3. Call `POST /v1/api-keys` to create new keys
4. Call `DELETE /v1/api-keys/{id}` to delete keys

### MCP server (optional)

For tighter integration, you could build an MCP server that wraps these API calls, allowing Claude Code to use dedicated tools rather than curl commands.

## Troubleshooting

### API key not working

- Ensure the key starts with `fm_`
- Check that the key hasn't been deleted
- Verify the workspace has an active subscription

### Features not appearing

- Ensure you're using the correct project ID
- Check that features exist in the project (visible in the UI)

### Status updates not reflected

- Refresh the featmap UI
- Check the API response for errors
