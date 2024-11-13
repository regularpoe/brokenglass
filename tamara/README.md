# tAMARA

Small Ruby app with GUI controls for CMUS.

1. Set password and start cmus
```
:set passwd=password

cmus --listen 0.0.0.0
```

2. Check if cmus is running as it should 
```
sudo netstat -tulpn | grep LISTEN
```

3. Start GUI
```
ruby app.rb
```

4. $profit
