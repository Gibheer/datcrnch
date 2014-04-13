package filestore

import (
  "testing"
  "time"
)

func TestRamDataStore(t *testing.T) {
  s := CreateRamDataStore()
  d := RawDataPoint{time.Now(), "foo", 23}

  s.Write(&d)
  r := s.Read()
  if r[0] != d {
    t.Log("Storing or Reading data points does not work!")
    t.Fail()
  }
}
