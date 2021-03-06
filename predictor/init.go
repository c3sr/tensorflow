package predictor

// #cgo LDFLAGS: -ltensorflow
// #cgo CFLAGS: -I/opt/tensorflow/include
// #include "tensorflow/c/c_api.h"
import "C"

import (
	"os"

	"github.com/c3sr/config"
	"github.com/c3sr/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

var (
	log                  *logrus.Entry
	disableOptimizations bool
)

func init() {
	disableOptimizations = cast.ToBool(os.Getenv("CARML_TF_DISABLE_OPTIMIZATION")) == true
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "tensorflow/predictor")
	})
}
