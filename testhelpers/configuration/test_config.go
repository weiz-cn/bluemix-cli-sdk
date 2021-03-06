package configuration

import (
	"github.com/IBM-Bluemix/bluemix-cli-sdk/bluemix/configuration"
	"github.com/IBM-Bluemix/bluemix-cli-sdk/bluemix/configuration/core_config"
)

type FakePersistor struct{}

func (f *FakePersistor) Save(configuration.DataInterface) error { return nil }
func (f *FakePersistor) Load(configuration.DataInterface) error { return nil }
func (f *FakePersistor) Exists() bool                           { return true }

func NewFakeCoreConfig() core_config.ReadWriter {
	config := core_config.NewCoreConfigFromPersistor(new(FakePersistor), new(FakePersistor), func(err error) { panic(err) })
	config.SetAPIVersion("3")
	return config
}
