package telegram_users

import "github.com/pkg/errors"

var (
	ErrorBadRequest        = errors.New("bad request")
	ErrorUserNotFound      = errors.New("user not found")
	ErrorInitDataIsMissing = errors.New("init data is missing, it must be present in x-init-data header")
	ErrorInitDataNotFound  = errors.New("init data not found")
	ErrorInitDataIsInvalid = errors.New("init data is invalid")
	ErrorInitDataIsEmpty   = errors.New("init data is empty")
	ErrorWebAppNotFound    = errors.New("web app id not found")
	ErrorInternal          = errors.New("internal server error")

	AuthenticationErrors = []error{
		ErrorInitDataIsMissing,
		ErrorInitDataNotFound,
		ErrorInitDataIsInvalid,
		ErrorInitDataIsEmpty,
	}
)
