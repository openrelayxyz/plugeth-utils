package core

import (
  "encoding/json"
  "testing"
)

func TestJSONUnMarshalAddress(t *testing.T) {
  addr := Address{}
  if err := json.Unmarshal([]byte(`"0x01"`), &addr); err != nil {
    t.Fatalf(err.Error())
  }
  if addr != HexToAddress("0x01") {
    t.Errorf("Unexpected address, wanted %v got %v", HexToAddress("0x01"), addr)
  }
}

func TestJSONUnMarshalHash(t *testing.T) {
  addr := Hash{}
  if err := json.Unmarshal([]byte(`"0x01"`), &addr); err != nil {
    t.Fatalf(err.Error())
  }
  if addr != HexToHash("0x01") {
    t.Errorf("Unexpected address, wanted %v got %v", HexToHash("0x01"), addr)
  }
}
