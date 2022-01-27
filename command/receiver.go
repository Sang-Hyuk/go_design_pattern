package main

import "fmt"

type Light interface {
	On()
	Off()
}

type KitchenLight struct {

}

func NewKitchenLight() Light {
	return &KitchenLight{}
}

func (k KitchenLight) On() {
	fmt.Println("거실 불 켜짐.")
}

func (k KitchenLight) Off() {
	fmt.Println("거실 불 꺼짐.")
}

type Stereo interface {
	On()
	Off()
	SetCD()
	SetDvd()
	SetVolume(volume int)
}

type DefaultStereo struct {

}

func NewDefaultStereo() Stereo {
	return &DefaultStereo{}
}

func (d DefaultStereo) On() {
	fmt.Println("기본 스테레오 켜짐.")
}

func (d DefaultStereo) Off() {
	fmt.Println("기본 스테레오 꺼짐.")
}

func (d DefaultStereo) SetCD() {
	fmt.Println("기본 스테레오 CD 삽입.")
}

func (d DefaultStereo) SetDvd() {
	fmt.Println("기본 스테레오 DVD 삽입.")
}

func (d DefaultStereo) SetVolume(volume int) {
	fmt.Printf("기본 스테레오 볼륨 조절 ( %d )\n", volume)
}
