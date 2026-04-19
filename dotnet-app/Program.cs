using Amazon.S3;

var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

var s3 = new AmazonS3Client(Amazon.RegionEndpoint.USEast1);

app.MapGet("/", () => new { message = "Hello from .NET demo app!", status = "ok" });
app.MapGet("/health", () => new { status = "healthy" });
app.MapGet("/aws", () => new { sdk = "AWSSDK.S3", version = typeof(AmazonS3Client).Assembly.GetName().Version?.ToString(), service = "s3" });

app.Run();
