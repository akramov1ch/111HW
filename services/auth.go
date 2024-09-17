package services

import (
    "errors"
    "111HW/models"
    "time"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "gorm.io/gorm"
)

func RegisterUser(user *models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)

    result := models.DB.Create(&user)
    if result.Error != nil {
        return result.Error
    }
    return nil
}
func LoginUser(input models.LoginInput) (string, error) {
    var user models.User

    result := models.DB.Where("email = ?", input.Email).First(&user)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return "", errors.New("user not found")
    }

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
    if err != nil {
        return "", errors.New("incorrect password")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userID": user.ID,
        "exp":    time.Now().Add(time.Hour * 72).Unix(),
    })
    tokenString, err := token.SignedString([]byte("secret"))
    if err != nil {
        return "", err
    }
    return tokenString, nil
}
