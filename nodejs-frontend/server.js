import { S3Client } from "@aws-sdk/client-s3";
import express from "express";
import { join, dirname } from "path";
import { fileURLToPath } from "url";

const __dirname = dirname(fileURLToPath(import.meta.url));
const app = express();
const PORT = process.env.PORT ?? 3000;
const API_URL = process.env.API_URL ?? "http://localhost:8080";

const s3 = new S3Client({ region: process.env.AWS_REGION ?? "us-east-1" });

app.use(express.static(join(__dirname, "public")));

app.get("/health", (_req, res) => {
  res.json({ status: "healthy" });
});

app.get("/api-url", (_req, res) => {
  res.json({ url: API_URL });
});

app.get("/aws", (_req, res) => {
  res.json({ sdk: "@aws-sdk/client-s3", region: s3.config.region ?? "us-east-1" });
});

app.listen(PORT, () => {
  console.log(`Frontend listening on port ${PORT}`);
});
