package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	if order != "burger" && order != "chips" && order != "nuggets" {
		return 404
	}
	f := food{
		preptime: 0,
	}
	switch order {
	case "burger":
		f.preptime = 15
	case "chips":
		f.preptime = 10
	case "nuggets":
		f.preptime = 12
	}
	return f.preptime
}
