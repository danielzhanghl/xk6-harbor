package module

import (
	"net/url"

	"github.com/dop251/goja"
	operation "github.com/heww/xk6-harbor/pkg/harbor/client/repository"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
	"go.k6.io/k6/js/common"
)

func (h *Harbor) DeleteRepository(projectName, repositoryName string) {
        ctx := h.vu.Context()
	h.mustInitialized()

	params := operation.NewDeleteRepositoryParams()
	params.WithProjectName(projectName).WithRepositoryName(url.PathEscape(repositoryName))

	_, err := h.api.Repository.DeleteRepository(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to delete repository %s/%s", projectName, repositoryName)
}

func (h *Harbor) GetRepository(projectName, repositoryName string) *models.Repository {
        ctx := h.vu.Context()
	h.mustInitialized()

	params := operation.NewGetRepositoryParams()
	params.WithProjectName(projectName)
	params.WithRepositoryName(repositoryName)

	res, err := h.api.Repository.GetRepository(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to get repository %s/%s", projectName, repositoryName)

	return res.Payload
}

type ListRepositoriesResult struct {
	Repositories []*models.Repository `js:"repositories"`
	Total        int64                `js:"total"`
}

func (h *Harbor) ListRepositories(projectName string, args ...goja.Value) ListRepositoriesResult {
        ctx := h.vu.Context()
	h.mustInitialized()

	params := operation.NewListRepositoriesParams()
	params.WithProjectName(projectName)

	if len(args) > 0 {
		rt := h.vu.Runtime()
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(h.vu.Runtime(), err)
		}
	}

	res, err := h.api.Repository.ListRepositories(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to list repositories of %s", projectName)

	return ListRepositoriesResult{
		Repositories: res.Payload,
		Total:        res.XTotalCount,
	}
}
