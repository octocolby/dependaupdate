FROM python:3.14-slim@sha256:bc389f7dfcb21413e72a28f491985326994795e34d2b86c8ae2f417b4e7818aa

COPY --from=ghcr.io/astral-sh/uv:0.5.0@sha256:0e0fb77970aceaa106c1fee7103ec3e4885f6c4289e32a596123af06d2e9db9d /uv /uvx /bin/

WORKDIR /app

COPY pyproject.toml uv.lock ./
RUN uv sync --frozen --no-dev

COPY app.py .

EXPOSE 8080

ENV PATH="/app/.venv/bin:$PATH"
CMD ["gunicorn", "--bind", "0.0.0.0:8080", "app:app"]
