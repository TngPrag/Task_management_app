package userhandlers

import (
	"encoding/json"
	"log"
	"time"
	"user_manager/logic/core"
	"user_manager/logic/dto/requests"
	"user_manager/logic/dto/responses"
	"user_manager/logic/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func User_authenticate(c *fiber.Ctx) error {
	user_verified_profile := new(core.User)
	user_verified_profile.Id = c.Locals("user_id").(string)
	user_data, err := user_verified_profile.Get_user_by_uid()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"User not found": err.Error()})
	}
	if err := json.Unmarshal(user_data, &user_verified_profile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	user_verified_resp := &responses.UserAuthProfileDto{}
	user_verified_resp.UserID = user_verified_profile.Id
	user_verified_resp.Name = user_verified_profile.Name
	user_verified_resp.UserName = user_verified_profile.UserName
	user_verified_resp.Email = user_verified_profile.Email
	//println(user_verified_resp)
	//data, _ := json.Marshal(user_verified_resp)
	//println(string(data))
	return c.Status(fiber.StatusOK).JSON(user_verified_resp)
}

func User_signup(c *fiber.Ctx) error {
	user_dto := new(requests.CreateUserDto)
	if err := c.BodyParser(user_dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"Invalid input, create user": err.Error()})
	}
	if err := user_dto.ValidateCreateUserDto(); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"Error": err.Error()})
	}
	// Authenticate the user and check if the user found
	user_id := c.Locals("user_id").(string)
	log.Println(user_id)
	//user_name := c.Locals("user_name").(string)
	//email := c.Locals("email").(string)
	token := c.Locals("token").(string)
	user_req_model := new(core.User)
	user_req_model.Id = user_id
	user_data, err := user_req_model.Get_user_by_uid()
	if user_data == nil && err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error data base: user not found": err.Error()})
	}
	// check user authorization: 1. check role of the requesting user
	role, err := pkg.GetUserRole(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error your role is not found": err.Error()})
	}

	// check if this request making user can access this api
	status, _ := pkg.VerifyPolicy(token, role, "task_app/user_manager_service/api/v0.1/user", "POST")
	log.Println(status)
	if role == "super-admin" {
		// check policy permission
		if status {
			// if so create the user
			//hashedPassword, _ := core.HashPassword(user_dto.Password)
			user_req_model.Id = uuid.NewString()
			user_req_model.Name = user_dto.FirstName + user_dto.LastName
			user_req_model.UserName = user_dto.UserName
			user_req_model.Owner_id = user_id
			user_req_model.Password = user_dto.Password
			user_req_model.Email = user_dto.Email
			user_req_model.CreateAt = time.Now()
			user_req_model.UPdatedAt = time.Now()
			if err := user_req_model.Create_user(); err != nil {
				//log.Println("hello")
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})

			}
			if err := pkg.AssignRole(token, user_req_model.Id, "admin"); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})
			}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "user created successfully"})
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error your identitiy is not found": err.Error()})
		}
	} else if role == "admin" {
		// check policy permission
		if status {
			// if so create the user
			//hashedPassword, _ := core.HashPassword(user_dto.Password)
			user_req_model.Id = uuid.NewString()
			user_req_model.Name = user_dto.FirstName + user_dto.LastName
			user_req_model.UserName = user_dto.UserName
			user_req_model.Owner_id = user_id
			user_req_model.Password = user_dto.Password
			user_req_model.Email = user_dto.Email
			user_req_model.CreateAt = time.Now()
			user_req_model.UPdatedAt = time.Now()
			if err := user_req_model.Create_user(); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})
			}
			if err := pkg.AssignRole(token, user_req_model.Id, "user"); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})
			}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "user created successfully"})
		}
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error your identitiy is not found": err.Error()})

}

func User_login(c *fiber.Ctx) error {
	// Parse the request body into a login DTO
	userLoginDto := new(requests.LoginDto)
	if err := c.BodyParser(userLoginDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input", "details": err.Error()})
	}

	// Validate the login DTO
	if err := userLoginDto.ValidateLoginRequestDto(); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Invalid input", "details": err.Error()})
	}

	// Check if the user exists
	userModel := new(core.User)
	userModel.UserName = userLoginDto.UserName
	userModel.Email = userLoginDto.Email
	user, err := userModel.Get_user_by_email_userName()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found", "details": err.Error()})
	}

	// Deserialize user data
	if err := json.Unmarshal(user, userModel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal problem", "details": err.Error()})
	}

	// Check the password
	//log.Println("Password from request:", userLoginDto.Password)
	//log.Println("Password from DB:", userModel.Password)
	if !core.CheckPasswordHash(userLoginDto.Password, userModel.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Wrong password"})
	}

	// Generate the token
	token, err := core.GenerateToken(userModel.Id, userModel.UserName, userModel.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Token generation failed", "details": err.Error()})
	}

	// Set the token as a cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HTTPOnly = true
	c.Cookie(cookie)

	// Return success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}

func User_read_by_id(c *fiber.Ctx) error {

	uid := c.Params("user_id")
	user_model := new(core.User)
	user_model.Id = uid
	req_user_id := c.Locals("user_id").(string)
	req_token := c.Locals("token").(string)
	user_model.Owner_id = req_user_id
	// authorize the user
	role, _ := pkg.GetUserRole(req_token)
	status, _ := pkg.VerifyPolicy(req_token, role, "task_app/user_manager_service/api/v0.1/user", "GET")
	if status {
		user_data, err := user_model.Get_user_by_uid()
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"User not found": err.Error()})
		}
		if err := json.Unmarshal(user_data, &user_model); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(user_model)
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error": "your identitiy is not found"})

}

func User_remove_by_id(c *fiber.Ctx) error {
	uid := c.Params("user_id")
	user_model := new(core.User)
	user_model.Id = uid

	req_user_id := c.Locals("user_id").(string)
	req_token := c.Locals("token").(string)
	user_model.Owner_id = req_user_id
	// authorize the user
	role, _ := pkg.GetUserRole(req_token)
	status, _ := pkg.VerifyPolicy(req_token, role, "task_app/user_manager_service/api/v0.1/user", "DELETE")
	if status {
		err := user_model.Remove_user_by_id()
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"User not found": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"user removed successfully": uid})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error": "your identitiy is not found"})

}
func User_remove_by_owner_id(c *fiber.Ctx) error {
	user_model := new(core.User)
	req_user_id := c.Locals("user_id").(string)
	req_token := c.Locals("token").(string)
	user_model.Owner_id = req_user_id
	// authorize the user
	role, _ := pkg.GetUserRole(req_token)
	status, _ := pkg.VerifyPolicy(req_token, role, "task_app/task_manager_service/api/v0.1/user", "DELETE")
	if status {
		err := user_model.Remove_user_by_owner()
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"User not found": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "user removed successfully"})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error": "your identitiy is not found"})

}
func User_list_by_owner_id(c *fiber.Ctx) error {
	user_model := new(core.User)
	req_user_id := c.Locals("user_id").(string)
	req_token := c.Locals("token").(string)
	user_model.Owner_id = req_user_id
	// authorize the user
	role, _ := pkg.GetUserRole(req_token)
	status, _ := pkg.VerifyPolicy(req_token, role, "task_app/task_manager_service/api/v0.1/user", "GET")
	if status {
		if role == "super-admin" || role == "admin" {
			data, err := user_model.Get_user_by_owner_id()
			if err != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"User not found": err.Error()})
			}

			return c.Status(fiber.StatusOK).JSON(data)
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error": "No policy defined to list for user"})
		}
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error": "your identitiy is not found"})

}

func User_notify(c *fiber.Ctx) error {
	userNotifyDto := new(requests.UserNotifyDto)
	if err := c.BodyParser(userNotifyDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input", "details": err.Error()})
	}

	if err := userNotifyDto.ValidateUserNotifyDto(); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"Invalid input ": err.Error()})
	}
	req_token := c.Locals("token").(string)
	role, _ := pkg.GetUserRole(req_token)
	status, _ := pkg.VerifyPolicy(req_token, role, "task_app/user_manager_service/api/v0.1/user", "POST")
	if status {
		if role == "admin" {
			// get task user email
			user_data_model := new(core.User)
			user_data_model.Id = userNotifyDto.User_id
			user_data, err2 := user_data_model.Get_user_by_uid()
			if err2 != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Error, asigned user not found": err2.Error()})
			}
			if json.Unmarshal(user_data,&user_data_model); err2!= nil{
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error":err2})
			}
			send_data_model := new(pkg.EmailAdpater)
			send_data_model.To = user_data_model.Email
			send_data_model.Subject = userNotifyDto.Title
			send_data_model.Body = userNotifyDto.Description + " And the deadline is " + userNotifyDto.Deadline
			err := send_data_model.SendMessageViaEmail()
			if err != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Unable to send user email notification": err.Error()})
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "User notified successfully via email"})
		}
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error": "your identitiy is not found"})

}
