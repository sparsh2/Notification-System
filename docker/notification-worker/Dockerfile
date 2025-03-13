FROM python:3.9-slim

WORKDIR /app

# Copy requirements first to leverage Docker cache
COPY ../../src/workers/requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy the worker code
COPY ../../src/workers/notification_worker.py .

# Create .env file with default values (these should be overridden in production)
RUN echo "KAFKA_BOOTSTRAP_SERVERS=kafka:9092\n\
FCM_API_KEY=your_fcm_api_key\n\
SMTP_SERVER=smtp.gmail.com\n\
SMTP_PORT=587\n\
SMTP_USERNAME=your_email@gmail.com\n\
SMTP_PASSWORD=your_app_password" > .env

CMD ["python", "notification_worker.py"]