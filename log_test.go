package dalog_test

import (
	"github.com/craigivy/dalog"
	"os"
	"testing"
)

func Test(t *testing.T) {
	os.Setenv("DALOG", "ZAP")
	dalog.WithContext(dalog.WithID("A123"), dalog.WithHostname()).Infof("%s %s", "hello", "world")
	dalog.WithContext().Infof("%s %s", "hello", "world")
	dalog.WithContext(dalog.WithID("B123"), dalog.WithHostname()).Errorf("%s %s", "hello", "world")

	log := dalog.WithContext(dalog.WithID("123"), dalog.WithHostname())
	log.Infof("ok %s", "doka")
	log.Warnf("take %v", 5)

	os.Setenv("DALOG", "GOLOG")
	dalog.WithContext(dalog.WithID("A123"), dalog.WithHostname()).Infof("%s %s", "hello", "world")
	dalog.WithContext().Infof("%s %s", "hello", "world")
	dalog.WithContext(dalog.WithID("B123"), dalog.WithHostname()).Warnf("%s %s", "hello", "world")

	log = dalog.WithContext(dalog.WithID("123"), dalog.WithHostname())
	log.Infof("ok %s", "doka")
	log.Errorf("take %v", 5)

}
