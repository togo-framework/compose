// Package compose is the Compose togo plugin: Compose-by-prompt pipeline: a queued-post inbox + composer UI + dynamic insert. Owner briefs it in prose, an agent drafts and publishes, the brief closes with a result ref.
//
// It self-registers a provider on blank-import and mounts its routes onto the
// kernel. The concrete implementation is ported from the fadymondy.com app under
// internal/server — this scaffold wires the provider + a health route.
package compose

import (
	"net/http"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("compose", togo.PriorityService, func(k *togo.Kernel) error {
		k.Router.Get("/api/compose/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"plugin":"compose","status":"ok"}`))
		})
		return nil
	})
}
