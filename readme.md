# Featmap (Claude Code Edition)

> **This is a modified fork of [Featmap](https://github.com/amborle/featmap)** with added API endpoints for Claude Code integration. The original Featmap is a user story mapping tool created by amborle.

This fork adds the ability for Claude Code to fetch features from your story map and update their status as work progresses, enabling AI-driven feature implementation workflows.

![Featmap screenshot](screenshot.png)

## What's different in this fork

This fork adds:

- **Claude Code API** (`/v1/claude/*`) - Endpoints for fetching features and updating status
- **API key authentication** - Secure access for CLI/agent use without browser sessions
- **Behaviour instructions** - API responses include guidance for Claude to clarify requirements before implementing

The original Featmap functionality remains unchanged.

## Claude Code integration

### How it works

1. Create your user story map in Featmap (features, milestones, workflows)
2. Generate an API key
3. Claude Code fetches features via the API
4. The API response includes instructions telling Claude to use `AskUserQuestion` to clarify requirements
5. Claude implements features and updates their status
6. You see status changes reflected in the Featmap UI

### API endpoints

| Endpoint | Purpose |
|----------|---------|
| `GET /v1/claude/projects` | List all projects |
| `GET /v1/claude/projects/{id}/features` | Get features with full context |
| `POST /v1/claude/features/{id}/status` | Update feature status |
| `POST /v1/claude/features/{id}/annotations` | Update annotations |

All endpoints require the `X-API-Key` header.

### Example usage

```bash
# Set your API key
export FEATMAP_API_KEY="fm_your_key_here"

# Fetch features for a project
curl -s "http://localhost:5000/v1/claude/projects/PROJECT_ID/features" \
  -H "X-API-Key: $FEATMAP_API_KEY" | jq .

# Update feature status
curl -X POST "http://localhost:5000/v1/claude/features/FEATURE_ID/status" \
  -H "X-API-Key: $FEATMAP_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"status": "IN_PROGRESS"}'
```

### What Claude Code should do

When working with Featmap features, Claude Code should:

1. **Fetch features** from the API to understand what needs to be built
2. **Read the instructions** included in the API response
3. **Ask clarifying questions** using `AskUserQuestion` before implementing (the API instructs Claude to do this)
4. **Update status to IN_PROGRESS** when starting work on a feature
5. **Implement the feature** based on the description and clarified requirements
6. **Update status to CLOSED** when complete
7. **Move to the next feature** by priority

Valid status values: `OPEN`, `IN_PROGRESS`, `CLOSED`

### Full setup guide

See [CLAUDE_INTEGRATION.md](CLAUDE_INTEGRATION.md) for detailed setup instructions.

---

## Original Featmap documentation

Featmap is a user story mapping tool for product people to build, plan and communicate product backlogs.

- [Introduction](#introduction)
  - [Purpose](#purpose)
  - [Features](#features)
  - [Intended audience](#intended-audience)
  - [Motivation](#motivation)
- [Getting started](#getting-started)
- [Self hosting](#self-hosting)
  - [Requirements](#requirements)
  - [Download](#download)
  - [Configuration](#configuration)
  - [Run](#run)
  - [Upgrade](#upgrade)
  - [Building from source and running with docker-compose](#Building-from-source-and-running-with-docker-compose)
- [License](#license)

## Introduction
 Featmap is an open source user story mapping tool. It is built using React, Typescript and Go.
 ### Purpose
Featmap was built for product people to take advantage of a technique called *user story mapping*. User story mapping, or just story mapping, is an effective tool to create, plan and communicate your product backlog. Story mapping was created by Jeff Patton and its primary utility is providing us with an overview of the entire product and how user goals are broken down into a series of tasks. Finally, it helps us to define valuable product slices (releases) and prioritize between them.
### Features
* Personas
* Markdown editing
* Discuss user stories
* Share your user story maps with external stakeholders
* User story annotations
* User story estimates with roll-ups

### Intended audience
Featmap is great for product managers, product owners or just about anyone who is building products. Featmap can also be used as a light weight work item management system for development teams.

### Motivation
There are many user story mapping tools, however none are really focused on easy-of-use and simplicity. Featmap was built to fill that gap. We hope you will find it as useful as we found building it.
## Getting started
You have two choices when it comes to using Featmap.
1. Use the hosted service at https://www.featmap.com (original version without Claude integration).
2. Host this fork yourself by running it on your own server. Please refer to the [instructions](#self-hosting) for self-hosting.
## Self hosting
Featmap can be run on your own server.
### Requirements
Featmap runs on top of [PostgreSQL](https://www.postgresql.org/), so make sure you have it running on your system. At this step, make sure to setup the credentials and database that Featmap will use.
### Download
For the original version, [download](https://github.com/amborle/featmap/releases) binaries from the upstream repo.

For this fork with Claude integration, build from source using Docker (see below).
### Configuration
In the directory where you placed the binary, create a file called ```conf.json```.

Here's a sample  ```conf.json``` you can use:

```json
{
  "appSiteURL": "http://localhost:5000",
  "dbConnectionString": "postgresql://postgres:postgres@postgres:5432/postgres?sslmode=disable",
  "jwtSecret": "ChangeMeForProduction",
  "port": "5000",
  "emailFrom": "",
  "smtpServer": "",
  "smtpPort": "587",
  "smtpUser": "",
  "smtpPass": "",
  "environment": "development",
  "mode": "selfhosted"
}
```
Setting | Description
--- | ---
`appSiteURL` | The url to where you will be hosting the app.
`dbConnectionString` | The connection string to the PostgreSQL database that Featmap should connect to.
`jwtSecret` | This setting is used to secure the cookies produced by Featmap. Generate a random string and keep it safe!
`port` | The port that Featmap should run on.
`emailFrom` | The email adress that should be used as sender when sending invitation and password reset mails.
`smtpServer` | SMTP server for sending emails.
`smtpPort` | **Optional** Will default to port 587 if not specified.
`smtpUser` | SMTP server username.
`smtpPass` | SMTP server password.
`environment` |  **Optional** If set to `development`, Featmap assumes your are **not** running on **https** and the the backend will not serve secure cookies. Remove this setting if you have set it up to run https.
`mode` | **Optional** Set to `selfhosted` for self-hosted deployments.
### Run
Execute the binary.

```bash
./featmap-1.0.0-linux-amd64
Serving on port 5000
```
Open a browser to http://localhost:5000 and you are ready to go!
### Upgrading
Just download the latest release and swap out the executable. Remember to backup your database and the old executable.

## Building from source and running with docker-compose

Clone this repository

```bash
git clone https://github.com/edbyford/featmap.git
```

Navigate to the repository.

```bash
cd featmap
```

Let's copy the configuration files

```bash
cp config/.env .
cp config/conf.json .
```

Now let's build it.

```bash
docker-compose build
```

Startup the services, the app should now be available on the port you defined in you configuration files (default 5000).
```bash
docker-compose up -d
```

### Upgrading
Remember to backup your database (/data), just in case.

Pull down the latest source
```bash
git pull
```
Now let's rebuild it.
```bash
docker-compose build --no-cache
```
And finally run it.
```bash
docker-compose up -d
```

## License
Featmap is licensed under Business Source License 1.1. See [license](https://github.com/amborle/featmap/blob/master/LICENSE)

This fork maintains the same license. Modifications are permitted for personal and internal use.
