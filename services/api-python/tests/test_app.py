from fastapi.testclient import TestClient

from app.main import app

client = TestClient(app)


def test_health_endpoint() -> None:
    response = client.get("/health")
    assert response.status_code == 200
    assert response.json() == {"status": "ok", "service": "api-python"}


def test_posts_endpoint() -> None:
    response = client.get("/posts")
    assert response.status_code == 200
    payload = response.json()
    assert len(payload) >= 2
    assert payload[0]["id"] == 1
    assert payload[1]["id"] == 2


def test_post_detail_endpoint() -> None:
    response = client.get("/posts/1")
    assert response.status_code == 200
    assert response.json()["title"] == "Hello from FastAPI"


def test_missing_post_returns_404() -> None:
    response = client.get("/posts/999")
    assert response.status_code == 404
    assert response.json()["detail"] == "Post not found"
