import express from "express";
import { join, dirname } from "path";
import { fileURLToPath } from "url";

const __dirname = dirname(fileURLToPath(import.meta.url));
const app = express();
const PORT = process.env.PORT ?? 3000;
const API_URL = process.env.API_URL ?? "http://localhost:8080";

app.use(express.static(join(__dirname, "public")));

app.get("/health", (_req, res) => {
  res.json({ status: "healthy" });
});

app.get("/api-url", (_req, res) => {
  res.json({ url: API_URL });
});

app.listen(PORT, () => {
  console.log(`Frontend listening on port ${PORT}`);
});
