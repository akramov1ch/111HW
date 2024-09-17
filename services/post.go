package services

import (
    "errors"
    "111HW/models"
    "gorm.io/gorm"
)

func CreatePost(post *models.Post, userID uint) error {
    post.UserID = userID
    result := models.DB.Create(&post)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func GetAllPosts() ([]models.Post, error) {
    var posts []models.Post
    result := models.DB.Find(&posts)
    if result.Error != nil {
        return nil, result.Error
    }
    return posts, nil
}

func 	UpdatePost(id uint, post *models.Post, userID uint) error {
    var existingPost models.Post
    result := models.DB.First(&existingPost, id)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return result.Error
    }

    if existingPost.UserID != userID {
        return errors.New("unauthorized")
    }

    existingPost.Title = post.Title
    existingPost.Content = post.Content
    models.DB.Save(&existingPost)

    return nil
}

func 	DeletePost(id uint, userID uint) error {
    var post models.Post
    result := models.DB.First(&post, id)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return result.Error
    }

    if post.UserID != userID {
        return errors.New("unauthorized")
    }

    models.DB.Delete(&post)
    return nil
}
