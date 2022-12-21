package crypting

import (
	"strconv"
	"strings"
	"task_rest/middleware"
)

type symbol struct {
	st    string
	count int
}

// Encrypt function for crypting string
// input "AAABBACACACAC"
// output "3A2B4(AC)"
func Encrypt(str string) (result string) {
	middleware.Logs.Debug().Msgf("[crypting] Encrypt started")
	middleware.Logs.Debug().Str("value str=", str).Msgf("[crypting]")
	if str == "" {
		return ""
	}
	numRep, num := []symbol{{st: string(str[0]), count: 0}}, 0
	for _, char := range str {
		if numRep[num].st == string(char) {
			numRep[num].count++
		} else {
			numRep = append(numRep, symbol{st: string(char), count: 1})
			num++
		}
	}
	middleware.Logs.Debug().Interface("disassemble str=", numRep).Msgf("[crypting] Encrypt: second loop for build str")
	num = 2
	checkExit := false
	for !checkExit {
		checkExit = true
		for i := 0; i < len(numRep)-num; i++ {
			checkEqual := true
			for k := i; k < num+i; k++ {
				if numRep[k] != numRep[k+num] {
					checkEqual = false
					break
				}
			}
			if checkEqual {
				checkExit = false
				newStr := ""
				var lastRep []symbol
				for k := i; k < num+i; k++ {
					if numRep[k].count > 1 {
						newStr += strconv.Itoa(numRep[k].count)
					}
					if len(numRep[k].st) > 1 {
						newStr += "(" + numRep[k].st + ")"
					} else {
						newStr += numRep[k].st
					}
					lastRep = append(lastRep, numRep[k])
				}
				numRep[i].st, numRep[i].count = newStr, 1
				numRep = append(numRep[:i+1], numRep[i+num:]...)
				for checkEqual {
					if i+num >= len(numRep) {
						break
					}
					for k := i + 1; k < i+num+1; k++ {
						if numRep[k] != lastRep[k-i-1] {
							checkEqual = false
							break
						}
					}
					if checkEqual {
						numRep[i].count++
						numRep = append(numRep[:i+1], numRep[i+num+1:]...)
					}
				}
			}
		}
		if !checkExit && num < len(numRep) {
			num++
		}
	}
	middleware.Logs.Debug().Msgf("[crypting] Encrypt: last loop - build result")
	for i := 0; i < len(numRep); i++ {
		if numRep[i].count > 1 {
			result += strconv.Itoa(numRep[i].count)
		}
		if len(numRep[i].st) > 1 {
			result += "(" + numRep[i].st + ")"
		} else {
			result += numRep[i].st
		}
	}
	middleware.Logs.Debug().Str("str=", result).Msgf("[crypting] result")
	middleware.Logs.Debug().Msgf("[crypting] Encrypt finished")
	return
}

// Decrypt function for decrypting string
// input "2A3B4(AC)"
// output "AABBBACACACAC"
func Decrypt(str string) string {
	middleware.Logs.Debug().Msgf("[crypting] Decrypt started")
	var numRep []symbol
	var newStr, numChar, long, add = "", "", 0, false
	middleware.Logs.Debug().Str("str=", str).Msgf("[crypting] start loop for disassemble str")
	for _, char := range str {
		if strings.ContainsRune("(", char) {
			if long > 0 {
				newStr += string(char)
			}
			long += 1
		} else if strings.ContainsRune(")", char) {
			long -= 1
			if long > 0 {
				newStr += string(char)
			} else {
				newStr = Decrypt(newStr)
				add = true
			}
		} else if long > 0 {
			newStr += string(char)
		} else if strings.ContainsRune("0123456789", char) {
			numChar += string(char)
		} else {
			newStr = string(char)
			add = true
		}
		if add {
			num, err := strconv.Atoi(numChar)
			if err != nil {
				num = 1
			}
			numRep = append(numRep, symbol{st: string(newStr), count: num})
			newStr, numChar, add = "", "", false
		}
	}
	middleware.Logs.Debug().Msgf("[crypting] Decrypt: last loop - build result")
	for _, val := range numRep {
		newStr += strings.Repeat(val.st, val.count)
	}
	middleware.Logs.Debug().Str("str=", newStr).Msgf("[crypting] result")
	middleware.Logs.Debug().Msgf("[crypting] Decrypt finished")
	return newStr
}
