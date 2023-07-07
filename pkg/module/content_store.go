package module

import (
	"errors"
	"os"
	"sync"

	"github.com/containerd/containerd/content"
	"github.com/dop251/goja"
	"github.com/dustin/go-humanize"
	"github.com/heww/xk6-harbor/pkg/util"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	ants "github.com/panjf2000/ants/v2"
)

func (h *Harbor) XContentStore(call goja.ConstructorCall, rt *goja.Runtime) *goja.Object {

        name := "data"

        if len(call.Arguments) >0 {
            name = call.Arguments[0].String()
        }

	store := newContentStore(h.vu.Runtime(), name)

	return rt.ToValue(store).ToObject(rt)
}

func newContentStore(r *goja.Runtime, name string) *ContentStore {
	rootPath, store := newLocalStore(r, name)

	return &ContentStore{Store: store, RootPath: rootPath}
}

type ContentStore struct {
	Store    content.Store
	RootPath string
}

func (s *ContentStore) Generate(humanSize goja.Value) (*ocispec.Descriptor, error) {
	size, err := humanize.ParseBytes(humanSize.String())
	if err != nil {
		return nil, err
	}

	data, err := util.GenerateRandomBytes(int(size))
	if err != nil {
		return nil, err
	}

	dgt, err := writeBlob(s.RootPath, data)
	if err != nil {
		return nil, err
	}

	return &ocispec.Descriptor{
		MediaType: "k6-x-harbor",
		Digest:    dgt,
		Size:      int64(len(data)),
		Annotations: map[string]string{
			ocispec.AnnotationTitle: "raw",
		},
	}, nil
}

func (s *ContentStore) GenerateMany(humanSize goja.Value, count int) ([]*ocispec.Descriptor, error) {
	size, err := humanize.ParseBytes(humanSize.String())
	if err != nil {
		return nil, err
	}
	if size == 0 {
		return nil, errors.New("size must bigger than zero")
	}

	if count <= 0 {
		return nil, errors.New("count must bigger than zero")
	}

	descriptors := make([]*ocispec.Descriptor, count)
	errs := make([]error, count)

	var wg sync.WaitGroup

	poolSize := DefaultPoolSise
	if count < poolSize {
		poolSize = count
	}

	p, _ := ants.NewPoolWithFunc(poolSize, func(i interface{}) {
		defer wg.Done()

		ix := i.(int)
		descriptor, err := s.Generate(humanSize)
		if err != nil {
			errs[ix] = err
		} else {
			descriptors[ix] = descriptor
		}
	})
	defer p.Release()

	for i := 0; i < count; i++ {
		wg.Add(1)
		_ = p.Invoke(i)
	}

	wg.Wait()

	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}

	return descriptors, nil
}

func (s *ContentStore) Free(r *goja.Runtime) {
	err := os.RemoveAll(s.RootPath)
	if err != nil {
		panic(r.NewGoError(err))
	}
}
