package main_test

import (
	"bytes"
	"net/http"
	"testing"
)

func TestSelectAll(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/tasks/all")
	if err != nil {
		t.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("FAILED Status code is %d, expected %d", resp.StatusCode, 200)
	}
	t.Logf("PASSED Status code is %d", resp.StatusCode)
}

func TestSelectById(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/tasks/id/1")
	if err != nil {
		t.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("FAILED Status code is %d, expected %d", resp.StatusCode, 200)
		return
	}
	t.Logf("PASSED Status code is %d", resp.StatusCode)
}

func TestSelectForToday(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/tasks/today")
	if err != nil {
		t.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("FAILED Status code is %d, expected %d", resp.StatusCode, 200)
		return
	}
	t.Logf("PASSED Status code is %d", resp.StatusCode)
}

func TestSelectForTomorrow(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/tasks/tomorrow")
	if err != nil {
		t.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("FAILED Status code is %d, expected %d", resp.StatusCode, 200)
		return
	}
	t.Logf("PASSED Status code is %d", resp.StatusCode)
}

func TestSelectForThisWeek(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/tasks/thisweek")
	if err != nil {
		t.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("FAILED Status code is %d, expected %d", resp.StatusCode, 200)
		return
	}
	t.Logf("PASSED Status code is %d", resp.StatusCode)
}

func TestSelectByTitle(t *testing.T) {
	{
		jsonstr := []byte(`{"title": "Test","description":"Test","expiry_date":"2023-07-08T23:03:12.633+02:00","complete":1}`)
		_, err := http.Post("http://localhost:3000/tasks/insert", "application/json", bytes.NewBuffer(jsonstr))
		if err != nil {
			t.Error(err)
			return
		}
	}
	resp, err := http.Get("http://localhost:3000/tasks/title/test")
	if err != nil {
		t.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("FAILED Status code is %d, expected %d", resp.StatusCode, 200)
		return
	}
	t.Logf("PASSED Status code is %d", resp.StatusCode)
}

func TestInsert(t *testing.T) {
	jsonstr := []byte(`{"title": "Test","description":"Test","expiry_date":"2023-07-08T23:03:12.633+02:00","complete":1}`)
	resp, err := http.Post("http://localhost:3000/tasks/insert", "application/json", bytes.NewBuffer(jsonstr))
	if err != nil {
		t.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("FAILED Status code is %d, expected %d", resp.StatusCode, 200)
		return
	}
	t.Logf("PASSED Status code is %d", resp.StatusCode)
}
