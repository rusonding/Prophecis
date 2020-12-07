// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteUserByNameHandlerFunc turns a function with the right signature into a delete user by name handler
type DeleteUserByNameHandlerFunc func(DeleteUserByNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteUserByNameHandlerFunc) Handle(params DeleteUserByNameParams) middleware.Responder {
	return fn(params)
}

// DeleteUserByNameHandler interface for that can handle valid delete user by name params
type DeleteUserByNameHandler interface {
	Handle(DeleteUserByNameParams) middleware.Responder
}

// NewDeleteUserByName creates a new http.Handler for the delete user by name operation
func NewDeleteUserByName(ctx *middleware.Context, handler DeleteUserByNameHandler) *DeleteUserByName {
	return &DeleteUserByName{Context: ctx, Handler: handler}
}

/*DeleteUserByName swagger:route DELETE /cc/v1/users/name/{userName} Users deleteUserByName

Returns a user.

Optional extended description in Markdown.

*/
type DeleteUserByName struct {
	Context *middleware.Context
	Handler DeleteUserByNameHandler
}

func (o *DeleteUserByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteUserByNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}