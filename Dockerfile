FROM python:3.9-slim

COPY src /opt/src
WORKDIR /opt/src

RUN pip3 install -r requirements.txt

CMD ["./main.py"]
