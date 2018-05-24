#!/usr/bin/env bash
file=/var/www/html/municipios.gaido.net.br/deploy/cidades
if [ -e "$file" ]; then
    systemctl stop municipios.service
    mv "$file" /var/www/html/municipios.gaido.net.br
    chmod 0774 /var/www/html/municipios.gaido.net.br/cidades
    systemctl start municipios.service
fi