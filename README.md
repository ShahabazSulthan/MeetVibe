# MeetVibe

MeetVibe is a simple video conferencing application designed for seamless real-time communication. Built using **Go**, **WebSocket**, and **WebRTC**, it enables users to create and join virtual rooms for video calls. With a user-friendly interface and robust backend architecture, MeetVibe is perfect for small-scale, real-time video conferencing needs.

---

## Features

- **Create and Join Rooms**: Generate unique room IDs and invite participants.
- **Real-time Video Streaming**: High-quality video and audio communication using WebRTC.
- **Dynamic UI**: A responsive, interactive interface built with HTML, CSS, and JavaScript.
- **WebSocket Integration**: Efficient real-time messaging for signaling and broadcasts.
- **Room Management**: Automatic cleanup of inactive participants and empty rooms.
- **Custom Controls**: Mute/unmute audio, toggle video, and hang up.

---

## Tech Stack

- **Backend**: Go, Gin Framework, Gorilla WebSocket
- **Frontend**: HTML, CSS, JavaScript
- **Real-time Communication**: WebRTC
- **Server Management**: Room and participant handling with thread-safe operations

---

## Getting Started

### Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/doc/install)

---

### Setup Instructions

1. **Clone the Repository**

   ```bash
   git clone https://github.com/your-username/MeetVibe.git
   cd MeetVibe
   ```

2. **Run the Server**

   - Navigate to the root directory of the project.
   - Start the backend server:

     ```bash
     go run main.go
     ```

   - The application will run at: [http://localhost:8000/](http://localhost:8000/).

3. **Open in Browser**

   - Open [http://localhost:8000/](http://localhost:8000/) in your web browser.
   - Click **Create Room** to generate a new room or **Join Room** with an existing room ID.

---

## Usage

### Create a Room
- Send a POST request to `/create` to generate a unique room ID.
- Example response: 
  ```json
  {
      "room_id": "aBcDe123"
  }
  ```

### Join a Room
- Use the `/room/:roomID` route to join a room.
- The video call interface will load for real-time communication.

---

## File Structure

```
MeetVibe/
├── server/                  # Backend server code
│   ├── room.go              # Room and participant management
│   ├── handlers.go          # WebSocket and HTTP handlers
├── templates/               # Frontend HTML templates
│   ├── index.html           # Landing page
│   ├── room.html            # Room video call interface
├── main.go                  # Entry point for the Go application
├── go.mod                   # Go modules dependencies
└── README.md                # Project documentation
```

---

## Future Enhancements

- **Chat Functionality**: Add a group chat feature for text communication.
- **Screen Sharing**: Enable participants to share their screens.
- **Room Authentication**: Implement room passwords or token-based authentication.
- **Participant Limits**: Configure maximum participants per room.
- **Recording**: Allow users to record and download meetings.

---

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

1. Fork the repository.
2. Create a new branch for your feature/bug fix.
3. Commit and push your changes.
4. Open a pull request describing the changes.


---

Start your video conferencing journey with **MeetVibe** today! 🚀
