package gitlab

import (
	"github.com/mark3labs/mcp-go/server"
)

func RegisterTools(s *server.MCPServer, readonly bool) {
	if !readonly {
		registerPostGroupsIdAccessRequests(s)
	}
	registerGetGroupsIdAccessRequests(s)
	registerGetGroupsIdEpicsEpicIidAwardEmoji(s)
	// registerGetGroupsIdEpicsEpicIidAwardEmojiAwardId(s)
	// registerGetGroupsIdEpicsEpicIidNotesNoteIdAwardEmoji(s)
	// registerGetGroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardId(s)
	registerGetGroupsIdBadges(s)
	registerGetGroupsIdBadgesRender(s)
	registerGetGroupsIdBadgesBadgeId(s)
	registerGetGroupsIdCustomAttributes(s)
	registerGetGroupsIdCustomAttributesKey(s)
	registerGetGroups(s)
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
		registerPostGroupsIdLdapSync(s)
	}
	registerGetGroupsIdAuditEvents(s)
	registerGetGroupsIdAuditEventsAuditEventId(s)
	registerGetGroupsIdSamlUsers(s)
	registerGetGroupsIdProvisionedUsers(s)
	registerGetGroupsIdUsers(s)
	registerGetGroupsIdSshCertificates(s)
	registerGetGroupsIdRunners(s)
	if !readonly {
		registerPostGroupsIdRunnersResetRegistrationToken(s)
	}
	// registerGetGroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileName(s)
	registerGetGroupsIdDeployTokens(s)
	registerGetGroupsIdDeployTokensTokenId(s)
	registerGetGroupsIdAvatar(s)
	registerGetGroupsIdClusters(s)
	registerGetGroupsIdClustersClusterId(s)
	registerGetGroupsIdRegistryRepositories(s)
	registerGetGroupsIdDebianDistributions(s)
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
	registerGetGroupsIdVariablesKey(s)
	registerGetGroupsIdIntegrations(s)
	registerGetGroupsIdIntegrationsSlug(s)
	registerGetGroupsIdInvitations(s)
	registerGetGroupsIdUploads(s)
	registerGetGroupsIdUploadsUploadId(s)
	registerGetGroupsIdUploadsSecretFilename(s)
	registerGetGroupsIdMembers(s)
	registerGetGroupsIdMembersAll(s)
	registerGetGroupsIdMembersUserId(s)
	registerGetGroupsIdMembersAllUserId(s)
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
	registerGetGroupsIdWikisSlug(s)
	if !readonly {
		registerPostProjectsIdAccessRequests(s)
	}
	registerGetProjectsIdAccessRequests(s)
	// if !readonly { registerPostProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorize(s) }
	// registerGetProjectsIdAlertManagementAlertsAlertIidMetricImages(s)
	registerGetProjectsIdIssuesIssueIidAwardEmoji(s)
	// registerGetProjectsIdIssuesIssueIidAwardEmojiAwardId(s)
	// registerGetProjectsIdIssuesIssueIidNotesNoteIdAwardEmoji(s)
	// registerGetProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardId(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidAwardEmoji(s)
	// registerGetProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardId(s)
	// registerGetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmoji(s)
	// registerGetProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardId(s)
	registerGetProjectsIdSnippetsSnippetIdAwardEmoji(s)
	// registerGetProjectsIdSnippetsSnippetIdAwardEmojiAwardId(s)
	// registerGetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmoji(s)
	// registerGetProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardId(s)
	registerGetProjectsIdBadges(s)
	registerGetProjectsIdBadgesRender(s)
	registerGetProjectsIdBadgesBadgeId(s)
	registerGetProjectsIdRepositoryBranches(s)
	registerGetProjectsIdRepositoryBranchesBranch(s)
	if !readonly {
		registerPutProjectsIdRepositoryBranchesBranchUnprotect(s)
	}
	registerGetProjectsIdJobsArtifactsRefNameDownload(s)
	registerGetProjectsIdJobsJobIdArtifacts(s)
	if !readonly {
		registerPostProjectsIdJobsJobIdArtifactsKeep(s)
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
		registerPostProjectsIdRunnersResetRegistrationToken(s)
	}
	registerGetProjectsIdSecureFiles(s)
	registerGetProjectsIdSecureFilesSecureFileId(s)
	// registerGetProjectsIdSecureFilesSecureFileIdDownload(s)
	registerGetProjectsIdPipelines(s)
	registerGetProjectsIdPipelinesLatest(s)
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
	// registerGetProjectsIdPipelineSchedulesPipelineScheduleId(s)
	// registerGetProjectsIdPipelineSchedulesPipelineScheduleIdPipelines(s)
	// if !readonly { registerPostProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnership(s) }
	// if !readonly { registerPostProjectsIdPipelineSchedulesPipelineScheduleIdPlay(s) }
	registerGetProjectsIdTriggers(s)
	registerGetProjectsIdTriggersTriggerId(s)
	registerGetProjectsIdVariables(s)
	registerGetProjectsIdVariablesKey(s)
	registerGetProjectsIdClusterAgentsAgentIdTokens(s)
	// registerGetProjectsIdClusterAgentsAgentIdTokensTokenId(s)
	registerGetProjectsIdClusterAgents(s)
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
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFiles(s)
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileName(s)
	// if !readonly { registerPutProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameAuthorize(s) }
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearch(s)
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatest(s)
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisions(s)
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFiles(s)
	// registerGetProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileName(s)
	// if !readonly { registerPutProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameAuthorize(s) }
	// registerGetProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileName(s)
	registerGetProjectsIdDeployKeys(s)
	registerGetProjectsIdDeployKeysKeyId(s)
	if !readonly {
		registerPostProjectsIdDeployKeysKeyIdEnable(s)
	}
	registerGetProjectsIdDeployTokens(s)
	registerGetProjectsIdDeployTokensTokenId(s)
	registerGetProjectsIdDeployments(s)
	registerGetProjectsIdDeploymentsDeploymentId(s)
	registerGetProjectsIdDeploymentsDeploymentIdMergeRequests(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidDraftNotes(s)
	// registerGetProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(s)
	// if !readonly { registerPutProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublish(s) }
	// if !readonly { registerPostProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublish(s) }
	registerGetProjectsIdEnvironments(s)
	registerGetProjectsIdEnvironmentsEnvironmentId(s)
	if !readonly {
		registerPostProjectsIdErrorTrackingClientKeys(s)
	}
	registerGetProjectsIdErrorTrackingClientKeys(s)
	registerGetProjectsIdErrorTrackingSettings(s)
	registerGetProjectsIdFeatureFlags(s)
	registerGetProjectsIdFeatureFlagsFeatureFlagName(s)
	registerGetProjectsIdFeatureFlagsUserLists(s)
	registerGetProjectsIdFeatureFlagsUserListsIid(s)
	registerGetProjectsIdRepositoryFilesFilePathBlame(s)
	registerGetProjectsIdRepositoryFilesFilePathRaw(s)
	registerGetProjectsIdRepositoryFilesFilePath(s)
	registerGetProjectsIdFreezePeriods(s)
	registerGetProjectsIdFreezePeriodsFreezePeriodId(s)
	registerGetProjectsIdPackagesHelmChannelIndexYaml(s)
	// registerGetProjectsIdPackagesHelmChannelChartsFileNameTgz(s)
	// if !readonly { registerPostProjectsIdPackagesHelmApiChannelChartsAuthorize(s) }
	registerGetProjectsIdServices(s)
	registerGetProjectsIdServicesSlug(s)
	registerGetProjectsIdIntegrations(s)
	registerGetProjectsIdIntegrationsSlug(s)
	registerGetProjectsIdInvitations(s)
	registerGetProjectsIdIssuesIssueIidLinks(s)
	// registerGetProjectsIdIssuesIssueIidLinksIssueLinkId(s)
	registerGetProjectsIdCiLint(s)
	if !readonly {
		registerPostProjectsIdUploadsAuthorize(s)
	}
	registerGetProjectsIdUploads(s)
	registerGetProjectsIdUploadsUploadId(s)
	registerGetProjectsIdUploadsSecretFilename(s)
	registerGetProjectsIdMembers(s)
	registerGetProjectsIdMembersAll(s)
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
	registerGetProjectsIdMergeRequestsMergeRequestIid(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidParticipants(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidReviewers(s)
	registerGetProjectsIdMergeRequestsMergeRequestIidCommits(s)
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
	registerGetProjectsIdPages(s)
	registerGetProjectsIdPagesDomains(s)
	registerGetProjectsIdPagesDomainsDomain(s)
	if !readonly {
		registerPutProjectsIdPagesDomainsDomainVerify(s)
	}
	registerGetProjectsIdAvatar(s)
	registerGetProjectsIdClusters(s)
	registerGetProjectsIdClustersClusterId(s)
	registerGetProjectsIdRegistryRepositories(s)
	registerGetProjectsIdRegistryRepositoriesRepositoryIdTags(s)
	// registerGetProjectsIdRegistryRepositoriesRepositoryIdTagsTagName(s)
	registerGetProjectsIdRegistryProtectionRepositoryRules(s)
	registerGetProjectsIdDebianDistributions(s)
	registerGetProjectsIdDebianDistributionsCodename(s)
	// registerGetProjectsIdDebianDistributionsCodenameKeyAsc(s)
	registerGetProjectsIdEvents(s)
	registerGetProjectsIdExport(s)
	registerGetProjectsIdExportDownload(s)
	registerGetProjectsIdExportRelationsDownload(s)
	registerGetProjectsIdExportRelationsStatus(s)
	registerGetProjectsIdHooks(s)
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
	registerGetProjectsIdPackages(s)
	registerGetProjectsIdPackagesPackageId(s)
	registerGetProjectsIdPackagesPackageIdPipelines(s)
	registerGetProjectsIdPackagesProtectionRules(s)
	registerGetProjectsIdSnapshot(s)
	registerGetProjectsIdSnippets(s)
	registerGetProjectsIdSnippetsSnippetId(s)
	registerGetProjectsIdSnippetsSnippetIdRaw(s)
	// registerGetProjectsIdSnippetsSnippetIdFilesRefFilePathRaw(s)
	// registerGetProjectsIdSnippetsSnippetIdUserAgentDetail(s)
	registerGetProjectsIdStatistics(s)
	registerGetProjectsIdTemplatesType(s)
	registerGetProjectsIdTemplatesTypeName(s)
	registerGetProjectsIdCustomAttributes(s)
	registerGetProjectsIdCustomAttributesKey(s)
	if !readonly {
		registerPostProjectsIdRestore(s)
	}
	registerGetProjects(s)
	registerGetProjectsIdShareLocations(s)
	registerGetProjectsId(s)
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
	registerGetProjectsIdProtectedBranchesName(s)
	registerGetProjectsIdProtectedTags(s)
	registerGetProjectsIdProtectedTagsName(s)
	registerGetProjectsIdPackagesPypiSimple(s)
	if !readonly {
		registerPostProjectsIdPackagesPypiAuthorize(s)
	}
	registerGetProjectsIdReleases(s)
	registerGetProjectsIdReleasesTagName(s)
	if !readonly {
		registerPostProjectsIdReleasesTagNameEvidence(s)
	}
	registerGetProjectsIdReleasesTagNameAssetsLinks(s)
	// registerGetProjectsIdReleasesTagNameAssetsLinksLinkId(s)
	registerGetProjectsIdRemoteMirrors(s)
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
	registerGetProjectsIdRepositoryTagsTagName(s)
	registerGetProjectsIdRepositoryTagsTagNameSignature(s)
	// registerGetProjectsIdPackagesTerraformModulesModuleNameModuleSystem(s)
	if !readonly {
		registerPostProjectsIdTerraformStateName(s)
	}
	registerGetProjectsIdTerraformStateName(s)
	// registerGetProjectsIdTerraformStateNameVersionsSerial(s)
	registerGetProjectsIdWikis(s)
	registerGetProjectsIdWikisSlug(s)
	registerGetAdminBatchedBackgroundMigrationsId(s)
	registerGetAdminBatchedBackgroundMigrations(s)
	registerGetAdminCiVariables(s)
	registerGetAdminCiVariablesKey(s)
	// registerGetAdminDatabasesDatabaseNameDictionaryTablesTableName(s)
	registerGetAdminClusters(s)
	registerGetAdminClustersClusterId(s)
	registerGetBroadcastMessages(s)
	registerGetBroadcastMessagesId(s)
	registerGetApplications(s)
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
	registerGetRunners(s)
	registerGetRunnersAll(s)
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
	registerGetNamespacesStorageLimitExclusions(s)
	registerGetNamespaces(s)
	registerGetNamespacesIdExists(s)
	registerGetPagesDomains(s)
	registerGetPersonalAccessTokensSelf(s)
	registerGetPersonalAccessTokensSelfAssociations(s)
	registerGetPersonalAccessTokens(s)
	registerGetPersonalAccessTokensId(s)
	registerGetSnippets(s)
	registerGetSnippetsPublic(s)
	registerGetSnippetsAll(s)
	registerGetSnippetsId(s)
	registerGetSnippetsIdRaw(s)
	registerGetSnippetsIdFilesRefFilePathRaw(s)
	registerGetSnippetsIdUserAgentDetail(s)
	registerGetHooks(s)
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
}
