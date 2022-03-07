
FROM python:3.7-buster

RUN pip install --upgrade pip

WORKDIR /home/worker

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

COPY server.py server.py

CMD ["python", "server.py"]
