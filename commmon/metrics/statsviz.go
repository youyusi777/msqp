package metrics

import (
	"github.com/arl/statsviz"
	"net/http"
)

// Server可视化监控
func Serve(addr string) error {
	mux := http.NewServeMux()
	statsviz.Register(mux)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		return err
	}
	return nil
}
