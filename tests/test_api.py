import requests

BASE_URL = "http://localhost:8080"

def test_health_check():
    response = requests.get(f"{BASE_URL}/health")
    assert response.status_code == 200
    assert response.text == "OK"

def test_get_balance():
    response = requests.get(f"{BASE_URL}/api/v1/balance")
    assert response.status_code == 200
    data = response.json()
    assert "balance" in data
    assert data["balance"] == 1000.50
