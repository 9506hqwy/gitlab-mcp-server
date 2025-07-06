package gitlab

import (
	"context"
	"math"
	"strconv"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetIssues(s *server.MCPServer) {
	tool := mcp.NewTool("issues_",
		mcp.WithDescription("List issues"),
		mcp.WithNumber("assignee_id",
			mcp.Description("Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee."),
		),
		mcp.WithString("assignee_username",
			mcp.Description("Return issues assigned to the given username. Similar to assignee_id and mutually exclusive with assignee_id. In GitLab CE, the assignee_username array should only contain a single value. Otherwise, an invalid parameter error is returned."),
		),
		mcp.WithNumber("author_id",
			mcp.Description("Return issues created by the given user id. Mutually exclusive with author_username. Combine with scope=all or scope=assigned_to_me."),
		),
		mcp.WithString("author_username",
			mcp.Description("Return issues created by the given username. Similar to author_id and mutually exclusive with author_id."),
		),
		mcp.WithBoolean("confidential",
			mcp.Description("Filter confidential or public issues."),
		),
		mcp.WithString("created_after",
			mcp.Description("Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)."),
		),
		mcp.WithString("created_before",
			mcp.Description("Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)."),
		),
		mcp.WithString("due_date",
			mcp.Description("Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next_month_and_previous_two_weeks."),
		),
		mcp.WithNumber("epic_id",
			mcp.Description("Return issues associated with the given epic ID. None returns issues that are not associated with an epic. Any returns issues that are associated with an epic. Premium and Ultimate only."),
		),
		mcp.WithString("health_status",
			mcp.Description("Return issues with the specified health_status. (Introduced in GitLab 15.4). In GitLab 15.5 and later, None returns issues with no health status assigned, and Any returns issues with a health status assigned. Ultimate only."),
		),
		mcp.WithString("iids[]",
			mcp.Description("Return only the issues having the given iid."),
		),
		mcp.WithString("in",
			mcp.Description("Modify the scope of the search attribute. title, description, or a string joining them with comma. Default is title,description."),
		),
		mcp.WithString("issue_type",
			mcp.Description("Filter to a given type of issue. One of issue, incident, test_case or task."),
		),
		mcp.WithNumber("iteration_id",
			mcp.Description("Return issues assigned to the given iteration ID. None returns issues that do not belong to an iteration. Any returns issues that belong to an iteration. Mutually exclusive with iteration_title. Premium and Ultimate only."),
		),
		mcp.WithString("iteration_title",
			mcp.Description("Return issues assigned to the iteration with the given title. Similar to iteration_id and mutually exclusive with iteration_id. Premium and Ultimate only."),
		),
		mcp.WithString("labels",
			mcp.Description("Comma-separated list of label names, issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive."),
		),
		mcp.WithString("milestone_id",
			mcp.Description("Returns issues assigned to milestones with a given timebox value (None, Any, Upcoming, and Started). None lists all issues with no milestone. Any lists all issues that have an assigned milestone. Upcoming lists all issues assigned to milestones due in the future. Started lists all issues assigned to open, started milestones. The logic for Upcoming and Started differs from the logic used in the GraphQL API. milestone and milestone_id are mutually exclusive."),
		),
		mcp.WithString("milestone",
			mcp.Description("The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone. Using None or Any will be deprecated in the future. Use milestone_id attribute instead. milestone and milestone_id are mutually exclusive."),
		),
		mcp.WithString("my_reaction_emoji",
			mcp.Description("Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction."),
		),
		mcp.WithBoolean("non_archived",
			mcp.Description("Return issues only from non-archived projects. If false, the response returns issues from both archived and non-archived projects. Default is true."),
		),
		mcp.WithString("not",
			mcp.Description("Return issues that do not match the parameters supplied. Accepts: assignee_id, assignee_username, author_id, author_username, iids, iteration_id, iteration_title, labels, milestone, milestone_id and weight."),
		),
		mcp.WithString("order_by",
			mcp.Description("Return issues ordered by created_at, due_date, label_priority, milestone_due, popularity, priority, relative_position, title, updated_at, or weight fields. Default is created_at."),
		),
		mcp.WithString("scope",
			mcp.Description("Return issues for the given scope: created_by_me, assigned_to_me or all. Defaults to created_by_me."),
		),
		mcp.WithString("search",
			mcp.Description("Search issues against their title and description."),
		),
		mcp.WithString("sort",
			mcp.Description("Return issues sorted in asc or desc order. Default is desc."),
		),
		mcp.WithString("state",
			mcp.Description("Return all issues or just those that are opened or closed."),
		),
		mcp.WithString("updated_after",
			mcp.Description("Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)."),
		),
		mcp.WithString("updated_before",
			mcp.Description("Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z)."),
		),
		mcp.WithNumber("weight",
			mcp.Description("Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned. Premium and Ultimate only."),
		),
		mcp.WithBoolean("with_labels_details",
			mcp.Description("If true, the response returns more details for each label in labels field: :name, :color, :description, :description_html, :text_color. Default is false."),
		),
	)

	s.AddTool(tool, getIssuesHandler)
}

func getIssuesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetIssues(request)
	return toResult(c.GetApiV4Issues(ctx, &params, authorizationHeader))
}

func parseGetIssues(request mcp.CallToolRequest) client.GetApiV4IssuesParams {
	params := client.GetApiV4IssuesParams{}

	assignee_id := request.GetInt("assignee_id", math.MinInt)
	if assignee_id != math.MinInt {

		params.AssigneeId = &assignee_id
	}

	assignee_username := request.GetString("assignee_username", "")
	if assignee_username != "" {
		assignee_username := strings.Split(assignee_username, ",")
		params.AssigneeUsername = &assignee_username
	}

	author_id := request.GetInt("author_id", math.MinInt)
	if author_id != math.MinInt {

		params.AuthorId = &author_id
	}

	author_username := request.GetString("author_username", "")
	if author_username != "" {

		params.AuthorUsername = &author_username
	}

	confidential := request.GetBool("confidential", false)
	params.Confidential = &confidential

	created_after := request.GetString("created_after", "")
	if created_after != "" {

		params.CreatedAfter = &created_after
	}

	created_before := request.GetString("created_before", "")
	if created_before != "" {

		params.CreatedBefore = &created_before
	}

	due_date := request.GetString("due_date", "")
	if due_date != "" {

		params.DueDate = &due_date
	}

	epic_id := request.GetInt("epic_id", math.MinInt)
	if epic_id != math.MinInt {

		params.EpicId = &epic_id
	}

	health_status := request.GetString("health_status", "")
	if health_status != "" {

		params.HealthStatus = &health_status
	}

	iids__ := request.GetString("iids[]", "")
	if iids__ != "" {
		iids__ := strings.Split(iids__, ",")
		var intSlice []int
		for _, v := range iids__ {
			intValue, _ := strconv.Atoi(v)
			intSlice = append(intSlice, intValue)
		}
		params.Iids = &intSlice
	}

	in := request.GetString("in", "")
	if in != "" {

		params.In = &in
	}

	issue_type := request.GetString("issue_type", "")
	if issue_type != "" {

		params.IssueType = &issue_type
	}

	iteration_id := request.GetInt("iteration_id", math.MinInt)
	if iteration_id != math.MinInt {

		params.IterationId = &iteration_id
	}

	iteration_title := request.GetString("iteration_title", "")
	if iteration_title != "" {

		params.IterationTitle = &iteration_title
	}

	labels := request.GetString("labels", "")
	if labels != "" {

		params.Labels = &labels
	}

	milestone_id := request.GetString("milestone_id", "")
	if milestone_id != "" {

		params.MilestoneId = &milestone_id
	}

	milestone := request.GetString("milestone", "")
	if milestone != "" {

		params.Milestone = &milestone
	}

	my_reaction_emoji := request.GetString("my_reaction_emoji", "")
	if my_reaction_emoji != "" {

		params.MyReactionEmoji = &my_reaction_emoji
	}

	non_archived := request.GetBool("non_archived", false)
	params.NonArchived = &non_archived

	not := request.GetString("not", "")
	if not != "" {

		params.Not = &not
	}

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	scope := request.GetString("scope", "")
	if scope != "" {

		params.Scope = &scope
	}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	state := request.GetString("state", "")
	if state != "" {

		params.State = &state
	}

	updated_after := request.GetString("updated_after", "")
	if updated_after != "" {

		params.UpdatedAfter = &updated_after
	}

	updated_before := request.GetString("updated_before", "")
	if updated_before != "" {

		params.UpdatedBefore = &updated_before
	}

	weight := request.GetInt("weight", math.MinInt)
	if weight != math.MinInt {

		params.Weight = &weight
	}

	with_labels_details := request.GetBool("with_labels_details", false)
	params.WithLabelsDetails = &with_labels_details

	return params
}

func registerGetIssuesId(s *server.MCPServer) {
	tool := mcp.NewTool("issues_id",
		mcp.WithDescription("Single issue"),
		mcp.WithNumber("id",
			mcp.Description("The ID of the issue."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, getIssuesIdHandler)
}

func getIssuesIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)

	return toResult(c.GetApiV4IssuesId(ctx, id, authorizationHeader))
}
