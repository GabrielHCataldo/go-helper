package helper

import (
	"log"
	"testing"
)

func TestMinInt(t *testing.T) {
	for _, tt := range initListTestIntMin() {
		t.Run(tt.name, func(t *testing.T) {
			result := MinInt(tt.value, tt.min)
			log.Println("MinInt:", result)
		})
	}
}

func TestMinInt32(t *testing.T) {
	for _, tt := range initListTestIntMin() {
		t.Run(tt.name, func(t *testing.T) {
			result := MinInt32(int32(tt.value), int32(tt.min))
			log.Println("MinInt32:", result)
		})
	}
}

func TestMinInt64(t *testing.T) {
	for _, tt := range initListTestIntMin() {
		t.Run(tt.name, func(t *testing.T) {
			result := MinInt64(int64(tt.value), int64(tt.min))
			log.Println("MinInt64:", result)
		})
	}
}

func TestMaxInt(t *testing.T) {
	for _, tt := range initListTestIntMax() {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxInt(tt.value, tt.max)
			log.Println("MaxInt:", result)
		})
	}
}

func TestMaxInt32(t *testing.T) {
	for _, tt := range initListTestIntMax() {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxInt32(int32(tt.value), int32(tt.max))
			log.Println("MaxInt32:", result)
		})
	}
}

func TestMaxInt64(t *testing.T) {
	for _, tt := range initListTestIntMax() {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxInt64(int64(tt.value), int64(tt.max))
			log.Println("MaxInt64:", result)
		})
	}
}

func TestRandomNumber(t *testing.T) {
	result := RandomNumber(1, 100)
	log.Println("RandomNumber:", result)
}
