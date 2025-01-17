package login

import (
	"net/http"

	"github.com/ory/kratos/session"
)

func RequiresAAL2ForTest(e HookExecutor, r *http.Request, s *session.Session) (bool, error) {
	return e.requiresAAL2(r, s, nil) // *login.Flow is nil to avoid an import cycle
}
