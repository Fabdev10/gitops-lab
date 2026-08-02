from fastapi import FastAPI, HTTPException

app = FastAPI(title="api-python", version="1.0.0")

POSTS = [
    {"id": 1, "title": "Hello from FastAPI", "content": "A simple post served by the Python service."},
    {"id": 2, "title": "DevOps learning", "content": "This sample payload demonstrates a second fake post."},
]


@app.get("/health")
def health() -> dict[str, str]:
    return {"status": "ok", "service": "api-python"}


@app.get("/posts")
def list_posts() -> list[dict[str, object]]:
    return POSTS


@app.get("/posts/{post_id}")
def get_post(post_id: int) -> dict[str, object]:
    post = next((item for item in POSTS if item["id"] == post_id), None)
    if post is None:
        raise HTTPException(status_code=404, detail="Post not found")
    return post
