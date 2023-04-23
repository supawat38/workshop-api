package controllers

import (
	"app/platform/storage"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/juunini/simple-go-line-notify/notify"
)

func ReadUserIP(c *fiber.Ctx) string {
	IPAddress := ""
	if c.Get("X-Forwarded-For") != "" {
		IPAddress = strings.Split(c.Get("X-Forwarded-For"), ",")[0]
	} else {
		IPAddress = GetLocalIP()
	}
	return IPAddress
}

// อัพโหลดรูปภาพ
func UploadFileImageS3(fileImage *multipart.FileHeader, tag string, path_import_name string) (path_image string, err error) {
	file_type := fileImage.Header.Get("content-type")

	f, err2 := fileImage.Open()
	if err2 != nil {
		return "", err2
	}
	defer f.Close()

	buffer, err3 := ioutil.ReadAll(f)
	if err3 != nil {
		return "", err3
	}
	//reader := bytes.NewReader(file_byte)
	reader := bytes.NewReader(buffer)
	// generate new uuid for image name
	uniqueId := uuid.New()
	// remove "- from imageName"
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	// extract image extension from original file filename
	// fileExt := strings.Split(fileImage.Filename, ".")[1]
	fileExt := filepath.Ext(fileImage.Filename)

	// generate image from filename and extension
	image_name := fmt.Sprintf("%s%s", filename, fileExt)
	storage.InitS3AWS()
	s, err4 := storage.ConnectToS3()
	if err4 != nil {
		return "", err4
	}
	// tag := "ads"
	path_s3 := path_import_name + image_name

	_, s3err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(storage.S3_BUCKET),
		Key:           aws.String(path_s3),
		Body:          reader,
		ContentLength: aws.Int64(fileImage.Size),
		ContentType:   aws.String(file_type),
		Tagging:       &tag,
	})
	if s3err != nil {
		return "", s3err
	}
	return path_s3, nil
}

// อัพโหลดไฟล์
func UploadFileToS3(Filename string) (status bool) {

	id := os.Getenv("AWS_ACCESS_KEY_ID")
	key := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("BUCKET_NAME")
	token := ""

	creds := credentials.NewStaticCredentials(id, key, token)
	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
	}
	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	path := "tempsitemap"
	filename := Filename
	file, err := os.Open(path + "/" + filename)
	if err != nil {
		fmt.Printf("err opening file: %s", err)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)

	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	pathSave := "/seo/" + filename
	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(pathSave),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}
	A, err := svc.PutObject(params)
	fmt.Println(A)

	if err != nil {
		return false
	}

	return true
}

// ลบรูปภาพ
func DeleteImageS3(path string) (err error) {
	storage.InitS3AWS()
	s, err := storage.ConnectToS3()
	if err != nil {
		return err
	}

	path_s3 := path

	rersp, err := s3.New(s).DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(storage.S3_BUCKET),
		Key:    aws.String(path_s3),
	})
	fmt.Println(rersp, err)
	return
}

// แจ้งเตือนผ่าน Line
func SendLineNotify(sendtext string, totken string) (status bool) {

	accessToken := totken
	//ข้อความที่จะส่งไป
	message := sendtext

	if err := notify.SendText(accessToken, message); err != nil {
		status = false
		return
	}
	status = true
	return
}

// สร้าง รหัส secret key
func RandStringBytes(n int) string {
	// rand.Seed(time.Now().UnixNano())
	// letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@$.()_+-!*"
	// b := make([]byte, n)
	// for i := range b {
	// 	b[i] = letterBytes[rand.Intn(len(letterBytes))]
	// }
	// return string(b)
	return "5"
}

// เข้ารหัส (ของใหม่ การ encrypt จะไม่ถูกเปลี่ยน)
func EncryptGenV2(stringToEncrypt string) (encryptedString string) {
	bKey := []byte(os.Getenv("AES_256_KEY"))
	bIV := []byte(os.Getenv("AES_256_IV"))
	bPlaintext := PKCS5PaddingV2([]byte(stringToEncrypt), aes.BlockSize, len(stringToEncrypt))
	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}

// เข้ารหัส สำหรับ V2
func PKCS5PaddingV2(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// ถอดรหัส (ของใหม่ การ Decrypt)
func DecryptGenV2(encryptedString string) (decryptedString string) {
	bKey := []byte(os.Getenv("AES_256_KEY"))
	bIV := []byte(os.Getenv("AES_256_IV"))
	cipherTextDecoded, err := hex.DecodeString(encryptedString)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	return string(cipherTextDecoded)
}

// Get Local IP
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// Get ค่าต่างๆ หลังจาก login
func GetTokenFromSession(c *fiber.Ctx) (token string, err error) {
	authBearer := string(c.Request().Header.Peek("Authorization"))
	splitToken := strings.Split(authBearer, "Bearer ")
	tokenString := strings.TrimSpace(splitToken[1])

	claims := jwt.MapClaims{}

	_, er := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		secret := os.Getenv("JWT_SECRET_KEY")
		return []byte(secret), nil
	})
	if er != nil {
		err = er
		return
	}

	if claims["id"] == nil {
		err = fmt.Errorf("user not found")
		return
	}

	token = tokenString
	return
}

// แปลง Json เป็น Map
func JsonToMap(b []byte) (m map[string]interface{}) {
	json.Unmarshal(b, &m)
	return
}

func JsonToArray(b []byte) (i []interface{}) {
	json.Unmarshal(b, &i)
	return
}

// แปลง String เป็น Int
func StrToInt(str string) (i int) {
	i, _ = strconv.Atoi(str)
	return
}

// แปลง Int เป็น String
func IntToStr(i int) (str string) {
	str = strconv.Itoa(i)
	return
}

// แปลง Map เป็น  Json
func MapToJson(m interface{}) (b []byte) {
	b, _ = json.Marshal(m)
	return
}

func InterfaceToMap(b interface{}) (m map[string]interface{}) {
	bytedata, _ := json.Marshal(b)
	json.Unmarshal(bytedata, &m)

	return
}

func GetDayofWeek(day string) (dayofweek int) {
	switch day {
	case "sunday":
		dayofweek = 0
		break
	case "monday":
		dayofweek = 1
		break
	case "tuesday":
		dayofweek = 2
		break
	case "wednesday":
		dayofweek = 3
		break
	case "thursday":
		dayofweek = 4
		break
	case "friday":
		dayofweek = 5
		break
	case "saturday":
		dayofweek = 6
		break

	}
	return dayofweek
}

func GetDateofWeek(dayofweek int) (date string) {
	week := map[int]string{
		0: "sunday",
		1: "monday",
		2: "tuesday",
		3: "wednesday",
		4: "thursday",
		5: "firday",
		6: "saturday",
	}
	return week[dayofweek]
}

func CheckArray(a []int, i interface{}) (found bool) {
	for index := range a {
		if a[index] == i {
			return true
		}
	}
	return
}

func CheckArrayInt64(a []int64, i interface{}) (found bool) {
	for index := range a {
		if a[index] == i {
			return true
		}
	}
	return
}

func CheckArrayString(a []string, i interface{}) (found bool) {
	for index := range a {
		if a[index] == i {
			return true
		}
	}
	return
}

// ปัดทศนิยม
func FloatPrecision(num float64, precision int) float64 {
	p := math.Pow10(precision)
	value := float64(int(num*p)) / p
	return value
}

// รับค่าได้เฉพาะอักษร
func Is_alphanum(word string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9ก-๛]*$`).MatchString(word)
}

// เช็คเวลาว่าดิฟกันเท่าไร
func DateDiff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

// รับค่าได้ตัวเลขที่เป็น str
func Is_StrNumber(word string) bool {
	return regexp.MustCompile(`^[0-9]*$`).MatchString(word)
}

// ตัดคำ
func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}

// ## Open sort

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

// Main for sort
func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// ## End sort

// ตัวที่ซ้ำใน Array String
func CutStringDupicateArray(key []string) (clean []string) {
	for _, value := range key {
		if !stringInSlice(value, clean) {
			clean = append(clean, value)
		}
	}
	return
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// CalTotalPage
func CalTotalPage(Total int, Limit int) (TotalPage int) {
	TotalPage = (Total / Limit)
	mod := (Total % Limit)
	if mod != 0 {
		TotalPage += 1
	}

	if TotalPage == 0 {
		TotalPage = 1
	}
	return
}
