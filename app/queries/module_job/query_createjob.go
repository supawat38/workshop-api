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

// แก้ไขข้อมูล
func UpdateJob(req *models.UpdateJob) (success bool) {

	query := `UPDATE job SET 
				position = @position , url = @url , 
				location = @location , salary = @salary ,
				job_reject = @job_reject , type = @type ,
				remark = @remark 
			  WHERE id = @id `

	var dataUpdate = map[string]interface{}{
		"id":         req.Id,
		"position":   req.Position,
		"url":        req.Url,
		"location":   req.Location,
		"salary":     req.Salary,
		"job_reject": req.JobReject,
		"type":       req.Type,
		"remark":     req.Remark,
	}

	err := database.DBConn.Exec(query, dataUpdate).Error
	if err != nil {
		success = false
		return
	}
	success = true
	return
}

// ลบข้อมูล
func DeleteJob(req *models.DeleteJob) (status bool, err_str string) {
	tx := database.DBConn.Begin()

	tablename := models.Job{}
	if err := tx.Where("id = ?", req.Id).Delete(&tablename).Error; err != nil {
		err_str = err.Error()
		tx.Rollback()
		return
	}

	tx.Commit()
	status = true

	return
}
