// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
)

func encodeEnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest(req EnterpriseAdminSetGithubActionsPermissionsEnterpriseApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseRequest(req EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminSetAllowedActionsEnterpriseRequest(req SelectedActions) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest(req EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest(req *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest(req EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest(req EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeGistsCreateCommentRequest(req GistsCreateCommentApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeGistsUpdateCommentRequest(req GistsUpdateCommentApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActivityMarkNotificationsAsReadRequest(req *ActivityMarkNotificationsAsReadApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActivitySetThreadSubscriptionRequest(req *ActivitySetThreadSubscriptionApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsSetGithubActionsPermissionsOrganizationRequest(req ActionsSetGithubActionsPermissionsOrganizationApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsSetSelectedRepositoriesEnabledGithubActionsOrganizationRequest(req ActionsSetSelectedRepositoriesEnabledGithubActionsOrganizationApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsSetAllowedActionsOrganizationRequest(req *SelectedActions) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsCreateSelfHostedRunnerGroupForOrgRequest(req ActionsCreateSelfHostedRunnerGroupForOrgApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsUpdateSelfHostedRunnerGroupForOrgRequest(req ActionsUpdateSelfHostedRunnerGroupForOrgApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsSetRepoAccessToSelfHostedRunnerGroupInOrgRequest(req ActionsSetRepoAccessToSelfHostedRunnerGroupInOrgApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsSetSelfHostedRunnersInGroupForOrgRequest(req ActionsSetSelfHostedRunnersInGroupForOrgApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsCreateOrUpdateOrgSecretRequest(req ActionsCreateOrUpdateOrgSecretApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsSetSelectedReposForOrgSecretRequest(req ActionsSetSelectedReposForOrgSecretApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeProjectsCreateForOrgRequest(req ProjectsCreateForOrgApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsUpdateInOrgRequest(req *TeamsUpdateInOrgApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsCreateDiscussionInOrgRequest(req TeamsCreateDiscussionInOrgApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsUpdateDiscussionInOrgRequest(req *TeamsUpdateDiscussionInOrgApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsCreateDiscussionCommentInOrgRequest(req TeamsCreateDiscussionCommentInOrgApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsUpdateDiscussionCommentInOrgRequest(req TeamsUpdateDiscussionCommentInOrgApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReactionsCreateForTeamDiscussionCommentInOrgRequest(req ReactionsCreateForTeamDiscussionCommentInOrgApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReactionsCreateForTeamDiscussionInOrgRequest(req ReactionsCreateForTeamDiscussionInOrgApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsAddOrUpdateMembershipForUserInOrgRequest(req *TeamsAddOrUpdateMembershipForUserInOrgApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsAddOrUpdateProjectPermissionsInOrgRequest(req *TeamsAddOrUpdateProjectPermissionsInOrgApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsAddOrUpdateRepoPermissionsInOrgRequest(req *TeamsAddOrUpdateRepoPermissionsInOrgApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest(req TeamsCreateOrUpdateIdpGroupConnectionsInOrgApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeProjectsUpdateCardRequest(req *ProjectsUpdateCardApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeProjectsUpdateColumnRequest(req ProjectsUpdateColumnApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeProjectsMoveColumnRequest(req ProjectsMoveColumnApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeProjectsUpdateRequest(req *ProjectsUpdateApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeProjectsCreateColumnRequest(req ProjectsCreateColumnApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsSetGithubActionsPermissionsRepositoryRequest(req ActionsSetGithubActionsPermissionsRepositoryApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsSetAllowedActionsRepositoryRequest(req *SelectedActions) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsCreateOrUpdateRepoSecretRequest(req ActionsCreateOrUpdateRepoSecretApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposUpdateBranchProtectionRequest(req ReposUpdateBranchProtectionApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeChecksCreateSuiteRequest(req ChecksCreateSuiteApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeChecksSetSuitesPreferencesRequest(req ChecksSetSuitesPreferencesApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeCodeScanningUploadSarifRequest(req CodeScanningUploadSarifApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposUpdateCommitCommentRequest(req ReposUpdateCommitCommentApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeMigrationsUpdateImportRequest(req *MigrationsUpdateImportApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeInteractionsSetRestrictionsForRepoRequest(req InteractionLimit) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposUpdateInvitationRequest(req *ReposUpdateInvitationApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeIssuesUpdateLabelRequest(req *IssuesUpdateLabelApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposMergeUpstreamRequest(req ReposMergeUpstreamApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeIssuesUpdateMilestoneRequest(req *IssuesUpdateMilestoneApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActivityMarkRepoNotificationsAsReadRequest(req *ActivityMarkRepoNotificationsAsReadApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeProjectsCreateForRepoRequest(req ProjectsCreateForRepoApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodePullsUpdateReviewCommentRequest(req PullsUpdateReviewCommentApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodePullsCreateReplyForReviewCommentRequest(req PullsCreateReplyForReviewCommentApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodePullsCreateReviewRequest(req *PullsCreateReviewApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodePullsUpdateReviewRequest(req PullsUpdateReviewApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodePullsDismissReviewRequest(req PullsDismissReviewApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodePullsSubmitReviewRequest(req PullsSubmitReviewApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposUpdateReleaseAssetRequest(req *ReposUpdateReleaseAssetApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposUpdateReleaseRequest(req *ReposUpdateReleaseApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeSecretScanningUpdateAlertRequest(req SecretScanningUpdateAlertApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposCreateCommitStatusRequest(req ReposCreateCommitStatusApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActivitySetRepoSubscriptionRequest(req *ActivitySetRepoSubscriptionApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposReplaceAllTopicsRequest(req ReposReplaceAllTopicsApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposTransferRequest(req ReposTransferApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposCreateUsingTemplateRequest(req ReposCreateUsingTemplateApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsCreateOrUpdateEnvironmentSecretRequest(req ActionsCreateOrUpdateEnvironmentSecretApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminProvisionAndInviteEnterpriseGroupRequest(req EnterpriseAdminProvisionAndInviteEnterpriseGroupApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminSetInformationForProvisionedEnterpriseGroupRequest(req EnterpriseAdminSetInformationForProvisionedEnterpriseGroupApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminProvisionAndInviteEnterpriseUserRequest(req EnterpriseAdminProvisionAndInviteEnterpriseUserApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminSetInformationForProvisionedEnterpriseUserRequest(req EnterpriseAdminSetInformationForProvisionedEnterpriseUserApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminUpdateAttributeForEnterpriseUserRequest(req EnterpriseAdminUpdateAttributeForEnterpriseUserApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsCreateDiscussionLegacyRequest(req TeamsCreateDiscussionLegacyApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsUpdateDiscussionLegacyRequest(req *TeamsUpdateDiscussionLegacyApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsCreateDiscussionCommentLegacyRequest(req TeamsCreateDiscussionCommentLegacyApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsUpdateDiscussionCommentLegacyRequest(req TeamsUpdateDiscussionCommentLegacyApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReactionsCreateForTeamDiscussionCommentLegacyRequest(req ReactionsCreateForTeamDiscussionCommentLegacyApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReactionsCreateForTeamDiscussionLegacyRequest(req ReactionsCreateForTeamDiscussionLegacyApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsAddOrUpdateMembershipForUserLegacyRequest(req *TeamsAddOrUpdateMembershipForUserLegacyApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeProjectsCreateForAuthenticatedUserRequest(req ProjectsCreateForAuthenticatedUserApplicationJSONRequest) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}
