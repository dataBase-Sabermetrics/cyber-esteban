## Discord Bot for dataBase Sabermetrics
Cyber-Esteban lives on a raspberry pi 5 hosted in my office. It serves as an api that the dataBase Sabermetrics related apps and services send data to when certain events happen and then Cyber-Esteban will send alerts and messages in our dataBase Sabermetrics Discord channel. 

### Getting Strted
```
# fill in .env with DISCORD_KEY and CHANNELID
make start
```

### Push Image to ghcr: 
```
make push 
```
