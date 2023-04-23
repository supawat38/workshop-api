package utils

func ResponseCode() (response map[string]map[string]int) {
	response = map[string]map[string]int{
		"api": {
			"success":           200,
			"cannot_insert":     1000,
			"cannot_update":     1001,
			"cannot_delete":     1001,
			"format_fail":       1002,
			"invalid_data_type": 1003,
		},
	}
	return
}

func ResponseMessage() (response map[string]map[string]map[string]string) {
	response = map[string]map[string]map[string]string{
		"api": {
			"success": {
				"th": "สำเร็จ",
				"en": "Success",
			},
			"cannot_insert": {
				"th": "ไม่สามารถเพิ่มข้อมูลได้",
				"en": "can't insert value !",
			},
			"cannot_update": {
				"th": "ไม่สามารถแก้ไขข้อมูลได้",
				"en": "can't update value !",
			},
			"cannot_delete": {
				"th": "ไม่สามารถลบข้อมูลได้",
				"en": "can't delete value !",
			},
			"format_fail": {
				"th": "ข้อความที่ส่งมาไม่ถูกต้อง (ตรวจพบอักขระพิเศษ)",
				"en": "Invalid message sent (Special character detected)",
			},
			"invalid_data_type": {
				"th": "ประเภทข้อมูลไม่ถูกต้อง",
				"en": "invalid data type",
			},
		},
	}
	return
}
