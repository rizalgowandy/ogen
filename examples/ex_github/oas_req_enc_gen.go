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

func encodeEnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest(req EnterpriseAdminSetGithubActionsPermissionsEnterpriseApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseRequest(req EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminSetAllowedActionsEnterpriseRequest(req SelectedActions) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest(req EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest(req *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest(req EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest(req EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeGistsCreateCommentRequest(req GistsCreateCommentApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeGistsUpdateCommentRequest(req GistsUpdateCommentApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeMarkdownRenderRequest(req MarkdownRenderApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActivityMarkNotificationsAsReadRequest(req *ActivityMarkNotificationsAsReadApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActivitySetThreadSubscriptionRequest(req *ActivitySetThreadSubscriptionApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsSetGithubActionsPermissionsOrganizationRequest(req ActionsSetGithubActionsPermissionsOrganizationApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsSetSelectedRepositoriesEnabledGithubActionsOrganizationRequest(req ActionsSetSelectedRepositoriesEnabledGithubActionsOrganizationApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsSetAllowedActionsOrganizationRequest(req *SelectedActions) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsCreateSelfHostedRunnerGroupForOrgRequest(req ActionsCreateSelfHostedRunnerGroupForOrgApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsUpdateSelfHostedRunnerGroupForOrgRequest(req ActionsUpdateSelfHostedRunnerGroupForOrgApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsSetRepoAccessToSelfHostedRunnerGroupInOrgRequest(req ActionsSetRepoAccessToSelfHostedRunnerGroupInOrgApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsSetSelfHostedRunnersInGroupForOrgRequest(req ActionsSetSelfHostedRunnersInGroupForOrgApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsCreateOrUpdateOrgSecretRequest(req ActionsCreateOrUpdateOrgSecretApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsSetSelectedReposForOrgSecretRequest(req ActionsSetSelectedReposForOrgSecretApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeProjectsCreateForOrgRequest(req ProjectsCreateForOrgApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsUpdateInOrgRequest(req *TeamsUpdateInOrgApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsCreateDiscussionInOrgRequest(req TeamsCreateDiscussionInOrgApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsUpdateDiscussionInOrgRequest(req *TeamsUpdateDiscussionInOrgApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsCreateDiscussionCommentInOrgRequest(req TeamsCreateDiscussionCommentInOrgApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsUpdateDiscussionCommentInOrgRequest(req TeamsUpdateDiscussionCommentInOrgApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReactionsCreateForTeamDiscussionCommentInOrgRequest(req ReactionsCreateForTeamDiscussionCommentInOrgApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReactionsCreateForTeamDiscussionInOrgRequest(req ReactionsCreateForTeamDiscussionInOrgApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsAddOrUpdateMembershipForUserInOrgRequest(req *TeamsAddOrUpdateMembershipForUserInOrgApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsAddOrUpdateProjectPermissionsInOrgRequest(req *TeamsAddOrUpdateProjectPermissionsInOrgApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsAddOrUpdateRepoPermissionsInOrgRequest(req *TeamsAddOrUpdateRepoPermissionsInOrgApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest(req TeamsCreateOrUpdateIdpGroupConnectionsInOrgApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeProjectsUpdateCardRequest(req *ProjectsUpdateCardApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeProjectsUpdateColumnRequest(req ProjectsUpdateColumnApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeProjectsMoveColumnRequest(req ProjectsMoveColumnApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeProjectsUpdateRequest(req *ProjectsUpdateApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeProjectsCreateColumnRequest(req ProjectsCreateColumnApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsSetGithubActionsPermissionsRepositoryRequest(req ActionsSetGithubActionsPermissionsRepositoryApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsSetAllowedActionsRepositoryRequest(req *SelectedActions) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsCreateOrUpdateRepoSecretRequest(req ActionsCreateOrUpdateRepoSecretApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposUpdateBranchProtectionRequest(req ReposUpdateBranchProtectionApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeChecksCreateSuiteRequest(req ChecksCreateSuiteApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeChecksSetSuitesPreferencesRequest(req ChecksSetSuitesPreferencesApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeCodeScanningUploadSarifRequest(req CodeScanningUploadSarifApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposUpdateCommitCommentRequest(req ReposUpdateCommitCommentApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeMigrationsUpdateImportRequest(req *MigrationsUpdateImportApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeInteractionsSetRestrictionsForRepoRequest(req InteractionLimit) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposUpdateInvitationRequest(req *ReposUpdateInvitationApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeIssuesUpdateLabelRequest(req *IssuesUpdateLabelApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposMergeUpstreamRequest(req ReposMergeUpstreamApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeIssuesUpdateMilestoneRequest(req *IssuesUpdateMilestoneApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActivityMarkRepoNotificationsAsReadRequest(req *ActivityMarkRepoNotificationsAsReadApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeProjectsCreateForRepoRequest(req ProjectsCreateForRepoApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodePullsUpdateReviewCommentRequest(req PullsUpdateReviewCommentApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodePullsCreateReplyForReviewCommentRequest(req PullsCreateReplyForReviewCommentApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodePullsCreateReviewRequest(req *PullsCreateReviewApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodePullsUpdateReviewRequest(req PullsUpdateReviewApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodePullsDismissReviewRequest(req PullsDismissReviewApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodePullsSubmitReviewRequest(req PullsSubmitReviewApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposUpdateReleaseAssetRequest(req *ReposUpdateReleaseAssetApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposUpdateReleaseRequest(req *ReposUpdateReleaseApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeSecretScanningUpdateAlertRequest(req SecretScanningUpdateAlertApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposCreateCommitStatusRequest(req ReposCreateCommitStatusApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActivitySetRepoSubscriptionRequest(req *ActivitySetRepoSubscriptionApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposReplaceAllTopicsRequest(req ReposReplaceAllTopicsApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposTransferRequest(req ReposTransferApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReposCreateUsingTemplateRequest(req ReposCreateUsingTemplateApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeActionsCreateOrUpdateEnvironmentSecretRequest(req ActionsCreateOrUpdateEnvironmentSecretApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminProvisionAndInviteEnterpriseGroupRequest(req EnterpriseAdminProvisionAndInviteEnterpriseGroupApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminSetInformationForProvisionedEnterpriseGroupRequest(req EnterpriseAdminSetInformationForProvisionedEnterpriseGroupApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminProvisionAndInviteEnterpriseUserRequest(req EnterpriseAdminProvisionAndInviteEnterpriseUserApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminSetInformationForProvisionedEnterpriseUserRequest(req EnterpriseAdminSetInformationForProvisionedEnterpriseUserApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeEnterpriseAdminUpdateAttributeForEnterpriseUserRequest(req EnterpriseAdminUpdateAttributeForEnterpriseUserApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsCreateDiscussionLegacyRequest(req TeamsCreateDiscussionLegacyApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsUpdateDiscussionLegacyRequest(req *TeamsUpdateDiscussionLegacyApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsCreateDiscussionCommentLegacyRequest(req TeamsCreateDiscussionCommentLegacyApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsUpdateDiscussionCommentLegacyRequest(req TeamsUpdateDiscussionCommentLegacyApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReactionsCreateForTeamDiscussionCommentLegacyRequest(req ReactionsCreateForTeamDiscussionCommentLegacyApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeReactionsCreateForTeamDiscussionLegacyRequest(req ReactionsCreateForTeamDiscussionLegacyApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeTeamsAddOrUpdateMembershipForUserLegacyRequest(req *TeamsAddOrUpdateMembershipForUserLegacyApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}

func encodeProjectsCreateForAuthenticatedUserRequest(req ProjectsCreateForAuthenticatedUserApplicationJSONReq) (data []byte, contentType string, err error) {
	return json.Encode(req), "application/json", nil
}
