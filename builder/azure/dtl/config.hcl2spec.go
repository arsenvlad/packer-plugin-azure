// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package dtl

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatArtifactParameter is an auto-generated flat version of ArtifactParameter.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatArtifactParameter struct {
	Name  *string `mapstructure:"name" cty:"name" hcl:"name"`
	Value *string `mapstructure:"value" cty:"value" hcl:"value"`
	Type  *string `mapstructure:"type" cty:"type" hcl:"type"`
}

// FlatMapstructure returns a new FlatArtifactParameter.
// FlatArtifactParameter is an auto-generated flat version of ArtifactParameter.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*ArtifactParameter) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatArtifactParameter)
}

// HCL2Spec returns the hcl spec of a ArtifactParameter.
// This spec is used by HCL to read the fields of ArtifactParameter.
// The decoded values from this spec will then be applied to a FlatArtifactParameter.
func (*FlatArtifactParameter) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"name":  &hcldec.AttrSpec{Name: "name", Type: cty.String, Required: false},
		"value": &hcldec.AttrSpec{Name: "value", Type: cty.String, Required: false},
		"type":  &hcldec.AttrSpec{Name: "type", Type: cty.String, Required: false},
	}
	return s
}

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName                     *string                            `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType                   *string                            `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion                   *string                            `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug                         *bool                              `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce                         *bool                              `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError                       *string                            `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars                      map[string]string                  `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars                 []string                           `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	CloudEnvironmentName                *string                            `mapstructure:"cloud_environment_name" required:"false" cty:"cloud_environment_name" hcl:"cloud_environment_name"`
	MetadataHost                        *string                            `mapstructure:"metadata_host" required:"false" cty:"metadata_host" hcl:"metadata_host"`
	ClientID                            *string                            `mapstructure:"client_id" cty:"client_id" hcl:"client_id"`
	ClientSecret                        *string                            `mapstructure:"client_secret" cty:"client_secret" hcl:"client_secret"`
	ClientCertPath                      *string                            `mapstructure:"client_cert_path" cty:"client_cert_path" hcl:"client_cert_path"`
	ClientCertExpireTimeout             *string                            `mapstructure:"client_cert_token_timeout" required:"false" cty:"client_cert_token_timeout" hcl:"client_cert_token_timeout"`
	ClientJWT                           *string                            `mapstructure:"client_jwt" cty:"client_jwt" hcl:"client_jwt"`
	ObjectID                            *string                            `mapstructure:"object_id" cty:"object_id" hcl:"object_id"`
	TenantID                            *string                            `mapstructure:"tenant_id" required:"false" cty:"tenant_id" hcl:"tenant_id"`
	SubscriptionID                      *string                            `mapstructure:"subscription_id" cty:"subscription_id" hcl:"subscription_id"`
	UseAzureCLIAuth                     *bool                              `mapstructure:"use_azure_cli_auth" required:"false" cty:"use_azure_cli_auth" hcl:"use_azure_cli_auth"`
	CaptureNamePrefix                   *string                            `mapstructure:"capture_name_prefix" cty:"capture_name_prefix" hcl:"capture_name_prefix"`
	CaptureContainerName                *string                            `mapstructure:"capture_container_name" cty:"capture_container_name" hcl:"capture_container_name"`
	SharedGallery                       *FlatSharedImageGallery            `mapstructure:"shared_image_gallery" cty:"shared_image_gallery" hcl:"shared_image_gallery"`
	SharedGalleryDestination            *FlatSharedImageGalleryDestination `mapstructure:"shared_image_gallery_destination" cty:"shared_image_gallery_destination" hcl:"shared_image_gallery_destination"`
	SharedGalleryTimeout                *string                            `mapstructure:"shared_image_gallery_timeout" cty:"shared_image_gallery_timeout" hcl:"shared_image_gallery_timeout"`
	CustomImageCaptureTimeout           *string                            `mapstructure:"custom_image_capture_timeout" cty:"custom_image_capture_timeout" hcl:"custom_image_capture_timeout"`
	ImagePublisher                      *string                            `mapstructure:"image_publisher" cty:"image_publisher" hcl:"image_publisher"`
	ImageOffer                          *string                            `mapstructure:"image_offer" cty:"image_offer" hcl:"image_offer"`
	ImageSku                            *string                            `mapstructure:"image_sku" cty:"image_sku" hcl:"image_sku"`
	ImageVersion                        *string                            `mapstructure:"image_version" cty:"image_version" hcl:"image_version"`
	ImageUrl                            *string                            `mapstructure:"image_url" cty:"image_url" hcl:"image_url"`
	CustomManagedImageResourceGroupName *string                            `mapstructure:"custom_managed_image_resource_group_name" cty:"custom_managed_image_resource_group_name" hcl:"custom_managed_image_resource_group_name"`
	CustomManagedImageName              *string                            `mapstructure:"custom_managed_image_name" cty:"custom_managed_image_name" hcl:"custom_managed_image_name"`
	Location                            *string                            `mapstructure:"location" cty:"location" hcl:"location"`
	VMSize                              *string                            `mapstructure:"vm_size" cty:"vm_size" hcl:"vm_size"`
	ManagedImageResourceGroupName       *string                            `mapstructure:"managed_image_resource_group_name" required:"true" cty:"managed_image_resource_group_name" hcl:"managed_image_resource_group_name"`
	ManagedImageName                    *string                            `mapstructure:"managed_image_name" required:"true" cty:"managed_image_name" hcl:"managed_image_name"`
	ManagedImageStorageAccountType      *string                            `mapstructure:"managed_image_storage_account_type" required:"false" cty:"managed_image_storage_account_type" hcl:"managed_image_storage_account_type"`
	AzureTags                           map[string]*string                 `mapstructure:"azure_tags" required:"false" cty:"azure_tags" hcl:"azure_tags"`
	PlanID                              *string                            `mapstructure:"plan_id" required:"false" cty:"plan_id" hcl:"plan_id"`
	PollingDurationTimeout              *string                            `mapstructure:"polling_duration_timeout" required:"false" cty:"polling_duration_timeout" hcl:"polling_duration_timeout"`
	OSType                              *string                            `mapstructure:"os_type" required:"false" cty:"os_type" hcl:"os_type"`
	OSDiskSizeGB                        *int32                             `mapstructure:"os_disk_size_gb" required:"false" cty:"os_disk_size_gb" hcl:"os_disk_size_gb"`
	AdditionalDiskSize                  []int32                            `mapstructure:"disk_additional_size" required:"false" cty:"disk_additional_size" hcl:"disk_additional_size"`
	DiskCachingType                     *string                            `mapstructure:"disk_caching_type" required:"false" cty:"disk_caching_type" hcl:"disk_caching_type"`
	StorageType                         *string                            `mapstructure:"storage_type" cty:"storage_type" hcl:"storage_type"`
	LabVirtualNetworkName               *string                            `mapstructure:"lab_virtual_network_name" cty:"lab_virtual_network_name" hcl:"lab_virtual_network_name"`
	LabName                             *string                            `mapstructure:"lab_name" required:"true" cty:"lab_name" hcl:"lab_name"`
	LabSubnetName                       *string                            `mapstructure:"lab_subnet_name" required:"true" cty:"lab_subnet_name" hcl:"lab_subnet_name"`
	LabResourceGroupName                *string                            `mapstructure:"lab_resource_group_name" required:"true" cty:"lab_resource_group_name" hcl:"lab_resource_group_name"`
	DtlArtifacts                        []FlatDtlArtifact                  `mapstructure:"dtl_artifacts" cty:"dtl_artifacts" hcl:"dtl_artifacts"`
	VMName                              *string                            `mapstructure:"vm_name" cty:"vm_name" hcl:"vm_name"`
	DisallowPublicIP                    *bool                              `mapstructure:"disallow_public_ip" required:"false" cty:"disallow_public_ip" hcl:"disallow_public_ip"`
	SkipSysprep                         *bool                              `mapstructure:"skip_sysprep" required:"false" cty:"skip_sysprep" hcl:"skip_sysprep"`
	Password                            *string                            `cty:"password" hcl:"password"`
	Type                                *string                            `mapstructure:"communicator" cty:"communicator" hcl:"communicator"`
	PauseBeforeConnect                  *string                            `mapstructure:"pause_before_connecting" cty:"pause_before_connecting" hcl:"pause_before_connecting"`
	SSHHost                             *string                            `mapstructure:"ssh_host" cty:"ssh_host" hcl:"ssh_host"`
	SSHPort                             *int                               `mapstructure:"ssh_port" cty:"ssh_port" hcl:"ssh_port"`
	SSHUsername                         *string                            `mapstructure:"ssh_username" cty:"ssh_username" hcl:"ssh_username"`
	SSHPassword                         *string                            `mapstructure:"ssh_password" cty:"ssh_password" hcl:"ssh_password"`
	SSHKeyPairName                      *string                            `mapstructure:"ssh_keypair_name" undocumented:"true" cty:"ssh_keypair_name" hcl:"ssh_keypair_name"`
	SSHTemporaryKeyPairName             *string                            `mapstructure:"temporary_key_pair_name" undocumented:"true" cty:"temporary_key_pair_name" hcl:"temporary_key_pair_name"`
	SSHTemporaryKeyPairType             *string                            `mapstructure:"temporary_key_pair_type" cty:"temporary_key_pair_type" hcl:"temporary_key_pair_type"`
	SSHTemporaryKeyPairBits             *int                               `mapstructure:"temporary_key_pair_bits" cty:"temporary_key_pair_bits" hcl:"temporary_key_pair_bits"`
	SSHCiphers                          []string                           `mapstructure:"ssh_ciphers" cty:"ssh_ciphers" hcl:"ssh_ciphers"`
	SSHClearAuthorizedKeys              *bool                              `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys" hcl:"ssh_clear_authorized_keys"`
	SSHKEXAlgos                         []string                           `mapstructure:"ssh_key_exchange_algorithms" cty:"ssh_key_exchange_algorithms" hcl:"ssh_key_exchange_algorithms"`
	SSHPrivateKeyFile                   *string                            `mapstructure:"ssh_private_key_file" undocumented:"true" cty:"ssh_private_key_file" hcl:"ssh_private_key_file"`
	SSHCertificateFile                  *string                            `mapstructure:"ssh_certificate_file" cty:"ssh_certificate_file" hcl:"ssh_certificate_file"`
	SSHPty                              *bool                              `mapstructure:"ssh_pty" cty:"ssh_pty" hcl:"ssh_pty"`
	SSHTimeout                          *string                            `mapstructure:"ssh_timeout" cty:"ssh_timeout" hcl:"ssh_timeout"`
	SSHWaitTimeout                      *string                            `mapstructure:"ssh_wait_timeout" undocumented:"true" cty:"ssh_wait_timeout" hcl:"ssh_wait_timeout"`
	SSHAgentAuth                        *bool                              `mapstructure:"ssh_agent_auth" undocumented:"true" cty:"ssh_agent_auth" hcl:"ssh_agent_auth"`
	SSHDisableAgentForwarding           *bool                              `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding" hcl:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts                *int                               `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts" hcl:"ssh_handshake_attempts"`
	SSHBastionHost                      *string                            `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host" hcl:"ssh_bastion_host"`
	SSHBastionPort                      *int                               `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port" hcl:"ssh_bastion_port"`
	SSHBastionAgentAuth                 *bool                              `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth" hcl:"ssh_bastion_agent_auth"`
	SSHBastionUsername                  *string                            `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username" hcl:"ssh_bastion_username"`
	SSHBastionPassword                  *string                            `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password" hcl:"ssh_bastion_password"`
	SSHBastionInteractive               *bool                              `mapstructure:"ssh_bastion_interactive" cty:"ssh_bastion_interactive" hcl:"ssh_bastion_interactive"`
	SSHBastionPrivateKeyFile            *string                            `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file" hcl:"ssh_bastion_private_key_file"`
	SSHBastionCertificateFile           *string                            `mapstructure:"ssh_bastion_certificate_file" cty:"ssh_bastion_certificate_file" hcl:"ssh_bastion_certificate_file"`
	SSHFileTransferMethod               *string                            `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method" hcl:"ssh_file_transfer_method"`
	SSHProxyHost                        *string                            `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host" hcl:"ssh_proxy_host"`
	SSHProxyPort                        *int                               `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port" hcl:"ssh_proxy_port"`
	SSHProxyUsername                    *string                            `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username" hcl:"ssh_proxy_username"`
	SSHProxyPassword                    *string                            `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password" hcl:"ssh_proxy_password"`
	SSHKeepAliveInterval                *string                            `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval" hcl:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout                 *string                            `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout" hcl:"ssh_read_write_timeout"`
	SSHRemoteTunnels                    []string                           `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels" hcl:"ssh_remote_tunnels"`
	SSHLocalTunnels                     []string                           `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels" hcl:"ssh_local_tunnels"`
	SSHPublicKey                        []byte                             `mapstructure:"ssh_public_key" undocumented:"true" cty:"ssh_public_key" hcl:"ssh_public_key"`
	SSHPrivateKey                       []byte                             `mapstructure:"ssh_private_key" undocumented:"true" cty:"ssh_private_key" hcl:"ssh_private_key"`
	WinRMUser                           *string                            `mapstructure:"winrm_username" cty:"winrm_username" hcl:"winrm_username"`
	WinRMPassword                       *string                            `mapstructure:"winrm_password" cty:"winrm_password" hcl:"winrm_password"`
	WinRMHost                           *string                            `mapstructure:"winrm_host" cty:"winrm_host" hcl:"winrm_host"`
	WinRMNoProxy                        *bool                              `mapstructure:"winrm_no_proxy" cty:"winrm_no_proxy" hcl:"winrm_no_proxy"`
	WinRMPort                           *int                               `mapstructure:"winrm_port" cty:"winrm_port" hcl:"winrm_port"`
	WinRMTimeout                        *string                            `mapstructure:"winrm_timeout" cty:"winrm_timeout" hcl:"winrm_timeout"`
	WinRMUseSSL                         *bool                              `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl" hcl:"winrm_use_ssl"`
	WinRMInsecure                       *bool                              `mapstructure:"winrm_insecure" cty:"winrm_insecure" hcl:"winrm_insecure"`
	WinRMUseNTLM                        *bool                              `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm" hcl:"winrm_use_ntlm"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":                        &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":                      &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":                      &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":                             &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                             &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":                          &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":                    &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":               &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"cloud_environment_name":                   &hcldec.AttrSpec{Name: "cloud_environment_name", Type: cty.String, Required: false},
		"metadata_host":                            &hcldec.AttrSpec{Name: "metadata_host", Type: cty.String, Required: false},
		"client_id":                                &hcldec.AttrSpec{Name: "client_id", Type: cty.String, Required: false},
		"client_secret":                            &hcldec.AttrSpec{Name: "client_secret", Type: cty.String, Required: false},
		"client_cert_path":                         &hcldec.AttrSpec{Name: "client_cert_path", Type: cty.String, Required: false},
		"client_cert_token_timeout":                &hcldec.AttrSpec{Name: "client_cert_token_timeout", Type: cty.String, Required: false},
		"client_jwt":                               &hcldec.AttrSpec{Name: "client_jwt", Type: cty.String, Required: false},
		"object_id":                                &hcldec.AttrSpec{Name: "object_id", Type: cty.String, Required: false},
		"tenant_id":                                &hcldec.AttrSpec{Name: "tenant_id", Type: cty.String, Required: false},
		"subscription_id":                          &hcldec.AttrSpec{Name: "subscription_id", Type: cty.String, Required: false},
		"use_azure_cli_auth":                       &hcldec.AttrSpec{Name: "use_azure_cli_auth", Type: cty.Bool, Required: false},
		"capture_name_prefix":                      &hcldec.AttrSpec{Name: "capture_name_prefix", Type: cty.String, Required: false},
		"capture_container_name":                   &hcldec.AttrSpec{Name: "capture_container_name", Type: cty.String, Required: false},
		"shared_image_gallery":                     &hcldec.BlockSpec{TypeName: "shared_image_gallery", Nested: hcldec.ObjectSpec((*FlatSharedImageGallery)(nil).HCL2Spec())},
		"shared_image_gallery_destination":         &hcldec.BlockSpec{TypeName: "shared_image_gallery_destination", Nested: hcldec.ObjectSpec((*FlatSharedImageGalleryDestination)(nil).HCL2Spec())},
		"shared_image_gallery_timeout":             &hcldec.AttrSpec{Name: "shared_image_gallery_timeout", Type: cty.String, Required: false},
		"custom_image_capture_timeout":             &hcldec.AttrSpec{Name: "custom_image_capture_timeout", Type: cty.String, Required: false},
		"image_publisher":                          &hcldec.AttrSpec{Name: "image_publisher", Type: cty.String, Required: false},
		"image_offer":                              &hcldec.AttrSpec{Name: "image_offer", Type: cty.String, Required: false},
		"image_sku":                                &hcldec.AttrSpec{Name: "image_sku", Type: cty.String, Required: false},
		"image_version":                            &hcldec.AttrSpec{Name: "image_version", Type: cty.String, Required: false},
		"image_url":                                &hcldec.AttrSpec{Name: "image_url", Type: cty.String, Required: false},
		"custom_managed_image_resource_group_name": &hcldec.AttrSpec{Name: "custom_managed_image_resource_group_name", Type: cty.String, Required: false},
		"custom_managed_image_name":                &hcldec.AttrSpec{Name: "custom_managed_image_name", Type: cty.String, Required: false},
		"location":                                 &hcldec.AttrSpec{Name: "location", Type: cty.String, Required: false},
		"vm_size":                                  &hcldec.AttrSpec{Name: "vm_size", Type: cty.String, Required: false},
		"managed_image_resource_group_name":        &hcldec.AttrSpec{Name: "managed_image_resource_group_name", Type: cty.String, Required: false},
		"managed_image_name":                       &hcldec.AttrSpec{Name: "managed_image_name", Type: cty.String, Required: false},
		"managed_image_storage_account_type":       &hcldec.AttrSpec{Name: "managed_image_storage_account_type", Type: cty.String, Required: false},
		"azure_tags":                               &hcldec.AttrSpec{Name: "azure_tags", Type: cty.Map(cty.String), Required: false},
		"plan_id":                                  &hcldec.AttrSpec{Name: "plan_id", Type: cty.String, Required: false},
		"polling_duration_timeout":                 &hcldec.AttrSpec{Name: "polling_duration_timeout", Type: cty.String, Required: false},
		"os_type":                                  &hcldec.AttrSpec{Name: "os_type", Type: cty.String, Required: false},
		"os_disk_size_gb":                          &hcldec.AttrSpec{Name: "os_disk_size_gb", Type: cty.Number, Required: false},
		"disk_additional_size":                     &hcldec.AttrSpec{Name: "disk_additional_size", Type: cty.List(cty.Number), Required: false},
		"disk_caching_type":                        &hcldec.AttrSpec{Name: "disk_caching_type", Type: cty.String, Required: false},
		"storage_type":                             &hcldec.AttrSpec{Name: "storage_type", Type: cty.String, Required: false},
		"lab_virtual_network_name":                 &hcldec.AttrSpec{Name: "lab_virtual_network_name", Type: cty.String, Required: false},
		"lab_name":                                 &hcldec.AttrSpec{Name: "lab_name", Type: cty.String, Required: false},
		"lab_subnet_name":                          &hcldec.AttrSpec{Name: "lab_subnet_name", Type: cty.String, Required: false},
		"lab_resource_group_name":                  &hcldec.AttrSpec{Name: "lab_resource_group_name", Type: cty.String, Required: false},
		"dtl_artifacts":                            &hcldec.BlockListSpec{TypeName: "dtl_artifacts", Nested: hcldec.ObjectSpec((*FlatDtlArtifact)(nil).HCL2Spec())},
		"vm_name":                                  &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
		"disallow_public_ip":                       &hcldec.AttrSpec{Name: "disallow_public_ip", Type: cty.Bool, Required: false},
		"skip_sysprep":                             &hcldec.AttrSpec{Name: "skip_sysprep", Type: cty.Bool, Required: false},
		"password":                                 &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"communicator":                             &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":                  &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                                 &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                                 &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                             &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                             &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":                         &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":                  &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"temporary_key_pair_type":                  &hcldec.AttrSpec{Name: "temporary_key_pair_type", Type: cty.String, Required: false},
		"temporary_key_pair_bits":                  &hcldec.AttrSpec{Name: "temporary_key_pair_bits", Type: cty.Number, Required: false},
		"ssh_ciphers":                              &hcldec.AttrSpec{Name: "ssh_ciphers", Type: cty.List(cty.String), Required: false},
		"ssh_clear_authorized_keys":                &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_key_exchange_algorithms":              &hcldec.AttrSpec{Name: "ssh_key_exchange_algorithms", Type: cty.List(cty.String), Required: false},
		"ssh_private_key_file":                     &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_certificate_file":                     &hcldec.AttrSpec{Name: "ssh_certificate_file", Type: cty.String, Required: false},
		"ssh_pty":                                  &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                              &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_wait_timeout":                         &hcldec.AttrSpec{Name: "ssh_wait_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":                           &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding":             &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":                   &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":                         &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":                         &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":                   &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":                     &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":                     &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_interactive":                  &hcldec.AttrSpec{Name: "ssh_bastion_interactive", Type: cty.Bool, Required: false},
		"ssh_bastion_private_key_file":             &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_bastion_certificate_file":             &hcldec.AttrSpec{Name: "ssh_bastion_certificate_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":                 &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":                           &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":                           &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":                       &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":                       &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":                  &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":                   &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":                       &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":                        &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":                           &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":                          &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":                           &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":                           &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                               &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_no_proxy":                           &hcldec.AttrSpec{Name: "winrm_no_proxy", Type: cty.Bool, Required: false},
		"winrm_port":                               &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                            &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                            &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":                           &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":                           &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
	}
	return s
}

// FlatDtlArtifact is an auto-generated flat version of DtlArtifact.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatDtlArtifact struct {
	ArtifactName   *string                 `mapstructure:"artifact_name" cty:"artifact_name" hcl:"artifact_name"`
	RepositoryName *string                 `mapstructure:"repository_name" cty:"repository_name" hcl:"repository_name"`
	ArtifactId     *string                 `mapstructure:"artifact_id" cty:"artifact_id" hcl:"artifact_id"`
	Parameters     []FlatArtifactParameter `mapstructure:"parameters" cty:"parameters" hcl:"parameters"`
}

// FlatMapstructure returns a new FlatDtlArtifact.
// FlatDtlArtifact is an auto-generated flat version of DtlArtifact.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*DtlArtifact) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatDtlArtifact)
}

// HCL2Spec returns the hcl spec of a DtlArtifact.
// This spec is used by HCL to read the fields of DtlArtifact.
// The decoded values from this spec will then be applied to a FlatDtlArtifact.
func (*FlatDtlArtifact) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"artifact_name":   &hcldec.AttrSpec{Name: "artifact_name", Type: cty.String, Required: false},
		"repository_name": &hcldec.AttrSpec{Name: "repository_name", Type: cty.String, Required: false},
		"artifact_id":     &hcldec.AttrSpec{Name: "artifact_id", Type: cty.String, Required: false},
		"parameters":      &hcldec.BlockListSpec{TypeName: "parameters", Nested: hcldec.ObjectSpec((*FlatArtifactParameter)(nil).HCL2Spec())},
	}
	return s
}

// FlatSharedImageGallery is an auto-generated flat version of SharedImageGallery.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatSharedImageGallery struct {
	Subscription  *string `mapstructure:"subscription" cty:"subscription" hcl:"subscription"`
	ResourceGroup *string `mapstructure:"resource_group" cty:"resource_group" hcl:"resource_group"`
	GalleryName   *string `mapstructure:"gallery_name" cty:"gallery_name" hcl:"gallery_name"`
	ImageName     *string `mapstructure:"image_name" cty:"image_name" hcl:"image_name"`
	ImageVersion  *string `mapstructure:"image_version" cty:"image_version" hcl:"image_version"`
}

// FlatMapstructure returns a new FlatSharedImageGallery.
// FlatSharedImageGallery is an auto-generated flat version of SharedImageGallery.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*SharedImageGallery) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatSharedImageGallery)
}

// HCL2Spec returns the hcl spec of a SharedImageGallery.
// This spec is used by HCL to read the fields of SharedImageGallery.
// The decoded values from this spec will then be applied to a FlatSharedImageGallery.
func (*FlatSharedImageGallery) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"subscription":   &hcldec.AttrSpec{Name: "subscription", Type: cty.String, Required: false},
		"resource_group": &hcldec.AttrSpec{Name: "resource_group", Type: cty.String, Required: false},
		"gallery_name":   &hcldec.AttrSpec{Name: "gallery_name", Type: cty.String, Required: false},
		"image_name":     &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"image_version":  &hcldec.AttrSpec{Name: "image_version", Type: cty.String, Required: false},
	}
	return s
}

// FlatSharedImageGalleryDestination is an auto-generated flat version of SharedImageGalleryDestination.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatSharedImageGalleryDestination struct {
	SigDestinationResourceGroup      *string  `mapstructure:"resource_group" cty:"resource_group" hcl:"resource_group"`
	SigDestinationGalleryName        *string  `mapstructure:"gallery_name" cty:"gallery_name" hcl:"gallery_name"`
	SigDestinationImageName          *string  `mapstructure:"image_name" cty:"image_name" hcl:"image_name"`
	SigDestinationImageVersion       *string  `mapstructure:"image_version" cty:"image_version" hcl:"image_version"`
	SigDestinationReplicationRegions []string `mapstructure:"replication_regions" cty:"replication_regions" hcl:"replication_regions"`
}

// FlatMapstructure returns a new FlatSharedImageGalleryDestination.
// FlatSharedImageGalleryDestination is an auto-generated flat version of SharedImageGalleryDestination.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*SharedImageGalleryDestination) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatSharedImageGalleryDestination)
}

// HCL2Spec returns the hcl spec of a SharedImageGalleryDestination.
// This spec is used by HCL to read the fields of SharedImageGalleryDestination.
// The decoded values from this spec will then be applied to a FlatSharedImageGalleryDestination.
func (*FlatSharedImageGalleryDestination) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"resource_group":      &hcldec.AttrSpec{Name: "resource_group", Type: cty.String, Required: false},
		"gallery_name":        &hcldec.AttrSpec{Name: "gallery_name", Type: cty.String, Required: false},
		"image_name":          &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"image_version":       &hcldec.AttrSpec{Name: "image_version", Type: cty.String, Required: false},
		"replication_regions": &hcldec.AttrSpec{Name: "replication_regions", Type: cty.List(cty.String), Required: false},
	}
	return s
}
