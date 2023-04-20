package internal

import (
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/internal/broker"
	"github.com/zgwit/iot-master/v3/internal/config"
	"github.com/zgwit/iot-master/v3/internal/device"
	"github.com/zgwit/iot-master/v3/internal/product"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/vconn"
	"net"
	"net/url"
)

func Open() error {

	err := config.Load()
	if err != nil {
		return err
	}

	err = log.Open(config.Config.Log)
	if err != nil {
		return err
	}

	//加载数据库
	err = db.Open(config.Config.Database)
	if err != nil {
		return err
	}

	//同步表结构
	err = db.Engine.Sync2(
		new(model.User), new(model.Password), new(model.Role),
		new(model.Broker), new(model.Gateway), new(model.Product),
		new(model.Device), new(model.DeviceArea), new(model.DeviceGroup), new(model.DeviceType),
		new(model.App), new(model.Plugin),
	)
	if err != nil {
		return err
	}
	err = broker.Open(config.Config.Broker)
	if err != nil {
		return err
	}

	if broker.Server != nil {
		err = mqtt.OpenBy(
			func(uri *url.URL, options paho.ClientOptions) (net.Conn, error) {
				c1, c2 := vconn.New()
				//EstablishConnection会读取connect，导致拥堵
				go func() {
					err := broker.Server.EstablishConnection("internal", c1)
					if err != nil {
						log.Error(err)
					}
				}()
				return c2, nil
			})
		if err != nil {
			return err
		}
	} else {
		//MQTT总线
		err = mqtt.Open(config.Config.Mqtt)
		if err != nil {
			return err
		}
	}

	err = product.LoadProducts()
	if err != nil {
		return err
	}

	//err = LoadDevices()
	//if err != nil {
	//	return err
	//}

	//webServe(fmt.Sprintf(":%d", config.Config.Web))
	err = device.SubscribeMaster()
	if err != nil {
		return err
	}

	err = device.SubscribeEvent()
	if err != nil {
		return err
	}

	err = device.SubscribeProperty()
	if err != nil {
		return err
	}

	err = device.SubscribePropertyStrict()
	if err != nil {
		return err
	}

	err = device.SubscribeOnline()
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	//TODO clear gateways devices data
	_ = db.Close()
	broker.Close()
	mqtt.Close()
}
