// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/sapcc/kubernikus/pkg/api/models"
)

// GetClusterEventsHandlerFunc turns a function with the right signature into a get cluster events handler
type GetClusterEventsHandlerFunc func(GetClusterEventsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetClusterEventsHandlerFunc) Handle(params GetClusterEventsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetClusterEventsHandler interface for that can handle valid get cluster events params
type GetClusterEventsHandler interface {
	Handle(GetClusterEventsParams, *models.Principal) middleware.Responder
}

// NewGetClusterEvents creates a new http.Handler for the get cluster events operation
func NewGetClusterEvents(ctx *middleware.Context, handler GetClusterEventsHandler) *GetClusterEvents {
	return &GetClusterEvents{Context: ctx, Handler: handler}
}

/*
	GetClusterEvents swagger:route GET /api/v1/clusters/{name}/events getClusterEvents

Get recent events about the cluster
*/
type GetClusterEvents struct {
	Context *middleware.Context
	Handler GetClusterEventsHandler
}

func (o *GetClusterEvents) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetClusterEventsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
