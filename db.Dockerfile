# PostgreSQL Docker image following best practices
# Reference: https://sliplane.io/blog/best-practices-for-postgres-in-docker
FROM postgres:15.11-alpine

# Default environment variables (use Docker secrets in production)
# For production, use POSTGRES_PASSWORD_FILE instead of POSTGRES_PASSWORD
ENV POSTGRES_DB=hrms
ENV POSTGRES_USER=postgres

# Enable data checksums for integrity verification
ENV POSTGRES_INITDB_ARGS="--data-checksums"

# Copy initialization scripts
# Scripts in /docker-entrypoint-initdb.d/ run automatically on first startup
COPY scripts/db/ /docker-entrypoint-initdb.d/

# Declare volume for data persistence
# Data stored here survives container restarts when using named volumes
VOLUME /var/lib/postgresql/data

# Health check to ensure the database is ready
# Runs every 30 seconds, times out after 5 seconds, retries 3 times
HEALTHCHECK --interval=30s --timeout=5s --start-period=30s --retries=3 \
    CMD pg_isready -U ${POSTGRES_USER} -d ${POSTGRES_DB} || exit 1

EXPOSE 5432

CMD ["postgres"]
