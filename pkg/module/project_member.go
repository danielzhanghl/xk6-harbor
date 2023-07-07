package module

import (
	"github.com/dop251/goja"
	operation "github.com/heww/xk6-harbor/pkg/harbor/client/member"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
	"go.k6.io/k6/js/common"
)

func (h *Harbor) CreateProjectMember(projectName string, userID int64, roleIDs ...int64) string {
        ctx := h.vu.Context()
	h.mustInitialized()

	roleID := int64(1)
	if len(roleIDs) > 0 {
		roleID = roleIDs[0]
	}

	params := operation.NewCreateProjectMemberParams()
	params.WithProjectNameOrID(projectName).WithXIsResourceName(&varTrue).WithProjectMember(&models.ProjectMember{
		MemberUser: &models.UserEntity{UserID: userID},
		RoleID:     roleID,
	})

	res, err := h.api.Member.CreateProjectMember(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to create project member for project %s", projectName)

	return res.Location
}

type ListProjectMembersResult struct {
	ProjectMembers []*models.ProjectMemberEntity `js:"projectMembers"`
	Total          int64                         `js:"total"`
}

func (h *Harbor) ListProjectMembers(projectName string, args ...goja.Value) ListProjectMembersResult {
        ctx := h.vu.Context()
	h.mustInitialized()

	params := operation.NewListProjectMembersParams()
	params.WithProjectNameOrID(projectName).WithXIsResourceName(&varTrue)

	if len(args) > 0 {
		rt := h.vu.Runtime()
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(h.vu.Runtime(), err)
		}
	}

	res, err := h.api.Member.ListProjectMembers(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to list project members of project %s", projectName)

	return ListProjectMembersResult{
		ProjectMembers: res.Payload,
		Total:          res.XTotalCount,
	}
}
