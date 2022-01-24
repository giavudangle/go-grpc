package sample

import (
	"fmt"

	"github.com/giavudangle/go-grpc/pb"
	"github.com/golang/protobuf/ptypes"
)

/* Function that return a sample Keyboard using randomization methodology*/
func NewKeyBoard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

/* Function that return a sample CPU using randomization methodology*/
func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberOfCores := randomInt(2, 8)
	numberOfThreads := randomInt(numberOfCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 8.0)

	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberOfCores),
		NumberThreads: uint32(numberOfThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	fmt.Println(cpu)
	return cpu
}

/* Function that return a sample GPU using randomization methodology*/
func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 3.0)
	maxGhz := randomFloat64(minGhz, 5.0)

	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}

	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return gpu

}

/* Function that return a sample RAM using randomization methodology*/
func newRAM() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit:  pb.Memory_GIGABYTE,
	}
	return ram
}

/* Function that return a sample SSD using randomization methodology*/
func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return ssd
}

/* Function that return a sample HDD using randomization methodology*/
func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
	return hdd
}

func NewScreen() *pb.Screen {
	screen := &pb.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}
	return screen
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pb.Laptop{
		Brand: brand,
		Name:  name,
		Id:    randomID(),
		Cpu:   NewCPU(),
		Ram:   newRAM(),
		Gpus: []*pb.GPU{
			NewGPU(),
		},
		Storages: []*pb.Storage{
			NewSSD(),
			NewHDD(),
		},
		Screen:   NewScreen(),
		Keyboard: NewKeyBoard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2018, 2021)),
		UpdatedAt:   ptypes.TimestampNow(),
	}
	return laptop
}
