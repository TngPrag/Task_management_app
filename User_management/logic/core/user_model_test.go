package core

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserSchema(t *testing.T) {
	err := CreateUserSchema()
	assert.NoError(t, err)
}

func TestInitSuperAdminUser(t *testing.T) {
	err := InitSuperAdminUser()
	assert.NoError(t, err)
}

func TestCreateUser(t *testing.T) {
	user := User{
		Id:        uuid.NewString(),
		Owner_id:  uuid.NewString(),
		Name:      "John Doe",
		UserName:  "johndoe",
		Password:  "Password123!",
		Email:     "john.doe@example.com",
		CreateAt:  time.Now(),
		UPdatedAt: time.Now(),
	}

	err := user.Create_user()
	assert.NoError(t, err)

	// Verify the user was created
	data, err := user.Get_user_by_uid()
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}

func TestGetUserByUID(t *testing.T) {
	user := User{
		Id:       uuid.NewString(),
		Owner_id: uuid.NewString(),
	}

	// Create user first
	err := user.Create_user()
	assert.NoError(t, err)

	// Retrieve user by UID
	data, err := user.Get_user_by_uid()
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}

func TestGetUserByEmailUserName(t *testing.T) {
	user := User{
		Id:        uuid.NewString(),
		Owner_id:  uuid.NewString(),
		Name:      "John Doe",
		UserName:  "johndoe",
		Password:  "Password123!",
		Email:     "john.doe@example.com",
		CreateAt:  time.Now(),
		UPdatedAt: time.Now(),
	}

	// Create user first
	err := user.Create_user()
	assert.NoError(t, err)

	// Retrieve user by email and username
	data, err := user.Get_user_by_email_userName()
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}

func TestGetUserByOwnerID(t *testing.T) {
	user := User{
		Id:        uuid.NewString(),
		Owner_id:  uuid.NewString(),
		Name:      "John Doe",
		UserName:  "johndoe",
		Password:  "Password123!",
		Email:     "john.doe@example.com",
		CreateAt:  time.Now(),
		UPdatedAt: time.Now(),
	}

	// Create user first
	err := user.Create_user()
	assert.NoError(t, err)

	// Retrieve users by owner ID
	data, err := user.Get_user_by_owner_id()
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}

func TestRemoveUserByID(t *testing.T) {
	user := User{
		Id:        uuid.NewString(),
		Owner_id:  uuid.NewString(),
		Name:      "John Doe",
		UserName:  "johndoe",
		Password:  "Password123!",
		Email:     "john.doe@example.com",
		CreateAt:  time.Now(),
		UPdatedAt: time.Now(),
	}

	// Create user first
	err := user.Create_user()
	assert.NoError(t, err)

	// Remove user by ID
	err = user.Remove_user_by_id()
	assert.NoError(t, err)

	// Verify user was removed
	data, err := user.Get_user_by_uid()
	assert.Error(t, err)
	assert.Empty(t, data)
}

func TestRemoveUserByOwner(t *testing.T) {
	user := User{
		Id:        uuid.NewString(),
		Owner_id:  uuid.NewString(),
		Name:      "John Doe",
		UserName:  "johndoe",
		Password:  "Password123!",
		Email:     "john.doe@example.com",
		CreateAt:  time.Now(),
		UPdatedAt: time.Now(),
	}

	// Create user first
	err := user.Create_user()
	assert.NoError(t, err)

	// Remove users by owner ID
	err = user.Remove_user_by_owner()
	assert.NoError(t, err)

	// Verify users were removed
	data, err := user.Get_user_by_owner_id()
	assert.Error(t, err)
	assert.Empty(t, data)
}
