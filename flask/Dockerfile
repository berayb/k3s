FROM python:3.6-alpine

RUN mkdir -p /deploy/app

COPY requirements.txt /deploy/
RUN apk --update add --virtual build-base \
  && apt-get install -y --no-install-recommends gcc and-build-dependencies \
  && python3 -m ensurepip \
  && pip3 install --upgrade pip \
  && pip3 install --upgrade setuptools \
  && pip3 install -r /deploy/requirements.txt 

WORKDIR /deploy/app
COPY . /deploy/

EXPOSE 5000

CMD ["gunicorn", "--config", "/deploy/gunicorn.cfg", "server:app"]