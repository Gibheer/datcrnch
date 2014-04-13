package filestore

type DataStore interface {
  Read() []DataPoint
  Write(DataPoint)
}

type RamDataStore struct {
  num_data int
  datapoints []DataPoint
}

func (d *RamDataStore) Read() []DataPoint {
  return d.datapoints[0:d.num_data]
}

func (d *RamDataStore) Write(new_point DataPoint) {
  if len(d.datapoints) < d.num_data - 1 {
    new_array := make([]DataPoint, len(d.datapoints) * 2)
    copy(new_array, d.datapoints)
    d.datapoints = new_array
  }
  d.num_data += 1
  d.datapoints[d.num_data] = new_point
}

func CreateRamDataStore() *RamDataStore {
  return &RamDataStore{0, make([]DataPoint, 16)}
}
