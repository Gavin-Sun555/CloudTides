// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserDetailsHandlerFunc turns a function with the right signature into a user details handler
type UserDetailsHandlerFunc func(UserDetailsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserDetailsHandlerFunc) Handle(params UserDetailsParams) middleware.Responder {
	return fn(params)
}

// UserDetailsHandler interface for that can handle valid user details params
type UserDetailsHandler interface {
	Handle(UserDetailsParams) middleware.Responder
}

// NewUserDetails creates a new http.Handler for the user details operation
func NewUserDetails(ctx *middleware.Context, handler UserDetailsHandler) *UserDetails {
	return &UserDetails{Context: ctx, Handler: handler}
}

/*UserDetails swagger:route GET /users/get_details user userDetails

get user profile

*/
type UserDetails struct {
	Context *middleware.Context
	Handler UserDetailsHandler
}

func (o *UserDetails) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserDetailsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UserDetailsNotFoundBody user details not found body
//
// swagger:model UserDetailsNotFoundBody
type UserDetailsNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this user details not found body
func (o *UserDetailsNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UserDetailsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserDetailsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res UserDetailsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UserDetailsOKBody user details o k body
//
// swagger:model UserDetailsOKBody
type UserDetailsOKBody struct {

	// message
	Message string `json:"message,omitempty"`

	// results
	Results *UserDetailsOKBodyResults `json:"results,omitempty"`
}

// Validate validates this user details o k body
func (o *UserDetailsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserDetailsOKBody) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(o.Results) { // not required
		return nil
	}

	if o.Results != nil {
		if err := o.Results.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userDetailsOK" + "." + "results")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserDetailsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserDetailsOKBody) UnmarshalBinary(b []byte) error {
	var res UserDetailsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UserDetailsOKBodyResults user details o k body results
//
// swagger:model UserDetailsOKBodyResults
type UserDetailsOKBodyResults struct {

	// city
	City string `json:"city,omitempty"`

	// company name
	CompanyName string `json:"companyName,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// position
	Position string `json:"position,omitempty"`
}

// Validate validates this user details o k body results
func (o *UserDetailsOKBodyResults) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UserDetailsOKBodyResults) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserDetailsOKBodyResults) UnmarshalBinary(b []byte) error {
	var res UserDetailsOKBodyResults
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
