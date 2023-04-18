package controllers

import (
	models "app/app/models/module_job"
	queries "app/app/queries/module_job"

	"app/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// เพิ่ม
func CreateJob(c *fiber.Ctx) error {

	//รับค่า
	filter := models.InputJsonStructJob{}
	if err := c.BodyParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":   utils.ResponseCode()["api"]["invalid_data_type"],
			"msg":    utils.ResponseMessage()["api"]["invalid_data_type"],
			"msglog": err,
		})
	}

	//เพิ่มข้อมูล
	if err := queries.InsertJob(&filter); !err {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":   utils.ResponseCode()["api"]["cannot_insert"],
			"msg":    utils.ResponseMessage()["api"]["cannot_insert"],
			"msglog": err,
		})
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"data": filter,
		"msg":  utils.ResponseMessage()["api"]["success"],
	})
}

// รายการ
func ListJob(c *fiber.Ctx) error {
	data := queries.GetListJob()
	return c.JSON(fiber.Map{
		"code": 200,
		"data": data,
		"msg":  utils.ResponseMessage()["api"]["success"],
	})
}

// รายการตามรหัส
func Detailjob(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": utils.ResponseCode()["api"]["cannot_insert"],
			"msg":  "ไม่พบ parameter",
		})
	}

	data := queries.GetListJobByID(id)

	return c.JSON(fiber.Map{
		"code": 200,
		"data": data,
		"msg":  utils.ResponseMessage()["api"]["success"],
	})
}

// แก้ไขข้อมูล
func UpdateJob(c *fiber.Ctx) error {
	//รับค่า
	filter := models.UpdateJob{}
	if err := c.BodyParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":   utils.ResponseCode()["api"]["invalid_data_type"],
			"msg":    utils.ResponseMessage()["api"]["invalid_data_type"],
			"msglog": err,
		})
	}

	//แก้ไขข้อมูล
	if err := queries.UpdateJob(&filter); !err {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":   utils.ResponseCode()["api"]["cannot_insert"],
			"msg":    utils.ResponseMessage()["api"]["cannot_insert"],
			"msglog": err,
		})
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"data": filter,
		"msg":  utils.ResponseMessage()["api"]["success"],
	})
}

// ลบข้อมูล
func DeleteJob(c *fiber.Ctx) error {
	//รับค่า
	filter := models.DeleteJob{}
	if err := c.BodyParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":   utils.ResponseCode()["api"]["invalid_data_type"],
			"msg":    utils.ResponseMessage()["api"]["invalid_data_type"],
			"msglog": err,
		})
	}

	//ลบข้อมูล
	status, err := queries.DeleteJob(&filter)
	if status != true {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":   utils.ResponseCode()["api"]["cannot_delete"],
			"msg":    utils.ResponseMessage()["api"]["cannot_delete"],
			"msglog": err,
		})
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"data": filter,
		"msg":  utils.ResponseMessage()["api"]["success"],
	})
}
