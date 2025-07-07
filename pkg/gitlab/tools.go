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
		registerDeleteGroupsIdAccessRequestsUserId(s)
	}
	registerGetGroupsIdEpicsEpicIidAwardEmoji(s)
	// if !readonly { registerDeleteGroupsIdEpicsEpicIidAwardEmojiAwardId(s) }
	// registerGetGroupsIdEpicsEpicIidAwardEmojiAwardId(s)
	// registerGetGroupsIdEpicsEpicIidNotesNoteIdAwardEmoji(s)
	// if !readonly { registerDeleteGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardId(s) }
	// registerGetGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardId(s)
	registerGetGroupsIdBadges(s)
	registerGetGroupsIdBadgesRender(s)
	if !readonly {
		registerDeleteGroupsIdBadgesBadgeId(s)
	}
	registerGetGroupsIdBadgesBadgeId(s)
	registerGetGroupsIdCustomAttributes(s)
	if !readonly {
		registerDeleteGroupsIdCustomAttributesKey(s)
	}
	registerGetGroupsIdCustomAttributesKey(s)
	registerGetGroups(s)
	if !readonly {
		registerDeleteGroupsId(s)
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
		registerDeleteGroupsIdShareGroupId(s)
	}
	if !readonly {
		registerPostGroupsIdLdapSync(s)
	}
	registerGetGroupsIdAuditEvents(s)
	registerGetGroupsIdAuditEventsAuditEventId(s)
	registerGetGroupsIdSamlUsers(s)
	registerGetGroupsIdProvisionedUsers(s)
	registerGetGroupsIdUsers(s)
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
	registerGetGroupsIdClustersClusterId(s)
	registerGetGroupsIdRegistryRepositories(s)
	registerGetGroupsIdDebianDistributions(s)
	if !readonly {
		registerDeleteGroupsIdDebianDistributionsCodename(s)
	}
	registerGetGroupsIdDebianDistributionsCodename(s)
	// registerGetGroupsIdDebianDistributionsCodenameKeyAsc(s)
	registerGetGroupsIdExportDownload(s)
	if !readonly {
		registerPostGroupsIdExport(s)
	}
	registerGetGroupsIdExportRelationsDownload(s)
	registerGetGroupsIdExportRelationsStatus(s)
	if !readonly {
		registerPostGroupsImportAuthorize(s)
	}
	registerGetGroupsIdPackages(s)
	registerGetGroupsIdPlaceholderReassignments(s)
	// if !readonly { registerPostGroupsIdPlaceholderReassignmentsAuthorize(s) }
	registerGetGroupsIdVariables(s)
	if !readonly {
		registerDeleteGroupsIdVariablesKey(s)
	}
	registerGetGroupsIdVariablesKey(s)
	registerGetGroupsIdIntegrations(s)
	if !readonly {
		registerDeleteGroupsIdIntegrationsSlug(s)
	}
	registerGetGroupsIdIntegrationsSlug(s)
	registerGetGroupsIdInvitations(s)
	if !readonly {
		registerDeleteGroupsIdInvitationsEmail(s)
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
	registerGetGroupsIdMembers(s)
	registerGetGroupsIdMembersAll(s)
	if !readonly {
		registerDeleteGroupsIdMembersUserId(s)
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
	registerGetGroupsIdWikis(s)
	if !readonly {
		registerDeleteGroupsIdWikisSlug(s)
	}
	registerGetGroupsIdWikisSlug(s)
	if !readonly {
		registerPostProjectsIdAccessRequests(s)
	}
	registerGetProjectsIdAccessRequests(s)
	if !readonly {
		registerDeleteProjectsIdAccessRequestsUserId(s)
	}
	// if !readonly { registerPostProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorize(s) }
	// registerGetProjectsIdAlertManagementAlertsAlertIidMetricImages(s)
	// if !readonly { registerDeleteProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageId(s) }
	registerGetProjectsIdIssuesIssueIidAwardEmoji(s)
	// if !readonly { registerDeleteProjectsIdIssuesIssueIidAwardEmojiAwardId(s) }
	// registerGetProjectsIdIssuesIssueIidAwardEmojiAwardId(s)
	// registerGetProjectsIdIssuesIssueIidNotesNoteIdAwardEmoji(s)
	// if !readonly { registerDeleteProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardId(s) }
	// registerGetProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardId(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidAwardEmoji(s)
	// if !readonly { registerDeleteProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardId(s) }
	// registerGetProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardId(s)
	// registerGetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmoji(s)
	// if !readonly { registerDeleteProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardId(s) }
	// registerGetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardId(s)
	registerGetProjectsIdSnippetsSnippetIdAwardEmoji(s)
	// if !readonly { registerDeleteProjectsIdSnippetsSnippetIdAwardEmojiAwardId(s) }
	// registerGetProjectsIdSnippetsSnippetIdAwardEmojiAwardId(s)
	// registerGetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmoji(s)
	// if !readonly { registerDeleteProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardId(s) }
	// registerGetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardId(s)
	registerGetProjectsIdBadges(s)
	registerGetProjectsIdBadgesRender(s)
	if !readonly {
		registerDeleteProjectsIdBadgesBadgeId(s)
	}
	registerGetProjectsIdBadgesBadgeId(s)
	registerGetProjectsIdRepositoryBranches(s)
	if !readonly {
		registerDeleteProjectsIdRepositoryBranchesBranch(s)
	}
	registerGetProjectsIdRepositoryBranchesBranch(s)
	if !readonly {
		registerPutProjectsIdRepositoryBranchesBranchUnprotect(s)
	}
	if !readonly {
		registerDeleteProjectsIdRepositoryMergedBranches(s)
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
		registerPostProjectsIdJobsJobIdRetry(s)
	}
	if !readonly {
		registerPostProjectsIdJobsJobIdErase(s)
	}
	registerGetProjectsIdResourceGroups(s)
	registerGetProjectsIdResourceGroupsKey(s)
	registerGetProjectsIdResourceGroupsKeyUpcomingJobs(s)
	registerGetProjectsIdRunners(s)
	if !readonly {
		registerDeleteProjectsIdRunnersRunnerId(s)
	}
	if !readonly {
		registerPostProjectsIdRunnersResetRegistrationToken(s)
	}
	registerGetProjectsIdSecureFiles(s)
	if !readonly {
		registerDeleteProjectsIdSecureFilesSecureFileId(s)
	}
	registerGetProjectsIdSecureFilesSecureFileId(s)
	// registerGetProjectsIdSecureFilesSecureFileIdDownload(s)
	registerGetProjectsIdPipelines(s)
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
		registerPostProjectsIdPipelinesPipelineIdRetry(s)
	}
	if !readonly {
		registerPostProjectsIdPipelinesPipelineIdCancel(s)
	}
	registerGetProjectsIdPipelineSchedules(s)
	// if !readonly { registerDeleteProjectsIdPipelineSchedulesPipelineScheduleId(s) }
	// registerGetProjectsIdPipelineSchedulesPipelineScheduleId(s)
	// registerGetProjectsIdPipelineSchedulesPipelineScheduleIdPipelines(s)
	// if !readonly { registerPostProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnership(s) }
	// if !readonly { registerPostProjectsIdPipelineSchedulesPipelineScheduleIdPlay(s) }
	// if !readonly { registerDeleteProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey(s) }
	registerGetProjectsIdTriggers(s)
	if !readonly {
		registerDeleteProjectsIdTriggersTriggerId(s)
	}
	registerGetProjectsIdTriggersTriggerId(s)
	registerGetProjectsIdVariables(s)
	if !readonly {
		registerDeleteProjectsIdVariablesKey(s)
	}
	registerGetProjectsIdVariablesKey(s)
	registerGetProjectsIdClusterAgentsAgentIdTokens(s)
	// if !readonly { registerDeleteProjectsIdClusterAgentsAgentIdTokensTokenId(s) }
	// registerGetProjectsIdClusterAgentsAgentIdTokensTokenId(s)
	registerGetProjectsIdClusterAgents(s)
	if !readonly {
		registerDeleteProjectsIdClusterAgentsAgentId(s)
	}
	registerGetProjectsIdClusterAgentsAgentId(s)
	registerGetProjectsIdPackagesCargoConfigJson(s)
	registerGetProjectsIdRepositoryCommits(s)
	registerGetProjectsIdRepositoryCommitsSha(s)
	registerGetProjectsIdRepositoryCommitsShaDiff(s)
	registerGetProjectsIdRepositoryCommitsShaComments(s)
	registerGetProjectsIdRepositoryCommitsShaSequence(s)
	registerGetProjectsIdRepositoryCommitsShaRefs(s)
	registerGetProjectsIdRepositoryCommitsShaMergeRequests(s)
	registerGetProjectsIdRepositoryCommitsShaSignature(s)
	registerGetProjectsIdRepositoryCommitsShaStatuses(s)
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
	// registerGetProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(s)
	// if !readonly { registerPutProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize(s) }
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
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileName(s)
	// if !readonly { registerPutProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameAuthorize(s) }
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearch(s)
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatest(s)
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisions(s)
	// if !readonly { registerDeleteProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevision(s) }
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFiles(s)
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileName(s)
	// if !readonly { registerPutProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameAuthorize(s) }
	// registerGetProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileName(s)
	registerGetProjectsIdDeployKeys(s)
	if !readonly {
		registerDeleteProjectsIdDeployKeysKeyId(s)
	}
	registerGetProjectsIdDeployKeysKeyId(s)
	if !readonly {
		registerPostProjectsIdDeployKeysKeyIdEnable(s)
	}
	registerGetProjectsIdDeployTokens(s)
	if !readonly {
		registerDeleteProjectsIdDeployTokensTokenId(s)
	}
	registerGetProjectsIdDeployTokensTokenId(s)
	registerGetProjectsIdDeployments(s)
	if !readonly {
		registerDeleteProjectsIdDeploymentsDeploymentId(s)
	}
	registerGetProjectsIdDeploymentsDeploymentId(s)
	registerGetProjectsIdDeploymentsDeploymentIdMergeRequests(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidDraftNotes(s)
	// if !readonly { registerDeleteProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(s) }
	// registerGetProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(s)
	// if !readonly { registerPutProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublish(s) }
	// if !readonly { registerPostProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublish(s) }
	registerGetProjectsIdEnvironments(s)
	if !readonly {
		registerDeleteProjectsIdEnvironmentsEnvironmentId(s)
	}
	registerGetProjectsIdEnvironmentsEnvironmentId(s)
	if !readonly {
		registerDeleteProjectsIdEnvironmentsReviewApps(s)
	}
	if !readonly {
		registerPostProjectsIdErrorTrackingClientKeys(s)
	}
	registerGetProjectsIdErrorTrackingClientKeys(s)
	// if !readonly { registerDeleteProjectsIdErrorTrackingClientKeysKeyId(s) }
	registerGetProjectsIdErrorTrackingSettings(s)
	registerGetProjectsIdFeatureFlags(s)
	if !readonly {
		registerDeleteProjectsIdFeatureFlagsFeatureFlagName(s)
	}
	registerGetProjectsIdFeatureFlagsFeatureFlagName(s)
	registerGetProjectsIdFeatureFlagsUserLists(s)
	if !readonly {
		registerDeleteProjectsIdFeatureFlagsUserListsIid(s)
	}
	registerGetProjectsIdFeatureFlagsUserListsIid(s)
	registerGetProjectsIdRepositoryFilesFilePathBlame(s)
	registerGetProjectsIdRepositoryFilesFilePathRaw(s)
	if !readonly {
		registerDeleteProjectsIdRepositoryFilesFilePath(s)
	}
	registerGetProjectsIdRepositoryFilesFilePath(s)
	registerGetProjectsIdFreezePeriods(s)
	if !readonly {
		registerDeleteProjectsIdFreezePeriodsFreezePeriodId(s)
	}
	registerGetProjectsIdFreezePeriodsFreezePeriodId(s)
	registerGetProjectsIdPackagesHelmChannelIndexYaml(s)
	// registerGetProjectsIdPackagesHelmChannelChartsFileNameTgz(s)
	// if !readonly { registerPostProjectsIdPackagesHelmApiChannelChartsAuthorize(s) }
	registerGetProjectsIdServices(s)
	if !readonly {
		registerDeleteProjectsIdServicesSlug(s)
	}
	registerGetProjectsIdServicesSlug(s)
	registerGetProjectsIdIntegrations(s)
	if !readonly {
		registerDeleteProjectsIdIntegrationsSlug(s)
	}
	registerGetProjectsIdIntegrationsSlug(s)
	registerGetProjectsIdInvitations(s)
	if !readonly {
		registerDeleteProjectsIdInvitationsEmail(s)
	}
	registerGetProjectsIdIssuesIssueIidLinks(s)
	// if !readonly { registerDeleteProjectsIdIssuesIssueIidLinksIssueLinkId(s) }
	// registerGetProjectsIdIssuesIssueIidLinksIssueLinkId(s)
	registerGetProjectsIdCiLint(s)
	if !readonly {
		registerPostProjectsIdUploadsAuthorize(s)
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
	registerGetProjectsIdMembers(s)
	registerGetProjectsIdMembersAll(s)
	if !readonly {
		registerDeleteProjectsIdMembersUserId(s)
	}
	registerGetProjectsIdMembersUserId(s)
	registerGetProjectsIdMembersAllUserId(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidApprovals(s)
	if !readonly {
		registerPostProjectsIdMergeRequestsMergeRequestIidUnapprove(s)
	}
	// if !readonly { registerPutProjectsIdMergeRequestsMergeRequestIidResetApprovals(s) }
	// registerGetProjectsIdMergeRequestsMergeRequestIidApprovalState(s)
	if !readonly {
		registerPostProjectsIdCreateCiConfig(s)
	}
	// if !readonly { registerPostProjectsIdMergeRequestsMergeRequestIidResetTimeEstimate(s) }
	// if !readonly { registerPostProjectsIdMergeRequestsMergeRequestIidResetSpentTime(s) }
	registerGetProjectsIdMergeRequestsMergeRequestIidTimeStats(s)
	registerGetProjectsIdMergeRequests(s)
	if !readonly {
		registerDeleteProjectsIdMergeRequestsMergeRequestIid(s)
	}
	registerGetProjectsIdMergeRequestsMergeRequestIid(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidParticipants(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidReviewers(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidCommits(s)
	// if !readonly { registerDeleteProjectsIdMergeRequestsMergeRequestIidContextCommits(s) }
	// registerGetProjectsIdMergeRequestsMergeRequestIidContextCommits(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidChanges(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidDiffs(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidRawDiffs(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidPipelines(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidMergeRef(s)
	// if !readonly { registerPostProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceeds(s) }
	registerGetProjectsIdMergeRequestsMergeRequestIidClosesIssues(s)
	// registerGetProjectsIdMergeRequestsMergeRequestIidRelatedIssues(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidVersions(s)
	// registerGetProjectsIdMergeRequestsMergeRequestIidVersionsVersionId(s)
	// if !readonly { registerPostProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulk(s) }
	// if !readonly { registerPostProjectsIdPackagesNpmNpmV1SecurityAuditsQuick(s) }
	registerGetProjectsIdPackagesNugetIndex(s)
	registerGetProjectsIdPackagesNugetV2(s)
	registerGetProjectsIdPackagesNugetV2Metadata(s)
	registerGetProjectsIdPackagesNugetQuery(s)
	if !readonly {
		registerPutProjectsIdPackagesNugetAuthorize(s)
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
	registerGetProjectsIdPagesDomains(s)
	if !readonly {
		registerDeleteProjectsIdPagesDomainsDomain(s)
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
	registerGetProjectsIdClustersClusterId(s)
	registerGetProjectsIdRegistryRepositories(s)
	if !readonly {
		registerDeleteProjectsIdRegistryRepositoriesRepositoryId(s)
	}
	// if !readonly { registerDeleteProjectsIdRegistryRepositoriesRepositoryIdTags(s) }
	registerGetProjectsIdRegistryRepositoriesRepositoryIdTags(s)
	// if !readonly { registerDeleteProjectsIdRegistryRepositoriesRepositoryIdTagsTagName(s) }
	// registerGetProjectsIdRegistryRepositoriesRepositoryIdTagsTagName(s)
	registerGetProjectsIdRegistryProtectionRepositoryRules(s)
	// if !readonly { registerDeleteProjectsIdRegistryProtectionRepositoryRulesProtectionRuleId(s) }
	registerGetProjectsIdDebianDistributions(s)
	if !readonly {
		registerDeleteProjectsIdDebianDistributionsCodename(s)
	}
	registerGetProjectsIdDebianDistributionsCodename(s)
	// registerGetProjectsIdDebianDistributionsCodenameKeyAsc(s)
	registerGetProjectsIdEvents(s)
	registerGetProjectsIdExport(s)
	registerGetProjectsIdExportDownload(s)
	registerGetProjectsIdExportRelationsDownload(s)
	registerGetProjectsIdExportRelationsStatus(s)
	if !readonly {
		registerDeleteProjectsIdHooksHookIdUrlVariablesKey(s)
	}
	if !readonly {
		registerDeleteProjectsIdHooksHookIdCustomHeadersKey(s)
	}
	registerGetProjectsIdHooks(s)
	if !readonly {
		registerDeleteProjectsIdHooksHookId(s)
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
	registerGetProjectsIdImport(s)
	if !readonly {
		registerPostProjectsImportRelationAuthorize(s)
	}
	registerGetProjectsIdRelationImports(s)
	registerGetProjectsIdJobTokenScope(s)
	registerGetProjectsIdJobTokenScopeAllowlist(s)
	registerGetProjectsIdJobTokenScopeGroupsAllowlist(s)
	// if !readonly { registerDeleteProjectsIdJobTokenScopeGroupsAllowlistTargetGroupId(s) }
	// if !readonly { registerDeleteProjectsIdJobTokenScopeAllowlistTargetProjectId(s) }
	registerGetProjectsIdPackages(s)
	if !readonly {
		registerDeleteProjectsIdPackagesPackageId(s)
	}
	registerGetProjectsIdPackagesPackageId(s)
	registerGetProjectsIdPackagesPackageIdPipelines(s)
	registerGetProjectsIdPackagesProtectionRules(s)
	// if !readonly { registerDeleteProjectsIdPackagesProtectionRulesPackageProtectionRuleId(s) }
	registerGetProjectsIdSnapshot(s)
	registerGetProjectsIdSnippets(s)
	if !readonly {
		registerDeleteProjectsIdSnippetsSnippetId(s)
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
	registerGetProjectsIdCustomAttributesKey(s)
	if !readonly {
		registerPostProjectsIdRestore(s)
	}
	registerGetProjects(s)
	registerGetProjectsIdShareLocations(s)
	if !readonly {
		registerDeleteProjectsId(s)
	}
	registerGetProjectsId(s)
	if !readonly {
		registerDeleteProjectsIdFork(s)
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
		registerDeleteProjectsIdShareGroupId(s)
	}
	if !readonly {
		registerPostProjectsIdImportProjectMembersProjectId(s)
	}
	registerGetProjectsIdUsers(s)
	registerGetProjectsIdGroups(s)
	registerGetProjectsIdInvitedGroups(s)
	if !readonly {
		registerPostProjectsIdRepositorySize(s)
	}
	registerGetProjectsIdTransferLocations(s)
	registerGetProjectsIdStorage(s)
	registerGetProjectsIdAuditEvents(s)
	registerGetProjectsIdAuditEventsAuditEventId(s)
	registerGetProjectsIdProtectedBranches(s)
	if !readonly {
		registerDeleteProjectsIdProtectedBranchesName(s)
	}
	registerGetProjectsIdProtectedBranchesName(s)
	registerGetProjectsIdProtectedTags(s)
	if !readonly {
		registerDeleteProjectsIdProtectedTagsName(s)
	}
	registerGetProjectsIdProtectedTagsName(s)
	registerGetProjectsIdPackagesPypiSimple(s)
	if !readonly {
		registerPostProjectsIdPackagesPypiAuthorize(s)
	}
	registerGetProjectsIdReleases(s)
	if !readonly {
		registerDeleteProjectsIdReleasesTagName(s)
	}
	registerGetProjectsIdReleasesTagName(s)
	if !readonly {
		registerPostProjectsIdReleasesTagNameEvidence(s)
	}
	registerGetProjectsIdReleasesTagNameAssetsLinks(s)
	// if !readonly { registerDeleteProjectsIdReleasesTagNameAssetsLinksLinkId(s) }
	// registerGetProjectsIdReleasesTagNameAssetsLinksLinkId(s)
	registerGetProjectsIdRemoteMirrors(s)
	if !readonly {
		registerDeleteProjectsIdRemoteMirrorsMirrorId(s)
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
	registerGetProjectsIdRepositoryChangelog(s)
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
	registerGetProjectsIdPackagesRubygemsApiV1Dependencies(s)
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
	// if !readonly { registerDeleteProjectsIdTerraformStateNameVersionsSerial(s) }
	// registerGetProjectsIdTerraformStateNameVersionsSerial(s)
	registerGetProjectsIdWikis(s)
	if !readonly {
		registerDeleteProjectsIdWikisSlug(s)
	}
	registerGetProjectsIdWikisSlug(s)
	registerGetAdminBatchedBackgroundMigrationsId(s)
	registerGetAdminBatchedBackgroundMigrations(s)
	registerGetAdminCiVariables(s)
	if !readonly {
		registerDeleteAdminCiVariablesKey(s)
	}
	registerGetAdminCiVariablesKey(s)
	// registerGetAdminDatabasesDatabaseNameDictionaryTablesTableName(s)
	registerGetAdminClusters(s)
	if !readonly {
		registerDeleteAdminClustersClusterId(s)
	}
	registerGetAdminClustersClusterId(s)
	registerGetBroadcastMessages(s)
	if !readonly {
		registerDeleteBroadcastMessagesId(s)
	}
	registerGetBroadcastMessagesId(s)
	registerGetApplications(s)
	if !readonly {
		registerDeleteApplicationsId(s)
	}
	if !readonly {
		registerPostApplicationsIdRenewSecret(s)
	}
	registerGetAvatar(s)
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
	registerGetRunners(s)
	if !readonly {
		registerDeleteRunnersManagers(s)
	}
	registerGetRunnersAll(s)
	if !readonly {
		registerDeleteRunnersId(s)
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
	// registerGetPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(s)
	// if !readonly { registerPutPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize(s) }
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
	registerGetGeoProxy(s)
	registerGetGeoRetrieveReplicableNameReplicableId(s)
	registerGetGeoRepositoriesGlRepositoryPipelineRefs(s)
	if !readonly {
		registerPostGeoNodeProxyIdGraphql(s)
	}
	if !readonly {
		registerPostIntegrationsSlackInteractions(s)
	}
	if !readonly {
		registerPostIntegrationsSlackOptions(s)
	}
	registerGetKeysId(s)
	registerGetKeys(s)
	registerGetMergeRequests(s)
	registerGetNamespacesId(s)
	registerGetNamespacesIdGitlabSubscription(s)
	if !readonly {
		registerDeleteNamespacesIdStorageLimitExclusion(s)
	}
	registerGetNamespacesStorageLimitExclusions(s)
	registerGetNamespaces(s)
	registerGetNamespacesIdExists(s)
	registerGetPagesDomains(s)
	if !readonly {
		registerDeletePersonalAccessTokensSelf(s)
	}
	registerGetPersonalAccessTokensSelf(s)
	registerGetPersonalAccessTokensSelfAssociations(s)
	registerGetPersonalAccessTokens(s)
	if !readonly {
		registerDeletePersonalAccessTokensId(s)
	}
	registerGetPersonalAccessTokensId(s)
	registerGetSnippets(s)
	registerGetSnippetsPublic(s)
	registerGetSnippetsAll(s)
	if !readonly {
		registerDeleteSnippetsId(s)
	}
	registerGetSnippetsId(s)
	registerGetSnippetsIdRaw(s)
	registerGetSnippetsIdFilesRefFilePathRaw(s)
	registerGetSnippetsIdUserAgentDetail(s)
	if !readonly {
		registerDeleteHooksHookIdUrlVariablesKey(s)
	}
	if !readonly {
		registerDeleteHooksHookIdCustomHeadersKey(s)
	}
	registerGetHooks(s)
	if !readonly {
		registerDeleteHooksHookId(s)
	}
	if !readonly {
		registerPostHooksHookId(s)
	}
	registerGetHooksHookId(s)
	registerGetFeatureFlagsUnleashProjectId(s)
	registerGetFeatureFlagsUnleashProjectIdFeatures(s)
	// registerGetFeatureFlagsUnleashProjectIdClientFeatures(s)
	registerGetUsageDataMetricDefinitions(s)
	registerGetUsageDataServicePing(s)
	registerGetUsageDataNonSqlMetrics(s)
	registerGetUsageDataQueries(s)
	registerGetUserCounts(s)
	registerGetApplicationPlanLimits(s)
	registerGetApplicationAppearance(s)
	registerGetApplicationStatistics(s)
	registerGetDiscoverCertBasedClusters(s)
	registerGetDeployKeys(s)
	registerGetDeployTokens(s)
	registerGetMetadata(s)
	registerGetVersion(s)
	registerGetTopics(s)
	if !readonly {
		registerDeleteTopicsId(s)
	}
	registerGetTopicsId(s)
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
	registerGetProjectsIdIssuesIssueIidMetricImages(s)
	// if !readonly { registerDeleteProjectsIdIssuesIssueIidMetricImagesImageId(s) }
}
