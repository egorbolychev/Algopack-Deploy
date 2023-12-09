FROM python:3.11

WORKDIR /app

COPY ./api .

RUN pip3 install --upgrade pip
RUN pip3 install -r ./requirements.txt

CMD ["python3", "main.py"]