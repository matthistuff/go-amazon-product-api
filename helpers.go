package amazonproduct
import (
	"fmt"
	"strconv"
)

func idType(id string) string {
	switch len(id) {
	case 10:
		return "ASIN"
	case 14:
		return "CartItemId"
	}
	
	return "OfferListingId"
}

func CreateItemList(items map[string]int) map[string]string {
	params := make(map[string]string)

	i := 1
	for k, v := range items {
		if i < 51 {
			key := fmt.Sprintf("Item.%d.%s", i, idType(k))
			params[key] = string(k)

			key = fmt.Sprintf("Item.%d.Quantity", i)
			params[key] = strconv.Itoa(v)

			i++
		} else {
			break
		}
	}

	return params
}