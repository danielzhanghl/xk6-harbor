package module

import (
	"context"

	"github.com/dop251/goja"
	operation "github.com/heww/xk6-harbor/pkg/harbor/client/registry"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
	"go.k6.io/k6/js/common"
)

func (h *Harbor) CreateRegistry(ctx context.Context, r models.Registry) int64 {
	h.mustInitialized(ctx)

	params := operation.NewCreateRegistryParams().WithRegistry(&r)

	res, err := h.api.Registry.CreateRegistry(ctx, params)
	Checkf(h.vu.Runtime(), ctx, err, "failed to create registry %s", params.Registry.Name)

	return IDFromLocation(h.vu.Runtime(), ctx, res.Location)
}

func (h *Harbor) DeleteRegistry(ctx context.Context, id int64) {
	h.mustInitialized(ctx)

	params := operation.NewDeleteRegistryParams().WithID(id)

	_, err := h.api.Registry.DeleteRegistry(ctx, params)
	Checkf(h.vu.Runtime(), ctx, err, "failed to delete registry %d", id)
}

type ListRegistriesResult struct {
	Registries []*models.Registry `js:"registries"`
}

func (h *Harbor) ListRegistries(ctx context.Context, args ...goja.Value) ListRegistriesResult {
	h.mustInitialized(ctx)

	params := operation.NewListRegistriesParams()
	if len(args) > 0 {
		rt := h.vu.Runtime()
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(h.vu.Runtime(), err)
		}
	}

	res, err := h.api.Registry.ListRegistries(ctx, params)
	Checkf(h.vu.Runtime(), ctx, err, "failed to list registries")

	return ListRegistriesResult{
		Registries: res.Payload,
	}
}
