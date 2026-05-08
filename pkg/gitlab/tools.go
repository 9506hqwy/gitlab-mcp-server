package gitlab

import (
	"github.com/mark3labs/mcp-go/server"
)

func RegisterTools(s *server.MCPServer, readonly bool) {
	if !readonly {
		registerPostGroupsIdAccessRequests(s)
	}
	registerGetGroupsIdAccessRequests(s)
	if !readonly {
		registerPutGroupsIdAccessRequestsUserIdApprove(s)
	}
	if !readonly {
		registerDeleteGroupsIdAccessRequestsUserId(s)
	}
	if !readonly {
		registerPostGroupsIdEpicsEpicIidAwardEmoji(s)
	}
	registerGetGroupsIdEpicsEpicIidAwardEmoji(s)
	// if !readonly { registerDeleteGroupsIdEpicsEpicIidAwardEmojiAwardId(s) }
	// registerGetGroupsIdEpicsEpicIidAwardEmojiAwardId(s)
	// if !readonly { registerPostGroupsIdEpicsEpicIidNotesNoteIdAwardEmoji(s) }
	// registerGetGroupsIdEpicsEpicIidNotesNoteIdAwardEmoji(s)
	// if !readonly { registerDeleteGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardId(s) }
	// registerGetGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardId(s)
	if !readonly {
		registerPostGroupsIdBadges(s)
	}
	registerGetGroupsIdBadges(s)
	registerGetGroupsIdBadgesRender(s)
	if !readonly {
		registerDeleteGroupsIdBadgesBadgeId(s)
	}
	if !readonly {
		registerPutGroupsIdBadgesBadgeId(s)
	}
	registerGetGroupsIdBadgesBadgeId(s)
	registerGetGroupsIdCustomAttributes(s)
	if !readonly {
		registerDeleteGroupsIdCustomAttributesKey(s)
	}
	if !readonly {
		registerPutGroupsIdCustomAttributesKey(s)
	}
	registerGetGroupsIdCustomAttributesKey(s)
	if !readonly {
		registerPostGroups(s)
	}
	registerGetGroups(s)
	if !readonly {
		registerDeleteGroupsId(s)
	}
	if !readonly {
		registerPutGroupsId(s)
	}
	registerGetGroupsId(s)
	if !readonly {
		registerPostGroupsIdArchive(s)
	}
	if !readonly {
		registerPostGroupsIdUnarchive(s)
	}
	if !readonly {
		registerPostGroupsIdRestore(s)
	}
	registerGetGroupsIdGroupsShared(s)
	registerGetGroupsIdInvitedGroups(s)
	registerGetGroupsIdProjects(s)
	registerGetGroupsIdProjectsShared(s)
	registerGetGroupsIdSubgroups(s)
	registerGetGroupsIdDescendantGroups(s)
	if !readonly {
		registerPostGroupsIdProjectsProjectId(s)
	}
	registerGetGroupsIdTransferLocations(s)
	if !readonly {
		registerPostGroupsIdTransfer(s)
	}
	if !readonly {
		registerPostGroupsIdShare(s)
	}
	if !readonly {
		registerDeleteGroupsIdShareGroupId(s)
	}
	if !readonly {
		registerPostGroupsIdTokensRevoke(s)
	}
	if !readonly {
		registerPostGroupsIdLdapSync(s)
	}
	registerGetGroupsIdAuditEvents(s)
	registerGetGroupsIdAuditEventsAuditEventId(s)
	registerGetGroupsIdSamlUsers(s)
	registerGetGroupsIdProvisionedUsers(s)
	registerGetGroupsIdUsers(s)
	if !readonly {
		registerPostGroupsIdSshCertificates(s)
	}
	registerGetGroupsIdSshCertificates(s)
	// if !readonly { registerDeleteGroupsIdSshCertificatesSshCertificatesId(s) }
	registerGetGroupsIdRunners(s)
	if !readonly {
		registerPostGroupsIdRunnersResetRegistrationToken(s)
	}
	// registerGetGroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileName(s)
	if !readonly {
		registerDeleteGroupsIdDependencyProxyCache(s)
	}
	if !readonly {
		registerPostGroupsIdDeployTokens(s)
	}
	registerGetGroupsIdDeployTokens(s)
	if !readonly {
		registerDeleteGroupsIdDeployTokensTokenId(s)
	}
	registerGetGroupsIdDeployTokensTokenId(s)
	registerGetGroupsIdAvatar(s)
	registerGetGroupsIdClusters(s)
	if !readonly {
		registerDeleteGroupsIdClustersClusterId(s)
	}
	if !readonly {
		registerPutGroupsIdClustersClusterId(s)
	}
	registerGetGroupsIdClustersClusterId(s)
	if !readonly {
		registerPostGroupsIdClustersUser(s)
	}
	registerGetGroupsIdRegistryRepositories(s)
	if !readonly {
		registerPostGroupsIdDebianDistributions(s)
	}
	registerGetGroupsIdDebianDistributions(s)
	if !readonly {
		registerDeleteGroupsIdDebianDistributionsCodename(s)
	}
	if !readonly {
		registerPutGroupsIdDebianDistributionsCodename(s)
	}
	registerGetGroupsIdDebianDistributionsCodename(s)
	// registerGetGroupsIdDebianDistributionsCodenameKeyAsc(s)
	registerGetGroupsIdExportDownload(s)
	if !readonly {
		registerPostGroupsIdExport(s)
	}
	if !readonly {
		registerPostGroupsIdExportRelations(s)
	}
	registerGetGroupsIdExportRelationsDownload(s)
	registerGetGroupsIdExportRelationsStatus(s)
	if !readonly {
		registerPostGroupsImportAuthorize(s)
	}
	// if !readonly { registerPostGroupsImport(s) }
	registerGetGroupsIdPackages(s)
	if !readonly {
		registerPostGroupsIdPlaceholderReassignments(s)
	}
	registerGetGroupsIdPlaceholderReassignments(s)
	// if !readonly { registerPostGroupsIdPlaceholderReassignmentsAuthorize(s) }
	if !readonly {
		registerPostGroupsIdVariables(s)
	}
	registerGetGroupsIdVariables(s)
	if !readonly {
		registerDeleteGroupsIdVariablesKey(s)
	}
	if !readonly {
		registerPutGroupsIdVariablesKey(s)
	}
	registerGetGroupsIdVariablesKey(s)
	registerGetGroupsIdIntegrations(s)
	if !readonly {
		registerPutGroupsIdIntegrationsAppleAppStore(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsAsana(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsAssembla(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsBamboo(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsBugzilla(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsBuildkite(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsCampfire(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsConfluence(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsCustomIssueTracker(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsDatadog(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsDiffblueCover(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsDiscord(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsDroneCi(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsEmailsOnPush(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsExternalWiki(s)
	}
	// if !readonly { registerPutGroupsIdIntegrationsGitlabSlackApplication(s) }
	if !readonly {
		registerPutGroupsIdIntegrationsGooglePlay(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsHangoutsChat(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsHarbor(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsIrker(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsJenkins(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsJira(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsJiraCloudApp(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsMatrix(s)
	}
	// if !readonly { registerPutGroupsIdIntegrationsMattermostSlashCommands(s) }
	if !readonly {
		registerPutGroupsIdIntegrationsSlackSlashCommands(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsPackagist(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsPhorge(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsPipelinesEmail(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsPivotaltracker(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsPumble(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsPushover(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsRedmine(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsEwm(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsYoutrack(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsClickup(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsSlack(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsMicrosoftTeams(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsMattermost(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsTeamcity(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsTelegram(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsUnifyCircuit(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsWebexTeams(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsZentao(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsSquashTm(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsGithub(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsGitGuardian(s)
	}
	// if !readonly { registerPutGroupsIdIntegrationsGoogleCloudPlatformArtifactRegistry(s) }
	// if !readonly { registerPutGroupsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederation(s) }
	if !readonly {
		registerPutGroupsIdIntegrationsMockCi(s)
	}
	if !readonly {
		registerPutGroupsIdIntegrationsMockMonitoring(s)
	}
	if !readonly {
		registerDeleteGroupsIdIntegrationsSlug(s)
	}
	registerGetGroupsIdIntegrationsSlug(s)
	if !readonly {
		registerPostGroupsIdInvitations(s)
	}
	registerGetGroupsIdInvitations(s)
	if !readonly {
		registerDeleteGroupsIdInvitationsEmail(s)
	}
	if !readonly {
		registerPutGroupsIdInvitationsEmail(s)
	}
	registerGetGroupsIdUploads(s)
	if !readonly {
		registerDeleteGroupsIdUploadsUploadId(s)
	}
	registerGetGroupsIdUploadsUploadId(s)
	if !readonly {
		registerDeleteGroupsIdUploadsSecretFilename(s)
	}
	registerGetGroupsIdUploadsSecretFilename(s)
	if !readonly {
		registerPostGroupsIdMembers(s)
	}
	registerGetGroupsIdMembers(s)
	registerGetGroupsIdMembersAll(s)
	if !readonly {
		registerDeleteGroupsIdMembersUserId(s)
	}
	if !readonly {
		registerPutGroupsIdMembersUserId(s)
	}
	registerGetGroupsIdMembersUserId(s)
	registerGetGroupsIdMembersAllUserId(s)
	if !readonly {
		registerDeleteGroupsIdMembersUserIdOverride(s)
	}
	if !readonly {
		registerPostGroupsIdMembersUserIdOverride(s)
	}
	if !readonly {
		registerPutGroupsIdMembersMemberIdApprove(s)
	}
	if !readonly {
		registerPostGroupsIdMembersApproveAll(s)
	}
	registerGetGroupsIdPendingMembers(s)
	registerGetGroupsIdBillableMembers(s)
	if !readonly {
		registerPutGroupsIdMembersUserIdState(s)
	}
	// registerGetGroupsIdBillableMembersUserIdMemberships(s)
	registerGetGroupsIdBillableMembersUserIdIndirect(s)
	if !readonly {
		registerDeleteGroupsIdBillableMembersUserId(s)
	}
	registerGetGroupsIdMergeRequests(s)
	// if !readonly { registerPostGroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulk(s) }
	// if !readonly { registerPostGroupsIdPackagesNpmNpmV1SecurityAuditsQuick(s) }
	registerGetGroupsIdPackagesNugetIndex(s)
	registerGetGroupsIdPackagesNugetV2(s)
	registerGetGroupsIdPackagesNugetV2Metadata(s)
	registerGetGroupsIdPackagesNugetQuery(s)
	registerGetGroupsIdPackagesPypiSimple(s)
	registerGetGroupsIdReleases(s)
	if !readonly {
		registerPostGroupsIdAccessTokensSelfRotate(s)
	}
	if !readonly {
		registerPostGroupsIdWikis(s)
	}
	registerGetGroupsIdWikis(s)
	if !readonly {
		registerDeleteGroupsIdWikisSlug(s)
	}
	if !readonly {
		registerPutGroupsIdWikisSlug(s)
	}
	registerGetGroupsIdWikisSlug(s)
	if !readonly {
		registerPostGroupsIdWikisAttachments(s)
	}
	if !readonly {
		registerPostProjectsIdAccessRequests(s)
	}
	registerGetProjectsIdAccessRequests(s)
	if !readonly {
		registerPutProjectsIdAccessRequestsUserIdApprove(s)
	}
	if !readonly {
		registerDeleteProjectsIdAccessRequestsUserId(s)
	}
	// if !readonly { registerPostProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorize(s) }
	// if !readonly { registerPostProjectsIdAlertManagementAlertsAlertIidMetricImages(s) }
	// registerGetProjectsIdAlertManagementAlertsAlertIidMetricImages(s)
	// if !readonly { registerDeleteProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageId(s) }
	// if !readonly { registerPutProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageId(s) }
	if !readonly {
		registerPostProjectsIdIssuesIssueIidAwardEmoji(s)
	}
	registerGetProjectsIdIssuesIssueIidAwardEmoji(s)
	// if !readonly { registerDeleteProjectsIdIssuesIssueIidAwardEmojiAwardId(s) }
	// registerGetProjectsIdIssuesIssueIidAwardEmojiAwardId(s)
	// if !readonly { registerPostProjectsIdIssuesIssueIidNotesNoteIdAwardEmoji(s) }
	// registerGetProjectsIdIssuesIssueIidNotesNoteIdAwardEmoji(s)
	// if !readonly { registerDeleteProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardId(s) }
	// registerGetProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardId(s)
	if !readonly {
		registerPostProjectsIdMergeRequestsMergeRequestIidAwardEmoji(s)
	}
	registerGetProjectsIdMergeRequestsMergeRequestIidAwardEmoji(s)
	// if !readonly { registerDeleteProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardId(s) }
	// registerGetProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardId(s)
	// if !readonly { registerPostProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmoji(s) }
	// registerGetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmoji(s)
	// if !readonly { registerDeleteProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardId(s) }
	// registerGetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardId(s)
	if !readonly {
		registerPostProjectsIdSnippetsSnippetIdAwardEmoji(s)
	}
	registerGetProjectsIdSnippetsSnippetIdAwardEmoji(s)
	// if !readonly { registerDeleteProjectsIdSnippetsSnippetIdAwardEmojiAwardId(s) }
	// registerGetProjectsIdSnippetsSnippetIdAwardEmojiAwardId(s)
	// if !readonly { registerPostProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmoji(s) }
	// registerGetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmoji(s)
	// if !readonly { registerDeleteProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardId(s) }
	// registerGetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardId(s)
	if !readonly {
		registerPostProjectsIdBadges(s)
	}
	registerGetProjectsIdBadges(s)
	registerGetProjectsIdBadgesRender(s)
	if !readonly {
		registerDeleteProjectsIdBadgesBadgeId(s)
	}
	if !readonly {
		registerPutProjectsIdBadgesBadgeId(s)
	}
	registerGetProjectsIdBadgesBadgeId(s)
	if !readonly {
		registerPostProjectsIdRepositoryBranches(s)
	}
	registerGetProjectsIdRepositoryBranches(s)
	if !readonly {
		registerDeleteProjectsIdRepositoryBranchesBranch(s)
	}
	registerGetProjectsIdRepositoryBranchesBranch(s)
	if !readonly {
		registerPutProjectsIdRepositoryBranchesBranchProtect(s)
	}
	if !readonly {
		registerPutProjectsIdRepositoryBranchesBranchUnprotect(s)
	}
	if !readonly {
		registerDeleteProjectsIdRepositoryMergedBranches(s)
	}
	if !readonly {
		registerPostProjectsIdCatalogPublish(s)
	}
	registerGetProjectsIdJobsArtifactsRefNameDownload(s)
	if !readonly {
		registerDeleteProjectsIdJobsJobIdArtifacts(s)
	}
	registerGetProjectsIdJobsJobIdArtifacts(s)
	if !readonly {
		registerPostProjectsIdJobsJobIdArtifactsKeep(s)
	}
	if !readonly {
		registerDeleteProjectsIdArtifacts(s)
	}
	registerGetProjectsIdJobs(s)
	registerGetProjectsIdJobsJobId(s)
	registerGetProjectsIdJobsJobIdTrace(s)
	if !readonly {
		registerPostProjectsIdJobsJobIdCancel(s)
	}
	if !readonly {
		registerPostProjectsIdJobsJobIdRetry(s)
	}
	if !readonly {
		registerPostProjectsIdJobsJobIdErase(s)
	}
	if !readonly {
		registerPostProjectsIdJobsJobIdPlay(s)
	}
	registerGetProjectsIdResourceGroups(s)
	if !readonly {
		registerPutProjectsIdResourceGroupsKey(s)
	}
	registerGetProjectsIdResourceGroupsKey(s)
	registerGetProjectsIdResourceGroupsKeyUpcomingJobs(s)
	if !readonly {
		registerPostProjectsIdRunners(s)
	}
	registerGetProjectsIdRunners(s)
	if !readonly {
		registerDeleteProjectsIdRunnersRunnerId(s)
	}
	if !readonly {
		registerPostProjectsIdRunnersResetRegistrationToken(s)
	}
	if !readonly {
		registerPostProjectsIdSecureFiles(s)
	}
	registerGetProjectsIdSecureFiles(s)
	if !readonly {
		registerDeleteProjectsIdSecureFilesSecureFileId(s)
	}
	registerGetProjectsIdSecureFilesSecureFileId(s)
	// registerGetProjectsIdSecureFilesSecureFileIdDownload(s)
	registerGetProjectsIdPipelines(s)
	if !readonly {
		registerPostProjectsIdPipeline(s)
	}
	registerGetProjectsIdPipelinesLatest(s)
	if !readonly {
		registerDeleteProjectsIdPipelinesPipelineId(s)
	}
	registerGetProjectsIdPipelinesPipelineId(s)
	registerGetProjectsIdPipelinesPipelineIdJobs(s)
	registerGetProjectsIdPipelinesPipelineIdBridges(s)
	registerGetProjectsIdPipelinesPipelineIdVariables(s)
	registerGetProjectsIdPipelinesPipelineIdTestReport(s)
	registerGetProjectsIdPipelinesPipelineIdTestReportSummary(s)
	if !readonly {
		registerPutProjectsIdPipelinesPipelineIdMetadata(s)
	}
	if !readonly {
		registerPostProjectsIdPipelinesPipelineIdRetry(s)
	}
	if !readonly {
		registerPostProjectsIdPipelinesPipelineIdCancel(s)
	}
	if !readonly {
		registerPostProjectsIdPipelineSchedules(s)
	}
	registerGetProjectsIdPipelineSchedules(s)
	// if !readonly { registerDeleteProjectsIdPipelineSchedulesPipelineScheduleId(s) }
	// if !readonly { registerPutProjectsIdPipelineSchedulesPipelineScheduleId(s) }
	// registerGetProjectsIdPipelineSchedulesPipelineScheduleId(s)
	// registerGetProjectsIdPipelineSchedulesPipelineScheduleIdPipelines(s)
	// if !readonly { registerPostProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnership(s) }
	// if !readonly { registerPostProjectsIdPipelineSchedulesPipelineScheduleIdPlay(s) }
	// if !readonly { registerPostProjectsIdPipelineSchedulesPipelineScheduleIdVariables(s) }
	// if !readonly { registerDeleteProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey(s) }
	// if !readonly { registerPutProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey(s) }
	if !readonly {
		registerPostProjectsIdRefRefTriggerPipeline(s)
	}
	if !readonly {
		registerPostProjectsIdTriggers(s)
	}
	registerGetProjectsIdTriggers(s)
	if !readonly {
		registerDeleteProjectsIdTriggersTriggerId(s)
	}
	if !readonly {
		registerPutProjectsIdTriggersTriggerId(s)
	}
	registerGetProjectsIdTriggersTriggerId(s)
	if !readonly {
		registerPostProjectsIdVariables(s)
	}
	registerGetProjectsIdVariables(s)
	if !readonly {
		registerDeleteProjectsIdVariablesKey(s)
	}
	if !readonly {
		registerPutProjectsIdVariablesKey(s)
	}
	registerGetProjectsIdVariablesKey(s)
	if !readonly {
		registerPostProjectsIdClusterAgentsAgentIdTokens(s)
	}
	registerGetProjectsIdClusterAgentsAgentIdTokens(s)
	// if !readonly { registerDeleteProjectsIdClusterAgentsAgentIdTokensTokenId(s) }
	// registerGetProjectsIdClusterAgentsAgentIdTokensTokenId(s)
	if !readonly {
		registerPostProjectsIdClusterAgents(s)
	}
	registerGetProjectsIdClusterAgents(s)
	if !readonly {
		registerDeleteProjectsIdClusterAgentsAgentId(s)
	}
	registerGetProjectsIdClusterAgentsAgentId(s)
	registerGetProjectsIdPackagesCargoConfigJson(s)
	if !readonly {
		registerPostProjectsIdRepositoryCommits(s)
	}
	registerGetProjectsIdRepositoryCommits(s)
	registerGetProjectsIdRepositoryCommitsSha(s)
	registerGetProjectsIdRepositoryCommitsShaDiff(s)
	if !readonly {
		registerPostProjectsIdRepositoryCommitsShaComments(s)
	}
	registerGetProjectsIdRepositoryCommitsShaComments(s)
	registerGetProjectsIdRepositoryCommitsShaSequence(s)
	if !readonly {
		registerPostProjectsIdRepositoryCommitsShaCherryPick(s)
	}
	if !readonly {
		registerPostProjectsIdRepositoryCommitsShaRevert(s)
	}
	registerGetProjectsIdRepositoryCommitsShaRefs(s)
	registerGetProjectsIdRepositoryCommitsShaMergeRequests(s)
	registerGetProjectsIdRepositoryCommitsShaSignature(s)
	registerGetProjectsIdRepositoryCommitsShaStatuses(s)
	if !readonly {
		registerPostProjectsIdStatusesSha(s)
	}
	if !readonly {
		registerPostProjectsIdPackagesComposer(s)
	}
	registerGetProjectsIdPackagesConanV1UsersAuthenticate(s)
	// registerGetProjectsIdPackagesConanV1UsersCheckCredentials(s)
	registerGetProjectsIdPackagesConanV1ConansSearch(s)
	// registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(s)
	registerGetProjectsIdPackagesConanV1Ping(s)
	// registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference(s)
	// if !readonly { registerDeleteProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(s) }
	// registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(s)
	// registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest(s)
	// registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest(s)
	// registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls(s)
	// registerGetProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls(s)
	// if !readonly { registerPostProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrls(s) }
	// if !readonly { registerPostProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrls(s) }
	// if !readonly { registerPutProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(s) }
	// registerGetProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(s)
	// if !readonly { registerPutProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize(s) }
	// if !readonly { registerPutProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(s) }
	// registerGetProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(s)
	// if !readonly { registerPutProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorize(s) }
	registerGetProjectsIdPackagesConanV2UsersAuthenticate(s)
	// registerGetProjectsIdPackagesConanV2UsersCheckCredentials(s)
	registerGetProjectsIdPackagesConanV2ConansSearch(s)
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(s)
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatest(s)
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisions(s)
	// if !readonly { registerDeleteProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevision(s) }
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFiles(s)
	// if !readonly { registerPutProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileName(s) }
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileName(s)
	// if !readonly { registerPutProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameAuthorize(s) }
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearch(s)
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatest(s)
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisions(s)
	// if !readonly { registerDeleteProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevision(s) }
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFiles(s)
	// if !readonly { registerPutProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileName(s) }
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileName(s)
	// if !readonly { registerPutProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameAuthorize(s) }
	// registerGetProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileName(s)
	if !readonly {
		registerPutProjectsIdPackagesDebianFileName(s)
	}
	if !readonly {
		registerPutProjectsIdPackagesDebianFileNameAuthorize(s)
	}
	if !readonly {
		registerPostProjectsIdDeployKeys(s)
	}
	registerGetProjectsIdDeployKeys(s)
	if !readonly {
		registerDeleteProjectsIdDeployKeysKeyId(s)
	}
	if !readonly {
		registerPutProjectsIdDeployKeysKeyId(s)
	}
	registerGetProjectsIdDeployKeysKeyId(s)
	if !readonly {
		registerPostProjectsIdDeployKeysKeyIdEnable(s)
	}
	if !readonly {
		registerPostProjectsIdDeployTokens(s)
	}
	registerGetProjectsIdDeployTokens(s)
	if !readonly {
		registerDeleteProjectsIdDeployTokensTokenId(s)
	}
	registerGetProjectsIdDeployTokensTokenId(s)
	if !readonly {
		registerPostProjectsIdDeployments(s)
	}
	registerGetProjectsIdDeployments(s)
	if !readonly {
		registerDeleteProjectsIdDeploymentsDeploymentId(s)
	}
	if !readonly {
		registerPutProjectsIdDeploymentsDeploymentId(s)
	}
	registerGetProjectsIdDeploymentsDeploymentId(s)
	registerGetProjectsIdDeploymentsDeploymentIdMergeRequests(s)
	if !readonly {
		registerPostProjectsIdDeploymentsDeploymentIdApproval(s)
	}
	if !readonly {
		registerPostProjectsIdMergeRequestsMergeRequestIidDraftNotes(s)
	}
	registerGetProjectsIdMergeRequestsMergeRequestIidDraftNotes(s)
	// if !readonly { registerDeleteProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(s) }
	// if !readonly { registerPutProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(s) }
	// registerGetProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(s)
	// if !readonly { registerPutProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublish(s) }
	// if !readonly { registerPostProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublish(s) }
	if !readonly {
		registerPostProjectsIdEnvironments(s)
	}
	registerGetProjectsIdEnvironments(s)
	if !readonly {
		registerDeleteProjectsIdEnvironmentsEnvironmentId(s)
	}
	if !readonly {
		registerPutProjectsIdEnvironmentsEnvironmentId(s)
	}
	registerGetProjectsIdEnvironmentsEnvironmentId(s)
	if !readonly {
		registerDeleteProjectsIdEnvironmentsReviewApps(s)
	}
	if !readonly {
		registerPostProjectsIdEnvironmentsEnvironmentIdStop(s)
	}
	if !readonly {
		registerPostProjectsIdEnvironmentsStopStale(s)
	}
	if !readonly {
		registerPostProjectsIdErrorTrackingClientKeys(s)
	}
	registerGetProjectsIdErrorTrackingClientKeys(s)
	// if !readonly { registerDeleteProjectsIdErrorTrackingClientKeysKeyId(s) }
	if !readonly {
		registerPutProjectsIdErrorTrackingSettings(s)
	}
	registerGetProjectsIdErrorTrackingSettings(s)
	if !readonly {
		registerPostProjectsIdFeatureFlags(s)
	}
	registerGetProjectsIdFeatureFlags(s)
	if !readonly {
		registerDeleteProjectsIdFeatureFlagsFeatureFlagName(s)
	}
	if !readonly {
		registerPutProjectsIdFeatureFlagsFeatureFlagName(s)
	}
	registerGetProjectsIdFeatureFlagsFeatureFlagName(s)
	if !readonly {
		registerPostProjectsIdFeatureFlagsUserLists(s)
	}
	registerGetProjectsIdFeatureFlagsUserLists(s)
	if !readonly {
		registerDeleteProjectsIdFeatureFlagsUserListsIid(s)
	}
	if !readonly {
		registerPutProjectsIdFeatureFlagsUserListsIid(s)
	}
	registerGetProjectsIdFeatureFlagsUserListsIid(s)
	registerGetProjectsIdRepositoryFilesFilePathBlame(s)
	registerGetProjectsIdRepositoryFilesFilePathRaw(s)
	if !readonly {
		registerDeleteProjectsIdRepositoryFilesFilePath(s)
	}
	if !readonly {
		registerPostProjectsIdRepositoryFilesFilePath(s)
	}
	if !readonly {
		registerPutProjectsIdRepositoryFilesFilePath(s)
	}
	registerGetProjectsIdRepositoryFilesFilePath(s)
	if !readonly {
		registerPostProjectsIdFreezePeriods(s)
	}
	registerGetProjectsIdFreezePeriods(s)
	if !readonly {
		registerDeleteProjectsIdFreezePeriodsFreezePeriodId(s)
	}
	if !readonly {
		registerPutProjectsIdFreezePeriodsFreezePeriodId(s)
	}
	registerGetProjectsIdFreezePeriodsFreezePeriodId(s)
	registerGetProjectsIdPackagesHelmChannelIndexYaml(s)
	// registerGetProjectsIdPackagesHelmChannelChartsFileNameTgz(s)
	// if !readonly { registerPostProjectsIdPackagesHelmApiChannelChartsAuthorize(s) }
	if !readonly {
		registerPostProjectsIdPackagesHelmApiChannelCharts(s)
	}
	registerGetProjectsIdServices(s)
	if !readonly {
		registerPutProjectsIdServicesAppleAppStore(s)
	}
	if !readonly {
		registerPutProjectsIdServicesAsana(s)
	}
	if !readonly {
		registerPutProjectsIdServicesAssembla(s)
	}
	if !readonly {
		registerPutProjectsIdServicesBamboo(s)
	}
	if !readonly {
		registerPutProjectsIdServicesBugzilla(s)
	}
	if !readonly {
		registerPutProjectsIdServicesBuildkite(s)
	}
	if !readonly {
		registerPutProjectsIdServicesCampfire(s)
	}
	if !readonly {
		registerPutProjectsIdServicesConfluence(s)
	}
	if !readonly {
		registerPutProjectsIdServicesCustomIssueTracker(s)
	}
	if !readonly {
		registerPutProjectsIdServicesDatadog(s)
	}
	if !readonly {
		registerPutProjectsIdServicesDiffblueCover(s)
	}
	if !readonly {
		registerPutProjectsIdServicesDiscord(s)
	}
	if !readonly {
		registerPutProjectsIdServicesDroneCi(s)
	}
	if !readonly {
		registerPutProjectsIdServicesEmailsOnPush(s)
	}
	if !readonly {
		registerPutProjectsIdServicesExternalWiki(s)
	}
	if !readonly {
		registerPutProjectsIdServicesGitlabSlackApplication(s)
	}
	if !readonly {
		registerPutProjectsIdServicesGooglePlay(s)
	}
	if !readonly {
		registerPutProjectsIdServicesHangoutsChat(s)
	}
	if !readonly {
		registerPutProjectsIdServicesHarbor(s)
	}
	if !readonly {
		registerPutProjectsIdServicesIrker(s)
	}
	if !readonly {
		registerPutProjectsIdServicesJenkins(s)
	}
	if !readonly {
		registerPutProjectsIdServicesJira(s)
	}
	if !readonly {
		registerPutProjectsIdServicesJiraCloudApp(s)
	}
	if !readonly {
		registerPutProjectsIdServicesMatrix(s)
	}
	if !readonly {
		registerPutProjectsIdServicesMattermostSlashCommands(s)
	}
	if !readonly {
		registerPutProjectsIdServicesSlackSlashCommands(s)
	}
	if !readonly {
		registerPutProjectsIdServicesPackagist(s)
	}
	if !readonly {
		registerPutProjectsIdServicesPhorge(s)
	}
	if !readonly {
		registerPutProjectsIdServicesPipelinesEmail(s)
	}
	if !readonly {
		registerPutProjectsIdServicesPivotaltracker(s)
	}
	if !readonly {
		registerPutProjectsIdServicesPumble(s)
	}
	if !readonly {
		registerPutProjectsIdServicesPushover(s)
	}
	if !readonly {
		registerPutProjectsIdServicesRedmine(s)
	}
	if !readonly {
		registerPutProjectsIdServicesEwm(s)
	}
	if !readonly {
		registerPutProjectsIdServicesYoutrack(s)
	}
	if !readonly {
		registerPutProjectsIdServicesClickup(s)
	}
	if !readonly {
		registerPutProjectsIdServicesSlack(s)
	}
	if !readonly {
		registerPutProjectsIdServicesMicrosoftTeams(s)
	}
	if !readonly {
		registerPutProjectsIdServicesMattermost(s)
	}
	if !readonly {
		registerPutProjectsIdServicesTeamcity(s)
	}
	if !readonly {
		registerPutProjectsIdServicesTelegram(s)
	}
	if !readonly {
		registerPutProjectsIdServicesUnifyCircuit(s)
	}
	if !readonly {
		registerPutProjectsIdServicesWebexTeams(s)
	}
	if !readonly {
		registerPutProjectsIdServicesZentao(s)
	}
	if !readonly {
		registerPutProjectsIdServicesSquashTm(s)
	}
	if !readonly {
		registerPutProjectsIdServicesGithub(s)
	}
	if !readonly {
		registerPutProjectsIdServicesGitGuardian(s)
	}
	// if !readonly { registerPutProjectsIdServicesGoogleCloudPlatformArtifactRegistry(s) }
	// if !readonly { registerPutProjectsIdServicesGoogleCloudPlatformWorkloadIdentityFederation(s) }
	if !readonly {
		registerPutProjectsIdServicesMockCi(s)
	}
	if !readonly {
		registerPutProjectsIdServicesMockMonitoring(s)
	}
	if !readonly {
		registerDeleteProjectsIdServicesSlug(s)
	}
	registerGetProjectsIdServicesSlug(s)
	// if !readonly { registerPostProjectsIdServicesMattermostSlashCommandsTrigger(s) }
	// if !readonly { registerPostProjectsIdServicesSlackSlashCommandsTrigger(s) }
	registerGetProjectsIdIntegrations(s)
	if !readonly {
		registerPutProjectsIdIntegrationsAppleAppStore(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsAsana(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsAssembla(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsBamboo(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsBugzilla(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsBuildkite(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsCampfire(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsConfluence(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsCustomIssueTracker(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsDatadog(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsDiffblueCover(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsDiscord(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsDroneCi(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsEmailsOnPush(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsExternalWiki(s)
	}
	// if !readonly { registerPutProjectsIdIntegrationsGitlabSlackApplication(s) }
	if !readonly {
		registerPutProjectsIdIntegrationsGooglePlay(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsHangoutsChat(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsHarbor(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsIrker(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsJenkins(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsJira(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsJiraCloudApp(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsMatrix(s)
	}
	// if !readonly { registerPutProjectsIdIntegrationsMattermostSlashCommands(s) }
	if !readonly {
		registerPutProjectsIdIntegrationsSlackSlashCommands(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsPackagist(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsPhorge(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsPipelinesEmail(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsPivotaltracker(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsPumble(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsPushover(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsRedmine(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsEwm(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsYoutrack(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsClickup(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsSlack(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsMicrosoftTeams(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsMattermost(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsTeamcity(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsTelegram(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsUnifyCircuit(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsWebexTeams(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsZentao(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsSquashTm(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsGithub(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsGitGuardian(s)
	}
	// if !readonly { registerPutProjectsIdIntegrationsGoogleCloudPlatformArtifactRegistry(s) }
	// if !readonly { registerPutProjectsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederation(s) }
	if !readonly {
		registerPutProjectsIdIntegrationsMockCi(s)
	}
	if !readonly {
		registerPutProjectsIdIntegrationsMockMonitoring(s)
	}
	if !readonly {
		registerDeleteProjectsIdIntegrationsSlug(s)
	}
	registerGetProjectsIdIntegrationsSlug(s)
	// if !readonly { registerPostProjectsIdIntegrationsMattermostSlashCommandsTrigger(s) }
	// if !readonly { registerPostProjectsIdIntegrationsSlackSlashCommandsTrigger(s) }
	if !readonly {
		registerPostProjectsIdInvitations(s)
	}
	registerGetProjectsIdInvitations(s)
	if !readonly {
		registerDeleteProjectsIdInvitationsEmail(s)
	}
	if !readonly {
		registerPutProjectsIdInvitationsEmail(s)
	}
	if !readonly {
		registerPostProjectsIdIssuesIssueIidLinks(s)
	}
	registerGetProjectsIdIssuesIssueIidLinks(s)
	// if !readonly { registerDeleteProjectsIdIssuesIssueIidLinksIssueLinkId(s) }
	// registerGetProjectsIdIssuesIssueIidLinksIssueLinkId(s)
	if !readonly {
		registerPostProjectsIdCiLint(s)
	}
	registerGetProjectsIdCiLint(s)
	if !readonly {
		registerPostProjectsIdUploadsAuthorize(s)
	}
	if !readonly {
		registerPostProjectsIdUploads(s)
	}
	registerGetProjectsIdUploads(s)
	if !readonly {
		registerDeleteProjectsIdUploadsUploadId(s)
	}
	registerGetProjectsIdUploadsUploadId(s)
	if !readonly {
		registerDeleteProjectsIdUploadsSecretFilename(s)
	}
	registerGetProjectsIdUploadsSecretFilename(s)
	if !readonly {
		registerPostProjectsIdMembers(s)
	}
	registerGetProjectsIdMembers(s)
	registerGetProjectsIdMembersAll(s)
	if !readonly {
		registerDeleteProjectsIdMembersUserId(s)
	}
	if !readonly {
		registerPutProjectsIdMembersUserId(s)
	}
	registerGetProjectsIdMembersUserId(s)
	registerGetProjectsIdMembersAllUserId(s)
	if !readonly {
		registerPostProjectsIdMergeRequestsMergeRequestIidApprovals(s)
	}
	registerGetProjectsIdMergeRequestsMergeRequestIidApprovals(s)
	if !readonly {
		registerPostProjectsIdMergeRequestsMergeRequestIidApprove(s)
	}
	if !readonly {
		registerPostProjectsIdMergeRequestsMergeRequestIidUnapprove(s)
	}
	// if !readonly { registerPutProjectsIdMergeRequestsMergeRequestIidResetApprovals(s) }
	// registerGetProjectsIdMergeRequestsMergeRequestIidApprovalState(s)
	if !readonly {
		registerPostProjectsIdCreateCiConfig(s)
	}
	// if !readonly { registerPostProjectsIdMergeRequestsMergeRequestIidTimeEstimate(s) }
	// if !readonly { registerPostProjectsIdMergeRequestsMergeRequestIidResetTimeEstimate(s) }
	// if !readonly { registerPostProjectsIdMergeRequestsMergeRequestIidAddSpentTime(s) }
	// if !readonly { registerPostProjectsIdMergeRequestsMergeRequestIidResetSpentTime(s) }
	registerGetProjectsIdMergeRequestsMergeRequestIidTimeStats(s)
	if !readonly {
		registerPostProjectsIdMergeRequests(s)
	}
	registerGetProjectsIdMergeRequests(s)
	if !readonly {
		registerDeleteProjectsIdMergeRequestsMergeRequestIid(s)
	}
	if !readonly {
		registerPutProjectsIdMergeRequestsMergeRequestIid(s)
	}
	registerGetProjectsIdMergeRequestsMergeRequestIid(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidParticipants(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidReviewers(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidCommits(s)
	// if !readonly { registerDeleteProjectsIdMergeRequestsMergeRequestIidContextCommits(s) }
	// if !readonly { registerPostProjectsIdMergeRequestsMergeRequestIidContextCommits(s) }
	// registerGetProjectsIdMergeRequestsMergeRequestIidContextCommits(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidChanges(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidDiffs(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidRawDiffs(s)
	if !readonly {
		registerPostProjectsIdMergeRequestsMergeRequestIidPipelines(s)
	}
	registerGetProjectsIdMergeRequestsMergeRequestIidPipelines(s)
	if !readonly {
		registerPutProjectsIdMergeRequestsMergeRequestIidMerge(s)
	}
	registerGetProjectsIdMergeRequestsMergeRequestIidMergeRef(s)
	// if !readonly { registerPostProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceeds(s) }
	if !readonly {
		registerPutProjectsIdMergeRequestsMergeRequestIidRebase(s)
	}
	registerGetProjectsIdMergeRequestsMergeRequestIidClosesIssues(s)
	// registerGetProjectsIdMergeRequestsMergeRequestIidRelatedIssues(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidVersions(s)
	// registerGetProjectsIdMergeRequestsMergeRequestIidVersionsVersionId(s)
	// if !readonly { registerPostProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulk(s) }
	// if !readonly { registerPostProjectsIdPackagesNpmNpmV1SecurityAuditsQuick(s) }
	if !readonly {
		registerPutProjectsIdPackagesNpmPackageName(s)
	}
	registerGetProjectsIdPackagesNugetIndex(s)
	if !readonly {
		registerPutProjectsIdPackagesNugetV2(s)
	}
	registerGetProjectsIdPackagesNugetV2(s)
	registerGetProjectsIdPackagesNugetV2Metadata(s)
	registerGetProjectsIdPackagesNugetQuery(s)
	if !readonly {
		registerPutProjectsIdPackagesNuget(s)
	}
	if !readonly {
		registerPutProjectsIdPackagesNugetAuthorize(s)
	}
	if !readonly {
		registerPutProjectsIdPackagesNugetSymbolpackage(s)
	}
	if !readonly {
		registerPutProjectsIdPackagesNugetSymbolpackageAuthorize(s)
	}
	if !readonly {
		registerPutProjectsIdPackagesNugetV2Authorize(s)
	}
	registerGetProjectsIdPackagesPackageIdPackageFiles(s)
	// if !readonly { registerDeleteProjectsIdPackagesPackageIdPackageFilesPackageFileId(s) }
	if !readonly {
		registerDeleteProjectsIdPages(s)
	}
	registerGetProjectsIdPages(s)
	if !readonly {
		registerPostProjectsIdPagesDomains(s)
	}
	registerGetProjectsIdPagesDomains(s)
	if !readonly {
		registerDeleteProjectsIdPagesDomainsDomain(s)
	}
	if !readonly {
		registerPutProjectsIdPagesDomainsDomain(s)
	}
	registerGetProjectsIdPagesDomainsDomain(s)
	if !readonly {
		registerPutProjectsIdPagesDomainsDomainVerify(s)
	}
	registerGetProjectsIdAvatar(s)
	registerGetProjectsIdClusters(s)
	if !readonly {
		registerDeleteProjectsIdClustersClusterId(s)
	}
	if !readonly {
		registerPutProjectsIdClustersClusterId(s)
	}
	registerGetProjectsIdClustersClusterId(s)
	if !readonly {
		registerPostProjectsIdClustersUser(s)
	}
	registerGetProjectsIdRegistryRepositories(s)
	if !readonly {
		registerDeleteProjectsIdRegistryRepositoriesRepositoryId(s)
	}
	// if !readonly { registerDeleteProjectsIdRegistryRepositoriesRepositoryIdTags(s) }
	registerGetProjectsIdRegistryRepositoriesRepositoryIdTags(s)
	// if !readonly { registerDeleteProjectsIdRegistryRepositoriesRepositoryIdTagsTagName(s) }
	// registerGetProjectsIdRegistryRepositoriesRepositoryIdTagsTagName(s)
	if !readonly {
		registerPostProjectsIdRegistryProtectionRepositoryRules(s)
	}
	registerGetProjectsIdRegistryProtectionRepositoryRules(s)
	// if !readonly { registerDeleteProjectsIdRegistryProtectionRepositoryRulesProtectionRuleId(s) }
	if !readonly {
		registerPostProjectsIdDebianDistributions(s)
	}
	registerGetProjectsIdDebianDistributions(s)
	if !readonly {
		registerDeleteProjectsIdDebianDistributionsCodename(s)
	}
	if !readonly {
		registerPutProjectsIdDebianDistributionsCodename(s)
	}
	registerGetProjectsIdDebianDistributionsCodename(s)
	// registerGetProjectsIdDebianDistributionsCodenameKeyAsc(s)
	registerGetProjectsIdEvents(s)
	if !readonly {
		registerPostProjectsIdExport(s)
	}
	registerGetProjectsIdExport(s)
	registerGetProjectsIdExportDownload(s)
	if !readonly {
		registerPostProjectsIdExportRelations(s)
	}
	registerGetProjectsIdExportRelationsDownload(s)
	registerGetProjectsIdExportRelationsStatus(s)
	if !readonly {
		registerDeleteProjectsIdHooksHookIdUrlVariablesKey(s)
	}
	if !readonly {
		registerPutProjectsIdHooksHookIdUrlVariablesKey(s)
	}
	if !readonly {
		registerDeleteProjectsIdHooksHookIdCustomHeadersKey(s)
	}
	if !readonly {
		registerPutProjectsIdHooksHookIdCustomHeadersKey(s)
	}
	if !readonly {
		registerPostProjectsIdHooks(s)
	}
	registerGetProjectsIdHooks(s)
	if !readonly {
		registerDeleteProjectsIdHooksHookId(s)
	}
	if !readonly {
		registerPutProjectsIdHooksHookId(s)
	}
	registerGetProjectsIdHooksHookId(s)
	registerGetProjectsIdHooksHookIdEvents(s)
	if !readonly {
		registerPostProjectsIdHooksHookIdTestTrigger(s)
	}
	// if !readonly { registerPostProjectsIdHooksHookIdEventsHookLogIdResend(s) }
	if !readonly {
		registerPostProjectsImportAuthorize(s)
	}
	// if !readonly { registerPostProjectsImport(s) }
	registerGetProjectsIdImport(s)
	// if !readonly { registerPostProjectsRemoteImport(s) }
	if !readonly {
		registerPostProjectsImportRelationAuthorize(s)
	}
	// if !readonly { registerPostProjectsImportRelation(s) }
	registerGetProjectsIdRelationImports(s)
	// if !readonly { registerPostProjectsRemoteImportS3(s) }
	registerGetProjectsIdJobTokenScope(s)
	if !readonly {
		registerPostProjectsIdJobTokenScopeAllowlist(s)
	}
	registerGetProjectsIdJobTokenScopeAllowlist(s)
	if !readonly {
		registerPostProjectsIdJobTokenScopeGroupsAllowlist(s)
	}
	registerGetProjectsIdJobTokenScopeGroupsAllowlist(s)
	// if !readonly { registerDeleteProjectsIdJobTokenScopeGroupsAllowlistTargetGroupId(s) }
	// if !readonly { registerDeleteProjectsIdJobTokenScopeAllowlistTargetProjectId(s) }
	registerGetProjectsIdPackages(s)
	if !readonly {
		registerDeleteProjectsIdPackagesPackageId(s)
	}
	registerGetProjectsIdPackagesPackageId(s)
	registerGetProjectsIdPackagesPackageIdPipelines(s)
	if !readonly {
		registerPostProjectsIdPackagesProtectionRules(s)
	}
	registerGetProjectsIdPackagesProtectionRules(s)
	// if !readonly { registerDeleteProjectsIdPackagesProtectionRulesPackageProtectionRuleId(s) }
	registerGetProjectsIdSnapshot(s)
	if !readonly {
		registerPostProjectsIdSnippets(s)
	}
	registerGetProjectsIdSnippets(s)
	if !readonly {
		registerDeleteProjectsIdSnippetsSnippetId(s)
	}
	if !readonly {
		registerPutProjectsIdSnippetsSnippetId(s)
	}
	registerGetProjectsIdSnippetsSnippetId(s)
	registerGetProjectsIdSnippetsSnippetIdRaw(s)
	// registerGetProjectsIdSnippetsSnippetIdFilesRefFilePathRaw(s)
	// registerGetProjectsIdSnippetsSnippetIdUserAgentDetail(s)
	registerGetProjectsIdStatistics(s)
	registerGetProjectsIdTemplatesType(s)
	registerGetProjectsIdTemplatesTypeName(s)
	registerGetProjectsIdCustomAttributes(s)
	if !readonly {
		registerDeleteProjectsIdCustomAttributesKey(s)
	}
	if !readonly {
		registerPutProjectsIdCustomAttributesKey(s)
	}
	registerGetProjectsIdCustomAttributesKey(s)
	if !readonly {
		registerPostProjectsIdRestore(s)
	}
	if !readonly {
		registerPostProjects(s)
	}
	registerGetProjects(s)
	if !readonly {
		registerPostProjectsUserUserId(s)
	}
	registerGetProjectsIdShareLocations(s)
	if !readonly {
		registerDeleteProjectsId(s)
	}
	if !readonly {
		registerPutProjectsId(s)
	}
	registerGetProjectsId(s)
	if !readonly {
		registerDeleteProjectsIdFork(s)
	}
	if !readonly {
		registerPostProjectsIdFork(s)
	}
	registerGetProjectsIdForks(s)
	registerGetProjectsIdPagesAccess(s)
	if !readonly {
		registerPostProjectsIdArchive(s)
	}
	if !readonly {
		registerPostProjectsIdUnarchive(s)
	}
	if !readonly {
		registerPostProjectsIdStar(s)
	}
	if !readonly {
		registerPostProjectsIdUnstar(s)
	}
	registerGetProjectsIdStarrers(s)
	registerGetProjectsIdLanguages(s)
	if !readonly {
		registerPostProjectsIdForkForkedFromId(s)
	}
	if !readonly {
		registerPostProjectsIdShare(s)
	}
	if !readonly {
		registerDeleteProjectsIdShareGroupId(s)
	}
	if !readonly {
		registerPostProjectsIdImportProjectMembersProjectId(s)
	}
	registerGetProjectsIdUsers(s)
	registerGetProjectsIdGroups(s)
	registerGetProjectsIdInvitedGroups(s)
	if !readonly {
		registerPostProjectsIdHousekeeping(s)
	}
	if !readonly {
		registerPostProjectsIdRepositorySize(s)
	}
	if !readonly {
		registerPutProjectsIdTransfer(s)
	}
	registerGetProjectsIdTransferLocations(s)
	registerGetProjectsIdStorage(s)
	registerGetProjectsIdAuditEvents(s)
	registerGetProjectsIdAuditEventsAuditEventId(s)
	if !readonly {
		registerPostProjectsIdProtectedBranches(s)
	}
	registerGetProjectsIdProtectedBranches(s)
	if !readonly {
		registerDeleteProjectsIdProtectedBranchesName(s)
	}
	registerGetProjectsIdProtectedBranchesName(s)
	if !readonly {
		registerPostProjectsIdProtectedTags(s)
	}
	registerGetProjectsIdProtectedTags(s)
	if !readonly {
		registerDeleteProjectsIdProtectedTagsName(s)
	}
	registerGetProjectsIdProtectedTagsName(s)
	registerGetProjectsIdPackagesPypiSimple(s)
	if !readonly {
		registerPostProjectsIdPackagesPypi(s)
	}
	if !readonly {
		registerPostProjectsIdPackagesPypiAuthorize(s)
	}
	if !readonly {
		registerPostProjectsIdReleases(s)
	}
	registerGetProjectsIdReleases(s)
	if !readonly {
		registerDeleteProjectsIdReleasesTagName(s)
	}
	if !readonly {
		registerPutProjectsIdReleasesTagName(s)
	}
	registerGetProjectsIdReleasesTagName(s)
	if !readonly {
		registerPostProjectsIdReleasesTagNameEvidence(s)
	}
	if !readonly {
		registerPostProjectsIdReleasesTagNameAssetsLinks(s)
	}
	registerGetProjectsIdReleasesTagNameAssetsLinks(s)
	// if !readonly { registerDeleteProjectsIdReleasesTagNameAssetsLinksLinkId(s) }
	// if !readonly { registerPutProjectsIdReleasesTagNameAssetsLinksLinkId(s) }
	// registerGetProjectsIdReleasesTagNameAssetsLinksLinkId(s)
	if !readonly {
		registerPostProjectsIdRemoteMirrors(s)
	}
	registerGetProjectsIdRemoteMirrors(s)
	if !readonly {
		registerDeleteProjectsIdRemoteMirrorsMirrorId(s)
	}
	if !readonly {
		registerPutProjectsIdRemoteMirrorsMirrorId(s)
	}
	registerGetProjectsIdRemoteMirrorsMirrorId(s)
	if !readonly {
		registerPostProjectsIdRemoteMirrorsMirrorIdSync(s)
	}
	registerGetProjectsIdRemoteMirrorsMirrorIdPublicKey(s)
	registerGetProjectsIdRepositoryTree(s)
	registerGetProjectsIdRepositoryBlobsShaRaw(s)
	registerGetProjectsIdRepositoryBlobsSha(s)
	registerGetProjectsIdRepositoryArchive(s)
	registerGetProjectsIdRepositoryCompare(s)
	registerGetProjectsIdRepositoryHealth(s)
	registerGetProjectsIdRepositoryContributors(s)
	registerGetProjectsIdRepositoryMergeBase(s)
	if !readonly {
		registerPostProjectsIdRepositoryChangelog(s)
	}
	registerGetProjectsIdRepositoryChangelog(s)
	if !readonly {
		registerPostProjectsIdAccessTokensSelfRotate(s)
	}
	// registerGetProjectsIdIssuesEventableIdResourceMilestoneEvents(s)
	// registerGetProjectsIdIssuesEventableIdResourceMilestoneEventsEventId(s)
	// registerGetProjectsIdMergeRequestsEventableIdResourceMilestoneEvents(s)
	// registerGetProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventId(s)
	if !readonly {
		registerPostProjectsIdPackagesRpm(s)
	}
	if !readonly {
		registerPostProjectsIdPackagesRpmAuthorize(s)
	}
	registerGetProjectsIdPackagesRubygemsFileName(s)
	// registerGetProjectsIdPackagesRubygemsQuickMarshal48FileName(s)
	registerGetProjectsIdPackagesRubygemsGemsFileName(s)
	// if !readonly { registerPostProjectsIdPackagesRubygemsApiV1GemsAuthorize(s) }
	if !readonly {
		registerPostProjectsIdPackagesRubygemsApiV1Gems(s)
	}
	registerGetProjectsIdPackagesRubygemsApiV1Dependencies(s)
	if !readonly {
		registerPutProjectsIdRepositorySubmodulesSubmodule(s)
	}
	if !readonly {
		registerPostProjectsIdRepositoryTags(s)
	}
	registerGetProjectsIdRepositoryTags(s)
	if !readonly {
		registerDeleteProjectsIdRepositoryTagsTagName(s)
	}
	registerGetProjectsIdRepositoryTagsTagName(s)
	registerGetProjectsIdRepositoryTagsTagNameSignature(s)
	// registerGetProjectsIdPackagesTerraformModulesModuleNameModuleSystem(s)
	if !readonly {
		registerDeleteProjectsIdTerraformStateName(s)
	}
	if !readonly {
		registerPostProjectsIdTerraformStateName(s)
	}
	registerGetProjectsIdTerraformStateName(s)
	if !readonly {
		registerDeleteProjectsIdTerraformStateNameLock(s)
	}
	if !readonly {
		registerPostProjectsIdTerraformStateNameLock(s)
	}
	// if !readonly { registerDeleteProjectsIdTerraformStateNameVersionsSerial(s) }
	// registerGetProjectsIdTerraformStateNameVersionsSerial(s)
	if !readonly {
		registerPostProjectsIdWikis(s)
	}
	registerGetProjectsIdWikis(s)
	if !readonly {
		registerDeleteProjectsIdWikisSlug(s)
	}
	if !readonly {
		registerPutProjectsIdWikisSlug(s)
	}
	registerGetProjectsIdWikisSlug(s)
	if !readonly {
		registerPostProjectsIdWikisAttachments(s)
	}
	registerGetAdminBatchedBackgroundMigrationsId(s)
	// if !readonly { registerPutAdminBatchedBackgroundMigrationsIdResume(s) }
	// if !readonly { registerPutAdminBatchedBackgroundMigrationsIdPause(s) }
	registerGetAdminBatchedBackgroundMigrations(s)
	if !readonly {
		registerPostAdminCiVariables(s)
	}
	registerGetAdminCiVariables(s)
	if !readonly {
		registerDeleteAdminCiVariablesKey(s)
	}
	if !readonly {
		registerPutAdminCiVariablesKey(s)
	}
	registerGetAdminCiVariablesKey(s)
	// registerGetAdminDatabasesDatabaseNameDictionaryTablesTableName(s)
	registerGetAdminClusters(s)
	if !readonly {
		registerDeleteAdminClustersClusterId(s)
	}
	if !readonly {
		registerPutAdminClustersClusterId(s)
	}
	registerGetAdminClustersClusterId(s)
	if !readonly {
		registerPostAdminClustersAdd(s)
	}
	if !readonly {
		registerPostAdminMigrationsTimestampMark(s)
	}
	if !readonly {
		registerPostBroadcastMessages(s)
	}
	registerGetBroadcastMessages(s)
	if !readonly {
		registerDeleteBroadcastMessagesId(s)
	}
	if !readonly {
		registerPutBroadcastMessagesId(s)
	}
	registerGetBroadcastMessagesId(s)
	if !readonly {
		registerPostApplications(s)
	}
	registerGetApplications(s)
	if !readonly {
		registerDeleteApplicationsId(s)
	}
	if !readonly {
		registerPostApplicationsIdRenewSecret(s)
	}
	registerGetAvatar(s)
	// if !readonly { registerPostBulkImports(s) }
	registerGetBulkImports(s)
	registerGetBulkImportsEntities(s)
	registerGetBulkImportsImportId(s)
	registerGetBulkImportsImportIdEntities(s)
	registerGetBulkImportsImportIdEntitiesEntityId(s)
	// registerGetBulkImportsImportIdEntitiesEntityIdFailures(s)
	if !readonly {
		registerPostBulkImportsImportIdCancel(s)
	}
	registerGetJob(s)
	registerGetJobAllowedAgents(s)
	if !readonly {
		registerDeleteRunners(s)
	}
	if !readonly {
		registerPostRunners(s)
	}
	registerGetRunners(s)
	if !readonly {
		registerDeleteRunnersManagers(s)
	}
	if !readonly {
		registerPostRunnersVerify(s)
	}
	if !readonly {
		registerPostRunnersResetAuthenticationToken(s)
	}
	registerGetRunnersAll(s)
	if !readonly {
		registerDeleteRunnersId(s)
	}
	if !readonly {
		registerPutRunnersId(s)
	}
	registerGetRunnersId(s)
	registerGetRunnersIdManagers(s)
	registerGetRunnersIdJobs(s)
	if !readonly {
		registerPostRunnersIdResetAuthenticationToken(s)
	}
	if !readonly {
		registerPostRunnersResetRegistrationToken(s)
	}
	if !readonly {
		registerPostJobsRequest(s)
	}
	if !readonly {
		registerPutJobsId(s)
	}
	if !readonly {
		registerPostJobsIdArtifactsAuthorize(s)
	}
	if !readonly {
		registerPostJobsIdArtifacts(s)
	}
	registerGetJobsIdArtifacts(s)
	registerGetGroupIdPackagesComposerPackages(s)
	registerGetGroupIdPackagesComposerPSha(s)
	registerGetPackagesConanV1UsersAuthenticate(s)
	registerGetPackagesConanV1UsersCheckCredentials(s)
	registerGetPackagesConanV1ConansSearch(s)
	// registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(s)
	registerGetPackagesConanV1Ping(s)
	// registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference(s)
	// if !readonly { registerDeletePackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(s) }
	// registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(s)
	// registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest(s)
	// registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest(s)
	// registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls(s)
	// registerGetPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls(s)
	// if !readonly { registerPostPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrls(s) }
	// if !readonly { registerPostPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrls(s) }
	// if !readonly { registerPutPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(s) }
	// registerGetPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(s)
	// if !readonly { registerPutPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize(s) }
	// if !readonly { registerPutPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(s) }
	// registerGetPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(s)
	// if !readonly { registerPutPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorize(s) }
	if !readonly {
		registerPostPackagesNpmNpmV1SecurityAdvisoriesBulk(s)
	}
	if !readonly {
		registerPostPackagesNpmNpmV1SecurityAuditsQuick(s)
	}
	// registerGetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersions(s)
	// registerGetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownload(s)
	// registerGetPackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystem(s)
	if !readonly {
		registerPostContainerRegistryEventEvents(s)
	}
	registerGetRegistryRepositoriesId(s)
	registerGetEvents(s)
	registerGetUsersIdEvents(s)
	registerGetUsersUserIdProjects(s)
	registerGetUsersUserIdContributedProjects(s)
	registerGetUsersUserIdStarredProjects(s)
	registerGetFeatures(s)
	registerGetFeaturesDefinitions(s)
	if !readonly {
		registerDeleteFeaturesName(s)
	}
	if !readonly {
		registerPostFeaturesName(s)
	}
	registerGetGeoProxy(s)
	registerGetGeoRetrieveReplicableNameReplicableId(s)
	registerGetGeoRepositoriesGlRepositoryPipelineRefs(s)
	if !readonly {
		registerPostGeoStatus(s)
	}
	if !readonly {
		registerPostGeoProxyGitSshInfoRefsUploadPack(s)
	}
	if !readonly {
		registerPostGeoProxyGitSshUploadPack(s)
	}
	if !readonly {
		registerPostGeoProxyGitSshInfoRefsReceivePack(s)
	}
	if !readonly {
		registerPostGeoProxyGitSshReceivePack(s)
	}
	if !readonly {
		registerPostGeoNodeProxyIdGraphql(s)
	}
	if !readonly {
		registerPostIntegrationsSlackEvents(s)
	}
	if !readonly {
		registerPostIntegrationsSlackInteractions(s)
	}
	if !readonly {
		registerPostIntegrationsSlackOptions(s)
	}
	if !readonly {
		registerPostIntegrationsJiraConnectSubscriptions(s)
	}
	registerGetKeysId(s)
	registerGetKeys(s)
	if !readonly {
		registerPostMarkdown(s)
	}
	registerGetMergeRequests(s)
	if !readonly {
		registerPutNamespacesId(s)
	}
	registerGetNamespacesId(s)
	registerGetNamespacesIdGitlabSubscription(s)
	if !readonly {
		registerDeleteNamespacesIdStorageLimitExclusion(s)
	}
	if !readonly {
		registerPostNamespacesIdStorageLimitExclusion(s)
	}
	registerGetNamespacesStorageLimitExclusions(s)
	registerGetNamespaces(s)
	registerGetNamespacesIdExists(s)
	if !readonly {
		registerPostOrganizations(s)
	}
	registerGetPagesDomains(s)
	if !readonly {
		registerDeletePersonalAccessTokensSelf(s)
	}
	registerGetPersonalAccessTokensSelf(s)
	registerGetPersonalAccessTokensSelfAssociations(s)
	if !readonly {
		registerPostPersonalAccessTokensSelfRotate(s)
	}
	registerGetPersonalAccessTokens(s)
	if !readonly {
		registerDeletePersonalAccessTokensId(s)
	}
	registerGetPersonalAccessTokensId(s)
	if !readonly {
		registerPostPersonalAccessTokensIdRotate(s)
	}
	if !readonly {
		registerPostSnippets(s)
	}
	registerGetSnippets(s)
	registerGetSnippetsPublic(s)
	registerGetSnippetsAll(s)
	if !readonly {
		registerDeleteSnippetsId(s)
	}
	if !readonly {
		registerPutSnippetsId(s)
	}
	registerGetSnippetsId(s)
	registerGetSnippetsIdRaw(s)
	registerGetSnippetsIdFilesRefFilePathRaw(s)
	registerGetSnippetsIdUserAgentDetail(s)
	if !readonly {
		registerPutSuggestionsIdApply(s)
	}
	if !readonly {
		registerPutSuggestionsBatchApply(s)
	}
	if !readonly {
		registerDeleteHooksHookIdUrlVariablesKey(s)
	}
	if !readonly {
		registerPutHooksHookIdUrlVariablesKey(s)
	}
	if !readonly {
		registerDeleteHooksHookIdCustomHeadersKey(s)
	}
	if !readonly {
		registerPutHooksHookIdCustomHeadersKey(s)
	}
	if !readonly {
		registerPostHooks(s)
	}
	registerGetHooks(s)
	if !readonly {
		registerDeleteHooksHookId(s)
	}
	if !readonly {
		registerPostHooksHookId(s)
	}
	if !readonly {
		registerPutHooksHookId(s)
	}
	registerGetHooksHookId(s)
	registerGetFeatureFlagsUnleashProjectId(s)
	registerGetFeatureFlagsUnleashProjectIdFeatures(s)
	// registerGetFeatureFlagsUnleashProjectIdClientFeatures(s)
	// if !readonly { registerPostFeatureFlagsUnleashProjectIdClientRegister(s) }
	// if !readonly { registerPostFeatureFlagsUnleashProjectIdClientMetrics(s) }
	if !readonly {
		registerPostUsageDataIncrementCounter(s)
	}
	if !readonly {
		registerPostUsageDataIncrementUniqueUsers(s)
	}
	if !readonly {
		registerPostUsageDataTrackEvents(s)
	}
	registerGetUsageDataMetricDefinitions(s)
	registerGetUsageDataServicePing(s)
	if !readonly {
		registerPostUsageDataTrackEvent(s)
	}
	registerGetUsageDataNonSqlMetrics(s)
	registerGetUsageDataQueries(s)
	registerGetUserCounts(s)
	if !readonly {
		registerPostUserRunners(s)
	}
	if !readonly {
		registerPutApplicationPlanLimits(s)
	}
	registerGetApplicationPlanLimits(s)
	// if !readonly { registerPutApplicationAppearance(s) }
	registerGetApplicationAppearance(s)
	registerGetApplicationStatistics(s)
	registerGetDiscoverCertBasedClusters(s)
	if !readonly {
		registerPostDeployKeys(s)
	}
	registerGetDeployKeys(s)
	registerGetDeployTokens(s)
	if !readonly {
		registerPostImportBitbucket(s)
	}
	if !readonly {
		registerPostImportBitbucketServer(s)
	}
	if !readonly {
		registerPostImportGithub(s)
	}
	if !readonly {
		registerPostImportGithubCancel(s)
	}
	if !readonly {
		registerPostImportGithubGists(s)
	}
	if !readonly {
		registerPostSlackTrigger(s)
	}
	registerGetMetadata(s)
	registerGetVersion(s)
	if !readonly {
		registerPostTopics(s)
	}
	registerGetTopics(s)
	if !readonly {
		registerDeleteTopicsId(s)
	}
	if !readonly {
		registerPutTopicsId(s)
	}
	registerGetTopicsId(s)
	if !readonly {
		registerPostTopicsMerge(s)
	}
	registerGetWebCommitsPublicKey(s)
	registerGetIssues(s)
	registerGetGroupsIdIssues(s)
	if !readonly {
		registerPostProjectsIdIssues(s)
	}
	registerGetProjectsIdIssues(s)
	registerGetIssuesId(s)
	if !readonly {
		registerDeleteProjectsIdIssuesIssueIid(s)
	}
	if !readonly {
		registerPutProjectsIdIssuesIssueIid(s)
	}
	registerGetProjectsIdIssuesIssueIid(s)
	if !readonly {
		registerPutProjectsIdIssuesIssueIidReorder(s)
	}
	// if !readonly { registerPostProjectsIdIssuesIssueIidMove(s) }
	if !readonly {
		registerPostProjectsIdIssuesIssueIidClone(s)
	}
	if !readonly {
		registerPostProjectsIdIssuesIssueIidSubscribe(s)
	}
	if !readonly {
		registerPostProjectsIdIssuesIssueIidUnsubscribe(s)
	}
	if !readonly {
		registerPostProjectsIdIssuesIssueIidTodo(s)
	}
	if !readonly {
		registerPostProjectsIdIssuesIssueIidNotes(s)
	}
	if !readonly {
		registerPostProjectsIdIssuesIssueIidTimeEstimate(s)
	}
	// if !readonly { registerPostProjectsIdIssuesIssueIidResetTimeEstimate(s) }
	if !readonly {
		registerPostProjectsIdIssuesIssueIidAddSpentTime(s)
	}
	if !readonly {
		registerPostProjectsIdIssuesIssueIidResetSpentTime(s)
	}
	registerGetProjectsIdIssuesIssueIidTimeStats(s)
	registerGetProjectsIdIssuesIssueIidRelatedMergeRequests(s)
	registerGetProjectsIdIssuesIssueIidClosedBy(s)
	registerGetProjectsIdIssuesIssueIidParticipants(s)
	registerGetProjectsIdIssuesIssueIidUserAgentDetail(s)
	// if !readonly { registerPostProjectsIdIssuesIssueIidMetricImages(s) }
	registerGetProjectsIdIssuesIssueIidMetricImages(s)
	// if !readonly { registerDeleteProjectsIdIssuesIssueIidMetricImagesImageId(s) }
	// if !readonly { registerPutProjectsIdIssuesIssueIidMetricImagesImageId(s) }
}
