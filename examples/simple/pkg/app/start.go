package app

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
	"github.com/b2wdigital/openbox-go/examples/simple/internal/pkg/options"
)

func init() {
	config.Add("local.message", "hello world!!!", "generator output path")
}

func Start(ctx context.Context) {
	log := logrus.FromContext(ctx)

	log.Info("starting application")

	o := new(options.Options)

	err := config.UnmarshalWithPath("local", o)
	if err != nil {
		log.Error(err)
	}

	log.Infof("with configurable message %s", o.Message)
}