package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/drivers/i2c"
	g "gobot.io/x/gobot/platforms/dexter/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
	"strconv"
	"time"
)

func robotRunLoop(gopigo3 *g.Driver, humiditySensor *aio.GroveTemperatureSensorDriver, lcd *i2c.GroveLcdDriver) {
	for {

		err := lcd.Clear()

		humiditySensorVal, err := humiditySensor.Read()

		if err != nil {
			fmt.Errorf("Error reading sensor %+v", err)
		}

		fmt.Println("Sensor Value is ", humiditySensorVal)

		err = lcd.Write(strconv.Itoa(humiditySensorVal))

		if err != nil {
			fmt.Errorf("Error printing to LCD %+v", err)
		}

		time.Sleep(time.Second)
	}
}

func main() {
	raspiAdaptor := raspi.NewAdaptor()
	gopigo3 := g.NewDriver(raspiAdaptor)

	humiditySensor := aio.NewGroveTemperatureSensorDriver(gopigo3, "AD_2_2")
	lcd := i2c.NewGroveLcdDriver(raspiAdaptor)

	mainRobotFunc := func() {
		robotRunLoop(gopigo3, humiditySensor, lcd)
	}

	robot := gobot.NewRobot("gopigo3HumiditySensor",
		[]gobot.Connection{raspiAdaptor},
		[]gobot.Device{gopigo3, humiditySensor, lcd},
		mainRobotFunc,
	)

	err := robot.Start()

	if err != nil {
		fmt.Errorf("Error starting the robot: %+v", err)
	}
}
