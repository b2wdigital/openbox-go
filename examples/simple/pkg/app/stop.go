package app

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/log/logrus"
)

func Stop(ctx context.Context) {
	log := logrus.FromContext(ctx)
	log.Info("stopping application")
}