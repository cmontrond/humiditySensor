package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/drivers/i2c"
	g "gobot.io/x/gobot/platforms/dexter/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

func robotRunLoop(gopigo3 *g.Driver, humiditySensor *aio.AnalogSensorDriver, lcd *i2c.GroveLcdDriver) {
	for {

		ultrasonicSensorVal, ultrasonicSensorErr := humiditySensor.Read()

		if ultrasonicSensorErr != nil {
			fmt.Errorf("Error reading sensor %+v", ultrasonicSensorErr)
		}

		fmt.Println("Sensor Value is ", ultrasonicSensorVal)
		//lcdPrintErr := lcd.Write(string(rune(ultrasonicSensorVal)))
		lcdPrintErr := lcd.Write("test")

		if lcdPrintErr != nil {
			fmt.Errorf("Error printing to LCD %+v", lcdPrintErr)
		}

		time.Sleep(time.Second)
	}
}

func main() {
	raspiAdaptor := raspi.NewAdaptor()
	gopigo3 := g.NewDriver(raspiAdaptor)

	humiditySensor := aio.NewAnalogSensorDriver(gopigo3, "AD_1_1")
	lcd := i2c.NewGroveLcdDriver(raspiAdaptor)

	mainRobotFunc := func() {
		robotRunLoop(gopigo3, humiditySensor, lcd)
	}

	robot := gobot.NewRobot("gopigo3HumiditySensor",
		[]gobot.Connection{raspiAdaptor},
		[]gobot.Device{gopigo3, humiditySensor},
		mainRobotFunc,
	)

	robot.Start()
}
