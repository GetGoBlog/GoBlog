#!/bin/bash
# Script written by Faraz
# Automates creating a new Journey Blog with nginx
SITENAME="$1" # Site Name (Without URL)
SITEURL="$2" # Site URL excluding http:// or www.
PORT="$3" # <- Random port passed in
if [ -z "$SITENAME" ] || [ -z "$SITEURL" ] || [ -z "$PORT" ] 
then
  echo 'Error, not enough arguments!'
  exit 2
fi
SITENAME=journey-$SITENAME
mkdir -p /var/www
cd /var/www
echo "This will delete the folder $SITENAME! Press CTRL + C to quit, or press enter to continue!"
read input
mkdir $SITENAME
cd $SITENAME
echo "Getting latest release!" # TODO make this a symlink for all except config and db files
wget https://github.com/kabukky/journey/releases/download/v0.1.9/journey-linux-amd64.zip
unzip journey-linux-amd64.zip
rm journey-linux-amd64.zip
mv journey-linux-amd64/ journey
cd journey
echo "Configuring config.json"
sed -i -e "s/8084/$PORT/g" config.json
sed -i -e "s/127.0.0.1:$PORT/$SITEURL:$PORT/g" config.json
echo "start on runlevel [2345]" >> /etc/init/$SITENAME.conf
echo "stop on runlevel [!2345]" >> /etc/init/$SITENAME.conf
echo "respawn" >> /etc/init/$SITENAME.conf
echo "console none" >> /etc/init/$SITENAME.conf
echo "exec /var/www/$SITENAME/journey/journey -log=/var/www/$SITENAME/journey/log.txt" >> /etc/init/$SITENAME.conf
cd /etc/nginx/sites-enabled
echo "Proxying port from 80 to $PORT"
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
echo "---------------------------------------------------"
echo "You may need to oonfigure your DNS Records if you used a custom domain!"
echo "ALL DONE! $SITEURL is viewable as a Journey blog!"
echo "Setup at $SITEURL/admin"