package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/rmluser"
	"github.com/mikestefanello/pagoda/pkg/services"
)

type AuthAPI struct {
	services.Container
}

func (h *AuthAPI) Init(c *services.Container) error {
	h.Container = *c
	return nil
}

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}

func (h *AuthAPI) Routes(g *echo.Group) {
	g.POST("/api/auth/register", h.Register)
	g.POST("/api/auth/login", h.Login)
	g.POST("/api/auth/reset-password", h.ResetPassword)
}

func (h *AuthAPI) ResetPassword(c echo.Context) error {
	var req ResetPasswordRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, AuthResponse{Success: false, Message: "Invalid request body"})
	}

	ctx := c.Request().Context()

	// Find the user by email
	user, err := h.Container.ORM.RMLUser.Query().
		Where(rmluser.EmailEQ(req.Email)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, AuthResponse{Success: false, Message: "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, AuthResponse{Success: false, Message: "Failed to find user"})
	}

	// Update the user's password
	_, err = user.Update().
		SetPassword(req.NewPassword). // Note: In a real application, you should hash the password
		Save(ctx)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, AuthResponse{Success: false, Message: "Failed to update password"})
	}

	return c.JSON(http.StatusOK, AuthResponse{Success: true, Message: "Password reset successfully"})
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	User    *UserOutput `json:"user,omitempty"`
}

type UserOutput struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func (h *AuthAPI) Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, AuthResponse{Success: false, Message: "Invalid request body"})
	}

	user, err := h.Container.ORM.RMLUser.Create().
		SetName(req.Name).
		SetEmail(req.Email).
		SetPassword(req.Password). // Note: In a real application, you should hash the password
		Save(c.Request().Context())

	if err != nil {
		if ent.IsConstraintError(err) {
			return c.JSON(http.StatusConflict, AuthResponse{Success: false, Message: "Email already exists"})
		}
		return c.JSON(http.StatusInternalServerError, AuthResponse{Success: false, Message: "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, AuthResponse{
		Success: true,
		Message: "User registered successfully",
		User: &UserOutput{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	})
}

func (h *AuthAPI) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, AuthResponse{Success: false, Message: "Invalid request body"})
	}

	user, err := h.Container.ORM.RMLUser.Query().
		Where(rmluser.Email(req.Email), rmluser.Password(req.Password)). // Note: In a real application, you should compare hashed passwords
		Only(c.Request().Context())

	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusUnauthorized, AuthResponse{Success: false, Message: "Invalid credentials"})
		}
		return c.JSON(http.StatusInternalServerError, AuthResponse{Success: false, Message: "Login failed"})
	}

	return c.JSON(http.StatusOK, AuthResponse{
		Success: true,
		User: &UserOutput{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	})
}

func init() {
	Register(new(AuthAPI))
}
