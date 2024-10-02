package main
import(
	"errors"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {

	if len(names) != len(phoneNumbers){
		return nil, errors.New("invalid sizes")
	}
	result := make(map[string]user, len(names))

	for i:=0; i<len(names); i++{
		result[names[i]]= user{name: names[i], phoneNumber: phoneNumbers[i]}
	}
	return result, nil
}

type user struct {
	name        string
	phoneNumber int
}

