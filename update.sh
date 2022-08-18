#!/bin/bash
set -ex

LICENCE_KEY=xxxx

curl "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=${LICENCE_KEY}&suffix=tar.gz" -o GeoLite2-City-latest.tar.gz

tar -xzvf GeoLite2-City-latest.tar.gz
version=$(tar -tf GeoLite2-City-latest.tar.gz | head -1)

cp -v $version/GeoLite2-City.mmdb GeoLite2-City.mmdb

docker restart hub-geoip
