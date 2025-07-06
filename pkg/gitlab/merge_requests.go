package gitlab

import (
	"context"
	"math"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/gitlab-client-go/pkg/gitlab"
)

func registerGetMergeRequests(s *server.MCPServer) {
	tool := mcp.NewTool("mrs_",
		mcp.WithDescription("Get all merge requests the authenticated user has access to. By default it returns only merge requests created by the current user. To get all merge requests, use parameter `scope=all`."),
		mcp.WithNumber("author_id",
			mcp.Description("Returns merge requests created by the given user `id`. Mutually exclusive with `author_username`. Combine with `scope=all` or `scope=assigned_to_me`."),
		),
		mcp.WithString("author_username",
			mcp.Description("Returns merge requests created by the given `username`. Mutually exclusive with `author_id`."),
		),
		mcp.WithNumber("assignee_id",
			mcp.Description("Returns merge requests assigned to the given user `id`. `None` returns unassigned merge requests. `Any` returns merge requests with an assignee."),
		),
		mcp.WithString("assignee_username",
			mcp.Description("Returns merge requests created by the given `username`. Mutually exclusive with `author_id`."),
		),
		mcp.WithString("reviewer_username",
			mcp.Description("Returns merge requests which have the user as a reviewer with the given `username`. `None` returns merge requests with no reviewers. `Any` returns merge requests with any reviewer. Mutually exclusive with `reviewer_id`. Introduced in GitLab 13.8."),
		),
		mcp.WithString("labels",
			mcp.Description("Returns merge requests matching a comma-separated list of labels. `None` lists all merge requests with no labels. `Any` lists all merge requests with at least one label. Predefined names are case-insensitive."),
		),
		mcp.WithString("milestone",
			mcp.Description("Returns merge requests for a specific milestone. `None` returns merge requests with no milestone. `Any` returns merge requests that have an assigned milestone."),
		),
		mcp.WithString("my_reaction_emoji",
			mcp.Description("Returns merge requests reacted by the authenticated user by the given `emoji`. `None` returns issues not given a reaction. `Any` returns issues given at least one reaction."),
		),
		mcp.WithNumber("reviewer_id",
			mcp.Description("Returns merge requests which have the user as a reviewer with the given user `id`. `None` returns merge requests with no reviewers. `Any` returns merge requests with any reviewer. Mutually exclusive with `reviewer_username`."),
		),
		mcp.WithString("state",
			mcp.Description("Returns `all` merge requests or just those that are `opened`, `closed`, `locked`, or `merged`. (default: all)"),

			mcp.Enum("opened", "closed", "locked", "merged", "all"),
		),
		mcp.WithString("order_by",
			mcp.Description("Returns merge requests ordered by `created_at`, `label_priority`, `milestone_due`, `popularity`, `priority`, `title`, `updated_at` or `merged_at` fields. Introduced in GitLab 14.8. (default: created_at)"),

			mcp.Enum("created_at", "label_priority", "milestone_due", "popularity", "priority", "title", "updated_at", "merged_at"),
		),
		mcp.WithString("sort",
			mcp.Description("Returns merge requests sorted in `asc` or `desc` order. (default: desc)"),

			mcp.Enum("asc", "desc"),
		),
		mcp.WithBoolean("with_labels_details",
			mcp.Description("If `true`, response returns more details for each label in labels field: `:name`,`:color`, `:description`, `:description_html`, `:text_color` (default: false)"),
		),
		mcp.WithBoolean("with_merge_status_recheck",
			mcp.Description("If `true`, this projection requests (but does not guarantee) that the `merge_status` field be recalculated asynchronously. Introduced in GitLab 13.0. (default: false)"),
		),
		mcp.WithString("created_after",
			mcp.Description("Returns merge requests created on or after the given time. Expected in ISO 8601 format. (example: 2019-03-15T08:00:00Z)"),
		),
		mcp.WithString("created_before",
			mcp.Description("Returns merge requests created on or before the given time. Expected in ISO 8601 format. (example: 2019-03-15T08:00:00Z)"),
		),
		mcp.WithString("updated_after",
			mcp.Description("Returns merge requests updated on or after the given time. Expected in ISO 8601 format. (example: 2019-03-15T08:00:00Z)"),
		),
		mcp.WithString("updated_before",
			mcp.Description("Returns merge requests updated on or before the given time. Expected in ISO 8601 format. (example: 2019-03-15T08:00:00Z)"),
		),
		mcp.WithString("view",
			mcp.Description("If simple, returns the `iid`, URL, title, description, and basic state of merge request"),

			mcp.Enum("simple"),
		),
		mcp.WithString("scope",
			mcp.Description("Returns merge requests for the given scope: `created_by_me`, `assigned_to_me` or `all` (default: created_by_me)"),

			mcp.Enum("created-by-me", "assigned-to-me", "created_by_me", "assigned_to_me", "all"),
		),
		mcp.WithString("source_branch",
			mcp.Description("Returns merge requests with the given source branch"),
		),
		mcp.WithNumber("source_project_id",
			mcp.Description("Returns merge requests with the given source project id"),
		),
		mcp.WithString("target_branch",
			mcp.Description("Returns merge requests with the given target branch"),
		),
		mcp.WithString("search",
			mcp.Description("Search merge requests against their `title` and `description`."),
		),
		mcp.WithString("in",
			mcp.Description("Modify the scope of the search attribute. `title`, `description`, or a string joining them with comma. (example: title,description)"),
		),
		mcp.WithString("wip",
			mcp.Description("Filter merge requests against their `wip` status. `yes` to return only draft merge requests, `no` to return non-draft merge requests."),

			mcp.Enum("yes", "no"),
		),
		mcp.WithNumber("not[author_id]",
			mcp.Description("`<Negated>` Returns merge requests created by the given user `id`. Mutually exclusive with `author_username`. Combine with `scope=all` or `scope=assigned_to_me`."),
		),
		mcp.WithString("not[author_username]",
			mcp.Description("`<Negated>` Returns merge requests created by the given `username`. Mutually exclusive with `author_id`."),
		),
		mcp.WithNumber("not[assignee_id]",
			mcp.Description("`<Negated>` Returns merge requests assigned to the given user `id`. `None` returns unassigned merge requests. `Any` returns merge requests with an assignee."),
		),
		mcp.WithString("not[assignee_username]",
			mcp.Description("`<Negated>` Returns merge requests created by the given `username`. Mutually exclusive with `author_id`."),
		),
		mcp.WithString("not[reviewer_username]",
			mcp.Description("`<Negated>` Returns merge requests which have the user as a reviewer with the given `username`. `None` returns merge requests with no reviewers. `Any` returns merge requests with any reviewer. Mutually exclusive with `reviewer_id`. Introduced in GitLab 13.8."),
		),
		mcp.WithString("not[labels]",
			mcp.Description("`<Negated>` Returns merge requests matching a comma-separated list of labels. `None` lists all merge requests with no labels. `Any` lists all merge requests with at least one label. Predefined names are case-insensitive."),
		),
		mcp.WithString("not[milestone]",
			mcp.Description("`<Negated>` Returns merge requests for a specific milestone. `None` returns merge requests with no milestone. `Any` returns merge requests that have an assigned milestone."),
		),
		mcp.WithString("not[my_reaction_emoji]",
			mcp.Description("`<Negated>` Returns merge requests reacted by the authenticated user by the given `emoji`. `None` returns issues not given a reaction. `Any` returns issues given at least one reaction."),
		),
		mcp.WithNumber("not[reviewer_id]",
			mcp.Description("`<Negated>` Returns merge requests which have the user as a reviewer with the given user `id`. `None` returns merge requests with no reviewers. `Any` returns merge requests with any reviewer. Mutually exclusive with `reviewer_username`."),
		),
		mcp.WithString("deployed_before",
			mcp.Description("Returns merge requests deployed before the given date/time. Expected in ISO 8601 format. (example: 2019-03-15T08:00:00Z)"),
		),
		mcp.WithString("deployed_after",
			mcp.Description("Returns merge requests deployed after the given date/time. Expected in ISO 8601 format (example: 2019-03-15T08:00:00Z)"),
		),
		mcp.WithString("environment",
			mcp.Description("Returns merge requests deployed to the given environment (example: 2019-03-15T08:00:00Z)"),
		),
		mcp.WithString("approved",
			mcp.Description("Filters merge requests by their `approved` status. `yes` returns only approved merge requests. `no` returns only non-approved merge requests."),

			mcp.Enum("yes", "no"),
		),
		mcp.WithNumber("merge_user_id",
			mcp.Description("Returns merge requests which have been merged by the user with the given user `id`. Mutually exclusive with `merge_user_username`."),
		),
		mcp.WithString("merge_user_username",
			mcp.Description("Returns merge requests which have been merged by the user with the given `username`. Mutually exclusive with `merge_user_id`."),
		),
		mcp.WithString("approver_ids",
			mcp.Description("Return merge requests which have specified the users with the given IDs as an individual approver"),
		),
		mcp.WithString("approved_by_ids",
			mcp.Description("Return merge requests which have been approved by the specified users with the given IDs"),
		),
		mcp.WithString("approved_by_usernames",
			mcp.Description("Return merge requests which have been approved by the specified users with the given usernames"),
		),
		mcp.WithNumber("page",
			mcp.Description("Current page number (example: 1) (default: 1)"),
		),
		mcp.WithNumber("per_page",
			mcp.Description("Number of items per page (example: 20) (default: 20)"),
		),
	)

	s.AddTool(tool, getMergeRequestsHandler)
}

func getMergeRequestsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGetMergeRequests(request)
	return toResult(c.GetApiV4MergeRequests(ctx, &params, authorizationHeader))
}

func parseGetMergeRequests(request mcp.CallToolRequest) client.GetApiV4MergeRequestsParams {
	params := client.GetApiV4MergeRequestsParams{}

	author_id := request.GetInt("author_id", math.MinInt)
	if author_id != math.MinInt {
		author_id := int32(author_id)
		params.AuthorId = &author_id
	}

	author_username := request.GetString("author_username", "")
	if author_username != "" {

		params.AuthorUsername = &author_username
	}

	assignee_id := request.GetInt("assignee_id", math.MinInt)
	if assignee_id != math.MinInt {
		assignee_id := int32(assignee_id)
		params.AssigneeId = &assignee_id
	}

	assignee_username := request.GetString("assignee_username", "")
	if assignee_username != "" {
		assignee_username := strings.Split(assignee_username, ",")
		params.AssigneeUsername = &assignee_username
	}

	reviewer_username := request.GetString("reviewer_username", "")
	if reviewer_username != "" {

		params.ReviewerUsername = &reviewer_username
	}

	labels := request.GetString("labels", "")
	if labels != "" {
		labels := strings.Split(labels, ",")
		params.Labels = &labels
	}

	milestone := request.GetString("milestone", "")
	if milestone != "" {

		params.Milestone = &milestone
	}

	my_reaction_emoji := request.GetString("my_reaction_emoji", "")
	if my_reaction_emoji != "" {

		params.MyReactionEmoji = &my_reaction_emoji
	}

	reviewer_id := request.GetInt("reviewer_id", math.MinInt)
	if reviewer_id != math.MinInt {
		reviewer_id := int32(reviewer_id)
		params.ReviewerId = &reviewer_id
	}

	state := request.GetString("state", "")
	if state != "" {

		params.State = &state
	}

	order_by := request.GetString("order_by", "")
	if order_by != "" {

		params.OrderBy = &order_by
	}

	sort := request.GetString("sort", "")
	if sort != "" {

		params.Sort = &sort
	}

	with_labels_details := request.GetBool("with_labels_details", false)
	params.WithLabelsDetails = &with_labels_details

	with_merge_status_recheck := request.GetBool("with_merge_status_recheck", false)
	params.WithMergeStatusRecheck = &with_merge_status_recheck

	created_after := request.GetString("created_after", "")
	if created_after != "" {
		created_after, _ := time.Parse(time.RFC3339, created_after)
		params.CreatedAfter = &created_after
	}

	created_before := request.GetString("created_before", "")
	if created_before != "" {
		created_before, _ := time.Parse(time.RFC3339, created_before)
		params.CreatedBefore = &created_before
	}

	updated_after := request.GetString("updated_after", "")
	if updated_after != "" {
		updated_after, _ := time.Parse(time.RFC3339, updated_after)
		params.UpdatedAfter = &updated_after
	}

	updated_before := request.GetString("updated_before", "")
	if updated_before != "" {
		updated_before, _ := time.Parse(time.RFC3339, updated_before)
		params.UpdatedBefore = &updated_before
	}

	view := request.GetString("view", "")
	if view != "" {

		params.View = &view
	}

	scope := request.GetString("scope", "")
	if scope != "" {

		params.Scope = &scope
	}

	source_branch := request.GetString("source_branch", "")
	if source_branch != "" {

		params.SourceBranch = &source_branch
	}

	source_project_id := request.GetInt("source_project_id", math.MinInt)
	if source_project_id != math.MinInt {
		source_project_id := int32(source_project_id)
		params.SourceProjectId = &source_project_id
	}

	target_branch := request.GetString("target_branch", "")
	if target_branch != "" {

		params.TargetBranch = &target_branch
	}

	search := request.GetString("search", "")
	if search != "" {

		params.Search = &search
	}

	in := request.GetString("in", "")
	if in != "" {

		params.In = &in
	}

	wip := request.GetString("wip", "")
	if wip != "" {

		params.Wip = &wip
	}

	not_author_id_ := request.GetInt("not[author_id]", math.MinInt)
	if not_author_id_ != math.MinInt {
		not_author_id_ := int32(not_author_id_)
		params.NotAuthorId = &not_author_id_
	}

	not_author_username_ := request.GetString("not[author_username]", "")
	if not_author_username_ != "" {

		params.NotAuthorUsername = &not_author_username_
	}

	not_assignee_id_ := request.GetInt("not[assignee_id]", math.MinInt)
	if not_assignee_id_ != math.MinInt {
		not_assignee_id_ := int32(not_assignee_id_)
		params.NotAssigneeId = &not_assignee_id_
	}

	not_assignee_username_ := request.GetString("not[assignee_username]", "")
	if not_assignee_username_ != "" {
		not_assignee_username_ := strings.Split(not_assignee_username_, ",")
		params.NotAssigneeUsername = &not_assignee_username_
	}

	not_reviewer_username_ := request.GetString("not[reviewer_username]", "")
	if not_reviewer_username_ != "" {

		params.NotReviewerUsername = &not_reviewer_username_
	}

	not_labels_ := request.GetString("not[labels]", "")
	if not_labels_ != "" {
		not_labels_ := strings.Split(not_labels_, ",")
		params.NotLabels = &not_labels_
	}

	not_milestone_ := request.GetString("not[milestone]", "")
	if not_milestone_ != "" {

		params.NotMilestone = &not_milestone_
	}

	not_my_reaction_emoji_ := request.GetString("not[my_reaction_emoji]", "")
	if not_my_reaction_emoji_ != "" {

		params.NotMyReactionEmoji = &not_my_reaction_emoji_
	}

	not_reviewer_id_ := request.GetInt("not[reviewer_id]", math.MinInt)
	if not_reviewer_id_ != math.MinInt {
		not_reviewer_id_ := int32(not_reviewer_id_)
		params.NotReviewerId = &not_reviewer_id_
	}

	deployed_before := request.GetString("deployed_before", "")
	if deployed_before != "" {

		params.DeployedBefore = &deployed_before
	}

	deployed_after := request.GetString("deployed_after", "")
	if deployed_after != "" {

		params.DeployedAfter = &deployed_after
	}

	environment := request.GetString("environment", "")
	if environment != "" {

		params.Environment = &environment
	}

	approved := request.GetString("approved", "")
	if approved != "" {

		params.Approved = &approved
	}

	merge_user_id := request.GetInt("merge_user_id", math.MinInt)
	if merge_user_id != math.MinInt {
		merge_user_id := int32(merge_user_id)
		params.MergeUserId = &merge_user_id
	}

	merge_user_username := request.GetString("merge_user_username", "")
	if merge_user_username != "" {

		params.MergeUserUsername = &merge_user_username
	}

	approver_ids := request.GetString("approver_ids", "")
	if approver_ids != "" {

		params.ApproverIds = &approver_ids
	}

	approved_by_ids := request.GetString("approved_by_ids", "")
	if approved_by_ids != "" {

		params.ApprovedByIds = &approved_by_ids
	}

	approved_by_usernames := request.GetString("approved_by_usernames", "")
	if approved_by_usernames != "" {

		params.ApprovedByUsernames = &approved_by_usernames
	}

	page := request.GetInt("page", 1)
	if page != math.MinInt {
		page := int32(page)
		params.Page = &page
	}

	per_page := request.GetInt("per_page", 20)
	if per_page != math.MinInt {
		per_page := int32(per_page)
		params.PerPage = &per_page
	}

	return params
}
