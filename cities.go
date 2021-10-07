package allthecities

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/encoding/protowire"
)

type City struct {
	ID          int64
	Name        string
	AltName     string
	Country     string
	Muni        string
	MuniSub     string
	FeatureCode string
	AdminCode   string
	Population  int
	Lon         float64
	Lat         float64
}

var Source string

func init() {
	Source, _ = filepath.Abs("./cities-v3.1.0.pbf")
}

func Load() ([]City, error) {
	var offset, lastLon, lastLat int64

	f, err := os.Open(Source)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	cities := make([]City, 0)
	for offset < int64(len(data)) {
		msgLength, n := protowire.ConsumeVarint(data[offset:])
		if n < 0 {
			return nil, fmt.Errorf("failed to read message length at offset %d", offset)
		}

		offset += int64(n)
		end := int64(msgLength) + offset

		var city City

		for offset < end {
			tag, _, n := protowire.ConsumeTag(data[offset:])
			offset += int64(n)
			switch tag {
			case 1:
				cityID, n := protowire.ConsumeVarint(data[offset:])
				if n < 0 {
					return nil, fmt.Errorf("failed to read `ID` at offset %d", offset)
				}
				offset += int64(n)
				city.ID = protowire.DecodeZigZag(cityID)
			case 2:
				name, n := protowire.ConsumeString(data[offset:])
				if n < 0 {
					return nil, fmt.Errorf("failed to read `Name` at offset %d", offset)
				}
				offset += int64(n)
				city.Name = name
			case 3:
				country, n := protowire.ConsumeString(data[offset:])
				if n < 0 {
					return nil, fmt.Errorf("failed to read `Country` at offset %d", offset)
				}
				offset += int64(n)
				city.Country = country
			case 4:
				altName, n := protowire.ConsumeString(data[offset:])
				if n < 0 {
					return nil, fmt.Errorf("failed to read `AltName` at offset %d", offset)
				}
				offset += int64(n)
				city.AltName = altName
			case 5:
				muni, n := protowire.ConsumeString(data[offset:])
				if n < 0 {
					return nil, fmt.Errorf("failed to read `Muni` at offset %d", offset)
				}
				offset += int64(n)
				city.Muni = muni
			case 6:
				muniSub, n := protowire.ConsumeString(data[offset:])
				if n < 0 {
					return nil, fmt.Errorf("failed to read `MuniSub` at offset %d", offset)
				}
				offset += int64(n)
				city.MuniSub = muniSub
			case 7:
				featureCode, n := protowire.ConsumeString(data[offset:])
				if n < 0 {
					return nil, fmt.Errorf("failed to read `FeatureCode` at offset %d", offset)
				}
				offset += int64(n)
				city.FeatureCode = featureCode
			case 8:
				adminCode, n := protowire.ConsumeString(data[offset:])
				if n < 0 {
					return nil, fmt.Errorf("failed to read `AdminCode` at offset %d", offset)
				}
				offset += int64(n)
				city.AdminCode = adminCode
			case 9:
				population, n := protowire.ConsumeVarint(data[offset:])
				if n < 0 {
					return nil, fmt.Errorf("failed to read `Population` at offset %d", offset)
				}
				offset += int64(n)
				city.Population = int(population)
			case 10:
				lon, n := protowire.ConsumeVarint(data[offset:])
				if n < 0 {
					return nil, fmt.Errorf("failed to read `Lon` at offset %d", offset)
				}
				offset += int64(n)
				lastLon += protowire.DecodeZigZag(lon)
				city.Lon = float64(lastLon) / 1e5
			case 11:
				lat, n := protowire.ConsumeVarint(data[offset:])
				if n < 0 {
					return nil, fmt.Errorf("failed to read `Lat` at offset %d", offset)
				}
				offset += int64(n)
				lastLat += protowire.DecodeZigZag(lat)
				city.Lat = float64(lastLat) / 1e5
			default:
				return nil, fmt.Errorf("unkown field at offset %d", offset)
			}
		}

		cities = append(cities, city)
	}

	return cities, nil
}
