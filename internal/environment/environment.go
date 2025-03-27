package environment

import (
	"crypto/rsa"
	"encoding/pem"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type Environment struct {
	CLIENT_ID				string
	CLIENT_SECRET			string

	PORT					int

	SECRET 					*rsa.PrivateKey

	TOKEN_EXPIRATION_TIME 	int
}

var instance *Environment = nil

func Get() *Environment {
	if instance == nil {
		port, _ := strconv.Atoi(os.Getenv("PORT"))
		tokenExpirationTime, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_TIME"))

		secret, _ := loadPrivateKey(os.Getenv("PRIVATE_KEY_PATH"))

		instance = &Environment{
			PORT:					port,
			SECRET: 				secret,
			CLIENT_ID: 				os.Getenv("CLIENT_ID"),
			CLIENT_SECRET: 			os.Getenv("CLIENT_SECRET"),
			TOKEN_EXPIRATION_TIME: 	tokenExpirationTime,
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

	// validate if all env vars are being set
	for i := range v.NumField() {
		field := v.Field(i)
		name := typeOfs.Field(i).Name

		switch field.Kind() {
		case reflect.String:
			ValidateString(field.String(), name, &missing)
		case reflect.Int:
			ValidateInt(int(field.Int()), name, &missing)
		case reflect.Ptr:
			ValidatePtr(field, name, &missing)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing env vars %v", missing)
	}

	return nil
}

func loadPrivateKey(path string) (*rsa.PrivateKey, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key: %w", err)
	}

	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, fmt.Errorf("no PEM data found")
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return key, nil
}

func ValidateString(val string, name string, missing *[]string) {
	if val == "" {
		*missing = append(*missing, name)
	}
}

func ValidateInt(val int, name string, missing *[]string) {
	if val == 0 {
		*missing = append(*missing, name)
	}
}

func ValidatePtr(val reflect.Value, name string, missing *[]string) {
	if val.IsNil() {
		*missing = append(*missing, name)
	}
}
