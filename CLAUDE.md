# Claude Code instructions for Featmap integration

This project has a Claude Code integration that allows you to fetch features from the Featmap story mapping tool and update their status as you work.

## How to use

When asked to work on features from Featmap, use these API endpoints:

### Fetch features

```bash
curl -s "http://localhost:5000/v1/claude/projects/$FEATMAP_PROJECT_ID/features" \
  -H "X-API-Key: $FEATMAP_API_KEY" | jq .
```

The response includes:
- `instructions` - Follow these when processing features
- `features` - Array of features with context (milestone, workflow, comments)

### Update feature status

When starting work on a feature:
```bash
curl -X POST "http://localhost:5000/v1/claude/features/FEATURE_ID/status" \
  -H "X-API-Key: $FEATMAP_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"status": "IN_PROGRESS"}'
```

When completing a feature:
```bash
curl -X POST "http://localhost:5000/v1/claude/features/FEATURE_ID/status" \
  -H "X-API-Key: $FEATMAP_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"status": "CLOSED"}'
```

Valid status values: `OPEN`, `IN_PROGRESS`, `CLOSED`

## Environment variables

The user should have these set:
- `FEATMAP_API_KEY` - API key starting with `fm_`
- `FEATMAP_PROJECT_ID` - The project ID to fetch features from

If these aren't set, ask the user to provide them.

## Workflow guidance

1. Fetch features from the API
2. Review the `instructions` field in the response
3. Use `AskUserQuestion` to clarify requirements before implementing
4. Update status to `IN_PROGRESS` when starting a feature
5. Implement the feature
6. Update status to `CLOSED` when complete
7. Move to the next feature

## List projects (if project ID unknown)

```bash
curl -s "http://localhost:5000/v1/claude/projects" \
  -H "X-API-Key: $FEATMAP_API_KEY" | jq .
```
