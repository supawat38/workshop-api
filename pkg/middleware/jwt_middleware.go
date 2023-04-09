package middleware

import (
	"app/pkg/utils"
	"os"

	funcCommon "app/app/controllers/common"

	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
)

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/jwt
func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey:   "jwt", // used in private routes
		ErrorHandler: jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":  4001,
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"code":  4001,
		"error": true,
		"msg":   err.Error(),
	})
}

//ตรวจสอบว่า Token นั้น เคยมีในระบบแล้วจริงๆ
func CheckAccessToken(c *fiber.Ctx) error {
	//ข้อมูล Token
	TokenForUser, _ := funcCommon.GetTokenFromSession(c)

	//Local IP
	LocalIP := funcCommon.ReadUserIP(c)

	//ตรวจสอบว่ามีการ login ในตาราง user_agent or sub_account จริงๆไหม (check token , ip address)
	StatusCheckToken := funcCommon.GetTokenInTable(TokenForUser, LocalIP)

	if StatusCheckToken != "" {
		return c.Next()
	}

	return c.JSON(fiber.Map{
		"code":     utils.ResponseCode()["api"]["token_auth_fail"],
		"msg":      utils.ResponseMessage()["api"]["token_auth_fail"],
		"platform": "Token ต้องเป็นการเข้าสู่ระบบ จากหลังบ้านเท่านั้น (agent & shareholder & company & subaccount)",
	})
}

//ตรวจสอบว่า Token นั้น เคยมีในระบบแล้วจริงๆ (และเป็น member ไหม)
func CheckAccessTokenMember(c *fiber.Ctx) error {
	//ข้อมูล Token
	TokenForUser, _ := funcCommon.GetTokenFromSession(c)

	//Local IP
	LocalIP := funcCommon.ReadUserIP(c)

	//ตรวจสอบว่ามีการ login ในตาราง user_members จริงๆไหม (check token , ip address)
	StatusCheckToken := funcCommon.GetTokenInTableMember(TokenForUser, LocalIP)

	if StatusCheckToken != "" {
		return c.Next()
	}

	return c.JSON(fiber.Map{
		"code":     utils.ResponseCode()["api"]["token_auth_fail"],
		"msg":      utils.ResponseMessage()["api"]["token_auth_fail"],
		"platform": "Token ต้องเป็นการเข้าสู่ระบบเป็น จากหน้าบ้านเท่านั้น (member)",
	})
}
