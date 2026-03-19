package entity

import "testing"

func TestNewTemperature(t *testing.T) {
	temp := NewTemperature(28.5)
	if temp.C != 28.5 {
		t.Errorf("Expected C to be 28.5, got %f", temp.C)
	}
	if temp.F != 83.3 {
		t.Errorf("Expected F to be 83.3, got %f", temp.F)
	}
	if temp.K != 301.5 {
		t.Errorf("Expected K to be 301.5, got %f", temp.K)
	}
}

func TestIsValidCEP(t *testing.T) {
	if !IsValidCEP("12345678") {
		t.Error("Expected 12345678 to be valid")
	}
	if IsValidCEP("1234567") {
		t.Error("Expected 1234567 to be invalid")
	}
	if IsValidCEP("123456789") {
		t.Error("Expected 123456789 to be invalid")
	}
	if IsValidCEP("1234567A") {
		t.Error("Expected 1234567A to be invalid")
	}
	if IsValidCEP("1234-567") {
		t.Error("Expected 1234-567 to be invalid")
	}
}
