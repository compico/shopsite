package dataworker

import "github.com/compico/shopsite/internal/config"

func InitData(conf config.Config, data interface{}) *Data {
	x := new(Data)
	x.Config = conf
	x.Data = data
	return x
}

func (d *Data) GetDataAndChangeTitle(title string) Data {
	x := *d
	x.Config.SetTitle(title + " - " + x.Config.GetSiteName())
	return x
}

func (d *Data) AddData(data interface{}) {
	d.Data = data
}
