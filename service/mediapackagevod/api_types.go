// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// A MediaPackage VOD Asset resource.
type AssetShallow struct {
	_ struct{} `type:"structure"`

	// The ARN of the Asset.
	Arn *string `locationName:"arn" type:"string"`

	// The time the Asset was initially submitted for Ingest.
	CreatedAt *string `locationName:"createdAt" type:"string"`

	// The unique identifier for the Asset.
	Id *string `locationName:"id" type:"string"`

	// The ID of the PackagingGroup for the Asset.
	PackagingGroupId *string `locationName:"packagingGroupId" type:"string"`

	// The resource ID to include in SPEKE key requests.
	ResourceId *string `locationName:"resourceId" type:"string"`

	// ARN of the source object in S3.
	SourceArn *string `locationName:"sourceArn" type:"string"`

	// The IAM role ARN used to access the source S3 bucket.
	SourceRoleArn *string `locationName:"sourceRoleArn" type:"string"`

	// A collection of tags associated with a resource
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s AssetShallow) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssetShallow) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedAt != nil {
		v := *s.CreatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdAt", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PackagingGroupId != nil {
		v := *s.PackagingGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "packagingGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceId != nil {
		v := *s.ResourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceArn != nil {
		v := *s.SourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceRoleArn != nil {
		v := *s.SourceRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sourceRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// CDN Authorization credentials
type Authorization struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) for the secret in AWS Secrets Manager that
	// is used for CDN authorization.
	//
	// CdnIdentifierSecret is a required field
	CdnIdentifierSecret *string `locationName:"cdnIdentifierSecret" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) for the IAM role that allows MediaPackage
	// to communicate with AWS Secrets Manager.
	//
	// SecretsRoleArn is a required field
	SecretsRoleArn *string `locationName:"secretsRoleArn" type:"string" required:"true"`
}

// String returns the string representation
func (s Authorization) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Authorization) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Authorization"}

	if s.CdnIdentifierSecret == nil {
		invalidParams.Add(aws.NewErrParamRequired("CdnIdentifierSecret"))
	}

	if s.SecretsRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecretsRoleArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Authorization) MarshalFields(e protocol.FieldEncoder) error {
	if s.CdnIdentifierSecret != nil {
		v := *s.CdnIdentifierSecret

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "cdnIdentifierSecret", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SecretsRoleArn != nil {
		v := *s.SecretsRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "secretsRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A CMAF encryption configuration.
type CmafEncryption struct {
	_ struct{} `type:"structure"`

	// A configuration for accessing an external Secure Packager and Encoder Key
	// Exchange (SPEKE) service that will provide encryption keys.
	//
	// SpekeKeyProvider is a required field
	SpekeKeyProvider *SpekeKeyProvider `locationName:"spekeKeyProvider" type:"structure" required:"true"`
}

// String returns the string representation
func (s CmafEncryption) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CmafEncryption) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CmafEncryption"}

	if s.SpekeKeyProvider == nil {
		invalidParams.Add(aws.NewErrParamRequired("SpekeKeyProvider"))
	}
	if s.SpekeKeyProvider != nil {
		if err := s.SpekeKeyProvider.Validate(); err != nil {
			invalidParams.AddNested("SpekeKeyProvider", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CmafEncryption) MarshalFields(e protocol.FieldEncoder) error {
	if s.SpekeKeyProvider != nil {
		v := s.SpekeKeyProvider

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "spekeKeyProvider", v, metadata)
	}
	return nil
}

// A CMAF packaging configuration.
type CmafPackage struct {
	_ struct{} `type:"structure"`

	// A CMAF encryption configuration.
	Encryption *CmafEncryption `locationName:"encryption" type:"structure"`

	// A list of HLS manifest configurations.
	//
	// HlsManifests is a required field
	HlsManifests []HlsManifest `locationName:"hlsManifests" type:"list" required:"true"`

	// Duration (in seconds) of each fragment. Actual fragments will berounded to
	// the nearest multiple of the source fragment duration.
	SegmentDurationSeconds *int64 `locationName:"segmentDurationSeconds" type:"integer"`
}

// String returns the string representation
func (s CmafPackage) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CmafPackage) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CmafPackage"}

	if s.HlsManifests == nil {
		invalidParams.Add(aws.NewErrParamRequired("HlsManifests"))
	}
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			invalidParams.AddNested("Encryption", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CmafPackage) MarshalFields(e protocol.FieldEncoder) error {
	if s.Encryption != nil {
		v := s.Encryption

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "encryption", v, metadata)
	}
	if s.HlsManifests != nil {
		v := s.HlsManifests

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "hlsManifests", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.SegmentDurationSeconds != nil {
		v := *s.SegmentDurationSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "segmentDurationSeconds", protocol.Int64Value(v), metadata)
	}
	return nil
}

// A Dynamic Adaptive Streaming over HTTP (DASH) encryption configuration.
type DashEncryption struct {
	_ struct{} `type:"structure"`

	// A configuration for accessing an external Secure Packager and Encoder Key
	// Exchange (SPEKE) service that will provide encryption keys.
	//
	// SpekeKeyProvider is a required field
	SpekeKeyProvider *SpekeKeyProvider `locationName:"spekeKeyProvider" type:"structure" required:"true"`
}

// String returns the string representation
func (s DashEncryption) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DashEncryption) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DashEncryption"}

	if s.SpekeKeyProvider == nil {
		invalidParams.Add(aws.NewErrParamRequired("SpekeKeyProvider"))
	}
	if s.SpekeKeyProvider != nil {
		if err := s.SpekeKeyProvider.Validate(); err != nil {
			invalidParams.AddNested("SpekeKeyProvider", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DashEncryption) MarshalFields(e protocol.FieldEncoder) error {
	if s.SpekeKeyProvider != nil {
		v := s.SpekeKeyProvider

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "spekeKeyProvider", v, metadata)
	}
	return nil
}

// A DASH manifest configuration.
type DashManifest struct {
	_ struct{} `type:"structure"`

	// Determines the position of some tags in the Media Presentation Description
	// (MPD). When set to FULL, elements like SegmentTemplate and ContentProtection
	// are included in each Representation. When set to COMPACT, duplicate elements
	// are combined and presented at the AdaptationSet level.
	ManifestLayout ManifestLayout `locationName:"manifestLayout" type:"string" enum:"true"`

	// An optional string to include in the name of the manifest.
	ManifestName *string `locationName:"manifestName" type:"string"`

	// Minimum duration (in seconds) that a player will buffer media before starting
	// the presentation.
	MinBufferTimeSeconds *int64 `locationName:"minBufferTimeSeconds" type:"integer"`

	// The Dynamic Adaptive Streaming over HTTP (DASH) profile type. When set to
	// "HBBTV_1_5", HbbTV 1.5 compliant output is enabled.
	Profile Profile `locationName:"profile" type:"string" enum:"true"`

	// A StreamSelection configuration.
	StreamSelection *StreamSelection `locationName:"streamSelection" type:"structure"`
}

// String returns the string representation
func (s DashManifest) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DashManifest) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ManifestLayout) > 0 {
		v := s.ManifestLayout

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "manifestLayout", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ManifestName != nil {
		v := *s.ManifestName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "manifestName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MinBufferTimeSeconds != nil {
		v := *s.MinBufferTimeSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "minBufferTimeSeconds", protocol.Int64Value(v), metadata)
	}
	if len(s.Profile) > 0 {
		v := s.Profile

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "profile", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StreamSelection != nil {
		v := s.StreamSelection

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "streamSelection", v, metadata)
	}
	return nil
}

// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
type DashPackage struct {
	_ struct{} `type:"structure"`

	// A list of DASH manifest configurations.
	//
	// DashManifests is a required field
	DashManifests []DashManifest `locationName:"dashManifests" type:"list" required:"true"`

	// A Dynamic Adaptive Streaming over HTTP (DASH) encryption configuration.
	Encryption *DashEncryption `locationName:"encryption" type:"structure"`

	// A list of triggers that controls when the outgoing Dynamic Adaptive Streaming
	// over HTTP (DASH)Media Presentation Description (MPD) will be partitioned
	// into multiple periods. If empty, the content will notbe partitioned into
	// more than one period. If the list contains "ADS", new periods will be created
	// wherethe Asset contains SCTE-35 ad markers.
	PeriodTriggers []PeriodTriggersElement `locationName:"periodTriggers" type:"list"`

	// Duration (in seconds) of each segment. Actual segments will berounded to
	// the nearest multiple of the source segment duration.
	SegmentDurationSeconds *int64 `locationName:"segmentDurationSeconds" type:"integer"`

	// Determines the type of SegmentTemplate included in the Media Presentation
	// Description (MPD). When set to NUMBER_WITH_TIMELINE, a full timeline is presented
	// in each SegmentTemplate, with $Number$ media URLs. When set to TIME_WITH_TIMELINE,
	// a full timeline is presented in each SegmentTemplate, with $Time$ media URLs.
	// When set to NUMBER_WITH_DURATION, only a duration is included in each SegmentTemplate,
	// with $Number$ media URLs.
	SegmentTemplateFormat SegmentTemplateFormat `locationName:"segmentTemplateFormat" type:"string" enum:"true"`
}

// String returns the string representation
func (s DashPackage) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DashPackage) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DashPackage"}

	if s.DashManifests == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashManifests"))
	}
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			invalidParams.AddNested("Encryption", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DashPackage) MarshalFields(e protocol.FieldEncoder) error {
	if s.DashManifests != nil {
		v := s.DashManifests

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "dashManifests", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Encryption != nil {
		v := s.Encryption

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "encryption", v, metadata)
	}
	if s.PeriodTriggers != nil {
		v := s.PeriodTriggers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "periodTriggers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.SegmentDurationSeconds != nil {
		v := *s.SegmentDurationSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "segmentDurationSeconds", protocol.Int64Value(v), metadata)
	}
	if len(s.SegmentTemplateFormat) > 0 {
		v := s.SegmentTemplateFormat

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "segmentTemplateFormat", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// The endpoint URL used to access an Asset using one PackagingConfiguration.
type EgressEndpoint struct {
	_ struct{} `type:"structure"`

	// The ID of the PackagingConfiguration being applied to the Asset.
	PackagingConfigurationId *string `locationName:"packagingConfigurationId" type:"string"`

	// The URL of the parent manifest for the repackaged Asset.
	Url *string `locationName:"url" type:"string"`
}

// String returns the string representation
func (s EgressEndpoint) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EgressEndpoint) MarshalFields(e protocol.FieldEncoder) error {
	if s.PackagingConfigurationId != nil {
		v := *s.PackagingConfigurationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "packagingConfigurationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Url != nil {
		v := *s.Url

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "url", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An HTTP Live Streaming (HLS) encryption configuration.
type HlsEncryption struct {
	_ struct{} `type:"structure"`

	// A constant initialization vector for encryption (optional).When not specified
	// the initialization vector will be periodically rotated.
	ConstantInitializationVector *string `locationName:"constantInitializationVector" type:"string"`

	// The encryption method to use.
	EncryptionMethod EncryptionMethod `locationName:"encryptionMethod" type:"string" enum:"true"`

	// A configuration for accessing an external Secure Packager and Encoder Key
	// Exchange (SPEKE) service that will provide encryption keys.
	//
	// SpekeKeyProvider is a required field
	SpekeKeyProvider *SpekeKeyProvider `locationName:"spekeKeyProvider" type:"structure" required:"true"`
}

// String returns the string representation
func (s HlsEncryption) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *HlsEncryption) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "HlsEncryption"}

	if s.SpekeKeyProvider == nil {
		invalidParams.Add(aws.NewErrParamRequired("SpekeKeyProvider"))
	}
	if s.SpekeKeyProvider != nil {
		if err := s.SpekeKeyProvider.Validate(); err != nil {
			invalidParams.AddNested("SpekeKeyProvider", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s HlsEncryption) MarshalFields(e protocol.FieldEncoder) error {
	if s.ConstantInitializationVector != nil {
		v := *s.ConstantInitializationVector

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "constantInitializationVector", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.EncryptionMethod) > 0 {
		v := s.EncryptionMethod

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "encryptionMethod", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.SpekeKeyProvider != nil {
		v := s.SpekeKeyProvider

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "spekeKeyProvider", v, metadata)
	}
	return nil
}

// An HTTP Live Streaming (HLS) manifest configuration.
type HlsManifest struct {
	_ struct{} `type:"structure"`

	// This setting controls how ad markers are included in the packaged OriginEndpoint."NONE"
	// will omit all SCTE-35 ad markers from the output."PASSTHROUGH" causes the
	// manifest to contain a copy of the SCTE-35 admarkers (comments) taken directly
	// from the input HTTP Live Streaming (HLS) manifest."SCTE35_ENHANCED" generates
	// ad markers and blackout tags based on SCTE-35messages in the input source.
	AdMarkers AdMarkers `locationName:"adMarkers" type:"string" enum:"true"`

	// When enabled, an I-Frame only stream will be included in the output.
	IncludeIframeOnlyStream *bool `locationName:"includeIframeOnlyStream" type:"boolean"`

	// An optional string to include in the name of the manifest.
	ManifestName *string `locationName:"manifestName" type:"string"`

	// The interval (in seconds) between each EXT-X-PROGRAM-DATE-TIME taginserted
	// into manifests. Additionally, when an interval is specifiedID3Timed Metadata
	// messages will be generated every 5 seconds using theingest time of the content.If
	// the interval is not specified, or set to 0, thenno EXT-X-PROGRAM-DATE-TIME
	// tags will be inserted into manifests and noID3Timed Metadata messages will
	// be generated. Note that irrespectiveof this parameter, if any ID3 Timed Metadata
	// is found in HTTP Live Streaming (HLS) input,it will be passed through to
	// HLS output.
	ProgramDateTimeIntervalSeconds *int64 `locationName:"programDateTimeIntervalSeconds" type:"integer"`

	// When enabled, the EXT-X-KEY tag will be repeated in output manifests.
	RepeatExtXKey *bool `locationName:"repeatExtXKey" type:"boolean"`

	// A StreamSelection configuration.
	StreamSelection *StreamSelection `locationName:"streamSelection" type:"structure"`
}

// String returns the string representation
func (s HlsManifest) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s HlsManifest) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.AdMarkers) > 0 {
		v := s.AdMarkers

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "adMarkers", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.IncludeIframeOnlyStream != nil {
		v := *s.IncludeIframeOnlyStream

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "includeIframeOnlyStream", protocol.BoolValue(v), metadata)
	}
	if s.ManifestName != nil {
		v := *s.ManifestName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "manifestName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProgramDateTimeIntervalSeconds != nil {
		v := *s.ProgramDateTimeIntervalSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "programDateTimeIntervalSeconds", protocol.Int64Value(v), metadata)
	}
	if s.RepeatExtXKey != nil {
		v := *s.RepeatExtXKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "repeatExtXKey", protocol.BoolValue(v), metadata)
	}
	if s.StreamSelection != nil {
		v := s.StreamSelection

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "streamSelection", v, metadata)
	}
	return nil
}

// An HTTP Live Streaming (HLS) packaging configuration.
type HlsPackage struct {
	_ struct{} `type:"structure"`

	// An HTTP Live Streaming (HLS) encryption configuration.
	Encryption *HlsEncryption `locationName:"encryption" type:"structure"`

	// A list of HLS manifest configurations.
	//
	// HlsManifests is a required field
	HlsManifests []HlsManifest `locationName:"hlsManifests" type:"list" required:"true"`

	// Duration (in seconds) of each fragment. Actual fragments will berounded to
	// the nearest multiple of the source fragment duration.
	SegmentDurationSeconds *int64 `locationName:"segmentDurationSeconds" type:"integer"`

	// When enabled, audio streams will be placed in rendition groups in the output.
	UseAudioRenditionGroup *bool `locationName:"useAudioRenditionGroup" type:"boolean"`
}

// String returns the string representation
func (s HlsPackage) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *HlsPackage) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "HlsPackage"}

	if s.HlsManifests == nil {
		invalidParams.Add(aws.NewErrParamRequired("HlsManifests"))
	}
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			invalidParams.AddNested("Encryption", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s HlsPackage) MarshalFields(e protocol.FieldEncoder) error {
	if s.Encryption != nil {
		v := s.Encryption

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "encryption", v, metadata)
	}
	if s.HlsManifests != nil {
		v := s.HlsManifests

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "hlsManifests", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.SegmentDurationSeconds != nil {
		v := *s.SegmentDurationSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "segmentDurationSeconds", protocol.Int64Value(v), metadata)
	}
	if s.UseAudioRenditionGroup != nil {
		v := *s.UseAudioRenditionGroup

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "useAudioRenditionGroup", protocol.BoolValue(v), metadata)
	}
	return nil
}

// A Microsoft Smooth Streaming (MSS) encryption configuration.
type MssEncryption struct {
	_ struct{} `type:"structure"`

	// A configuration for accessing an external Secure Packager and Encoder Key
	// Exchange (SPEKE) service that will provide encryption keys.
	//
	// SpekeKeyProvider is a required field
	SpekeKeyProvider *SpekeKeyProvider `locationName:"spekeKeyProvider" type:"structure" required:"true"`
}

// String returns the string representation
func (s MssEncryption) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MssEncryption) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MssEncryption"}

	if s.SpekeKeyProvider == nil {
		invalidParams.Add(aws.NewErrParamRequired("SpekeKeyProvider"))
	}
	if s.SpekeKeyProvider != nil {
		if err := s.SpekeKeyProvider.Validate(); err != nil {
			invalidParams.AddNested("SpekeKeyProvider", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s MssEncryption) MarshalFields(e protocol.FieldEncoder) error {
	if s.SpekeKeyProvider != nil {
		v := s.SpekeKeyProvider

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "spekeKeyProvider", v, metadata)
	}
	return nil
}

// A Microsoft Smooth Streaming (MSS) manifest configuration.
type MssManifest struct {
	_ struct{} `type:"structure"`

	// An optional string to include in the name of the manifest.
	ManifestName *string `locationName:"manifestName" type:"string"`

	// A StreamSelection configuration.
	StreamSelection *StreamSelection `locationName:"streamSelection" type:"structure"`
}

// String returns the string representation
func (s MssManifest) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s MssManifest) MarshalFields(e protocol.FieldEncoder) error {
	if s.ManifestName != nil {
		v := *s.ManifestName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "manifestName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamSelection != nil {
		v := s.StreamSelection

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "streamSelection", v, metadata)
	}
	return nil
}

// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
type MssPackage struct {
	_ struct{} `type:"structure"`

	// A Microsoft Smooth Streaming (MSS) encryption configuration.
	Encryption *MssEncryption `locationName:"encryption" type:"structure"`

	// A list of MSS manifest configurations.
	//
	// MssManifests is a required field
	MssManifests []MssManifest `locationName:"mssManifests" type:"list" required:"true"`

	// The duration (in seconds) of each segment.
	SegmentDurationSeconds *int64 `locationName:"segmentDurationSeconds" type:"integer"`
}

// String returns the string representation
func (s MssPackage) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MssPackage) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MssPackage"}

	if s.MssManifests == nil {
		invalidParams.Add(aws.NewErrParamRequired("MssManifests"))
	}
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			invalidParams.AddNested("Encryption", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s MssPackage) MarshalFields(e protocol.FieldEncoder) error {
	if s.Encryption != nil {
		v := s.Encryption

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "encryption", v, metadata)
	}
	if s.MssManifests != nil {
		v := s.MssManifests

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "mssManifests", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.SegmentDurationSeconds != nil {
		v := *s.SegmentDurationSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "segmentDurationSeconds", protocol.Int64Value(v), metadata)
	}
	return nil
}

// A MediaPackage VOD PackagingConfiguration resource.
type PackagingConfiguration struct {
	_ struct{} `type:"structure"`

	// The ARN of the PackagingConfiguration.
	Arn *string `locationName:"arn" type:"string"`

	// A CMAF packaging configuration.
	CmafPackage *CmafPackage `locationName:"cmafPackage" type:"structure"`

	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *DashPackage `locationName:"dashPackage" type:"structure"`

	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *HlsPackage `locationName:"hlsPackage" type:"structure"`

	// The ID of the PackagingConfiguration.
	Id *string `locationName:"id" type:"string"`

	// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
	MssPackage *MssPackage `locationName:"mssPackage" type:"structure"`

	// The ID of a PackagingGroup.
	PackagingGroupId *string `locationName:"packagingGroupId" type:"string"`

	// A collection of tags associated with a resource
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s PackagingConfiguration) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PackagingConfiguration) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CmafPackage != nil {
		v := s.CmafPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "cmafPackage", v, metadata)
	}
	if s.DashPackage != nil {
		v := s.DashPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "dashPackage", v, metadata)
	}
	if s.HlsPackage != nil {
		v := s.HlsPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "hlsPackage", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MssPackage != nil {
		v := s.MssPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "mssPackage", v, metadata)
	}
	if s.PackagingGroupId != nil {
		v := *s.PackagingGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "packagingGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// A MediaPackage VOD PackagingGroup resource.
type PackagingGroup struct {
	_ struct{} `type:"structure"`

	// The ARN of the PackagingGroup.
	Arn *string `locationName:"arn" type:"string"`

	// CDN Authorization credentials
	Authorization *Authorization `locationName:"authorization" type:"structure"`

	// The fully qualified domain name for Assets in the PackagingGroup.
	DomainName *string `locationName:"domainName" type:"string"`

	// The ID of the PackagingGroup.
	Id *string `locationName:"id" type:"string"`

	// A collection of tags associated with a resource
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s PackagingGroup) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PackagingGroup) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Authorization != nil {
		v := s.Authorization

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "authorization", v, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// A configuration for accessing an external Secure Packager and Encoder Key
// Exchange (SPEKE) service that will provide encryption keys.
type SpekeKeyProvider struct {
	_ struct{} `type:"structure"`

	// An Amazon Resource Name (ARN) of an IAM role that AWS ElementalMediaPackage
	// will assume when accessing the key provider service.
	//
	// RoleArn is a required field
	RoleArn *string `locationName:"roleArn" type:"string" required:"true"`

	// The system IDs to include in key requests.
	//
	// SystemIds is a required field
	SystemIds []string `locationName:"systemIds" type:"list" required:"true"`

	// The URL of the external key provider service.
	//
	// Url is a required field
	Url *string `locationName:"url" type:"string" required:"true"`
}

// String returns the string representation
func (s SpekeKeyProvider) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SpekeKeyProvider) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SpekeKeyProvider"}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}

	if s.SystemIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SystemIds"))
	}

	if s.Url == nil {
		invalidParams.Add(aws.NewErrParamRequired("Url"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SpekeKeyProvider) MarshalFields(e protocol.FieldEncoder) error {
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SystemIds != nil {
		v := s.SystemIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "systemIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Url != nil {
		v := *s.Url

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "url", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A StreamSelection configuration.
type StreamSelection struct {
	_ struct{} `type:"structure"`

	// The maximum video bitrate (bps) to include in output.
	MaxVideoBitsPerSecond *int64 `locationName:"maxVideoBitsPerSecond" type:"integer"`

	// The minimum video bitrate (bps) to include in output.
	MinVideoBitsPerSecond *int64 `locationName:"minVideoBitsPerSecond" type:"integer"`

	// A directive that determines the order of streams in the output.
	StreamOrder StreamOrder `locationName:"streamOrder" type:"string" enum:"true"`
}

// String returns the string representation
func (s StreamSelection) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StreamSelection) MarshalFields(e protocol.FieldEncoder) error {
	if s.MaxVideoBitsPerSecond != nil {
		v := *s.MaxVideoBitsPerSecond

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxVideoBitsPerSecond", protocol.Int64Value(v), metadata)
	}
	if s.MinVideoBitsPerSecond != nil {
		v := *s.MinVideoBitsPerSecond

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "minVideoBitsPerSecond", protocol.Int64Value(v), metadata)
	}
	if len(s.StreamOrder) > 0 {
		v := s.StreamOrder

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "streamOrder", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}