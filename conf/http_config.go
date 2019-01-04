package conf

import "net/http"

const (
	STATUS_CREATED               = http.StatusOK
	STATUS_BAD_REQUEST           = http.StatusBadRequest
	STATUS_INTERNAL_SERVER_ERROR = http.StatusInternalServerError
	STATUS_GONE                  = http.StatusGone
)
