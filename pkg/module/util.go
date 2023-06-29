package module

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/containerd/containerd/content"
	"github.com/containerd/containerd/content/local"
	"github.com/dop251/goja"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
	"github.com/opencontainers/go-digest"
	"go.k6.io/k6/js/common"
)

func newLocalStore(r *goja.Runtime, ctx context.Context, name string) (string, content.Store) {
	rootPath := filepath.Join(DefaultRootPath, name)

	store, err := local.NewStore(rootPath)
	Check(r, ctx, err)

	return rootPath, store
}

func isDigest(reference string) bool {
	i := strings.Index(reference, ":")
	return i > 0 && i+1 != len(reference)
}

func getDistrubtionRef(projectName, repositoryName, reference string) string {
	if isDigest(reference) {
		return fmt.Sprintf("%s/%s@%s", projectName, repositoryName, reference)
	}

	return fmt.Sprintf("%s/%s:%s", projectName, repositoryName, reference)
}

func writeBlob(rootPath string, data []byte) (digest.Digest, error) {
	dgt := digest.FromBytes(data)

	dir := path.Join(rootPath, "blobs", dgt.Algorithm().String())

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	filename := path.Join(dir, dgt.Hex())
	if err := ioutil.WriteFile(filename, data, 0664); err != nil {
		return "", err
	}

	return dgt, nil
}

func getErrors(i interface{}) *models.Errors {
	if v, ok := i.(interface {
		GetPayload() *models.Errors
	}); ok {
		return v.GetPayload()
	}

	return nil
}

func getErrorMessage(err error) string {
	if errs := getErrors(err); errs != nil && len(errs.Errors) > 0 {
		return errs.Errors[0].Message
	}

	return err.Error()
}

func Check(r *goja.Runtime, ctx context.Context, err error) {
	if err == nil {
		return
	}

	common.Throw(r, errors.New(getErrorMessage(err)))
}

func Checkf(r *goja.Runtime, ctx context.Context, err error, format string, a ...interface{}) {
	if err == nil {
		return
	}

	common.Throw(
		r,
		fmt.Errorf("%s, error: %s", fmt.Sprintf(format, a...), getErrorMessage(err)),
	)
}

func Throwf(r *goja.Runtime, ctx context.Context, format string, a ...interface{}) {
	common.Throw(r, fmt.Errorf(format, a...))
}

func IDFromLocation(r *goja.Runtime, ctx context.Context, loc string) int64 {
	parts := strings.Split(loc, "/")

	id, err := strconv.ParseInt(parts[len(parts)-1], 10, 64)
	Check(r, ctx, err)

	return id
}

func NameFromLocation(ctx context.Context, loc string) string {
	parts := strings.Split(loc, "/")

	return parts[len(parts)-1]
}
