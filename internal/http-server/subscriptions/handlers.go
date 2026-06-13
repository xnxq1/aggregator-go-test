package subscriptions

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type SubscriptionRouterManager struct {
}

// GetSubscriptions godoc
// @Summary      List subscriptions
// @Tags         subscriptions
// @Produce      json
// @Success      200
// @Router       /subscriptions [get]
func (m *SubscriptionRouterManager) GetSubscriptions(w http.ResponseWriter, r *http.Request) {

}

// CreateSubscription godoc
// @Summary      Create subscription
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Success      201
// @Router       /subscriptions [post]
func (m *SubscriptionRouterManager) CreateSubscription(w http.ResponseWriter, r *http.Request) {

}

// UpdateSubscription godoc
// @Summary      Update subscription
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Subscription ID"
// @Success      200
// @Router       /subscriptions/{id} [patch]
func (m *SubscriptionRouterManager) UpdateSubscription(w http.ResponseWriter, r *http.Request) {}

// GetSubscription godoc
// @Summary      Get subscription
// @Tags         subscriptions
// @Produce      json
// @Param        id  path  string  true  "Subscription ID"
// @Success      200
// @Router       /subscriptions/{id} [get]
func (m *SubscriptionRouterManager) GetSubscription(w http.ResponseWriter, r *http.Request) {}

// DeleteSubscription godoc
// @Summary      Delete subscription
// @Tags         subscriptions
// @Param        id  path  string  true  "Subscription ID"
// @Success      204
// @Router       /subscriptions/{id} [delete]
func (m *SubscriptionRouterManager) DeleteSubscription(w http.ResponseWriter, r *http.Request) {}

func (m *SubscriptionRouterManager) Init() *chi.Mux {
	router := chi.NewRouter()
	router.Route("/subscriptions", func(r chi.Router) {
		r.Get("/", m.GetSubscriptions)
		r.Post("/", m.CreateSubscription)
		r.Patch("/{id}", m.UpdateSubscription)
		r.Get("/{id}", m.GetSubscription)
		r.Delete("/{id}", m.DeleteSubscription)
	})
	return router
}
