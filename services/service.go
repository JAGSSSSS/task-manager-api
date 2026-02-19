package services

import (
	"errors"
	"os"
	"time"

	"task-manager/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var TaskQueue = make(chan uuid.UUID, 100)

func CreateTask(task *model.Task) error {
	task.ID = uuid.New()
	task.Status = "pending"
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	err := model.DB.Create(task).Error
	if err == nil {
		TaskQueue <- task.ID
	}
	return err
}

func GetTasksByUser(userID uuid.UUID) ([]model.Task, error) {
	var tasks []model.Task
	err := model.DB.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	err := model.DB.Find(&tasks).Error
	return tasks, err
}

func DeleteTask(id uuid.UUID) error {
	return model.DB.Delete(&model.Task{}, "id = ?", id).Error
}
func StartAutoCompleteWorker(delay time.Duration) {
	for id := range TaskQueue {
		go func(taskID uuid.UUID) {
			time.Sleep(delay)

			var task model.Task
			if err := model.DB.First(&task, "id = ?", taskID).Error; err != nil {
				return
			}

			if task.Status != "completed" {
				model.DB.Model(&task).Update("status", "completed")
			}
		}(id)
	}
}

func Register(email, password, role string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.User{
		ID:       uuid.New(),
		Email:    email,
		Password: string(hashed),
		Role:     role,
	}

	return model.DB.Create(&user).Error
}

func Login(email, password string) (string, error) {
	var user model.User

	err := model.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signed, nil
}
