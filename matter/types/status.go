package types

type StatusCode uint8

const (
	StatusCodeSuccess                   StatusCode = 0x00
	StatusCodeFailure                   StatusCode = 0x01
	StatusCodeInvalidSubscription       StatusCode = 0x7D
	StatusCodeUnsupportedAccess         StatusCode = 0x7E
	StatusCodeUnsupportedEndpoint       StatusCode = 0x7F
	StatusCodeInvalidAction             StatusCode = 0x80
	StatusCodeUnsupportedCommand        StatusCode = 0x81
	StatusCodeInvalidCommand            StatusCode = 0x85
	StatusCodeUnsupportedAttribute      StatusCode = 0x86
	StatusCodeConstraintError           StatusCode = 0x87
	StatusCodeUnsupportedWrite          StatusCode = 0x88
	StatusCodeResourceExhausted         StatusCode = 0x89
	StatusCodeNotFound                  StatusCode = 0x8B
	StatusCodeUnreportableAttribute     StatusCode = 0x8C
	StatusCodeInvalidDataType           StatusCode = 0x8D
	StatusCodeUnsupportedRead           StatusCode = 0x8F
	StatusCodeDataVersionMismatch       StatusCode = 0x92
	StatusCodeTimeout                   StatusCode = 0x94
	StatusCodeUnsupportedNode           StatusCode = 0x9B
	StatusCodeBusy                      StatusCode = 0x9C
	StatusCodeAccessRestricted          StatusCode = 0x9D
	StatusCodeUnsupportedCluster        StatusCode = 0xC3
	StatusCodeNoUpstreamSubscription    StatusCode = 0xC5
	StatusCodeNeedsTimedInteraction     StatusCode = 0xC6
	StatusCodeUnsupportedEvent          StatusCode = 0xC7
	StatusCodePathsExhausted            StatusCode = 0xC8
	StatusCodeTimedRequestMismatch      StatusCode = 0xC9
	StatusCodeFailsafeRequired          StatusCode = 0xCA
	StatusCodeInvalidInState            StatusCode = 0xCB
	StatusCodeNoCommandResponse         StatusCode = 0xCC
	StatusCodeTermsAndConditionsChanged StatusCode = 0xCD
	StatusCodeMaintenanceRequired       StatusCode = 0xCE
)

func (sc StatusCode) String() string {
	switch sc {
	case StatusCodeSuccess:
		return "SUCCESS"
	case StatusCodeFailure:
		return "FAILURE"
	case StatusCodeInvalidSubscription:
		return "INVALID_SUBSCRIPTION"
	case StatusCodeUnsupportedAccess:
		return "UNSUPPORTED_ACCESS"
	case StatusCodeUnsupportedEndpoint:
		return "UNSUPPORTED_ENDPOINT"
	case StatusCodeInvalidAction:
		return "INVALID_ACTION"
	case StatusCodeUnsupportedCommand:
		return "UNSUPPORTED_COMMAND"
	case StatusCodeInvalidCommand:
		return "INVALID_COMMAND"
	case StatusCodeUnsupportedAttribute:
		return "UNSUPPORTED_ATTRIBUTE"
	case StatusCodeConstraintError:
		return "CONSTRAINT_ERROR"
	case StatusCodeUnsupportedWrite:
		return "UNSUPPORTED_WRITE"
	case StatusCodeResourceExhausted:
		return "RESOURCE_EXHAUSTED"
	case StatusCodeNotFound:
		return "NOT_FOUND"
	case StatusCodeUnreportableAttribute:
		return "UNREPORTABLE_ATTRIBUTE"
	case StatusCodeInvalidDataType:
		return "INVALID_DATA_TYPE"
	case StatusCodeUnsupportedRead:
		return "UNSUPPORTED_READ"
	case StatusCodeDataVersionMismatch:
		return "DATA_VERSION_MISMATCH"
	case StatusCodeTimeout:
		return "TIMEOUT"
	case StatusCodeUnsupportedNode:
		return "UNSUPPORTED_NODE"
	case StatusCodeBusy:
		return "BUSY"
	case StatusCodeAccessRestricted:
		return "ACCESS_RESTRICTED"
	case StatusCodeUnsupportedCluster:
		return "UNSUPPORTED_CLUSTER"
	case StatusCodeNoUpstreamSubscription:
		return "NO_UPSTREAM_SUBSCRIPTION"
	case StatusCodeNeedsTimedInteraction:
		return "NEEDS_TIMED_INTERACTION"
	case StatusCodeUnsupportedEvent:
		return "UNSUPPORTED_EVENT"
	case StatusCodePathsExhausted:
		return "PATHS_EXHAUSTED"
	case StatusCodeTimedRequestMismatch:
		return "TIMED_REQUEST_MISMATCH"
	case StatusCodeFailsafeRequired:
		return "FAILSAFE_REQUIRED"
	case StatusCodeInvalidInState:
		return "INVALID_IN_STATE"
	case StatusCodeNoCommandResponse:
		return "NO_COMMAND_RESPONSE"
	case StatusCodeTermsAndConditionsChanged:
		return "TERMS_AND_CONDITIONS_CHANGED"
	case StatusCodeMaintenanceRequired:
		return "MAINTENANCE_REQUIRED"
	default:
		return "UNKNOWN_STATUS_CODE"
	}
}
