package queries

import (
	models "app/app/models/module_job"
	"time"

	"app/platform/database"
)

// เพิ่มข้อมูล
func InsertJob(req *models.InputJsonStructJob) (success bool) {

	//วันที่ปัจจุบัน
	dTimeNow := time.Now().Local()
	dTimeString := dTimeNow.Format("2006-01-02 15:04:05")

	sqlStatement := ` INSERT INTO job ( name , position , url , location , salary , type  , apply , created_at , job_reaction , job_reject) `
	sqlStatement += ` 	VALUES (@name , @position , @url , @location , @salary , @type , @apply , @created_at , false , false) `
	if database.DBConn.Exec(sqlStatement,
		map[string]interface{}{
			"name":       req.Name,
			"position":   req.Position,
			"url":        req.Url,
			"location":   req.Location,
			"salary":     req.Salary,
			"type":       req.Type,
			"apply":      req.Apply,
			"created_at": dTimeString,
		}).Error != nil {
		return
	}

	success = true
	return
}

// รายการ
func GetListJob() (result []models.InputJsonStructJob) {
	query := ` SELECT * FROM job ORDER BY created_at DESC `
	database.DBConn.Raw(query).Scan(&result)
	return
}

// รายการข้อมูลตาม ID
func GetListJobByID(ID string) (result models.InputJsonStructJob) {
	query := ` SELECT * FROM job WHERE id = @ID`
	database.DBConn.Raw(query, map[string]interface{}{"ID": ID}).Scan(&result)
	return
}
