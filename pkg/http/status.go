package http

type StatusCode int

const (
	StatusContinue StatusCode = iota + 100
	StatusSwitchingProtocols
	StatusProcessing
	StatusEarlyHints
)

const (
	StatusOK StatusCode = iota + 200
	StatusCreated
	StatusAccepted
	StatusNonAuthoritativeInfo
	StatusNoContent
	StatusResetContent
	StatusPartialContent
	StatusMultiStatus
	StatusAlreadyReported
	StatusIMUsed
)

const (
	StatusMultipleChoices StatusCode = iota + 300
	StatusMovedPermanently
	StatusFound
	StatusSeeOther
	StatusNotModified
	StatusUseProxy
	_ // RFC 9110, 15.4.7 (Unused)
	StatusTemporaryRedirect
	StatusPermanentRedirect
)

const (
	StatusBadRequest StatusCode = iota + 400 // RFC 9110, 15.5.1
	StatusUnauthorized
	StatusPaymentRequired
	StatusForbidden
	StatusNotFound
	StatusMethodNotAllowed
	StatusNotAcceptable
	StatusProxyAuthRequired
	StatusRequestTimeout
	StatusConflict
	StatusGone
	StatusLengthRequired
	StatusPreconditionFailed
	StatusRequestEntityTooLarge
	StatusRequestURITooLong
	StatusUnsupportedMediaType
	StatusRequestedRangeNotSatisfiable
	StatusExpectationFailed
	StatusTeapot
	StatusMisdirectedRequest
	StatusUnprocessableEntity
	StatusLocked
	StatusFailedDependency
	StatusTooEarly
	StatusUpgradeRequired
	StatusPreconditionRequired
	StatusTooManyRequests
	StatusRequestHeaderFieldsTooLarge
	StatusUnavailableForLegalReasons
)

const (
	StatusInternalServerError StatusCode = iota + 500
	StatusNotImplemented
	StatusBadGateway
	StatusServiceUnavailable
	StatusGatewayTimeout
	StatusHTTPVersionNotSupported
	StatusVariantAlsoNegotiates
	StatusInsufficientStorage
	StatusLoopDetected
	StatusNotExtended
	StatusNetworkAuthenticationRequired
)

func (sc StatusCode) Int() int {
	return int(sc)
}

func (sc StatusCode) IsInfo() bool {
	return sc >= StatusContinue && sc <= StatusEarlyHints
}

func (sc StatusCode) IsSuccess() bool {
	return sc >= StatusOK && sc <= StatusIMUsed
}

func (sc StatusCode) IsClientError() bool {
	return sc >= StatusBadRequest && sc <= StatusUnavailableForLegalReasons
}

func (sc StatusCode) IsServerError() bool {
	return sc >= StatusInternalServerError && sc <= StatusNetworkAuthenticationRequired
}
