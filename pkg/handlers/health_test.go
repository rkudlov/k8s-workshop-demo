package handlers

import (
	"net/http"
	"testing"

	"github.com/rkudlov/k8s-workshop-demo/pkg/config"
	"github.com/rkudlov/k8s-workshop-demo/pkg/logger"
	"github.com/rkudlov/k8s-workshop-demo/pkg/logger/standard"
	"github.com/rkudlov/k8s-workshop-demo/pkg/router/bitroute"
)

func TestHealth(t *testing.T) {
	h := New(standard.New(&logger.Config{}), new(config.Config))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(h.Health)(bitroute.NewControl(w, r))
	})

	testHandler(t, handler, http.StatusOK, http.StatusText(http.StatusOK))
}
