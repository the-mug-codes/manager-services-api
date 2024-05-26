package people

import (
	"fmt"
	"strconv"

	entity "github.com/kodit-tecnologia/service-manager/entities"
)

func ReadByPhoneNumber(person entity.PersonRepository, fullPhoneNumber string) (readData *entity.Person, err error) {
	countryCode, err := strconv.Atoi(fullPhoneNumber[:2])
	if err != nil {
		return readData, err
	}
	areaCode, err := strconv.Atoi(fullPhoneNumber[2:4])
	if err != nil {
		return readData, err
	}
	if len(fullPhoneNumber[4:]) == 8 {
		fullPhoneNumber = fmt.Sprintf("%s9%s", fullPhoneNumber[0:4], fullPhoneNumber[4:])
	}
	phoneNumber, err := strconv.Atoi(fullPhoneNumber[4:])
	if err != nil {
		return readData, err
	}
	return person.ReadByPhoneNumber(countryCode, areaCode, phoneNumber)
}
