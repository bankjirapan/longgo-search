package search

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Document struct {
	Contents []string `json:"contents"`
	URL      string   `json:"url"`
}

func Search() {

	data := `[
        {
            "contents": [
                "h1: การหาระยะห่างพิกัด",
                "p:  ตัวอย่างนี้จะแสดงวิธีการหาระยะห่างระหว่างพิกัด 2 จุด "
            ],
            "url": "http://localhost:3000/docs3/ios/etc/distance-and-displacement"
        },
        {
            "contents": [
                "h1: เริ่มต้นการใช้งาน Longdo Map API",
                "p:  บทความนี้จะแสดงวิธีการนำแผนที่ Longdo Map ไปใช้บนเว็บไซต์อย่างง่าย โดยนำไปประยุกต์ใช้ในงานของท่านได้ในหลากหลายรูปแบบ ไม่ว่าจะเป็นการนำ API (Application Programming Interface) ไปใช้ เรียกแสดงแผนที่ในเว็บหรือแอปพลิเคชันของท่าน, ติดตั้งผลิตภัณฑ์ Map Server ส่วนตัว ที่เรียกว่า Longdo Box ภายในเครือข่ายของท่าน (ดูหน้าผลิตภัณฑ์) เพื่อแยกการใช้งานแบบเป็นเอกเทศ ไม่ปะปนกับผู้ใช้รายอื่น, เพิ่มไอคอนของธุรกิจของท่านลงในแผนที่ รวมถึงป้องกันไม่ให้เกิดการแก้ไขจากบุคคลภายนอก ",
                "p:  เข้าหน้าเว็บไซต์ map.longdo.com เพื่อสมัครสมาชิกของ Longdo Map ก่อน ",
                "p:  เมื่อทำตามขั้นตอนทั้งหมดเสร็จสิ้น ท่านจะได้รับคีย์ยาว ๆ ในการใช้แพคเกจฟรี เช่น ",
                "p:  หลังจากนักพัฒนาได้สร้างแผนที่พื้นฐานเรียบร้อยแล้ว ลำดับถัดไปจะเป็นการเพิ่มฟังก์ชันต่าง ๆ ที่ต้องการบนแผนที่ โดยหัวข้อด้านล่างนี้ "
            ],
            "url": "http://localhost:3000/docs3/react-native/getting-start"
        }
    ]`

	var documents []Document
	err := json.Unmarshal([]byte(data), &documents)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Phrase to search for
	phrase := "ห่างระหว่างพิกัด"

	// Search for the phrase in the contents of each document
	for _, doc := range documents {
		for _, content := range doc.Contents {
			if strings.Contains(content, phrase) {
				fmt.Println("Found in URL:", doc.URL)
				fmt.Println("Content:", content)
			}
		}
	}

}
