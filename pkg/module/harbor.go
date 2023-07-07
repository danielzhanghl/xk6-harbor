package module

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/dop251/goja"
	rtclient "github.com/go-openapi/runtime/client"
	"github.com/heww/xk6-harbor/pkg/harbor/client"
	"github.com/heww/xk6-harbor/pkg/util"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

var DefaultRootPath = filepath.Join(os.TempDir(), "harbor")

func init() {
	modules.Register("k6/x/harbor", New())

	rootPath := os.Getenv("HARBOR_ROOT")
	if rootPath != "" {
		DefaultRootPath = rootPath
	}

	if err := os.MkdirAll(DefaultRootPath, 0755); err != nil {
		panic(err)
	}
}

type (
	// RootModule is the global module instance that will create module
	// instances for each VU.
	RootModule struct{}

	// ModuleInstance represents an instance of the JS module.
	ModuleInstance struct {
		// vu provides methods for accessing internal k6 objects for a VU
		vu modules.VU
		// comparator is the exported type
		harbor *Harbor
	}
)

// Ensure the interfaces are implemented correctly.
var (
	_ modules.Instance = &ModuleInstance{}
	_ modules.Module = &RootModule{}
)

// New returns a pointer to a new RootModule instance.
func New() *RootModule {
	return &RootModule{}
}

// NewModuleInstance implements the modules.Module interface returning a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ModuleInstance{
		vu: vu,
		harbor: &Harbor{vu: vu},
	}
}

// Exports implements the modules.Instance interface and returns the exported types for the JS module.
func (mi *ModuleInstance) Exports() modules.Exports {
    return modules.Exports{
        Default: mi.harbor,
    }
}

var (
	varTrue = true
)

type Option struct {
	Scheme   string // http or https
	Host     string
	Username string
	Password string
	Insecure bool // Allow insecure server connections when using SSL
}

type Harbor struct {
    vu modules.VU           // provides methods for accessing internal k6 objects

	httpClient  *http.Client
	api         *client.HarborAPI
	option      *Option
	initialized bool
	once        sync.Once
}

func (h *Harbor) Initialize(args ...goja.Value) {
	if h.initialized {
		common.Throw(h.vu.Runtime(), errors.New("harbor module initialized"))
	}
  
	h.once.Do(func() {
		opt := &Option{
			Scheme:   util.GetEnv("HARBOR_SCHEME", "https"),
			Host:     util.GetEnv("HARBOR_HOST", ""),
			Username: util.GetEnv("HARBOR_USERNAME", "admin"),
			Password: util.GetEnv("HARBOR_PASSWORD", "Harbor12345"),
			Insecure: false,
		}

		if len(args) > 0 {
			if args[0] != nil && !goja.IsUndefined(args[0]) && !goja.IsNull(args[0]) {
				rt := h.vu.Runtime()

				err := rt.ExportTo(args[0], opt)
				Checkf(h.vu.Runtime(), err, "failed to parse the option")
			}
		}

		if opt.Host == "" {
			h.vu.Runtime().Interrupt("harbor host is required in initialization")
			return
		}

		opt.Scheme = strings.ToLower(opt.Scheme)
		if opt.Scheme != "https" && opt.Scheme != "http" {
			h.vu.Runtime().Interrupt(fmt.Sprintf("invalid harbor scheme %s", opt.Scheme))
			return
		}

		opt.Host = strings.TrimSuffix(opt.Host, "/")

		rawURL := fmt.Sprintf("%s://%s/%s", opt.Scheme, opt.Host, client.DefaultBasePath)
		u, err := url.Parse(rawURL)
		if err != nil {
			common.Throw(h.vu.Runtime(), err)
		}

		config := client.Config{URL: u}

		if opt.Username != "" && opt.Password != "" {
			config.AuthInfo = rtclient.BasicAuth(opt.Username, opt.Password)
		}

		if opt.Insecure {
			config.Transport = util.NewInsecureTransport()
		} else {
			config.Transport = util.NewDefaultTransport()
		}

		h.api = client.New(config)
		h.option = opt
		h.httpClient = &http.Client{Transport: config.Transport}
		h.initialized = true
	})
}

func (h *Harbor) Free() {
	err := os.RemoveAll(DefaultRootPath)
	if err != nil {
		panic(h.vu.Runtime().NewGoError(err))
	}
}

func (h *Harbor) mustInitialized() {
	if !h.initialized {
		common.Throw(h.vu.Runtime(), errors.New("harbor module not initialized"))
	}
}
