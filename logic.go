package dynalo

import (
	"github.com/tidwall/gjson"
)

type DynamicLogic struct {
	InputQuery  string
	PolicyQuery string
	Input       interface{}
}

func ExtractInput(rootJson string, dList []DynamicLogic) error {
	if !jsonValid(rootJson) {
		return invalidJson()
	}

	if len(dList) == 0 {
		return nil
	}

	for i := range dList {
		path := dList[i].InputQuery

		value := gjson.Get(rootJson, path)
		if !value.Exists() {
			return invalidPath(rootJson, path)
		}
		dList[i].Input = value.Value()
	}

	return nil
}

func jsonValid(json string) bool {
	return gjson.Valid(json)
}
