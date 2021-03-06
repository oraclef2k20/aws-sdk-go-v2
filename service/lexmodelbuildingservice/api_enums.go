// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

type ChannelStatus string

// Enum values for ChannelStatus
const (
	ChannelStatusInProgress ChannelStatus = "IN_PROGRESS"
	ChannelStatusCreated    ChannelStatus = "CREATED"
	ChannelStatusFailed     ChannelStatus = "FAILED"
)

func (enum ChannelStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChannelStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChannelType string

// Enum values for ChannelType
const (
	ChannelTypeFacebook  ChannelType = "Facebook"
	ChannelTypeSlack     ChannelType = "Slack"
	ChannelTypeTwilioSms ChannelType = "Twilio-Sms"
	ChannelTypeKik       ChannelType = "Kik"
)

func (enum ChannelType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChannelType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ContentType string

// Enum values for ContentType
const (
	ContentTypePlainText     ContentType = "PlainText"
	ContentTypeSsml          ContentType = "SSML"
	ContentTypeCustomPayload ContentType = "CustomPayload"
)

func (enum ContentType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ContentType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Destination string

// Enum values for Destination
const (
	DestinationCloudwatchLogs Destination = "CLOUDWATCH_LOGS"
	DestinationS3             Destination = "S3"
)

func (enum Destination) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Destination) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ExportStatus string

// Enum values for ExportStatus
const (
	ExportStatusInProgress ExportStatus = "IN_PROGRESS"
	ExportStatusReady      ExportStatus = "READY"
	ExportStatusFailed     ExportStatus = "FAILED"
)

func (enum ExportStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ExportStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ExportType string

// Enum values for ExportType
const (
	ExportTypeAlexaSkillsKit ExportType = "ALEXA_SKILLS_KIT"
	ExportTypeLex            ExportType = "LEX"
)

func (enum ExportType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ExportType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FulfillmentActivityType string

// Enum values for FulfillmentActivityType
const (
	FulfillmentActivityTypeReturnIntent FulfillmentActivityType = "ReturnIntent"
	FulfillmentActivityTypeCodeHook     FulfillmentActivityType = "CodeHook"
)

func (enum FulfillmentActivityType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FulfillmentActivityType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ImportStatus string

// Enum values for ImportStatus
const (
	ImportStatusInProgress ImportStatus = "IN_PROGRESS"
	ImportStatusComplete   ImportStatus = "COMPLETE"
	ImportStatusFailed     ImportStatus = "FAILED"
)

func (enum ImportStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ImportStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Locale string

// Enum values for Locale
const (
	LocaleEnUs Locale = "en-US"
	LocaleEnGb Locale = "en-GB"
	LocaleDeDe Locale = "de-DE"
)

func (enum Locale) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Locale) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LogType string

// Enum values for LogType
const (
	LogTypeAudio LogType = "AUDIO"
	LogTypeText  LogType = "TEXT"
)

func (enum LogType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LogType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MergeStrategy string

// Enum values for MergeStrategy
const (
	MergeStrategyOverwriteLatest MergeStrategy = "OVERWRITE_LATEST"
	MergeStrategyFailOnConflict  MergeStrategy = "FAIL_ON_CONFLICT"
)

func (enum MergeStrategy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MergeStrategy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ObfuscationSetting string

// Enum values for ObfuscationSetting
const (
	ObfuscationSettingNone               ObfuscationSetting = "NONE"
	ObfuscationSettingDefaultObfuscation ObfuscationSetting = "DEFAULT_OBFUSCATION"
)

func (enum ObfuscationSetting) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ObfuscationSetting) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProcessBehavior string

// Enum values for ProcessBehavior
const (
	ProcessBehaviorSave  ProcessBehavior = "SAVE"
	ProcessBehaviorBuild ProcessBehavior = "BUILD"
)

func (enum ProcessBehavior) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProcessBehavior) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ReferenceType string

// Enum values for ReferenceType
const (
	ReferenceTypeIntent     ReferenceType = "Intent"
	ReferenceTypeBot        ReferenceType = "Bot"
	ReferenceTypeBotAlias   ReferenceType = "BotAlias"
	ReferenceTypeBotChannel ReferenceType = "BotChannel"
)

func (enum ReferenceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ReferenceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeBot      ResourceType = "BOT"
	ResourceTypeIntent   ResourceType = "INTENT"
	ResourceTypeSlotType ResourceType = "SLOT_TYPE"
)

func (enum ResourceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SlotConstraint string

// Enum values for SlotConstraint
const (
	SlotConstraintRequired SlotConstraint = "Required"
	SlotConstraintOptional SlotConstraint = "Optional"
)

func (enum SlotConstraint) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SlotConstraint) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SlotValueSelectionStrategy string

// Enum values for SlotValueSelectionStrategy
const (
	SlotValueSelectionStrategyOriginalValue SlotValueSelectionStrategy = "ORIGINAL_VALUE"
	SlotValueSelectionStrategyTopResolution SlotValueSelectionStrategy = "TOP_RESOLUTION"
)

func (enum SlotValueSelectionStrategy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SlotValueSelectionStrategy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Status string

// Enum values for Status
const (
	StatusBuilding          Status = "BUILDING"
	StatusReady             Status = "READY"
	StatusReadyBasicTesting Status = "READY_BASIC_TESTING"
	StatusFailed            Status = "FAILED"
	StatusNotBuilt          Status = "NOT_BUILT"
)

func (enum Status) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Status) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StatusType string

// Enum values for StatusType
const (
	StatusTypeDetected StatusType = "Detected"
	StatusTypeMissed   StatusType = "Missed"
)

func (enum StatusType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StatusType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
