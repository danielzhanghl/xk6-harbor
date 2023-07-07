package module

import (
	"time"

	operation "github.com/heww/xk6-harbor/pkg/harbor/client/gc"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

func (h *Harbor) StartGC() int64 {
        ctx := h.vu.Context()
	h.mustInitialized()

	params := operation.NewCreateGCScheduleParams().WithSchedule(&models.Schedule{
		Schedule: &models.ScheduleObj{Type: "Manual"},
		Parameters: map[string]interface{}{
			"dry_run":         false,
			"delete_untagged": true,
		},
	})

	res, err := h.api.GC.CreateGCSchedule(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to start gc")

	return IDFromLocation(h.vu.Runtime(), res.Location)
}

func (h *Harbor) GetGC(id int64) *models.GCHistory {
        ctx := h.vu.Context()
	h.mustInitialized()

	params := operation.NewGetGCParams().WithGCID(id)

	res, err := h.api.GC.GetGC(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to get gc %d", id)

	return res.Payload
}

func (h *Harbor) StartGCAndWait() {
	jobID := h.StartGC()

	for {
		gc := h.GetGC(jobID)

		if gc.JobStatus == "Success" {
			break
		} else if gc.JobStatus == "Error" || gc.JobStatus == "Stopped" {
			Throwf(h.vu.Runtime(), "expect Success but get %s for gc %d", gc.JobStatus, jobID)
		}

		time.Sleep(time.Second)
	}
}
