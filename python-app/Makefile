.PHONY: install run

install:
	uv sync --frozen --no-dev

run:
	uv run gunicorn --bind 0.0.0.0:8080 app:app
