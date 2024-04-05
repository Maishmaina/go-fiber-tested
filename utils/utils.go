// utils/utils.go

package utils

import (
    "crypto/sha256"
    "encoding/hex"
    "github.com/dgrijalva/jwt-go"
    "time"
)

const secretKey = "chamasoft12QWASX"

func HashPassword(password string) (string, error) {
    hash := sha256.New()
    hash.Write([]byte(password))
    hashedPassword := hex.EncodeToString(hash.Sum(nil))
    return hashedPassword, nil
}

func CheckPasswordHash(password, hash string) bool {
    hashedPassword, _ := HashPassword(password)
    return hashedPassword == hash
}

func GenerateJWT(userID int) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }

    return signedToken, nil
}
