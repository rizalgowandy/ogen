// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/errors"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
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
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
)

func encodeCreateAdmissionregistrationV1MutatingWebhookConfigurationRequest(req IoK8sAPIAdmissionregistrationV1MutatingWebhookConfiguration, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateAdmissionregistrationV1ValidatingWebhookConfigurationRequest(req IoK8sAPIAdmissionregistrationV1ValidatingWebhookConfiguration, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateApiextensionsV1CustomResourceDefinitionRequest(req IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateApiregistrationV1APIServiceRequest(req IoK8sKubeAggregatorPkgApisApiregistrationV1APIService, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateAuthenticationV1TokenReviewRequest(req IoK8sAPIAuthenticationV1TokenReview, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateAuthorizationV1SelfSubjectAccessReviewRequest(req IoK8sAPIAuthorizationV1SelfSubjectAccessReview, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateAuthorizationV1SelfSubjectRulesReviewRequest(req IoK8sAPIAuthorizationV1SelfSubjectRulesReview, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateAuthorizationV1SubjectAccessReviewRequest(req IoK8sAPIAuthorizationV1SubjectAccessReview, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateCertificatesV1CertificateSigningRequestRequest(req IoK8sAPICertificatesV1CertificateSigningRequest, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateCoreV1NamespaceRequest(req IoK8sAPICoreV1Namespace, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateCoreV1NodeRequest(req IoK8sAPICoreV1Node, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateCoreV1PersistentVolumeRequest(req IoK8sAPICoreV1PersistentVolume, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateFlowcontrolApiserverV1beta1FlowSchemaRequest(req IoK8sAPIFlowcontrolV1beta1FlowSchema, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateFlowcontrolApiserverV1beta1PriorityLevelConfigurationRequest(req IoK8sAPIFlowcontrolV1beta1PriorityLevelConfiguration, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateFlowcontrolApiserverV1beta2FlowSchemaRequest(req IoK8sAPIFlowcontrolV1beta2FlowSchema, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateFlowcontrolApiserverV1beta2PriorityLevelConfigurationRequest(req IoK8sAPIFlowcontrolV1beta2PriorityLevelConfiguration, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateInternalApiserverV1alpha1StorageVersionRequest(req IoK8sAPIApiserverinternalV1alpha1StorageVersion, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateNetworkingV1IngressClassRequest(req IoK8sAPINetworkingV1IngressClass, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateNodeV1RuntimeClassRequest(req IoK8sAPINodeV1RuntimeClass, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateNodeV1alpha1RuntimeClassRequest(req IoK8sAPINodeV1alpha1RuntimeClass, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateNodeV1beta1RuntimeClassRequest(req IoK8sAPINodeV1beta1RuntimeClass, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreatePolicyV1beta1PodSecurityPolicyRequest(req IoK8sAPIPolicyV1beta1PodSecurityPolicy, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateRbacAuthorizationV1ClusterRoleRequest(req IoK8sAPIRbacV1ClusterRole, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateRbacAuthorizationV1ClusterRoleBindingRequest(req IoK8sAPIRbacV1ClusterRoleBinding, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateSchedulingV1PriorityClassRequest(req IoK8sAPISchedulingV1PriorityClass, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateStorageV1CSIDriverRequest(req IoK8sAPIStorageV1CSIDriver, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateStorageV1CSINodeRequest(req IoK8sAPIStorageV1CSINode, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateStorageV1StorageClassRequest(req IoK8sAPIStorageV1StorageClass, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeCreateStorageV1VolumeAttachmentRequest(req IoK8sAPIStorageV1VolumeAttachment, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteAdmissionregistrationV1CollectionMutatingWebhookConfigurationRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteApiextensionsV1CollectionCustomResourceDefinitionRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteApiregistrationV1CollectionAPIServiceRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteCertificatesV1CollectionCertificateSigningRequestRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteCoreV1CollectionNodeRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteCoreV1CollectionPersistentVolumeRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteFlowcontrolApiserverV1beta1CollectionFlowSchemaRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteFlowcontrolApiserverV1beta1CollectionPriorityLevelConfigurationRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteFlowcontrolApiserverV1beta2CollectionFlowSchemaRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteFlowcontrolApiserverV1beta2CollectionPriorityLevelConfigurationRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteInternalApiserverV1alpha1CollectionStorageVersionRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteNetworkingV1CollectionIngressClassRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteNodeV1CollectionRuntimeClassRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteNodeV1alpha1CollectionRuntimeClassRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteNodeV1beta1CollectionRuntimeClassRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeletePolicyV1beta1CollectionPodSecurityPolicyRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteRbacAuthorizationV1CollectionClusterRoleRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteRbacAuthorizationV1CollectionClusterRoleBindingRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteSchedulingV1CollectionPriorityClassRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteStorageV1CollectionCSIDriverRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteStorageV1CollectionCSINodeRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteStorageV1CollectionStorageClassRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}

func encodeDeleteStorageV1CollectionVolumeAttachmentRequest(req IoK8sApimachineryPkgApisMetaV1DeleteOptions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	return nil, "", errors.New(`*/* encoder not implemented`)
}
