package eirini

import (
	"context"
	"fmt"
	"net/http"

	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/eirini/models/cf"
	"code.cloudfoundry.org/eirini/opi"
	"code.cloudfoundry.org/runtimeschema/cc_messages"
)

const (
	//Environment Variable Names
	EnvDownloadURL        = "DOWNLOAD_URL"
	EnvBuildpacks         = "BUILDPACKS"
	EnvDropletUploadURL   = "DROPLET_UPLOAD_URL"
	EnvAppID              = "APP_ID"
	EnvStagingGUID        = "STAGING_GUID"
	EnvCompletionCallback = "COMPLETION_CALLBACK"
	EnvCfUsername         = "CF_USERNAME"
	EnvCfPassword         = "CF_PASSWORD"
	EnvAPIAddress         = "API_ADDRESS"
	EnvEiriniAddress      = "EIRINI_ADDRESS"

	RegisteredRoutes = "routes"

	CCUploaderInternalURL = "cc-uploader.service.cf.internal"
	CCCertsMountPath      = "/etc/config/certs"
	CCCertsVolumeName     = "cc-certs-volume"
	CCAPICertName         = "cc-server-crt"
	CCAPIKeyName          = "cc-server-crt-key"
	CCUploaderCertName    = "cc-uploader-crt"
	CCUploaderKeyName     = "cc-uploader-crt-key"
	CCInternalCACertName  = "internal-ca-cert"

	UAACertName           = "uaa-server-cert"
	UAAKeyName            = "uaa-server-cert-key"
	UAAInternalCACertName = "uaa-internal-ca-cert"
)

//go:generate counterfeiter . CfClient
type CfClient interface {
	GetDropletByAppGuid(string) ([]byte, error)
	PushDroplet(string, string) error
	GetAppBitsByAppGuid(string) (*http.Response, error)
}

type Config struct {
	Properties Properties `yaml:"opi"`
}

type Properties struct {
	KubeConfig        string `yaml:"kube_config"`
	KubeNamespace     string `yaml:"kube_namespace"`
	KubeEndpoint      string `yaml:"kube_endpoint"`
	NatsPassword      string `yaml:"nats_password"`
	NatsIP            string `yaml:"nats_ip"`
	CcUploaderIP      string `yaml:"cc_uploader_ip"`
	CcAPI             string `yaml:"api_endpoint"`
	CcInternalAPI     string `yaml:"cc_internal_api"`
	Backend           string `yaml:"backend"`
	CfUsername        string `yaml:"cf_username"`
	CfPassword        string `yaml:"cf_password"`
	CcUser            string `yaml:"cc_internal_user"`
	CcPassword        string `yaml:"cc_internal_password"`
	CCCertsSecretName string `yaml:"cc_certs_secret_name"`
	RegistryAddress   string `yaml:"external_eirini_address"`
	EiriniAddress     string `yaml:"eirini_address"`
	StagerImageTag    string `yaml:"stager_image_tag"`
	UseIngress        bool   `yaml:"use_ingress"`

	MetricsSourceAddress string `yaml:"metrics_source_address"`

	LoggregatorAddress  string `yaml:"loggregator_address"`
	LoggregatorCertPath string `yaml:"loggergator_cert_path"`
	LoggregatorKeyPath  string `yaml:"loggregator_key_path"`
	LoggregatorCAPath   string `yaml:"loggregator_ca_path"`

	CCCertPath string `yaml:"cc_cert_path"`
	CCKeyPath  string `yaml:"cc_key_path"`
	CCCAPath   string `yaml:"cc_ca_path"`
}

//go:generate counterfeiter . RemoveCallbackFunc
type RemoveCallbackFunc func(string) error

//go:generate counterfeiter . Stager
type Stager interface {
	Stage(string, cc_messages.StagingRequestFromCC) error //stage
	CompleteStaging(*models.TaskCallbackResponse) error
}

type StagerConfig struct {
	CfUsername    string
	CfPassword    string
	APIAddress    string
	EiriniAddress string
	Image         string
}

//go:generate counterfeiter . Extractor
type Extractor interface {
	Extract(src, targetDir string) error
}

//go:generate counterfeiter . Bifrost
type Bifrost interface {
	Transfer(ctx context.Context, request cf.DesireLRPRequest) error
	List(ctx context.Context) ([]*models.DesiredLRPSchedulingInfo, error)
	Update(ctx context.Context, update cf.UpdateDesiredLRPRequest) error
	Stop(ctx context.Context, identifier opi.LRPIdentifier) error
	GetApp(ctx context.Context, identifier opi.LRPIdentifier) *models.DesiredLRP
	GetInstances(ctx context.Context, identifier opi.LRPIdentifier) ([]*cf.Instance, error)
}

func GetInternalServiceName(appName string) string {
	//Prefix service as the appName could start with numerical characters, which is not allowed
	return fmt.Sprintf("cf-%s", appName)
}

func GetInternalHeadlessServiceName(appName string) string {
	//Prefix service as the appName could start with numerical characters, which is not allowed
	return fmt.Sprintf("cf-%s-headless", appName)
}
