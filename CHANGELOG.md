# [1.1.0] - 2023-04-05
## Added
- Config struct to store the bot configuration

# [1.0.4] - 2023-04-05
## Added
- Added the `SendRawRequest` function to send a raw request to the WebSocket - connection.

## Fixed
- Fixed concurrent write to websocket connection

# [1.0.3] - 2023-04-05
## Added
- Added the SendPublicMessage function to send a public message to the chat.
- Added the SendPrivateMessage function to send a private message to a user.
- Added the SendEmote function to send an emote to the chat.
- Added the TeleportUser function to teleport a user to a specific location.
- Added the FloorHit function to hit the floor at a specific location.
  
The functions listed above can be used to perform various actions in the game using the WebSocket - connection. The functions are available in the `Bot` struct, which is returned by the `NewBot` function. The `NewBot` function takes a single argument, which is a function that will be called whenever the bot receives a message. The function takes two arguments: the bot instance and the message received. The function should return a string, which will be sent back to the sender.