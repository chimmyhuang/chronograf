package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/influxdata/mrfusion/models"
)

// NewPatchSourcesIDKapacitorsKapaIDParams creates a new PatchSourcesIDKapacitorsKapaIDParams object
// with the default values initialized.
func NewPatchSourcesIDKapacitorsKapaIDParams() PatchSourcesIDKapacitorsKapaIDParams {
	var ()
	return PatchSourcesIDKapacitorsKapaIDParams{}
}

// PatchSourcesIDKapacitorsKapaIDParams contains all the bound params for the patch sources ID kapacitors kapa ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters PatchSourcesIDKapacitorsKapaID
type PatchSourcesIDKapacitorsKapaIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*kapacitor configuration
	  Required: true
	  In: body
	*/
	Config *models.Kapacitor
	/*ID of the source
	  Required: true
	  In: path
	*/
	ID string
	/*ID of a kapacitor backend
	  Required: true
	  In: path
	*/
	KapaID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PatchSourcesIDKapacitorsKapaIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Kapacitor
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("config", "body"))
			} else {
				res = append(res, errors.NewParseError("config", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Config = &body
			}
		}

	} else {
		res = append(res, errors.Required("config", "body"))
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	rKapaID, rhkKapaID, _ := route.Params.GetOK("kapa_id")
	if err := o.bindKapaID(rKapaID, rhkKapaID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchSourcesIDKapacitorsKapaIDParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ID = raw

	return nil
}

func (o *PatchSourcesIDKapacitorsKapaIDParams) bindKapaID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.KapaID = raw

	return nil
}