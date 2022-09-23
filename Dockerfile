FROM python:3

COPY src /opt/src
WORKDIR /opt/src

RUN pip3 install -r requirements.txt

CMD ["./main.py"]
