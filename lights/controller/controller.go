package controller

type Controller interface {
	GreenOn()
	YellowOn()
	RedOn()
	YellowRedOn()
	AllOn()
	GreenOff()
	YellowOff()
	RedOff()
	YellowRedOff()
	AllOff()
}