package device

import (
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/pkg/log"
	"github.com/zgwit/iot-master/v4/pkg/mqtt"
	"github.com/zgwit/iot-master/v4/types"
	"strings"
)

func SubscribeOnline() error {
	mqtt.Subscribe[any]("online/+/+", func(topic string, _ *any) {
		topics := strings.Split(topic, "/")
		//pid := topics[1]
		id := topics[2]

		dev, err := Ensure(id)
		if err != nil {
			log.Error(err)
			return
		}
		dev.Online = true
		dev.Values["$online"] = true
	})

	mqtt.Subscribe[any]("offline/+/+", func(topic string, _ *any) {
		topics := strings.Split(topic, "/")
		pid := topics[1]
		id := topics[2]

		dev, err := Ensure(id)
		if err != nil {
			log.Error(err)
			return
		}
		dev.Online = false
		dev.Values["$online"] = false

		//产生日志
		alarm := types.AlarmEx{
			Alarm: types.Alarm{
				ProductId: pid,
				DeviceId:  id,
				Type:      "离线", //TODO 在 产品和设备 中配置
				Title:     "离线",
				Level:     3,
			},
			Product: dev.product.Name,
			Device:  dev.Name,
		}
		_, err = db.Engine.Insert(&alarm.Alarm)
		if err != nil {
			log.Error(err)
			//continue
		}

		//通知
		err = notify(&alarm)
		if err != nil {
			log.Error(err)
			//continue
		}

	})

	return nil
}
