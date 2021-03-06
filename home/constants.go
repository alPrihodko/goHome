package home

const (
	emailFrom         = "alPrihodko@gmail.com"
	emailTo           = "alPrihodko@gmail.com"
	googleAccountName = "alPrihodko@gmail.com"
	googleAPIKey      = "ahvpiuembqkvszpi"

	cloudURL         = "http://data.sparkfun.com/input/"
	phantPrivateTemp = "jkybBZw5j9F0xxmwWar1"
	phantPublicTemp  = "rozL4pN5jMC466MKl8AE"

	iotURL             = "http://myhome.iot.touchpanelcontrol.com/v1"
	iotCloudAPIKey     = "35d1d4d0-f80f-11e5-83f9-7d8483cd58bb"
	iotCloudHomeDevice = "pi"

	iotCloudInternalTempVar = "tempInternal"
	internalSensorID        = "28-0315a14596ff"
	reverseSensorID         = "28-0115a32c38ff"
	entryRoomSensorID       = "28-0115a3752aff"
	heaterSensorID          = "28-0115a32295ff"
	_waterBoilerSensorID    = "28-0415a1cbb3ff"
	_outsideSensorID        = "28-0315a1a9cfff"
	_recuperatorSensorID    = "28-0315a1d318ff"

	iotDevURL             = "http://al.dev-iot.touchpanelcontrol.com/v1"
	iotDevCloudAPIKey     = "ea85d840-2a96-11e6-b9d4-e7f016c54c9f"
	iotDevCloudHomeDevice = "r01"
	iotCloudEntry         = "entry"
	iotCloudHin           = "hin"
	iotCloudHout          = "hout"
	iotCloudKitchen       = "kitchen"
	iotCloudHeater        = "heater"
	iotCloudPump          = "hpump"
)

/*
HeaterOnThreshold - getting cold and need heater
*/
//const HeaterOnThreshold = 16.0

/*
DropboxClientid - client id
*/
const DropboxClientid = "mhde2yglerq0gcz"

/*
DropboxClientsecret - secret
*/
const DropboxClientsecret = "xz44arsbpb8yrtu"

/*
DropboxToken - token
*/
const DropboxToken = "wFGCDkkgII0AAAAAAAAHFuWXLOZ5uYabVctDUrRQk2XLEzEYihg2XjgyHlGVBIrE"

/*
HeaterPumpRunThreshold Run heater pump, stop heater pumb if temp 5 degree less
*/
const HeaterPumpRunThreshold = 30

/*
CommandTogglePumpr1 command to toggle
*/
const CommandTogglePumpr1 = "togglePumpr1"

/*
CommandTogglePumpr2  command to toggle
*/
const CommandTogglePumpr2 = "togglePumpr2"

/*
CommandToggleHeater1  command to toggle
*/
const CommandToggleHeater1 = "toggleHeaterr1"

/*
CommandToggleHeater2  command to toggle
*/
const CommandToggleHeater2 = "toggleHeaterr2"

/*
CommandOnPumpr1 command to toggle
*/
const CommandOnPumpr1 = "onPumpr1"

/*
CommandOnPumpr2  command to toggle
*/
const CommandOnPumpr2 = "onPumpr2"

/*
CommandOnHeater1  command to toggle
*/
const CommandOnHeater1 = "onHeaterr1"

/*
CommandOnHeater2  command to toggle
*/
const CommandOnHeater2 = "onHeaterr2"

/*
CommandOffPumpr1 command to toggle
*/
const CommandOffPumpr1 = "offPumpr1"

/*
CommandOffPumpr2  command to toggle
*/
const CommandOffPumpr2 = "offPumpr2"

/*
CommandOffHeater1  command to toggle
*/
const CommandOffHeater1 = "offHeaterr1"

/*
CommandOffHeater2  command to toggle
*/
const CommandOffHeater2 = "offHeaterr2"

/*Recipients is SMS recipients*/
var Recipients = []string{"380939760324"}

/*
ElectroOnFrom - timezone for electricity metter - from
*/
const ElectroOnFrom = 23

/*
ElectroOnTo - timezone for electricity metter - from
*/
const ElectroOnTo = 7
