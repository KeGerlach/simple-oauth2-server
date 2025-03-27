package environment

import (
	"fmt"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type Environment struct {
	PORT			string

	SECRET_KEY 		string
	CLIENT_ID		string
	CLIENT_SECRET	string
}

var instance *Environment = nil

func Get() *Environment {
	if instance == nil {
		instance = &Environment{
			PORT:			os.Getenv("PORT"),
			SECRET_KEY: 	os.Getenv("SECRET_KEY"),
			CLIENT_ID: 		os.Getenv("CLIENT_ID"),
			CLIENT_SECRET: 	os.Getenv("CLIENT_SECRET"),
		}
	}

	return instance
}

func Init() error {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	environment := Get()

	v := reflect.ValueOf(*environment)
	typeOfs := v.Type()

	var missing []string

	for i := range v.NumField() {
		if v.Field(i).Interface().(string) == "" {
			missing = append(missing, typeOfs.Field(i).Name)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing env vars %v", missing)
	}

	return nil
}
