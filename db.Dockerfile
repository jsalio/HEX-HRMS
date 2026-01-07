FROM postgres:15-alpine

# Set default environment variables
ENV POSTGRES_DB=hrms
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres

# Healthcheck to ensure the database is ready
HEALTHCHECK --interval=10s --timeout=5s --retries=5 \
    CMD pg_isready -U ${POSTGRES_USER} -d ${POSTGRES_DB} || exit 1

EXPOSE 5432

CMD ["postgres"]
