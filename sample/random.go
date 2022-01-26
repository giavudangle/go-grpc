package sample

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/giavudangle/go-grpc/pb"
	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_ASTRAZENICA
	}
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
	return resolution
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD", "Apple")
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"GTX 1060",
			"GTX 1050",
			"GTX 1010",
		)
	}
	return randomStringFromSet("AMD 7200K", "AMD 74200K", "AMD 17200K")
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
	fmt.Print(brand)
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro", "IMac")
	case "Dell":
		return randomStringFromSet("Latitude", "XPS", "Vostro")
	default:
		return randomStringFromSet("Thinkpad X1", "Thinkpad X3", "Thinkpad X5")
	}
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Core i9 Extreme Edition",
			"Core 2 Quad",
			"Core i7 Extreme Edition",
			"Core i7",
		)
	}
	if brand == "AMD" {
		return randomStringFromSet(
			"Ryzen 7 Pro",
			"Ryzen 9 Pro",
			"Ryzen 5 Pro",
		)
	}
	return randomStringFromSet("M1", "M1 PRO", "M1 PROMAX")
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomID() string {
	return uuid.New().String()
}
