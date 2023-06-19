package piscine

// func isAlphabet(r rune) bool {
// 	if isUpper(r) || isLower(r) {
// 		return true
// 	}
// 	return false
// }

// func isUpper(r rune) bool {
// 	if r >= 65 && r <= 90 {
// 		return true
// 	}
// 	return false
// }

// func isLower(r rune) bool {
// 	if r >= 97 && r <= 122 {
// 		return true
// 	}
// 	return false
// }

// func isExceeding(r rune, cases bool) bool {
// 	if cases {
// 		if int(r)+14 > 90 {
// 			return true
// 		}
// 		return false
// 	}
// 	if int(r)+14 > 122 {
// 		return true
// 	}
// 	return false
// }

// func Rot14(s string) string {
// 	res := ""
// 	for _, r := range s {
// 		if isAlphabet(r) {
// 			if isUpper(r) {
// 				if isExceeding(r, true) {
// 					shift := int(r+14-65) % 26
// 					res += string(rune(shift + 65))
// 				} else {
// 					res += string(r + 14)
// 				}
// 			} else if isLower(r) {
// 				if isExceeding(r, false) {
// 					shift := int(r+14-97) % 26
// 					res += string(rune(shift + 97))
// 				} else {
// 					res += string(r + 14)
// 				}
// 			}
// 		} else {
// 			res += string(r)
// 		}
// 	}
// 	return res
// }

func Rot14(s string) string {
	var runes []rune
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			r = (r-'a'+14)%('z'-'a'+1) + 'a'
		}
		if r >= 'A' && r <= 'Z' {
			r = (r-'A'+14)%('Z'-'A'+1) + 'A'
		}
		runes = append(runes, r)
	}
	return string(runes)
}

/*
A - 65
B - 66
C - 67
D - 68
E - 69
F - 70
G - 71
H - 72
I - 73
J - 74
K - 75
L - 76
M - 77

N - 78
O - 79
P - 80
Q - 81
R - 82
S - 83
T - 84
U - 85
V - 86
W - 87
X - 88
Y - 89
Z - 90
*/

/*
a - 97
b - 98
c - 99
d - 100
e - 101
f - 102
g - 103
h - 104
i - 105
j - 106
k - 107
l - 108
m - 109

n - 110
o - 111
p - 112
q - 113
r - 114
s - 115
t - 116
u - 117
v - 118
w - 119
x - 120
y - 121
z - 122
*/
