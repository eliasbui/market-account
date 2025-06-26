# Build stage
FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build

# Set working directory
WORKDIR /src

# Copy solution and project files
COPY ProductService.sln .
COPY Domain/ProductService.Domain.csproj Domain/
COPY Application/ProductService.Application.csproj Application/
COPY Infrastructure/ProductService.Infrastructure.csproj Infrastructure/
COPY WebApi/ProductService.WebApi.csproj WebApi/

# Restore dependencies
RUN dotnet restore ProductService.sln

# Copy all source code
COPY . .

# Build the application
WORKDIR /src/WebApi
RUN dotnet build ProductService.WebApi.csproj -c Release -o /app/build

# Publish stage
FROM build AS publish
RUN dotnet publish ProductService.WebApi.csproj -c Release -o /app/publish /p:UseAppHost=false

# Final stage
FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS final

# Install curl for health checks
RUN apt-get update && apt-get install -y curl && rm -rf /var/lib/apt/lists/*

# Create non-root user
RUN groupadd -r appgroup && useradd -r -g appgroup appuser

# Set working directory
WORKDIR /app

# Copy published application
COPY --from=publish /app/publish .

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Use non-root user
USER appuser

# Expose port
EXPOSE 8002

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD curl -f http://localhost:8002/health || exit 1

# Command to run
ENTRYPOINT ["dotnet", "ProductService.WebApi.dll"] 