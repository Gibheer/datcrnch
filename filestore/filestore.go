package filestore

import "os"
import "time"
import encoding "encoding/binary"

type DataPoint interface {
  Read(file *os.File) error
  Write(file *os.File) error
}

// raw data point
type RawDataPoint struct {
  Timestamp time.Time
  Key string
  Value int32
}

func (d *RawDataPoint) Read(file *os.File) error {
  error := encoding.Read(file, encoding.BigEndian, &d.Value)
  return error
}

func (d *RawDataPoint) Write(file *os.File) error {
  error := encoding.Write(file, encoding.BigEndian, d.Value)
  return error
}

// aggregated data point
type AggregatedDataPoint struct {
  Timestamp time.Time
  Key string
  Values IntValues
}

type IntValues struct {
  Count int32
  Average int32
  Min int32
  Max int32
  Percentile99 int32
}

func (d *AggregatedDataPoint) Read(file *os.File) error {
  error := encoding.Read(file, encoding.BigEndian, &d.Values)
  return error
}

func (d *AggregatedDataPoint) Write(file *os.File) error {
  error := encoding.Write(file, encoding.BigEndian, d.Values)
  return error
}

// file handling
func OpenForWrite(filename string) *os.File {
  f, _ := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0600)
  return f
}

func OpenForRead(filename string) *os.File {
  f, _ := os.Open(filename)
  return f
}
