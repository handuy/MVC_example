package helper

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func DecodeDataFromJsonFile(f *os.File, data interface{}) error {
	jsonParser := jsoniter.NewDecoder(f)
	err := jsonParser.Decode(&data)
	if err != nil {
		return err
	}

	return nil
}

func EqualIntSlice(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func CheckStringElementInSlice(list []string, str string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}
	return false
}

type MenuItems struct {
	Text string   `json:"text"`
	Link string   `json:"link"`
	Sub  []string `json:"sub"`
}

type FooterItems struct {
	Title string `json:"title"`
	Link  string `json:"link"`
	Type  string `json:"type"`
}

type FooterCol struct {
	ColumnName string        `json:"column_name"`
	Items      []FooterItems `json:"items"`
}

type widget struct {
	WidgetType string        `json:"widget_type"`
	Code       template.HTML `json:"code"`
}

type Layout struct {
	LayoutType      int32    `json:"layout_type"`
	Widgets         []widget `json:"widgets"`
	Incontainer     bool     `json:"in_container"`
	BackgroundColor string   `json:"background_color"`
	BackgroundImage string   `json:"background_image"`
}

func ListLayout() []Layout {
	jsonFile, err := os.Open("layout.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var arr []Layout
	err = json.Unmarshal([]byte([]byte(byteValue)), &arr)
	if err != nil {
		fmt.Println(err)
	}
	return arr
}

func ListMenu() []MenuItems {
	jsonFile, err := os.Open("menu.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var arr []MenuItems
	json.Unmarshal([]byte([]byte(byteValue)), &arr)

	return arr
}

// Lưu lại lỗi
func HandlerError(userId, path string, level string, message error) {
	var log = logrus.New()

	log.Out = os.Stdout

	file, err := os.OpenFile("log/logging.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err == nil {
		log.Out = file
	} else {
		panic(err)
	}

	if level == "error" {
		log.WithFields(logrus.Fields{
			"user_id": userId,
			"path":    path,
		}).Error(message)
	} else if level == "info" {
		log.WithFields(logrus.Fields{
			"user_id": userId,
			"path":    path,
		}).Info(message)
	} else if level == "debug" {
		log.WithFields(logrus.Fields{
			"user_id": userId,
			"path":    path,
		}).Debug(message)
	} else if level == "warn" {
		log.WithFields(logrus.Fields{
			"user_id": userId,
			"path":    path,
		}).Warn(message)
	}

}

// CheckFormData Kiểm tra các tham số có rỗng hay không, trả về mảng các tham số rỗng
func CheckFormData(ParseFormData map[string][]string, params ...string) []string {
	var result []string
	for _, param := range params {
		if len(ParseFormData[param]) < 1 || ParseFormData[param][0] == "" {
			result = append(result, param)
		}
	}
	return result
}
