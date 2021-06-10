package generate_items

import (
	crand "crypto/rand"
	"encoding/binary"
	"log"
	"math/rand"
)

// GenerateObjFloat32 - генерация значений float32.
type GenerateObjFloat32 struct {
	Value	float32
	Weight	int
}

// GenerateObjInt - генерация значений int.
type GenerateObjInt struct {
	Value  int
	Weight int
}

// WeightSumObjFloat32 - подсчёт суммы веса обьектов float32.
func WeightSumObjFloat32(objects []GenerateObjFloat32) int {
	var weight int

	for _, value := range objects {
		weight += value.Weight
	}

	return weight
}

// WeightSumObjInt - подсчёт суммы веса обьектов imt.
func WeightSumObjInt(objects []GenerateObjInt) int {
	var weight int

	for _, value := range objects {
		weight += value.Weight
	}

	return weight
}

// GenerateValueObjFloat32 - генерация значения обьекта float32.
func GenerateValueObjFloat32(weight int, objects []GenerateObjFloat32) float32 {
	var src CryptoSource
	crnd := rand.New(src)

	var count int
	rnd := crnd.Intn(weight) + 1

	for _, generateObjFloat32 := range objects {
		count += generateObjFloat32.Weight

		if count >= rnd {
			return generateObjFloat32.Value
		}
	}

	return 0
}

// GenerateValueObjInt - генерация значения обьекта int.
func GenerateValueObjInt(weight int, objects []GenerateObjInt) int {
	var src CryptoSource
	crnd := rand.New(src)

	var count int
	rnd := crnd.Intn(weight) + 1

	for _, generateObjInt := range objects {
		count += generateObjInt.Weight

		if count >= rnd {
			return generateObjInt.Value
		}
	}

	return 0
}

type CryptoSource struct{}

func (s CryptoSource) Seed(seed int64) {}

func (s CryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s CryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
