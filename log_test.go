package dalog_test

import (
	"os"
	"testing"

	goerr "errors"
	"github.com/craigivy/dalog"
	"github.com/pkg/errors"
)

func Test(t *testing.T) {
	//os.Setenv("DALOG_LOGGER", "GO")
	//os.Setenv("DALOG_STACK", "TRUE")

	ens := goerr.New("add stack")
	ens = errors.WithStack(ens)
	dalog.NoContext().Error(ens)

	os.Setenv("DALOG_LOGGER", "ZAP")
	// os.Setenv("DALOG_DEBUG", "TRUE") // once set and a message is logged this can not be changed (for now)
	dalog.WithContext(dalog.WithID("A123"), dalog.WithHostname()).Infof("%s %s", "hello", "world")
	dalog.WithContext().Infof("%s %s", "hello", "world")
	dalog.WithContext(dalog.WithID("B123"), dalog.WithHostname()).Error(errors.Errorf("%s %s", "hello", "world"))

	log := dalog.WithContext(dalog.WithID("123"), dalog.WithHostname())
	log.Infof("ok %s", "doka")
	log.Warnf("take %v", 5)
	log.Debugf("debug me %s", "in json")
	log.Stackf("debug stack")

	os.Setenv("DALOG_LOGGER", "GO")
	os.Setenv("DALOG_DEBUG", "TRUE")
	dalog.WithContext(dalog.WithID("A123"), dalog.WithHostname()).Infof("%s %s", "hello", "world")
	dalog.WithContext().Infof("%s %s", "hello", "world")
	dalog.WithContext(dalog.WithID("B123"), dalog.WithHostname()).Warnf("%s %s", "hello", "world")

	log = dalog.WithContext(dalog.WithID("123"), dalog.WithHostname())
	log.Infof("ok %s", "doka")
	log.Error(errors.Errorf("take %v", 5))
	log.Debugf("debug me %s", "now!")
	log.Stackf("debug stack")

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

	log2 := log.WithContext(dalog.WithKey("hello", "world"))
	log2.Info("we have context now!")

	log3 := log2.WithContext(dalog.WithKey("foo", "bar"))
	log3.Info("even more context!")
	log2.Info("but still keeps a separate context in this other logger")
}

func TestStack(t *testing.T) {
	os.Setenv("DALOG_LOGGER", "ZAP")
	os.Setenv("DALOG_STACK", "TRUE")

	e := errors.New("This is an error using pkg error")
	dalog.NoContext().Error(e)

	ens := goerr.New("no stack")
	dalog.NoContext().Error(ens)

	// add a stack to the exiting error
	esa := errors.WithStack(ens)
	dalog.NoContext().Error(esa)

	os.Setenv("DALOG_LOGGER", "GO")
	dalog.NoContext().Error(e)
	dalog.NoContext().Error(ens)
	dalog.NoContext().Error(esa)

}

func TestDebugContext(t *testing.T) {
	os.Setenv("DALOG_LOGGER", "ZAP")
	os.Setenv("DALOG_STACK", "TRUE")
	log := dalog.WithContext(dalog.WithDebugContext("component1"))
	log.Debugf("debug message with debug context")
	log.Debug("debug message with debug context")
	log.Warn("should not inlude debug context")

	os.Setenv("DALOG_LOGGER", "GO")

	log = dalog.WithContext(dalog.WithDebugContext("component1"))
	log.Debugf("debug message with debug context")
	log.Debug("debug message with debug context")
	log.Warn("should not inlude debug context")

	slog := log.WithContext(dalog.WithDebugContext("component2"))
	slog.Debugf("debug message with component2 debug context")

	log.Debugf("debug message still has component1 debug context")

}
