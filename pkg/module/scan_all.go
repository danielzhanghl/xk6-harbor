package module

import (
	operation "github.com/heww/xk6-harbor/pkg/harbor/client/scan_all"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

func (h *Harbor) StartScanAll() {
        ctx := h.vu.Context()
	h.mustInitialized()

	params := operation.NewCreateScanAllScheduleParams().
		WithSchedule(&models.Schedule{
			Schedule: &models.ScheduleObj{
				Type: models.ScheduleObjTypeManual,
			},
		})

	_, err := h.api.ScanAll.CreateScanAllSchedule(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to start scan all")
}

func (h *Harbor) GetScanAllMetrics() *models.Stats {
        ctx := h.vu.Context()
	h.mustInitialized()

	parmas := operation.NewGetLatestScanAllMetricsParams()

	res, err := h.api.ScanAll.GetLatestScanAllMetrics(ctx, parmas)
	Checkf(h.vu.Runtime(), err, "failed to get metrics of scan all")

	return res.Payload
}
