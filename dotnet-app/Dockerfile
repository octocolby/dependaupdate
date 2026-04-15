FROM mcr.microsoft.com/dotnet/sdk:9.0@sha256:e04aef5e32fe25cea4641ecd2f2d8d7eca009f5d937bd545250f5ee5f35cf3b5 AS build
WORKDIR /src

COPY dotnet-app.csproj .
RUN dotnet restore

COPY . .
RUN dotnet publish -c Release -o /app/publish

FROM mcr.microsoft.com/dotnet/aspnet:10.0@sha256:9271888dde1f4408dfb7ce75bc0f513c903d20ff2f1287ab153b641d4588ec7d AS runtime
WORKDIR /app

COPY --from=build /app/publish .

EXPOSE 8080
ENV ASPNETCORE_URLS=http://+:8080

ENTRYPOINT ["dotnet", "dotnet-app.dll"]
