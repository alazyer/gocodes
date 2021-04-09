package main

import (
	"fmt"
	"bytes"
	"io"
	"encoding/json"
	
	"compress/zlib"
)
var data = `[{
        "time":1617865851713072
    }
]`

type LogData struct {
	CreatedAt    int64             `json:"time"`
}

func main() {
	var s []LogData
	
	err := json.Unmarshal([]byte(data), &s)
	if err != nil {
		fmt.Println("Unmarshal error", err)
	}

	buf := bytes.NewBuffer(nil)
	writer := zlib.NewWriter(buf)
	
	enc := json.NewEncoder(writer)
	enc.Encode(s)
	
	// b, err := json.Marshal(s)
	// if err != nil {
	//	fmt.Println("Marshal error", err)
	// }
	// writer.Write(b)
	writer.Flush()

	b := append(buf.Bytes(), []byte("fasdfasfasdfsafffff")...)

	r := bytes.NewBuffer(b)
	reader, err := zlib.NewReader(r)
	// reader, err := zlib.NewReader(buf)
	var out bytes.Buffer
	io.Copy(&out, reader)

	var s1 []LogData
	err = json.Unmarshal(out.Bytes(), &s1)
	if err != nil {
		fmt.Println("Zlib data Unmarshal error", err)
	}
	for _, log1 := range s1 {
		fmt.Printf("Log: %+v\n", log1)
	}
	fmt.Println("Zlib Logs Count: ", len(s1))	
}
