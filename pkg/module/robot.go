package module

import (
	operation "github.com/heww/xk6-harbor/pkg/harbor/client/robot"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

func (h *Harbor) CreateRobot(robot models.RobotCreate) int64 {
        ctx := h.vu.Context()
	h.mustInitialized()

	params := operation.NewCreateRobotParams().WithRobot(&robot)

	res, err := h.api.Robot.CreateRobot(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to create robot %s", robot.Name)

	return IDFromLocation(h.vu.Runtime(), res.Location)
}
