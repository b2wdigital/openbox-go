package app

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
	"github.com/b2wdigital/openbox-go/examples/simple/internal/pkg/model/response"
)

var (
	options *response.Default
)

func init() {
	config.Add("local.message", "gen", "generator output path")
}

func Start(ctx context.Context) {
	log := logrus.FromContext(ctx)

	log.Info("starting application")

	options = new(response.Default)

	err := config.UnmarshalWithPath("local", &options)
	if err != nil {
		log.Error(err)
	}
}