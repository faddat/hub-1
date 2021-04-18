package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(ctx client.Context, router *mux.Router) {
	router.HandleFunc("/subscriptions", querySubscriptions(ctx)).
		Methods("GET")
	router.HandleFunc("/subscriptions/{id}", querySubscription(ctx)).
		Methods("GET")
	router.HandleFunc("/subscriptions/{id}/quotas", queryQuotas(ctx)).
		Methods("GET")
	router.HandleFunc("/subscriptions/{id}/quotas/{address}", queryQuota(ctx)).
		Methods("GET")
}

func registerTxRoutes(ctx client.Context, router *mux.Router) {
	router.HandleFunc("/subscriptions", txSubscribeToNode(ctx)).
		Methods("POST")
	router.HandleFunc("/subscriptions", txSubscribeToPlan(ctx)).
		Methods("POST")
	router.HandleFunc("/subscriptions", txCancel(ctx)).
		Methods("DELETE")
	router.HandleFunc("/subscriptions/{id}/quotas", txAddQuota(ctx)).
		Methods("POST")
	router.HandleFunc("/subscriptions/{id}/quotas/{address}", txUpdateQuota(ctx)).
		Methods("PUT")
}

func RegisterRoutes(ctx client.Context, router *mux.Router) {
	registerQueryRoutes(ctx, router)
	registerTxRoutes(ctx, router)
}
