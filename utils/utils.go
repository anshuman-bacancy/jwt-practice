package utils

import (
	"fmt"
	"jwt/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// LoadConfig reads and returns config
func LoadConfig(path string) (config models.Config, configErr error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	configErr = viper.ReadInConfig()
	if configErr != nil {
		return
	}
	configErr = viper.Unmarshal(&config)
	return
}

// HashPassword hashes the password
func HashPassword(password string) (string, error) {
	passbytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(passbytes), err
}

// CheckPasswordHash checks the validity of password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWTToken generates JWT token
func GenerateJWTToken(email, role string) (string, error) {
	fmt.Printf("Inside JWT : %s %s \n", email, role)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	config, _ := LoadConfig(".")
	skbyte := []byte(config.SigningKey)

	tokenString, _ := token.SignedString(skbyte)
	fmt.Println(tokenString)
	return tokenString, nil
}
