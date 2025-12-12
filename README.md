# Meme Bot Frontend (Discord Bot)

2025 Golang Final Project

https://hackmd.io/@penguin71630/go-final-project

Note:
- View API document: `make viewapi`

<!-- My friend wrote this login test script.  So our goal is to make bot generates a link and send this link in the form of Ephemeral message to user.  User can open the webpage with provided link (this way we can skip the hassles over Discord OAuth in webpage), and use some magic mechanism (my friend said cookie but I'm not familiar with it).

Here's my understanding of the entire workflow:
1. User types /web in discord chat room.
2. Bot calls POST /auth/get-login-url with payload {
  "userId": "string"
}, and receive a response: {
  "loginUrl": "http://localhost:8080/login?token=T2KsDcj-kiULQLdotioxmsZG"
}
3.  -->