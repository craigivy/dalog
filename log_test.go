package dalog_test

import (
	"github.com/craigivy/dalog"
	"os"
	"testing"
)

func Test(t *testing.T) {
	os.Setenv("DALOG_LOGGER", "ZAP")
	// os.Setenv("DALOG_DEBUG", "TRUE") // once set and a message is logged this can not be changed (for now)
	dalog.WithContext(dalog.WithID("A123"), dalog.WithHostname()).Infof("%s %s", "hello", "world")
	dalog.WithContext().Infof("%s %s", "hello", "world")
	dalog.WithContext(dalog.WithID("B123"), dalog.WithHostname()).Errorf("%s %s", "hello", "world")

	log := dalog.WithContext(dalog.WithID("123"), dalog.WithHostname())
	log.Infof("ok %s", "doka")
	log.Warnf("take %v", 5)
	log.Debugf("debug me %s", "in json")

	os.Setenv("DALOG_LOGGER", "GOLOG")
	os.Setenv("DALOG_DEBUG", "TRUE")
	dalog.WithContext(dalog.WithID("A123"), dalog.WithHostname()).Infof("%s %s", "hello", "world")
	dalog.WithContext().Infof("%s %s", "hello", "world")
	dalog.WithContext(dalog.WithID("B123"), dalog.WithHostname()).Warnf("%s %s", "hello", "world")

	log = dalog.WithContext(dalog.WithID("123"), dalog.WithHostname())
	log.Infof("ok %s", "doka")
	log.Errorf("take %v", 5)
	log.Debugf("debug me %s", "now!")

	os.Setenv("DALOG_DEBUG", "FALSE")
	log = dalog.WithContext()
	log.Debugf("DEBUG OFF! %s", "NOT LOGGED")

}
