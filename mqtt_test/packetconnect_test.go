package mqtt_test

import (
	"mqtt"
	"testing"
)

func TestPacketConnect(t *testing.T) {
	input := []byte{0x10, 0x0D, 0x00, 0x04, 0x4D, 0x51, 0x54, 0x54, 0x04, 0x00, 0x00, 0x10, 0x00, 0x01, 0xFF}

	pkt := mqtt.NewPacketConnect()
	if err := pkt.Parse(input); err != nil {
		t.Errorf(err.Error())
		return
	}

	output := pkt.Bytes()
	for i := 0; i < len(input); i++ {
		if input[i] != output[i] {
			t.Errorf("Mismatch %02x vs %02x\n", input[i], output[i])
		}
	}

	invalids := [][]byte{{0x10},
		{0x10, 0x02},
		{0x10, 0x02, 0x01},
		{0x10, 0x02, 0x01, 0xF0, 0x2f},
		{0xF0, 0x0D, 0x00, 0x04, 0x4D, 0x51, 0x54, 0x54, 0x04, 0x00, 0x00, 0x10, 0x00, 0x01, 0xFF},
		{0x12, 0x0D, 0x00, 0x04, 0x4D, 0x51, 0x54, 0x54, 0x04, 0x00, 0x00, 0x10, 0x00, 0x01, 0xFF},
		{0x10, 0x0D, 0x00, 0x04, 0x4D, 0x52, 0x54, 0x54, 0x04, 0x00, 0x00, 0x10, 0x00, 0x01, 0xFF},
		{0x10, 0x0C, 0x00, 0x04, 0x4D, 0x54, 0x54, 0x04, 0x00, 0x00, 0x10, 0x00, 0x01, 0xFF},
		{0x10, 0x0D, 0x00, 0x04, 0x4D, 0x51, 0x54, 0x54, 0x03, 0x00, 0x00, 0x10, 0x00, 0x01, 0xFF},
		{0x10, 0x0D, 0x00, 0x04, 0x4D, 0x51, 0x54, 0x54, 0x04, 0x01, 0x00, 0x10, 0x00, 0x01, 0xFF},
		{0x10, 0x0D, 0x00, 0x04, 0x4D, 0x51, 0x54, 0x54, 0x04, 0x1C, 0x00, 0x10, 0x00, 0x01, 0xFF},
		{0x10, 0x0D, 0x00, 0x04, 0x4D, 0x51, 0x54, 0x54, 0x04, 0x08, 0x00, 0x10, 0x00, 0x01, 0xFF},
		{0x10, 0x0D, 0x00, 0x04, 0x4D, 0x51, 0x54, 0x54, 0x04, 0x20, 0x00, 0x10, 0x00, 0x01, 0xFF}}
	for i := 0; i < len(invalids); i++ {
		if err := pkt.Parse(invalids[i]); err != nil {
			t.Logf(err.Error())
		} else {
			t.Logf("%v", invalids[i])
		}
	}
}
