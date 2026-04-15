FROM mcr.microsoft.com/dotnet/sdk:9.0@sha256:e04aef5e32fe25cea4641ecd2f2d8d7eca009f5d937bd545250f5ee5f35cf3b5 AS build
WORKDIR /src

COPY dotnet-app.csproj .
RUN dotnet restore

COPY . .
RUN dotnet publish -c Release -o /app/publish

FROM mcr.microsoft.com/dotnet/aspnet:9.0@sha256:7799364313c08f464193c93a9db935d86697b2b25c96c89fb2b629bbf9877016 AS runtime
WORKDIR /app

COPY --from=build /app/publish .

EXPOSE 8080
ENV ASPNETCORE_URLS=http://+:8080

ENTRYPOINT ["dotnet", "dotnet-app.dll"]
