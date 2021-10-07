package allthecities

import (
	"testing"
)

func TestCities(t *testing.T) {
	cities, err := Load()
	if err != nil {
		t.Error(err)
	}

	if len(cities) != 135233 {
		t.Fail()
	}

	for _, city := range cities {
		if len(city.Name) == 0 {
			t.Fail()
		}

		if city.Lon < -180 || city.Lon > 180 {
			t.Fail()
		}

		if city.Lat < -90 || city.Lat > 90 {
			t.Fail()
		}
	}
}

func BenchmarkCities(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Load()
	}
}
