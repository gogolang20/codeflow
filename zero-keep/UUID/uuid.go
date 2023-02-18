package main

import (
	"encoding/base64"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

func main() {
	// UUID2BinData
	dataTobinary := "ab0d21e1-714f-48fe-b97b-87914304f174"
	UUID2BinData(dataTobinary)
	// BinData2UUID
	dataTostring := "Dq3njr8FQ4mWMPLawkg5+w=="
	BinData2UUID(dataTostring)

	// timestamp format
	fmt.Println(time.Unix(1659508624, 0).Format(time.RFC3339))
	fmt.Println(time.Now().Format(time.RFC3339))
}

func UUID2BinData(dataTobinary string) {
	u, err := uuid.FromString(dataTobinary)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(u[:]))
}

func BinData2UUID(dataTostring string) {
	uuidB, err := base64.StdEncoding.DecodeString(dataTostring)
	if err != nil {
		fmt.Println(err)
	}
	uu := uuid.UUID{}
	copy(uu[:], uuidB)
	fmt.Println(uu.String())
}
