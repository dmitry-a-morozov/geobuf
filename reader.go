package geobuf_new

import (
	"github.com/murphy214/geobuf_new/geobuf_raw"
	"os"
	"bufio"
	//"io"
	"bytes"
	"fmt"
	"github.com/paulmach/go.geojson"
    "github.com/murphy214/protoscan"
)

// protobuf scanner implementation
type Reader struct {
	FileBool bool
	Reader *protoscan.ProtobufScanner
	Filename string
	Feature_Bytes []byte
}

// creates a reader for a byte array
func ReaderBytes(bytevals []byte) *Reader {
	return &Reader{Reader:protoscan.NewProtobufScanner(bytes.NewReader(bytevals)),FileBool:false}
}

// creates a reader for file
func ReaderFile(filename string) *Reader {
	file,err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	return &Reader{
			Reader:protoscan.NewProtobufScanner(bufio.NewReader(file)),
			Filename:filename,
			FileBool:true,
		}
}

// alias for the Scan method on reader
// next is a little more expressive
func (reader *Reader) Next() bool {
	return reader.Reader.Scan()
}

// alias for the Protobuf() method 
// again more expressive for our use case
func (reader *Reader) Bytes() []byte {
	return reader.Reader.Protobuf()
}

// alias for the Protobuf() method 
// again more expressive for our use case
func (reader *Reader) Feature() *geojson.Feature {
	return geobuf_raw.ReadFeature(reader.Bytes())
}


func ReadFeature(bytevals []byte) *geojson.Feature {
	return geobuf_raw.ReadFeature(bytevals)
}


