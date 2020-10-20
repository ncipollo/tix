package cmd

import "tix/logger"

const helpMessage =
`Tix -A command line utility for generating jira, etc tickets from a markdown document.
Usage: tix [OPTIONS] <markdown file>
-d	prints out ticket information instead of creating tickets (shorthand)
-dryrun
prints out ticket information instead of creating tickets
-h	prints help for tix (shorthand)
-help
prints help for tix
-q	suppresses all log output except errors (shorthand)
-quiet
suppresses all log output except errors
-v	enables verbose output (shorthand)
-verbose
enables verbose output
-version
prints tix version

# Settings

Tix expects to find a file called tix.yml in the same directory as the markdown document provided. This settings
file contains information needed by tix in order to communicate with ticketing systems.

The following is the expected format of the settings file:
// yml
github:
	no_projects: true // Indicates if tix should use projects or treat root tickets as issues. Defaults to false.
	owner: owner // The owner of the github repo (ex - ncipollo)
	repo: repo  // The github repo (ex - tix)
	tickets:
		default:
			default: default // Fields to be added to both projets and issues
		project:
			project: project // Fields to be added to projects
		issue:
			labels: [label1, label2] // Fields to be added to issues
jira:
	no_epics: false // Indicates if tix should use epics or treat root tickets as stories / issues. Defaults to false.
	url: https://url.to.your.jira.instance.com
	tickets:
	// All fields should be lower case. Field name spaces should be included (ex- epic name)
		default:
			field: value // Fields to be added to any kind of jira issue
		epic:
			field: value // Fields to be added to epics
		issue:
			field: value // Fields to be added to issues
		task:
			field: value // Fields to be added to tasks (sub-issues of issues)
		variables:
			key: value
		envKey: $ENVIRONMENT_VARIABLE
// tix will parse the markdown document and replace each occurance of "key" with it's value (or environment variable
// when a '$' preceeds the value)
// end yml

# Markdown

Tix interprets the content of heading elements as a tickets. The indent level of the heading element will indicate the
level of nesting for the ticket. For example:
// Markdown
# Root Ticket

This will be a root level ticket.

## Child Ticket

This ticket will be a child of root ticket.

### Grandchild Ticket

This ticket will be a child of child ticket.

// End Markdown

The parent-child relationship for tickets will translate into specific relationships for ticketing systems. In Jira,
for example, the root ticket could be an epic, the next level down a story, then the final level would be a task.

## Markdown Elements

Tix parses each markdown element into an abstract representation of that element. Tix will then utilize these
abstractions to generate representations which are specific to a ticketing system. For example:
- List type elements will would utilize the wiki-media style indentation markers for jira (--)
- Code blocks in jira will use the {code} markers.

## Special Blocks
Tix supports special code blocks which may be used to add fields for a ticket. For example:

%v

# Jira
Tix expects your jira account information to be stored in your environment. Specifically, it will look for the following
variables:

- JIRA_USERNAME: This should be set to your Jira user name (typically an email address).
- JIRA_API_TOKEN: This should be set to your Jira api key. This may be generated by following the instructions found
here: [Jira Api tokens](https://confluence.atlassian.com/cloud/api-tokens-938839638.html)

Jira tickets have the following relationship with heading indent levels:
- #: Epic, if epics are allowed via settings. Story otherwise.
- ##: Story / issue
- ###: Task

# Github
Tix expects your github account information to be stored in your environment. Specifically, it will look for the following variable:

- GITHUB_API_TOKEN: This should be set to your github access token. Access token's may be generated here: [Github API Tokens](https://github.com/settings/tokens)

Github tickets have the following relationship with heading indent levels:
- #: Project, if projects are allowed via settings. Issue otherwise.
- ##: Issue

## Github Settings
The following settings properties effect the issue type tickets.

- assignee: (string) Sets the assignee of the ticket. Should match a github account name (ex - ncipollo).
- assignees: (array) Sets multiple assignees for the ticket. Should match a github account name (ex - [ncipollo, etc]).
- column: (string) Specifies the column the issue should be placed in when it's within a project (ex - To do). Defaults to To do.
- labels: (array) Specifies the labels to apply to the issue (ex - [label1, label2]).
- milestone: (string) Specifies the milestone to add this issue to (ex - v1.0.0). Note: This will create a milestone if it doesn't yet exist. This can cause tix to fail if the milestone already exists but is closed.
- project: (number) The project the issue should be added to (ex - 13). The project should exist and be in the open state. This only takes effect if no_projects is true.

For up to date documentation check out https://github.com/ncipollo/tix/blob/main/README.md
`

const specialBlocks  =
	"```tix\n" +
		"// Adds fields to the ticket, regardless of ticket system\n" +
		"field: value\n" +
		"```\n" +
		"```github\n" +
		"// Adds field to the ticket only if the ticketing system is github.\n" +
		"field: value\n" +
		"```\n" +
		"```jira\n" +
		"// Adds field to the ticket only if the ticketing system is jira.\n" +
		"field: value\n" +
		"```"

type HelpCommand struct {

}

func NewHelpCommand() *HelpCommand {
	return &HelpCommand{}
}

func (h HelpCommand) Run() error {
	logger.Output(helpMessage, specialBlocks)
	return nil
}

