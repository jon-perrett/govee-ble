#!/bin/bash
set -e

cp .env /usr/local/env/govee-store.env

systemctl stop store.service
systemctl stop reader.service
cp ./out/* /usr/bin/
cp *.service /etc/systemd/system/

systemctl enable store.service
systemctl enable reader.service
systemctl restart store.service
systemctl restart reader.service
