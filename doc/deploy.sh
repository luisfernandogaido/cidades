#!/usr/bin/env bash
file=/var/www/html/municipios.gaido.net.br/deploy/service
if [ -e "$file" ]; then
    systemctl stop municipios.service
    mv "$file" /var/www/html/municipios.gaido.net.br
    chmod 0774 /var/www/html/municipios.gaido.net.br/service
    systemctl start municipios.service
fi