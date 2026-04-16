using Amazon.S3;

var builder = WebApplication.CreateBuilder(args);
builder.Services.AddAWSService<IAmazonS3>();
var app = builder.Build();

app.MapGet("/", () => new { message = "Hello from .NET demo app!", status = "ok" });
app.MapGet("/health", () => new { status = "healthy" });
app.MapGet("/aws", () => new { sdk = "AWSSDK.S3", service = "s3" });

app.Run();
