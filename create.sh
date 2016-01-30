#!/bin/bash
# Script written by Faraz
# Automates creating a new Journey Blog with nginx
SITENAME="$1" # Site Name (Without URL)
SITEURL="$2" # Site URL excluding http:// or www. (subdomain for now)
PORT="$3" # <- Port passed in
if [ -z "$SITENAME" ] || [ -z "$SITEURL" ] || [ -z "$PORT" ] 
then
  echo 'Error, not enough arguments!'
  exit 2
fi
SITENAME=journey-$SITENAME
mkdir -p /var/www/$SITENAME
cd /var/www/$SITENAME
cp -R /journey . # Recursively copy preexisting files
cd journey
sed -i -e "s/8084/$PORT/g" config.json
sed -i -e "s/127.0.0.1:$PORT/$SITEURL:$PORT/g" config.json
echo "start on runlevel [2345]" >> /etc/init/$SITENAME.conf
echo "stop on runlevel [!2345]" >> /etc/init/$SITENAME.conf
echo "respawn" >> /etc/init/$SITENAME.conf
echo "console none" >> /etc/init/$SITENAME.conf
echo "exec /var/www/$SITENAME/journey/journey -log=/var/www/$SITENAME/journey/log.txt" >> /etc/init/$SITENAME.conf
cd /etc/nginx/sites-enabled
echo "server {" >> $SITENAME.conf
echo "listen 0.0.0.0:80;" >> $SITENAME.conf
echo "server_name $SITEURL;" >> $SITENAME.conf
echo "access_log /var/log/nginx/$SITENAME.log;" >> $SITENAME.conf
echo "location / {" >> $SITENAME.conf
echo "proxy_pass http://127.0.0.1:$PORT;" >> $SITENAME.conf
echo " }" >> $SITENAME.conf
echo "}" >> $SITENAME.conf
echo "server {" >> $SITENAME.conf
echo "listen 0.0.0.0:80;" >> $SITENAME.conf
echo "server_name www.$SITEURL;" >> $SITENAME.conf
echo "access_log /var/log/nginx/$SITENAME.log;" >> $SITENAME.conf
echo "location / {" >> $SITENAME.conf
echo "proxy_pass http://127.0.0.1:$PORT;" >> $SITENAME.conf
echo " }" >> $SITENAME.conf
echo "}" >> $SITENAME.conf
service $SITENAME start
service nginx restart
service $SITENAME restart # <- required
echo "Created new Journey blog! -> $SITEURL"
echo "---------------------------------------------------"