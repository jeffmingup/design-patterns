package main

import "fmt"

type Subject interface {
	registerObserver(observer)
	removeObserver(observer)
	notifyObserver()
}
type observer interface {
	update(Subject)
}

type WeatherData struct {
	temperature float32 //温度
	humidity    float32 //湿度
	pressure    float32 //压强
	observers   []observer
}

func (data *WeatherData) registerObserver(ob observer) {
	data.observers = append(data.observers, ob)
}
func (data *WeatherData) removeObserver(ob observer) {
	for k, v := range data.observers {
		if v == ob {
			data.observers = append(data.observers[:k], data.observers[k+1:]...)
		}
	}
}
func (data *WeatherData) notifyObserver() {
	for _, ob := range data.observers {
		ob.update(data)
	}
}

//修改数据并通知订阅者
func (data *WeatherData) Setmeasurements(temperature, humidity, pressure float32) {
	data.temperature = temperature
	data.humidity = humidity
	data.pressure = pressure
	data.notifyObserver()
}

//订阅者 1
type currentConditionsDisplay struct {
	ID          int     //订阅者ID
	temperature float32 //温度
	humidity    float32 //湿度
}

func newObserver(id int) *currentConditionsDisplay {
	return &currentConditionsDisplay{
		ID: id,
	}
}

func (cur *currentConditionsDisplay) addSubscription(sub Subject) {
	sub.registerObserver(cur)
}
func (cur *currentConditionsDisplay) removeSubscription(sub Subject) {
	sub.removeObserver(cur)
}

func (cur *currentConditionsDisplay) update(sub Subject) {
	if sub, ok := sub.(*WeatherData); ok {
		cur.temperature = sub.temperature
		cur.humidity = sub.humidity
	}
	fmt.Printf("ID:%v ;current conditions:%v F degrees and %v humidity\n", cur.ID, cur.temperature, cur.humidity)
}

func main() {
	data := new(WeatherData)

	cur_1 := newObserver(1)

	cur_2 := newObserver(2)
	cur_2.addSubscription(data)
	cur_1.addSubscription(data)

	data.Setmeasurements(1, 2, 3)

	cur_1.removeSubscription(data)

	data.Setmeasurements(1, 2, 3)
}
