package render

import (
	"testing"
)

func TestTemplateData_Add(t *testing.T) {
	data := []struct {
		key, value string
	}{
		{"first", "one"},
		{"second", "two"},
	}
	d := NewTemplateData()
	for _, v := range data {
		d.Add(v.key, v.value)
	}
	if len(d) != 2 {
		t.Error("Expected 2 got %d", len(d))
	}
}

func TestTemplateData_Merge(t *testing.T) {
	data := []struct {
		key, value string
	}{
		{"first", "one"},
		{"second", "two"},
	}
	d1 := NewTemplateData()
	d2 := NewTemplateData()
	for _, v := range data {
		d1.Add(v.key, v.value)
		d2.Add(v.value, v.key)
	}
	d1.Merge(d2)
	if len(d1) != 4 {
		t.Errorf("Expected 4 got %d", len(d1))
	}
	if d1[data[0].value] != data[0].key {
		t.Errorf("Expected %s got %s", data[0].key, d1[data[0].value])
	}
}
