# Scratch
## Tools
### Relevant Links
- Atlassian dev community: https://community.developer.atlassian.com/
- About Atlassian API tokens: https://support.atlassian.com/atlassian-accountq/docs/manage-api-tokens-for-your-atlassian-account/
- Getting started with Connect: https://developer.atlassian.com/cloud/jira/platform/getting-started-with-connect/
- Jira REST APIs: https://developer.atlassian.com/server/jira/platform/rest-apis/
- Jira Software REST API Reference: https://docs.atlassian.com/jira-software/REST/7.3.1/
- Jira Cloud API: https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/
- Dynamic JSON parsing: https://github.com/Jeffail/gabs
- Make subcommands without framework: https://blog.ippon.tech/no-framework-approach-to-building-a-cli-with-go/

- TEST board: https://natx.atlassian.net/jira/software/projects/TEST/boards/1
- My test pofile: https://natx.atlassian.net/jira/people/5c9476dc627d3b2d0163ee3d

Discarded:
- I've discarded using
  [go-jira](https://github.com/andygrunwald/go-jira) as this is not
  providing some features I need straight away and it is too complex
  to contribute them to it.

### Snippets
#### cURL for sprint issues
```
curl -s --request GET \
  --url "${JIRA_URL}/rest/agile/1.0/sprint/883/issue" \
  --header "Authorization: Basic $(echo -n ${JIRA_USERNAME}:${JIRA_API_TOKEN} | base64)" \
  --header "Accept: application/json"
```

## Roadmap
- List projects [X]
- List sprints for a project [X]
- List issues for sprints [X]
- Print sprint stats command [X]
- Make Project name also key [X]
- Add Sprint state filter [X]
- Add command to move all issues not closed to the next sprint [X]
- Add complete sprint command, only if all issues are closed [X]
- Add start sprint command [X]
- Add create sprint command [X]
- Fix sprint name generation [X]
- Move all commands as subcommands of a single command [X]
- Add CI and fix all steps [X]
- Add CD [X]
- Add README.md with description, install, setup and usage instructions [X]
- Release v0.1.0 [X]
- Add ROADMAP.md []
- Add tests, and corresponding execution in CI/CD [X]
- Make complete command ensure unconcluded issues are moved to next sprint/backlog []
- Add sprint start and end dates options to create/start sprint command []
- Add goal option to create/start sprint command []
- Add origin board id and name opstions to create sprint command []
- Add environment variable for the list of issue statuses to move to next sprint []
- Research and fix possible issues with Issue state identification (Closed=done) []
- Add config subcommand to create configuration file with options []
- Add incomplete issues movement options to complete sprint command []
- Include count of sprints/issues []
- Add support for pagination and rate limiting []
- Add issue estimation feature []

## Final outputs
ITT Closing NUT Sprint 2022-7 and starting NUT Sprint 2022-8 :done:
NUT Sprint 2022-7 closed:
4 issues (9 sp) have been Closed
2 issues (8 sp) are still in Progress and will move to Sprint 8
12 issues (42 sp) have not been started and will also move to Sprint 8
Sprint 8 has been started, Sprint 9 created.

ITT Closing INF Sprint 2022-7 and starting INF Sprint 2022-8 :done:
INF Sprint 7 closed. @Elias Porter’s issue continues to be in Progress and has moved to Sprint 8, which has been started.
One task (5sp) was still in To Do and also moved to the next sprint.
INF Sprint 2022-9 has been created.

ITT Closing DVX Sprint 2022-7 and starting DVX Sprint 2022-8 :done:
We completed all 6 scheduled production deployments. Sprint closed.
Sprint 8 started. Sprint goal: Complete 8 production deployments.
Sprint 9 had already been created, I created also Sprint 10 just to get ahead of myself

:spinning-parrot: switching sprints , Sprint S7 stats will be flowing in soon. Status -> :done:
5 tickets were Closed, one of them added after the sprint was started .
7 tickets (one of them a story) are not completed, 2 of them being added after the sprint started.
4 -> In Progress
2 -> Resolved
1 -> Not Started
2 tickets were removed from the sprint because we realized we won't get to them anyway.
See full sprint stats https://mydevex.atlassian.net/jira/software/c/projects/BAN/boards/107/reports/sprint-retrospective?sprint=870
Screen Shot 2022-04-22 at 11.41.18 AM.png ￼
Sprint Report
The sprint goal "Execute the migration of the newsletter subscriptions" was not achieved since we did not run the migration in time.
The ticket will spill over and will be part of the sprint goal for S8.
Sprint 8 started. https://mydevex.atlassian.net/jira/software/c/projects/BAN/boards/107
