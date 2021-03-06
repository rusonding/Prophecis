// Code generated by go-swagger; DO NOT EDIT.

package namespaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAllNamespacesParams creates a new GetAllNamespacesParams object
// with the default values initialized.
func NewGetAllNamespacesParams() GetAllNamespacesParams {

	var (
		// initialize parameters with default values

		pageDefault = int64(0)
		sizeDefault = int64(0)
	)

	return GetAllNamespacesParams{
		Page: &pageDefault,

		Size: &sizeDefault,
	}
}

// GetAllNamespacesParams contains all the bound params for the get all namespaces operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetAllNamespaces
type GetAllNamespacesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*page number.
	  In: query
	  Default: 0
	*/
	Page *int64
	/*entity number per page.
	  In: query
	  Default: 0
	*/
	Size *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAllNamespacesParams() beforehand.
func (o *GetAllNamespacesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qPage, qhkPage, _ := qs.GetOK("page")
	if err := o.bindPage(qPage, qhkPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qSize, qhkSize, _ := qs.GetOK("size")
	if err := o.bindSize(qSize, qhkSize, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPage binds and validates parameter Page from query.
func (o *GetAllNamespacesParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAllNamespacesParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("page", "query", "int64", raw)
	}
	o.Page = &value

	return nil
}

// bindSize binds and validates parameter Size from query.
func (o *GetAllNamespacesParams) bindSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAllNamespacesParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("size", "query", "int64", raw)
	}
	o.Size = &value

	return nil
}
