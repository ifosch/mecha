# mecha

[![passing](https://github.com/ifosch/mecha/actions/workflows/test.yaml/badge.svg)](https://github.com/ifosch/mecha/actions/workflows/test.yaml) [![release](https://img.shields.io/github/release/ifosch/mecha.svg)](https://github.com/ifosch/mecha/releases/)

## About mecha

mecha is a CLI for Jira, written in go.

So far it is designed to be used both directly from the prompt or for
scripting, but it's not interactive.

It provides some basic features useful to handle some usual activities
while running Agile software development processes, like completing,
creating, or starting sprints, moving issues to another sprint
depending on status, or other similar operations.

As Jira's names is reported to be after Gojira (Godzilla), a kaiju, a
giant monster destroying the city in Japanese pop culture, mecha is
supposed to be a giant robot, which usually fight with kaijus.

## Install

The easiest way to get mecha installed is choosing the corresponding
binary for your system from any production-ready release in [this
repository's page releases area](releases).

Then, put the binary file in you execution path, with execution
permissions, and it should be ok.

## Setup

In order to access your Jira installation, you'll need to create the
following environment variables:
* `JIRA_API_TOKEN`: This must be a valid API token for your user in
  your Jira installation. For Atlassian hosted installations, probably
  it's enough to checkout [their documentation about API
  tokens](https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/). Self-hosted
  installations should be kind of similar.
* `JIRA_USERNAME`: This must be your username in your Jira
  installation. 
* `JIRA_URL`: This must be the whole URL to your Jira installation,
  including the HTTP/S protocol at the beginning.
* `JIRA_SP_FIELD`: This must be the Custom field ID your Jira
  installation is using to store the story points estimations. You can
  probably use the following cURL command to find out:
  
  ```console
  curl -s \
     --request GET \
     --url "${JIRA_URL}/rest/api/3/field" \
     --header "Authorization: Basic $(echo -n ${JIRA_USERNAME}:${JIRA_API_TOKEN} | base64)" \
     --header "Accept: application/json" | \
  jq '.[] | select(.name | contains("Story point estimate")).key'
  ```

## Usage

mecha is the main CLI entrypoint providing no functionality.

To have it doing stuff, you need to provide specific commands, which
sometimes have other subcommands, or also could have options.

If you run mecha without any subcommand, it will show you what's available:
```console
$ mecha 
mecha: You need to provide a command.

Usage: mecha command [options]

A Jira CLI focused in project management tasks.

Commands:
  list     Lists projects, sprints, and issues
  add      Adds a new sprint to specified project
  get      Gets stats for active and future sprints in specified project
  move     Moves all issues in currently active sprint to next one for the specified project
  start    Starts the next sprint for the specified project
  complete Completes the active sprint in the specified project
```

You can check any command's subcommands with the `-h` option:
```console
$ mecha list projects -h
Lists projects.

Usage: mecha list projects
```

### Get sprint stats
```console
$ mecha get --project TEST
```

### Move all non closed stories in current active sprint to the next sprint
```console
$ mecha move --project TEST --to-move "To Do,In Progress"
```

### Complete the currently active sprint
```console
$ mecha complete --project TEST
```

### Start the next sprint
```console
$ mecha start --project TEST
```

### Add a new sprint
```console
$ mecha add --project TEST
```
