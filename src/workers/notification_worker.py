import json
import os
from kafka import KafkaConsumer
from pyfcm import FCMNotification
import smtplib
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart
from dotenv import load_dotenv

# Load environment variables
load_dotenv()

class NotificationWorker:
    def __init__(self):
        # Kafka configuration
        self.consumer = KafkaConsumer(
            'notifications',
            bootstrap_servers=os.getenv('KAFKA_BOOTSTRAP_SERVERS', 'localhost:9092'),
            group_id='notification_worker_group',
            auto_offset_reset='earliest',
            enable_auto_commit=True,
            value_deserializer=lambda x: json.loads(x.decode('utf-8'))
        )

        # FCM configuration for Android push notifications
        self.push_service = FCMNotification(api_key=os.getenv('FCM_API_KEY'))

        # Email configuration
        self.smtp_server = os.getenv('SMTP_SERVER', 'smtp.gmail.com')
        self.smtp_port = int(os.getenv('SMTP_PORT', '587'))
        self.smtp_username = os.getenv('SMTP_USERNAME')
        self.smtp_password = os.getenv('SMTP_PASSWORD')

    def send_push_notification(self, registration_token, title, message):
        """Send push notification to Android device"""
        try:
            result = self.push_service.notify_single_device(
                registration_id=registration_token,
                message_title=title,
                message_body=message
            )
            print(f"Push notification sent: {result}")
            return True
        except Exception as e:
            print(f"Error sending push notification: {e}")
            return False

    def send_email(self, to_email, subject, body):
        """Send email notification"""
        try:
            msg = MIMEMultipart()
            msg['From'] = self.smtp_username
            msg['To'] = to_email
            msg['Subject'] = subject

            msg.attach(MIMEText(body, 'plain'))

            with smtplib.SMTP(self.smtp_server, self.smtp_port) as server:
                server.starttls()
                server.login(self.smtp_username, self.smtp_password)
                server.send_message(msg)

            print(f"Email sent to {to_email}")
            return True
        except Exception as e:
            print(f"Error sending email: {e}")
            return False

    def process_notification(self, notification):
        """Process a notification and send it via appropriate channel"""
        notification_type = notification.get('type')
        if notification_type == 'email':
            return self.send_email(
                notification['target'],
                notification['title'],
                notification['message']
            )
        elif notification_type == 'push':
            return self.send_push_notification(
                notification['target'],
                notification['title'],
                notification['message']
            )
        else:
            print(f"Unknown notification type: {notification_type}")
            return False

    def run(self):
        """Main loop to consume and process notifications"""
        print("Starting notification worker...")
        try:
            for message in self.consumer:
                notification = message.value
                print(f"Processing notification: {notification}")
                success = self.process_notification(notification)
                if success:
                    print("Notification processed successfully")
                else:
                    print("Failed to process notification")
        except Exception as e:
            print(f"Error in worker: {e}")
            self.consumer.close()

if __name__ == "__main__":
    worker = NotificationWorker()
    worker.run() 