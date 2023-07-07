package module

import (
	operation "github.com/heww/xk6-harbor/pkg/harbor/client/scanner"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

func (h *Harbor) CreateScanner(registration models.ScannerRegistrationReq) string {
        ctx := h.vu.Context()
	h.mustInitialized()

	params := operation.NewCreateScannerParams().WithRegistration(&registration)

	res, err := h.api.Scanner.CreateScanner(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to create scanner %s", *registration.Name)

	return NameFromLocation(res.Location)
}

func (h *Harbor) SetScannerAsDefault(registrationID string) {
        ctx := h.vu.Context()
	h.mustInitialized()

	params := operation.NewSetScannerAsDefaultParams().
		WithRegistrationID(registrationID).
		WithPayload(&models.IsDefault{IsDefault: true})

	_, err := h.api.Scanner.SetScannerAsDefault(ctx, params)

	Checkf(h.vu.Runtime(), err, "failed to set scanner %s as default", registrationID)
}
