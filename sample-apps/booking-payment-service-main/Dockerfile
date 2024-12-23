FROM python:3.9-slim

RUN pip install flask

WORKDIR /app
COPY booking_payment_service.py /app/
EXPOSE 5003
CMD ["python", "booking_payment_service.py"]
