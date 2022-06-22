// This file not necessary for main application
// Used only generating test data
package database

import (
	"encoding/json"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

var cities = map[int]string{
	1:   "New York City",
	2:   "Los Angeles",
	3:   "Chicago",
	4:   "Houston",
	5:   "Phoenix",
	6:   "San Antonio",
	7:   "Philadelphia",
	8:   "San Diego",
	9:   "Dallas",
	10:  "Austin",
	11:  "San Jose",
	12:  "Fort Worth",
	13:  "Jacksonville",
	14:  "Charlotte",
	15:  "Columbus",
	16:  "Indianapolis",
	17:  "San Francisco",
	18:  "Seattle",
	19:  "Denver",
	20:  "Washington",
	21:  "Boston",
	22:  "El Paso",
	23:  "Nashville",
	24:  "Oklahoma City",
	25:  "Las Vegas",
	26:  "Portland",
	27:  "Detroit",
	28:  "Memphis",
	29:  "Louisville",
	30:  "Milwaukee",
	31:  "Baltimore",
	32:  "Albuquerque",
	33:  "Tucson",
	34:  "Mesa",
	35:  "Fresno",
	36:  "Atlanta",
	37:  "Sacramento",
	38:  "Kansas City",
	39:  "Colorado Springs",
	40:  "Raleigh",
	41:  "Miami",
	42:  "Omaha",
	43:  "Long Beach",
	44:  "Virginia Beach",
	45:  "Oakland",
	46:  "Minneapolis",
	47:  "Tampa",
	48:  "Tulsa",
	49:  "Arlington",
	50:  "Aurora",
	51:  "Wichita",
	52:  "Bakersfield",
	53:  "New Orleans",
	54:  "Cleveland",
	55:  "Henderson",
	56:  "Anaheim",
	57:  "Honolulu",
	58:  "Riverside",
	59:  "Santa Ana",
	60:  "Corpus Christi",
	61:  "Lexington",
	62:  "San Juan",
	63:  "Stockton",
	64:  "St. Paul",
	65:  "Cincinnati",
	66:  "Irvine",
	67:  "Greensboro",
	68:  "Pittsburgh",
	69:  "Lincoln",
	70:  "Durham",
	71:  "Orlando",
	72:  "St. Louis",
	73:  "Chula Vista",
	74:  "Plano",
	75:  "Newark",
	76:  "Anchorage",
	77:  "Fort Wayne",
	78:  "Chandler",
	79:  "Reno",
	80:  "North Las Vegas",
	81:  "Scottsdale",
	82:  "St. Petersburg",
	83:  "Laredo",
	84:  "Gilbert",
	85:  "Toledo",
	86:  "Lubbock",
	87:  "Madison",
	88:  "Glendale",
	89:  "Jersey City",
	90:  "Buffalo",
	91:  "Chesapeake",
	92:  "Winston-Salem",
	93:  "Fremont",
	94:  "Norfolk",
	95:  "Frisco",
	96:  "Paradise",
	97:  "Irving",
	98:  "Garland",
	99:  "Richmond",
	100: "Arlington",
	101: "Boise",
	102: "Spokane",
	103: "Hialeah",
	104: "Moreno Valley",
	105: "Tacoma",
	106: "Port St. Lucie",
	107: "McKinney",
	108: "Fontana",
	109: "Modesto",
	110: "Fayetteville",
	111: "Baton Rouge",
	112: "San Bernardino",
	113: "Santa Clarita",
	114: "Cape Coral",
	115: "Des Moines",
	116: "Tempe",
	117: "Huntsville",
	118: "Oxnard",
	119: "Spring Valley",
	120: "Birmingham",
	121: "Rochester",
	122: "Overland Park",
	123: "Grand Rapids",
	124: "Yonkers",
	125: "Salt Lake City",
	126: "Columbus",
	127: "Augusta",
	128: "Amarillo",
	129: "Tallahassee",
	130: "Ontario",
	131: "Montgomery",
	132: "Little Rock",
	133: "Akron",
	134: "Huntington Beach",
	135: "Grand Prairie",
	136: "Glendale",
	137: "Sioux Falls",
	138: "Sunrise Manor",
	139: "Aurora",
	140: "Vancouver",
	141: "Knoxville",
	142: "Peoria",
	143: "Mobile",
	144: "Chattanooga",
	145: "Worcester",
	146: "Brownsville",
	147: "Fort Lauderdale",
	148: "Newport News",
	149: "Elk Grove",
	150: "Providence",
	151: "Shreveport",
	152: "Salem",
	153: "Pembroke Pines",
	154: "Eugene",
	155: "Rancho Cucamonga",
	156: "Cary",
	157: "Santa Rosa",
	158: "Fort Collins",
	159: "Oceanside",
	160: "Corona",
	161: "Enterprise",
	162: "Garden Grove",
	163: "Springfield",
	164: "Clarksville",
	165: "Murfreesboro",
	166: "Lakewood",
	167: "Bayamon",
	168: "Killeen",
	169: "Alexandria",
	170: "Midland",
	171: "Hayward",
	172: "Hollywood",
	173: "Salinas",
	174: "Lancaster",
	175: "Macon",
	176: "Surprise",
	177: "Kansas City",
	178: "Sunnyvale",
	179: "Palmdale",
	180: "Bellevue",
	181: "Springfield",
	182: "Denton",
	183: "Jackson",
	184: "Escondido",
	185: "Pomona",
	186: "Naperville",
	187: "Roseville",
	188: "Thornton",
	189: "Round Rock",
	190: "Pasadena",
	191: "Joliet",
	192: "Carrollton",
	193: "McAllen",
	194: "Paterson",
	195: "Rockford",
	196: "Waco",
	197: "Bridgeport",
	198: "Miramar",
	199: "Olathe",
	200: "Metairie",
}

func CreateTestOrders() error {
	for cityCode, cityName := range cities {
		orderUUID := uuid.NewString()
		userUUID := uuid.NewString()
		productUUID := uuid.NewString()

		now := time.Now()

		o := Order{OrderUUID: orderUUID,
			CustomerUUID: userUUID,
			Product: Product{
				ProductUUID: productUUID,
				Name:        "Toy car",
				ImageUrl:    "https://example.com/example.jpg",
			},
			CreatedAt: now,
			UpdatedAt: now,
			Address: Address{
				Country:     "United States",
				City:        cityName,
				CityCode:    cityCode,
				AddressLine: cityName + " x street, z apartment number:5/7",
			},
		}

		o.Price = rand.Float32()*100 + 50
		o.Quantity = rand.Intn(20-5) + 5

		orderStatus := rand.Intn(4)
		switch orderStatus {
		case 0:
			o.Status = "New order"
		case 1:
			o.Status = "Preparing"
		case 2:
			o.Status = "Shipped"
		case 3:
			o.Status = "Delivered"
		}
		_, err := o.CreateOrder()
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteTestOrders() error {
	var o Order

	r, e := o.GetOrder()
	if e != nil {
		return e
	}
	var ol []Order

	json.Unmarshal(r, &ol)

	for _, v := range ol {
		_, err := v.DeleteOrder()
		if err != nil {
			return err
		}
	}
	return nil
}
