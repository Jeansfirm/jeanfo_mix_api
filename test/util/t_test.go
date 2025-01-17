package util_test

import (
	"jeanfo_mix/util/log_util"
	"testing"
)

func TestT(t *testing.T) {
	log_util.Debug("any %s %d", "jeanfo", 3)
	log_util.Info("some")
	log_util.Warn("haha")
	log_util.Error("hehe %s", "kk")
	log_util.Error("%v coming %f", 3, 4.4)
}
