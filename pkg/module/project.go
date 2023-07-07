package module

import (
	"strings"

	"github.com/dop251/goja"
	operation "github.com/heww/xk6-harbor/pkg/harbor/client/project"
	"github.com/heww/xk6-harbor/pkg/harbor/client/repository"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
	"github.com/heww/xk6-harbor/pkg/util"
	log "github.com/sirupsen/logrus"
	"go.k6.io/k6/js/common"
)

func (h *Harbor) CreateProject(body goja.Value) string {
        ctx := h.vu.Context()
	h.mustInitialized()

	rt := h.vu.Runtime()
	var project models.ProjectReq
	err := rt.ExportTo(body, &project)
	Check(h.vu.Runtime(), err)

	params := operation.NewCreateProjectParams()
	params.WithProject(&project).WithXResourceNameInLocation(&varTrue)

	res, err := h.api.Project.CreateProject(ctx, params)

	Checkf(h.vu.Runtime(), err, "failed to create project %s", params.Project.ProjectName)

	return NameFromLocation(res.Location)
}

//func (h *Harbor) GetProject(projectName string) *models.Project {
func (h *Harbor) GetProject(projectName string) *models.Project {
        ctx := h.vu.Context()
	h.mustInitialized()

	params := operation.NewGetProjectParams()
	params.WithProjectNameOrID(projectName)

	res, err := h.api.Project.GetProject(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to get project %s", projectName)

	return res.Payload
}

func (h *Harbor) DeleteProject(projectName string, args ...goja.Value) {
        ctx := h.vu.Context()
	h.mustInitialized()

	var force bool
	if len(args) > 0 {
		force = args[0].ToBoolean()
	}

	if force {
		pageSize := 20

		params := repository.NewListRepositoriesParams().WithProjectName(projectName)
		params.Page = util.Int64(1)
		params.PageSize = util.Int64(int64(pageSize))

		for {
			resp, err := h.api.Repository.ListRepositories(ctx, params)
			Checkf(h.vu.Runtime(), err, "failed to list repositories for page %d", *params.Page)

			for _, repo := range resp.Payload {
				repoName := strings.TrimPrefix(repo.Name, projectName+"/")
				h.DeleteRepository(projectName, repoName)
			}

			if len(resp.Payload) < pageSize {
				break
			}
		}
	}

	params := operation.NewDeleteProjectParams()
	params.WithProjectNameOrID(projectName).WithXIsResourceName(&varTrue)

	_, err := h.api.Project.DeleteProject(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to delete project %s", projectName)
}

func (h *Harbor) DeleteAllProjects(excludeProjects []string) {
	h.mustInitialized()

	rt := h.vu.Runtime()

	m := make(map[string]bool, len(excludeProjects))
	for _, projectName := range excludeProjects {
		m[projectName] = true
	}

	page := 1
	pageSize := 10
	for {
		arg := map[string]interface{}{"page": page, "page_size": pageSize}

		result := h.ListProjects(rt.ToValue(arg))

		projects := result.Projects

		deleted := 0
		for _, project := range projects {
			if m[project.Name] {
				continue
			}

			func() {
				defer func() {
					if r := recover(); r != nil {
						log.Warnf("%v", r)
					}
				}()

				h.DeleteProject(project.Name, rt.ToValue(project.RepoCount+project.ChartCount > 0))
				deleted++
			}()
		}

		if len(projects) == 0 || len(projects) < pageSize {
			break
		}

		if deleted == 0 {
			page++
		}
	}
}

type ListProjectsResult struct {
	Projects []*models.Project `js:"projects"`
	Total    int64             `js:"total"`
}

func (h *Harbor) ListProjects(args ...goja.Value) ListProjectsResult {
        ctx := h.vu.Context()
	h.mustInitialized()

	params := operation.NewListProjectsParams()

	if len(args) > 0 {
		rt := h.vu.Runtime()
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(h.vu.Runtime(), err)
		}
	}

	res, err := h.api.Project.ListProjects(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to list projects")

	return ListProjectsResult{
		Projects: res.Payload,
		Total:    res.XTotalCount,
	}
}

type ListAuditLogsOfProjectResult struct {
	Logs  []*models.AuditLog `js:"logs"`
	Total int64              `js:"total"`
}

func (h *Harbor) ListAuditLogsOfProject(projectName string, args ...goja.Value) ListAuditLogsOfProjectResult {
        ctx := h.vu.Context()
	h.mustInitialized()

	params := operation.NewGetLogsParams().WithProjectName(projectName)

	if len(args) > 0 {
		rt := h.vu.Runtime()
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(h.vu.Runtime(), err)
		}
	}

	res, err := h.api.Project.GetLogs(ctx, params)
	Checkf(h.vu.Runtime(), err, "failed to list audit logs of project %s", projectName)

	return ListAuditLogsOfProjectResult{
		Logs:  res.Payload,
		Total: res.XTotalCount,
	}
}
