package tensorflow

import (
	"github.com/c3sr/config"
	"github.com/c3sr/logger"
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Entry = logger.New().WithField("pkg", "tensorflow")
)

func init() {
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "tensorflow")
	})
}
