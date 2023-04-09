package utils

func ResponseCode() (response map[string]map[string]int) {
	response = map[string]map[string]int{
		"api": {
			"success":              200,
			"cannot_insert":        1000,
			"cannot_update":        1001,
			"cannot_delete":        1001,
			"format_fail":          1002,
			"invalid_data_type":    1003,
			"user_invalid_login":   1004,
			"token_generate_error": 1005,
			"permission_denied":    1006,
			"token_auth_fail":      1007,
			"parameter_fail":       1008,

			//acount
			"password_not_correct": 2001,
			"char_more_exceeded":   2002,

			//management
			"username_duplicate":                3001,
			"thb_more_than_parents":             3050,
			"usd_more_than_parents":             3051,
			"cny_more_than_parents":             3052,
			"idr_more_than_parents":             3053,
			"vnd_more_than_parents":             3054,
			"lak_more_than_parents":             3055,
			"myr_more_than_parents":             3056,
			"mmk_more_than_parents":             3057,
			"khr_more_than_parents":             3058,
			"php_more_than_parents":             3059,
			"eur_more_than_parents":             3060,
			"twd_more_than_parents":             3061,
			"inr_more_than_parents":             3062,
			"gbp_more_than_parents":             3063,
			"krw_more_than_parents":             3064,
			"pt_more_than_parents":              3002,
			"max_more_than_parents":             3003,
			"min_more_than_parents":             3004,
			"payout_more_than_parents":          3005,
			"payout_discount_more_than_parents": 3006,
			"permision_denied_register":         3007,
			"account_number_duplicate":          3008,
			"uploadimage_fail":                  3009,
			"credit_not_enough":                 3010,
			"resellandmaster_isnull":            3011,
			"ptdis_than_child":                  3012,

			//payment
			"invalid_transaction_type": 6001,
			"invalid_payment_role":     6002,
			"balance_notenough":        6003,
			"job_lock_payment":         6004,

			//setting
			"create_drawagent_error":                 8001,
			"get_defaultclosedraw_config_error":      8001,
			"cal_result_error":                       8002,
			"update_result_error":                    8003,
			"reject_drawagent_error":                 8004,
			"list_drawagent_error":                   8005,
			"create_placeout_error":                  8020,
			"create_levellimit_error":                8021,
			"create_enablehotplaceout_error":         8022,
			"create_hotnumber_error":                 8023,
			"create_limitnumber_error":               8024,
			"list_placeout_error":                    8030,
			"list_hotplaceout_error":                 8031,
			"list_hotnumber_error":                   8032,
			"list_limit_error":                       8033,
			"waitapproveconfirm":                     8034,
			"cantapproveclosedraw":                   8035,
			"username_duplication_approveconfirm":    8036,
			"create_drawagent_error_invalid_lotto":   8037,
			"create_drawagent_error_lotto_pending":   8038,
			"create_drawagent_error_lotto_duplicate": 8039,
			"create_drawagent_error_lotto_match":     8040,

			//หวยที่ชอบ
			"update_favorite_lotto": 9001,
			//เลขทีชอบ
			"add_favorite_number": 9002,

			//yeekee ยิงเลข
			"add_yeekeepredictnumber":          9003,
			"list_yeekeepredictnumber":         9004,
			"add_yeekeepredictnumber_cooldown": 9005,

			//bet
			"get_drawagent_error": 1101,
			"bet_closetime_error": 1102,
			"bet_jobcheck_error":  1201,
			"bet_credit_enough":   1203,
			"bet_detail_zodiac":   1202,
			"bet_cancel_fail":     1204,
			"bet_cancel_timeout":  1205,
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
			"user_invalid_login": {
				"th": "ไม่พบข้อมูล ล็อกอิน",
				"en": "invalid user login",
			},
			"token_generate_error": {
				"th": "ไม่สามารถสร้างโทเค็นได้",
				"en": "can't create token",
			},
			"username_duplicate": {
				"th": "พบข้อมูลผู้ใช้ซ้ำ กรุณาลองใหม่อีกครั้ง",
				"en": "Username duplicate please try again.",
			},
			"parameter_fail": {
				"th": "ข้อมูลที่ส่งมาไม่ครบถ้วน",
				"en": "incomplete information",
			},
			"thb_more_than_parents": {
				"th": "สกุลเงิน THB ไม่สามารถกรอกเกินกว่าสาย กรุณาลองใหม่อีกครั้ง",
				"en": "THB more than parents!",
			},
			"usd_more_than_parents": {
				"th": "สกุลเงิน USD ไม่สามารถกรอกเกินกว่าสาย กรุณาลองใหม่อีกครั้ง",
				"en": "USD more than parents!",
			},
			"cny_more_than_parents": {
				"th": "สกุลเงิน CNY ไม่สามารถกรอกเกินกว่าสาย กรุณาลองใหม่อีกครั้ง",
				"en": "CNY more than parents!",
			},
			"idr_more_than_parents": {
				"th": "สกุลเงิน IDR ไม่สามารถกรอกเกินกว่าสาย กรุณาลองใหม่อีกครั้ง",
				"en": "IDR more than parents!",
			},
			"vnd_more_than_parents": {
				"th": "สกุลเงิน VND ไม่สามารถกรอกเกินกว่าสาย กรุณาลองใหม่อีกครั้ง",
				"en": "VND more than parents!",
			},
			"lak_more_than_parents": {
				"th": "สกุลเงิน LAK ไม่สามารถกรอกเกินกว่าสาย กรุณาลองใหม่อีกครั้ง",
				"en": "LAK more than parents!",
			},
			"myr_more_than_parents": {
				"th": "สกุลเงิน MYR ไม่สามารถกรอกเกินกว่าสาย กรุณาลองใหม่อีกครั้ง",
				"en": "MYR more than parents!",
			},
			"mmk_more_than_parents": {
				"th": "สกุลเงิน MMK ไม่สามารถกรอกเกินกว่าสาย กรุณาลองใหม่อีกครั้ง",
				"en": "MMK more than parents!",
			},
			"khr_more_than_parents": {
				"th": "สกุลเงิน KHR ไม่สามารถกรอกเกินกว่าสาย กรุณาลองใหม่อีกครั้ง",
				"en": "KHR more than parents!",
			},
			"php_more_than_parents": {
				"th": "สกุลเงิน PHP ไม่สามารถกรอกเกินกว่าสาย กรุณาลองใหม่อีกครั้ง",
				"en": "PHP more than parents!",
			},
			"eur_more_than_parents": {
				"th": "สกุลเงิน EUR ไม่สามารถกรอกเกินกว่าสาย กรุณาลองใหม่อีกครั้ง",
				"en": "EUR more than parents!",
			},
			"twd_more_than_parents": {
				"th": "สกุลเงิน TWD ไม่สามารถกรอกเกินกว่าสาย กรุณาลองใหม่อีกครั้ง",
				"en": "TWD more than parents!",
			},
			"inr_more_than_parents": {
				"th": "สกุลเงิน INR ไม่สามารถกรอกเกินกว่าสาย กรุณาลองใหม่อีกครั้ง",
				"en": "INR more than parents!",
			},
			"gbp_more_than_parents": {
				"th": "สกุลเงิน GBP ไม่สามารถกรอกเกินกว่าสาย กรุณาลองใหม่อีกครั้ง",
				"en": "GBP more than parents!",
			},
			"krw_more_than_parents": {
				"th": "สกุลเงิน KRW ไม่สามารถกรอกเกินกว่าสาย กรุณาลองใหม่อีกครั้ง",
				"en": "KRW more than parents!",
			},
			"credit_not_enough": {
				"th": "ค่าเงิน ของคนสร้างไม่เพียงพอที่จะสร้างสมาชิก",
				"en": "Credit not enough",
			},
			"create_drawagent_error": {
				"th": "ไม่สามารถสร้างรอบหวยได้",
				"en": "can't create drawagents",
			},
			"permission_denied": {
				"th": "คุณไม่มีสิทธื์ในการทำรายการ",
				"en": "you do not have permission to access",
			},
			"get_defaultclosedraw_config_error": {
				"th": "ไม่สามารถกำหนดค่าเรื่มต้นได้",
				"en": "can't default data config",
			},
			"cal_result_error": {
				"th": "ไม่สามารถคำนวณได้",
				"en": "can't cal result",
			},
			"pt_more_than_parents": {
				"th": "ตั้งค่า PT เกินกว่าเงื่อนไขที่กำหนด กรุณาลองใหม่อีกครั้ง",
				"en": "PT More then condition!",
			},
			"update_result_error": {
				"th": "เกิดข้อผิดพลาดในการปิดผล",
				"en": "can't close draw",
			},
			"token_auth_fail": {
				"th": "ไม่สามารถตรวจสอบรหัส Token ได้ กรุณาลองใหม่อีกครั้ง",
				"en": "Can't verify token code",
			},
			"max_more_than_parents": {
				"th": "ค่า มากสุด เกินกว่าเงื่อนไขที่กำหนด กรุณาลองใหม่อีกครั้ง",
				"en": "Value Max More then condition!",
			},
			"min_more_than_parents": {
				"th": "ค่า ต่ำสุด น้อยกว่าเงื่อนไขที่กำหนด กรุณาลองใหม่อีกครั้ง",
				"en": "Value Min More then condition!",
			},
			"payout_more_than_parents": {
				"th": "อัตราการจ่ายเงินหลังบวก 10% ห้ามกรอกเกินกว่าสาย กรุณาลองใหม่อีกครั้ง",
				"en": "Value Payout More then condition!",
			},
			"payout_discount_more_than_parents": {
				"th": "อัตราการจ่ายเงินหลังหักส่วนลด เกินกว่าเงื่อนไขที่กำหนด กรุณาลองใหม่อีกครั้ง",
				"en": "Value Payout More then condition!",
			},
			"reject_drawagent_error": {
				"th": "ไม่สามารถยกเลิกรอบของหวยได้",
				"en": "can't reject drawagent",
			},
			"list_drawagent_error": {
				"th": "ไม่สามารถเรียกดูรายการได้",
				"en": "can't list drawagent",
			},
			"permision_denied_register": {
				"th": "คุณไม่มีสิทธิ์สร้างข้อมูลผู้ใช้",
				"en": "You don't have permission to create member",
			},
			"account_number_duplicate": {
				"th": "รหัสบัญชีธนาคารซ้ำ ในสายเดียวกัน",
				"en": "Bank account is duplicate",
			},
			"invalid_transaction_type": {
				"th": "ประเภทการทำรายการไม่ถูกต้อง (deposit,withdraw)",
				"en": "Invalid transaction type (deposit,withdraw)",
			},
			"invalid_payment_role": {
				"th": "สิทธิในการทำข้อมูลให้ไม่ถูกต้อง (member,agent)",
				"en": "You don't have permission to deposite or withdraw (member,agent)",
			},
			"balance_notenough": {
				"th": "ยอดคงเหลือไม่เพียงพอต่อการทำรายการ",
				"en": "Balance not enough for transaction",
			},
			"resellandmaster_isnull": {
				"th": "กรุณาระบุประเภทของข้อมูล (master หรือ reseller)",
				"en": "Please select type data (master or reseller)",
			},
			"ptdis_than_child": {
				"th": "เปอร์เซ็นถือสู้ไม่สามารถตั้งได้ต่ำกว่า สายที่เคยกำหนดให้",
				"en": "PT Can't not less agent",
			},
			"password_not_correct": {
				"th": "รหัสผ่านเก่าไม่ถูกต้อง",
				"en": "Password old not correct",
			},
			"create_placeout_error": {
				"th": "ตั้งค่ารับกินล้มเหลว",
				"en": "Create Placeout failed",
			},
			"create_levellimit_error": {
				"th": "ตั้งค่ารับกินล้มเหลว",
				"en": "Create Placeout failed",
			},
			"create_enablehotplaceout_error": {
				"th": "ตั้งค่ารับกินล้มเหลว",
				"en": "Create Placeout failed",
			},
			"list_hotplaceout_error": {
				"th": "ตั้งค่ารับกินล้มเหลว",
				"en": "Create Placeout failed",
			},
			"job_lock_payment": {
				"th": "ติดข้อมูล joblock",
				"en": "Job lock -> joblock_payments",
			},
			"create_hotnumber_error": {
				"th": "อัพเดตเลขดังล้มเหลว",
				"en": "Update Hotnumber failed",
			},
			"list_hotnumber_error": {
				"th": "รายการเลขดังล้มเหลว",
				"en": "List Hotnumber failed",
			},
			"waitapproveconfirm": {
				"th": "รอการอนุมัติจาก บุคคลที่สอง",
				"en": "Waiting for approval from the second party",
			},
			"cantapproveclosedraw": {
				"th": "ข้อมูลการปิดผลรอบแรก และรอบที่สอง ไม่สอดคล้องกัน กรุณาดำเนินการใหม่ทั้งหมดอีกครั้ง",
				"en": "Information does not correspond to first person. Please do all over again.",
			},
			"username_duplication_approveconfirm": {
				"th": "ไม่อนุญาตให้ ผู้ใช้คนเดิมปิดผลรอบสอง",
				"en": "The same user is not allowed to turn off the second result.",
			},
			"create_limit_error": {
				"th": "อัพเดตเลขอั้าล้มเหลว",
				"en": "Update Limitnumber failed",
			},
			"list_limitnumber_error": {
				"th": "รายการเลขอั้นล้มเหลว",
				"en": "List Limitnumber failed",
			},
			"get_drawagent_error": {
				"th": "รายการงวดหวยล้มเหลว",
				"en": "List DrawAgent failed",
			},
			"bet_closetime_error": {
				"th": "หวยปิดรับแทงแล้ว กรุณาเลือกงวดใหม่",
				"en": "Draw closed",
			},
			"uploadimage_fail": {
				"th": "มีปัญหากับไฟล์รูปภาพ",
				"en": "There is a problem with the image file.",
			},
			"char_more_exceeded": {
				"th": "ตัวอักษรเกิน 15 ตัวอักษร",
				"en": "Nickname More than 15 characters",
			},
			"bet_credit_enough": {
				"th": "ยอดเงินในการเดิมพันไม่เพียงพอ",
				"en": "Credit balance is not enough",
			},
			"bet_jobcheck_error": {
				"th": "มีรายการแทงค้าง กรุณาลองใหม่อีกครั้ง",
				"en": "Bet processing, Please try again.",
			},
			"bet_detail_zodiac": {
				"th": "เกิดข้อผิดพลาด กรุณาลองใหม่อีกครั้ง",
				"en": "Error Please try again.",
			},
			"update_favorite_lotto": {
				"th": "ไม่สามารถเพิ่มหวยที่ชอบได้",
				"en": "can't add favorite lotto",
			},
			"add_favorite_number": {
				"th": "ไม่สามารถเพิ่มเลขที่ชอบได้",
				"en": "can't add favorite number",
			},
			"add_yeekeepredictnumber": {
				"th": "ไม่สามารถยิงเลขได้",
				"en": "can't add predict yeekeenumber",
			},
			"add_yeekeepredictnumber_cooldown": {
				"th": "โปรดรอสักครู๋สำหรับการยิงเลข",
				"en": "wait a minute for  yeekee predict",
			},
			"list_yeekeepredictnumber": {
				"th": "ไม่สามารถแสดงรายการได้",
				"en": "can't list predict yeekeenumber",
			},
			"create_drawagent_error_invalid_lotto": {
				"th": "ไม่มีประเภทหวยนี้ในระบบ",
				"en": "invalid lottotype",
			},
			"create_drawagent_error_lotto_pending": {
				"th": "มีรอบของหวยนี้ที่ยังไม่ได้ปิด",
				"en": "this lotto has pending",
			},
			"create_drawagent_error_lotto_duplicate": {
				"th": "หวยรอบนี้มีรอบที่เปิดอยู่แล้ว",
				"en": "this round has dulpicate",
			},
			"create_drawagent_error_lotto_match": {
				"th": "วันที่ของหวยกับเวลาปิดไม่ตรงกัน",
				"en": "datetime and closetime not match",
			},
			"bet_cancel_fail": {
				"th": "ไม่สามารถยกเลิกได้",
				"en": "can't cancel this ticket",
			},
			"bet_cancel_timeout": {
				"th": "ไม่สามารถยกเลิกได้ เนื่องจากหมดเวลาการยกเลิก",
				"en": "can't cancel this ticket because close time",
			},
		},
	}
	return
}
