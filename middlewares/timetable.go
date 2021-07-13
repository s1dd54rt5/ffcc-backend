package middlewares

import "errors"

func GetSlotPosition(Slot string) ([]int, error) {
	err := errors.New("invalid slot")
	switch Slot {
	case "L1+L2":
		return []int{0, 1}, nil
	case "L3+L4":
		return []int{2, 3}, nil
	case "L5+L6":
		return []int{4, 5}, nil
	case "L7+L8":
		return []int{6, 7}, nil
	case "L9+L10":
		return []int{8, 9}, nil
	case "L11+L12":
		return []int{10, 11}, nil
	case "L13+L14":
		return []int{12, 13}, nil
	case "L15+L16":
		return []int{14, 15}, nil
	case "L17+L18":
		return []int{16, 17}, nil
	case "L19+L20":
		return []int{18, 19}, nil
	case "L21+L22":
		return []int{20, 21}, nil
	case "L23+L24":
		return []int{22, 23}, nil
	case "L25+L26":
		return []int{24, 25}, nil
	case "L27+L28":
		return []int{26, 27}, nil
	case "L29+L30":
		return []int{28, 29}, nil
	case "L31+L32":
		return []int{30, 31}, nil
	case "L33+L34":
		return []int{32, 33}, nil
	case "L35+L36":
		return []int{34, 35}, nil
	case "L37+L38":
		return []int{36, 37}, nil
	case "L39+L40":
		return []int{38, 39}, nil
	case "L41+L42":
		return []int{40, 41}, nil
	case "L43+L44":
		return []int{42, 43}, nil
	case "L45+L46":
		return []int{44, 45}, nil
	case "L47+L48":
		return []int{46, 47}, nil
	case "L49+L50":
		return []int{48, 49}, nil
	case "L51+L52":
		return []int{50, 51}, nil
	case "L53+L54":
		return []int{52, 53}, nil
	case "L55+L56":
		return []int{54, 55}, nil
	case "L57+L58":
		return []int{56, 57}, nil
	case "L59+L60":
		return []int{58, 59}, nil
	default:
		return nil, err
	}
}
