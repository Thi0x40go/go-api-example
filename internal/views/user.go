package views

import "github.com/Thi0x40go/gin-postgres/internal/models"


func UserResponse(u models.User) map[string]any {
    return map[string]any{
        "id":   u.ID,
        "name": u.Name,
    }
}
