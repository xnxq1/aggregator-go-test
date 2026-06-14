package http_server

import (
	"aggregator-go-test/internal/domain"
	http_server "aggregator-go-test/internal/http-server"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type SubscriptionRouterManager struct {
	SubscriptionService SubscriptionServicePort
}

// GetSubscriptions godoc
// @Summary      List subscriptions
// @Tags         subscriptions
// @Produce      json
// @Success      200
// @Router       /subscriptions [get]
func (m *SubscriptionRouterManager) GetSubscriptions(w http.ResponseWriter, r *http.Request) {
}

var validate = validator.New()

// CreateSubscription godoc
// @Summary      Create subscription
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        request  body  CreateSubscriptionRequest  true  "Subscription payload"
// @Success      201  {object}  domain.Subscription
// @Failure      400
// @Failure      500
// @Router       /subscriptions [post]
func (m *SubscriptionRouterManager) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var schema CreateSubscriptionRequest
	if err := http_server.DecodeJson(r, &schema); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validate.Struct(schema); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := m.SubscriptionService.Create(r.Context(), &domain.Subscription{
		ServiceName: schema.ServiceName,
		Price:       schema.Price,
		UserId:      schema.UserId,
		StartDate:   schema.StartDate,
		EndDate:     schema.EndDate,
	})
	if err != nil {
		log.Printf("create subscription: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	http_server.WriteResponse(w, http.StatusCreated, http_server.SuccessResponse[domain.Subscription]{Data: *res})
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
func (m *SubscriptionRouterManager) GetSubscription(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	sub, err := m.SubscriptionService.GetById(r.Context(), id)
	if err != nil {
		log.Printf("get subscription: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	http_server.WriteResponse(w, http.StatusOK, http_server.SuccessResponse[domain.Subscription]{Data: *sub})
}

// DeleteSubscription godoc
// @Summary      Delete subscription
// @Tags         subscriptions
// @Param        id  path  string  true  "Subscription ID"
// @Success      204
// @Router       /subscriptions/{id} [delete]
func (m *SubscriptionRouterManager) DeleteSubscription(w http.ResponseWriter, r *http.Request) {}

func NewSubscriptionRouterManager(svc SubscriptionServicePort) *SubscriptionRouterManager {
	return &SubscriptionRouterManager{SubscriptionService: svc}
}

func (m *SubscriptionRouterManager) Routes() *chi.Mux {
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
