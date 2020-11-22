package uaalpha

//UaStringLess - compare to strings of ukrainian alphabet
func UaStringLess(r1, r2 string) bool {
	return uaRunesLess([]rune(r1), []rune(r2))
}

func uaRunesLess(r1, r2 []rune) bool {
	if len(r1) > 0 && len(r2) > 0 {
		if AlphaIndex[string(r1[0])] == AlphaIndex[string(r2[0])] {
			return uaRunesLess(r1[1:], r2[1:])
		}
		return (AlphaIndex[string(r1[0])] < AlphaIndex[string(r2[0])])
	}
	return len(r1) < len(r2)
}

//Alphabet - Ukrainian letters set
const Alphabet = " АБВГҐДЕЄЖЗИІЇЙКЛМНОПРСТУФХЦЧШЩЬЮЯабвгґдеєжзиіїйклмнопрстуфхцчшщьюя"

//AlphaIndex - map between letter and it's position in ukrainian alphabet
var AlphaIndex = map[string]int{
	" ": 0,
	"А": 1,
	"Б": 2,
	"В": 3,
	"Г": 4,
	"Ґ": 5,
	"Д": 6,
	"Е": 7,
	"Є": 8,
	"Ж": 9,
	"З": 10,
	"И": 11,
	"І": 12,
	"Ї": 13,
	"Й": 14,
	"К": 15,
	"Л": 16,
	"М": 17,
	"Н": 18,
	"О": 19,
	"П": 20,
	"Р": 21,
	"С": 22,
	"Т": 23,
	"У": 24,
	"Ф": 25,
	"Х": 26,
	"Ц": 27,
	"Ч": 28,
	"Ш": 29,
	"Щ": 30,
	"Ь": 31,
	"Ю": 32,
	"Я": 33,
	"а": 34,
	"б": 35,
	"в": 36,
	"г": 37,
	"ґ": 38,
	"д": 39,
	"е": 40,
	"є": 41,
	"ж": 42,
	"з": 43,
	"и": 44,
	"і": 45,
	"ї": 46,
	"й": 47,
	"к": 48,
	"л": 49,
	"м": 50,
	"н": 51,
	"о": 52,
	"п": 53,
	"р": 54,
	"с": 55,
	"т": 56,
	"у": 57,
	"ф": 58,
	"х": 59,
	"ц": 60,
	"ч": 61,
	"ш": 62,
	"щ": 63,
	"ь": 64,
	"ю": 65,
	"я": 66,
}
