package plugins

import (
	"net/http"

	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"github.com/apache/apisix-go-plugin-runner/pkg/log"
	"github.com/apache/apisix-go-plugin-runner/pkg/plugin"
)

func init() {
	err := plugin.RegisterPlugin(&Count{})
	if err != nil {
		log.Fatalf("failed to register plugin count: %s", err)
	}
}

type Count struct {
	plugin.DefaultPlugin
}

func (c *Count) Name() string {
	return "count"
}

func (c *Count) ParseConf(in []byte) (conf interface{}, err error) {
	return Count{}, nil
}

func (c *Count) RequestFilter(conf interface{}, w http.ResponseWriter, r pkgHTTP.Request) {
	r.Header().Set("x-rainbow-req-count", "1000")
}

func (c *Count) ResponseFilter(conf interface{}, w pkgHTTP.Response) {
	w.Header().Set("x-rainbow-resp-count", "2000")
}
