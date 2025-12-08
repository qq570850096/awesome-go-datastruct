package structdemo

import (
	"encoding/json"
	"testing"
)

func TestValueVsPointerReceiver(t *testing.T) {
	u := User{Name: "Alice", Age: 20}
	u.RenameValue("Bob")
	if u.Name != "Alice" {
		t.Fatalf("RenameValue should not change original, got %q", u.Name)
	}
	u.RenamePointer("Bob")
	if u.Name != "Bob" {
		t.Fatalf("RenamePointer should change original, got %q", u.Name)
	}

	ageUser := User{Name: "Tom", Age: 30}
	ChangeUserValue(ageUser)
	if ageUser.Age != 30 {
		t.Fatalf("ChangeUserValue should not change Age, got %d", ageUser.Age)
	}
	ChangeUserPointer(&ageUser)
	if ageUser.Age != 31 {
		t.Fatalf("ChangeUserPointer should change Age, got %d", ageUser.Age)
	}
}

func TestEmbeddingLogger(t *testing.T) {
	svc := NewService("auth")
	msg := svc.Info("ok")
	if msg != "auth: ok" {
		t.Fatalf("unexpected Info result: %q", msg)
	}
	// 直接访问嵌入字段的方法。
	if svc.Log("ping") != "auth: ping" {
		t.Fatalf("embedded Log not working")
	}
}

func TestJSONTags(t *testing.T) {
	a := Account{ID: 1, Name: "", Token: "secret"}
	data, err := EncodeAccount(a)
	if err != nil {
		t.Fatalf("EncodeAccount error: %v", err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("unmarshal encoded account error: %v", err)
	}
	if len(m) != 1 || m["id"] != float64(1) {
		t.Fatalf("unexpected json map: %#v", m)
	}
	if _, ok := m["name"]; ok {
		t.Fatalf("empty name should be omitted")
	}

	a2 := Account{ID: 2, Name: "Bob"}
	data2, err := EncodeAccount(a2)
	if err != nil {
		t.Fatalf("EncodeAccount error: %v", err)
	}
	var m2 map[string]any
	if err := json.Unmarshal(data2, &m2); err != nil {
		t.Fatalf("unmarshal encoded account error: %v", err)
	}
	if m2["name"] != "Bob" {
		t.Fatalf("expected name Bob, got %#v", m2["name"])
	}

	raw := []byte(`{"id":3,"name":"Carol","token":"ignored"}`)
	decoded, err := DecodeAccount(raw)
	if err != nil {
		t.Fatalf("DecodeAccount error: %v", err)
	}
	if decoded.ID != 3 || decoded.Name != "Carol" {
		t.Fatalf("unexpected decoded account: %#v", decoded)
	}
	if decoded.Token != "" {
		t.Fatalf("Token should not be filled from json, got %q", decoded.Token)
	}
}

