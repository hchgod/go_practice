package main

import "fmt"

// 自行车产品
type Bike struct {
	frame string
	tires string
}

func (b *Bike) setFrame(frame string) {
	b.frame = frame
}

func (b *Bike) setTires(tires string) {
	b.tires = tires
}

func (b *Bike) String() string {
	return b.frame + " " + b.tires
}

// 自行车建造者接口
type BikeBuilder interface {
	buildFrame()
	buildTires()
	getResult() *Bike
}

// 山地自行车建造者
type MountainBikeBuilder struct {
	bike *Bike
}

func NewMountainBikeBuilder() *MountainBikeBuilder {
	return &MountainBikeBuilder{
		bike: &Bike{},
	}
}

func (mbb *MountainBikeBuilder) buildFrame() {
	mbb.bike.setFrame("Aluminum Frame")
}

func (mbb *MountainBikeBuilder) buildTires() {
	mbb.bike.setTires("Knobby Tires")
}

func (mbb *MountainBikeBuilder) getResult() *Bike {
	return mbb.bike
}

func (mbb *MountainBikeBuilder) getResult1() *Bike {
	return mbb.bike
}

// 公路自行车建造者
type RoadBikeBuilder struct {
	bike *Bike
}

func NewRoadBikeBuilder() *RoadBikeBuilder {
	return &RoadBikeBuilder{
		bike: &Bike{},
	}
}

func (rbb *RoadBikeBuilder) buildFrame() {
	rbb.bike.setFrame("Carbon Frame")
}

func (rbb *RoadBikeBuilder) buildTires() {
	rbb.bike.setTires("Slim Tires")
}

func (rbb *RoadBikeBuilder) getResult() *Bike {
	return rbb.bike
}

// 自行车Director，负责构建自行车    专门搞个结构体  用来建造
type BikeDirector struct{}

func (bd *BikeDirector) construct(builder BikeBuilder) *Bike {
	builder.buildFrame()
	builder.buildTires()
	return builder.getResult()
}

func main() {
	var N int = 3
	director := &BikeDirector{}
	for i := 0; i < N; i++ {
		var bikeType string = "mountain"
		var builder BikeBuilder
		// 根据输入类别，创建不同类型的具体建造者
		if bikeType == "mountain" {
			builder = NewMountainBikeBuilder()
		} else {
			builder = NewRoadBikeBuilder()
		}
		// Director负责指导生产产品
		bike := director.construct(builder)
		fmt.Println(bike)
	}
}