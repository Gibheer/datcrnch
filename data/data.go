package data

import "time"

type DataItem struct {
  Timestamp time.Time
  Key string
  Value int32
}

func CreateDataItem(key string, value int32) *DataItem {
  return &DataItem{time.Now(), key, value}
}
