package day4

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// PassportData => stores fields for passport
type PassportData struct {
	countryID      int
	birthYear      int
	expirationYear int
	issueYear      int
	height         string
	hairColor      string
	eyeColor       string
	passportID     string
}

// CountValidPassports => Counts number of passports which have all required fields present
// NOTE: assumes validation has already been satisfied)
func CountValidPassports(passports []PassportData) int {
	count := 0
	for _, passport := range passports {
		//fmt.Printf("%+v\n", passport)
		if passport.birthYear != 0 && passport.expirationYear != 0 && passport.issueYear != 0 && passport.height != "" && passport.hairColor != "" && passport.passportID != "" && passport.eyeColor != "" {
			count++
		} else {
			fmt.Println("Passport not valid")
		}
	}
	return count
}

// ParsePassportData => parses passport data from file and applies validation
// (TODO : separate validation into separate function)
func ParsePassportData(path string) []PassportData {

	filename := fmt.Sprintf("%s/day4/day4.txt", path)
	//filename := fmt.Sprintf("%s/day4/test.txt", path)
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(err)
		os.Exit(1)
	}
	var passports []PassportData
	passportContents := strings.Split(string(contents), "\n\n")
	re := regexp.MustCompile(`(\w{3}):[\w#]+`)
	pidValidator := regexp.MustCompile(`(^\d{9}$)`)
	hairColorValidator := regexp.MustCompile(`^#(\w{6})$`)
	allowedEyeColors := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	for _, content := range passportContents {
		fields := re.FindAllString(content, -1)
		passportData := PassportData{}
		for _, field := range fields {
			data := strings.Split(field, ":")
			key, value := data[0], data[1]
			switch key {
			case "cid":
				passportData.countryID, _ = strconv.Atoi(value)
			case "byr":
				birthYear, _ := strconv.Atoi(value)
				if birthYear >= 1920 && birthYear <= 2002 {
					passportData.birthYear = birthYear
				}
			case "iyr":
				issueYear, _ := strconv.Atoi(value)
				if issueYear >= 2010 && issueYear <= 2020 {
					passportData.issueYear = issueYear
				}
			case "hgt":
				if len(value) >= 4 {
					val, unit := value[:len(value)-2], value[len(value)-2:]
					quantity, _ := strconv.Atoi(val)
					if unit == "cm" && quantity >= 150 && quantity <= 193 {
						passportData.height = value
					}
					if unit == "in" && quantity >= 59 && quantity <= 76 {
						passportData.height = value
					}
				}
			case "hcl":
				value = hairColorValidator.FindString(value)
				if value != "" {
					passportData.hairColor = value
				}
			case "ecl":
				if allowedEyeColors[value] {
					passportData.eyeColor = value
				}
			case "pid":
				value = pidValidator.FindString(value)
				if value != "" {
					passportData.passportID = value
				}
			case "eyr":
				expirationYear, _ := strconv.Atoi(value)
				if expirationYear >= 2020 && expirationYear <= 2030 {
					passportData.expirationYear = expirationYear
				}
			default:
				continue
			}
		}
		passports = append(passports, passportData)
	}
	return passports
}
