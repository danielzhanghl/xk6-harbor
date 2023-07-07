package module

import (
	"fmt"

	"github.com/dop251/goja"
	operation "github.com/heww/xk6-harbor/pkg/harbor/client/user"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
	"go.k6.io/k6/js/common"
)

func (h *Harbor) CreateUser(username string, passwords ...string) int64 {
        ctx := h.vu.Context()
        h.mustInitialized()

        password := "Harbor12345"
	if len(passwords) > 0 {
		password = passwords[0]
	}

	params := operation.NewCreateUserParams()
	params.WithUserReq(&models.UserCreationReq{
		Username: username,
		Email:    fmt.Sprintf("%s@goharbor.io", username),
		Password: password,
		Realname: username,
	})

	res, err := h.api.User.CreateUser(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to create user %s", username)

	return IDFromLocation(h.vu.Runtime(), res.Location)
}

func (h *Harbor) DeleteUser(userid int64) {
        ctx := h.vu.Context()
	h.mustInitialized()

	params := operation.NewDeleteUserParams()
	params.WithUserID(userid)

	_, err := h.api.User.DeleteUser(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to delete user %v", userid)
}

type ListUsersResult struct {
	Users []*models.UserResp `js:"users"`
	Total int64              `js:"total"`
}

func (h *Harbor) ListUsers(args ...goja.Value) ListUsersResult {
        ctx := h.vu.Context()
	h.mustInitialized()

	params := operation.NewListUsersParams()

	if len(args) > 0 {
		rt := h.vu.Runtime()
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(h.vu.Runtime(), err)
		}
	}

	res, err := h.api.User.ListUsers(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to list users")

	return ListUsersResult{
		Users: res.Payload,
		Total: res.XTotalCount,
	}
}

type SearchUsersResult struct {
	Users []*models.UserSearchRespItem `js:"users"`
	Total int64                        `js:"total"`
}

func (h *Harbor) SearchUsers(args ...goja.Value) SearchUsersResult {
        ctx := h.vu.Context()
	h.mustInitialized()

	params := operation.NewSearchUsersParams()

	if len(args) > 0 {
		rt := h.vu.Runtime()
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(h.vu.Runtime(), err)
		}
	}

	res, err := h.api.User.SearchUsers(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to list users")

	return SearchUsersResult{
		Users: res.Payload,
		Total: res.XTotalCount,
	}
}
