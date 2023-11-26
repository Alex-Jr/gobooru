FROM python:3.9.18-alpine3.18

WORKDIR /app

RUN apk update && \
    apk add --no-cache libpq-dev \
    build-base libffi libheif-dev libde265-dev

COPY . .

RUN pip install -r requirements.txt

EXPOSE 8000

CMD [ "python", "-u", "main.py" ]
