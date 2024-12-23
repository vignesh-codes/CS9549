FROM python:3.9-slim


RUN pip install flask
WORKDIR /app
COPY booking_order_service.py /app/
EXPOSE 5002
CMD ["python", "booking_order_service.py"]
