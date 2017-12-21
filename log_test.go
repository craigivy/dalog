package dalog_test

import (
	"errors"
	"os"
	"testing"

	"github.com/craigivy/dalog"
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

func TestSubLoggers(t *testing.T) {
	os.Setenv("DALOG_LOGGER", "ZAP")
	log := dalog.NoContext()
	log.Info("just a string without context")
	log.Warn("just a warning without context")
	log.Error(errors.New("just an error without context"))
	log.Debug("just a debug statement without context")

	log2 := log.WithContext(dalog.Context{Key: "hello", Value: "world"})
	log2.Info("we have context now!")

	log3 := log2.WithContext(dalog.Context{Key: "foo", Value: "bar"})
	log3.Info("even more context!")
	log2.Info("but still keeps a separate context in this other logger")
}
