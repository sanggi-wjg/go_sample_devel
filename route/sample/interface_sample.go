package sample

import (
	"fmt"
	"log"
)

// SendCompany Interface
type SendCompany interface {
	Send(parcel string) error
}

// KoreaPost Struct
type KoreaPost struct {
	// ...
}

func (k *KoreaPost) Send(parcel string) error {
	fmt.Println("우체국 택배 배송:", parcel)
	return nil
}

// FedexPost Struct
type FedexPost struct {
	// ...
}

func (f *FedexPost) Send(parcel string) error {
	fmt.Println("Fedex 택배 배송:", parcel)
	return nil
}

func SendParcel(parcel string, company SendCompany) {
	err := company.Send(parcel)
	if err != nil {
		log.Fatalln(err)
	}
}

func InterfaceSampleMain() {
	SendParcel("열혈 C++", &KoreaPost{})
	SendParcel("Dynamic Python", &FedexPost{})
}
