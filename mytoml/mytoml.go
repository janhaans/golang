package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type user struct {
	Fname string
	Bname string
}

func UserToml(fName string) (user, error) {
	var u user
	_, err := toml.DecodeFile(fName, &u)
	if err != nil {
		return user{}, err
	}
	return u, nil
}

type tokens struct {
	Cases []struct {
		Text   string
		Tokens []string
	}
}

func TokenCasesToml(fName string) (tokens, error) {
	var ts tokens
	_, err := toml.DecodeFile(fName, &ts)
	if err != nil {
		return tokens{}, err
	}
	return ts, nil

}

func main() {
	token_cases_toml := "tokenize_cases.toml"
	ts, err := TokenCasesToml(token_cases_toml)
	if err != nil {
		log.Fatalf("%#v", err)
	}
	fmt.Printf("%v\n", ts)
	for _, c := range ts.Cases {
		fmt.Printf("text = %s, tokens = %s\n", c.Text, c.Tokens)
	}

	user_toml := "user.toml"
	u, err := UserToml(user_toml)
	if err != nil {
		log.Fatalf("%#v", err)
	}
	fmt.Printf("user, first name = %s, family name = %s\n", u.Fname, u.Bname)
}
