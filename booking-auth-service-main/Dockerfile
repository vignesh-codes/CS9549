FROM python:3.9-slim

RUN pip install flask

WORKDIR /app
COPY booking_auth_service.py /app/
EXPOSE 5001
CMD ["python", "booking_auth_service.py"]
