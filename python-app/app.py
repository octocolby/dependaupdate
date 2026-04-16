import boto3
from botocore.config import Config
from flask import Flask

app = Flask(__name__)

_s3 = boto3.client(
    "s3",
    config=Config(signature_version="unsigned"),
    region_name="us-east-1",
)


@app.route("/")
def hello():
    return {"message": "Hello from Python demo app!", "status": "ok"}


@app.route("/health")
def health():
    return {"status": "healthy"}


@app.route("/aws")
def aws():
    return {"sdk": "boto3", "version": boto3.__version__, "service": "s3"}


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
