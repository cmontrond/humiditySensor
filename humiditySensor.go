package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	g "gobot.io/x/gobot/platforms/dexter/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

func robotRunLoop(gopigo3 *g.Driver, humiditySensor *i2c.SHT3xDriver, lcd *i2c.GroveLcdDriver) {
	for {

		err := lcd.Clear()

		temp, _, err := humiditySensor.Sample()

		if err != nil {
			fmt.Errorf("Error reading sensor %+v", err)
		}

		fmt.Println("Sensor Value is ", temp)

		err = lcd.Write(fmt.Sprintf("%f", temp))

		if err != nil {
			fmt.Errorf("Error printing to LCD %+v", err)
		}

		time.Sleep(time.Second)
	}
}

func main() {
	raspiAdaptor := raspi.NewAdaptor()
	gopigo3 := g.NewDriver(raspiAdaptor)

	//humiditySensor := aio.NewAnalogSensorDriver(gopigo3, "AD_1_1")
	lcd := i2c.NewGroveLcdDriver(raspiAdaptor)
	temperatureSensor := i2c.NewSHT3xDriver(raspiAdaptor)

	mainRobotFunc := func() {
		robotRunLoop(gopigo3, temperatureSensor, lcd)
	}

	robot := gobot.NewRobot("gopigo3HumiditySensor",
		[]gobot.Connection{raspiAdaptor},
		[]gobot.Device{gopigo3, temperatureSensor, lcd},
		mainRobotFunc,
	)

	err := robot.Start()

	if err != nil {
		fmt.Errorf("Error starting the robot: %+v", err)
	}
}
