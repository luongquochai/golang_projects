package main

import (
	"os"
	"testing"
)

func TestTaoBaiMoi(t *testing.T) {
	nhieuLaBai := taoBaiMoi()
	if len(nhieuLaBai) != 12 {
		t.Errorf("Expect: 12, Actual: %v", len(nhieuLaBai))
	}
}

func TestLuuVaTaoBaiTuFile(t *testing.T) {
	nhieuLaBai := taoBaiMoi()
	err := nhieuLaBai.luuFile("_testBai")

	if err != nil {
		t.Errorf("Luu file khong thanh cong ! (err=%v)", err)
	}

	nhieuLaBai, err = taoBaiTuFile("_testBai")
	if err != nil {
		t.Errorf("Tai file khong thanh cong ! (err=%v)", err)
	}

	if len(nhieuLaBai) != 12 {
		t.Errorf("Expect: 12, Actual: %v", len(nhieuLaBai))
	}

	os.Remove("_testBai")
}
