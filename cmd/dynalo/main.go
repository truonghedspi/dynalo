package main

import (
	"context"
	"dynalo"
	"dynalo/policy"
	"fmt"
)

const (
	POLICIES_PATH = "../../policies/"
)

func main() {
	policy, err := policy.NewPolicyEval(POLICIES_PATH)
	if err != nil {
		panic(err)
	}
	dlList := []dynalo.DynamicLogic{
		{
			InputQuery:  "input",
			PolicyQuery: "data.authz.allow",
		},
		{
			InputQuery:  "input",
			PolicyQuery: "data.authz.allow",
		},
	}

	json := `
			{
				"input": {
					"path": ["users"],
					"method": "POST"
				}
			}
	`
	err = dynalo.ExtractInput(json, dlList)
	if err != nil {
		panic(err)
	}
	for _, dlogic := range dlList {
		ctx := context.Background()
		res, err := policy.Eval(
			ctx,
			dlogic,
		)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("evaluate %s with input %v, result %v\n", dlogic.PolicyQuery, dlogic.Input, res)
	}
}
