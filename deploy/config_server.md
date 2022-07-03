## Deploy
#### сервис
После копирования файла `systemd/marusya.service` в директорию `/etc/systemd/system/` выполняем:
```shell
sudo systemctl start namaztime.service
sudo systemctl enable namaztime.service
```

#### nginx
копируем конфиг nginx `nginx/sites-available/marusya.nginx` в `/etc/nginx/sites-available/marusya` и выполняем
```shell
ln -s /etc/nginx/sites-available/namaztime /etc/nginx/sites-enabled/
nginx -t && nginx -s reload
```


