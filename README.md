# **Chat Application using RPC in Go**
### **Distributed System (ECE5654) - Project Documentation**

## **Overview**
This project implements a **chat application using Remote Procedure Call (RPC) in Go**, allowing multiple clients to join chat rooms and exchange messages. The chat messages are **streamed through a server**, which updates itself based on client actions.

## **Project Structure**
The project is organized into three main directories:


### **1. Server Directory**
The **Server** directory contains the implementation of the chatroom system and client management.

- **chatRoom.go**  
  - Implements the **ChatRoom** struct and manages chat rooms.
  - Functions: `AddChatRoom`, `GetChatRoom`, `GetChatRoomNames`, and `RemoveChatRoom`.
  - Uses **lock() and unlock()** methods for thread safety.

- **client.go**  
  - Manages clients within the chat application.
  - Functions: `AddClient`, `GetClient`, `RemoveClient`, and `NewClient`.
  - Uses **lock() and unlock()** methods for concurrent access control.

- **main.go**  
  - Sets up the **RPC server** using `net/rpc` and `net/http`.
  - Registers the **Receiver** as the RPC handler.
  - Listens for incoming connections on a **TCP network (port 3333)**.
  - Limits the number of clients to **10** (any additional clients receive an error message).

- **receiver.go**  
  - Implements **RPC methods** for:
    - Connecting clients
    - Sending messages
    - Creating and listing chat rooms
    - Joining and leaving chat rooms
    - Changing client names
    - Managing chat room communication

---

### **2. Client Directory**
The **Client** directory contains:
- **main.go**  
  - Establishes a connection to the RPC server.
  - Allows users to send commands and messages.
  - Displays messages received from the server.
  - Commands supported:
    - `/create name` → Create a chat room.
    - `/join name` → Join a chat room.
    - `/help` → Display available commands.

---

### **3. Shared Directory**
The **Shared** directory contains:
- **shared.go**  
  - Defines **constants and shared structures** used between the server and client.

---

## **Network Configuration**
- **Protocol**: TCP  
- **Port**: 3333  

---

## **How to Run the Chat Application**
### **1. Start the Server**


### **1. Server Directory**
The **Server** directory contains the implementation of the chatroom system and client management.

- **chatRoom.go**  
  - Implements the **ChatRoom** struct and manages chat rooms.
  - Functions: `AddChatRoom`, `GetChatRoom`, `GetChatRoomNames`, and `RemoveChatRoom`.
  - Uses **lock() and unlock()** methods for thread safety.

- **client.go**  
  - Manages clients within the chat application.
  - Functions: `AddClient`, `GetClient`, `RemoveClient`, and `NewClient`.
  - Uses **lock() and unlock()** methods for concurrent access control.

- **main.go**  
  - Sets up the **RPC server** using `net/rpc` and `net/http`.
  - Registers the **Receiver** as the RPC handler.
  - Listens for incoming connections on a **TCP network (port 3333)**.
  - Limits the number of clients to **10** (any additional clients receive an error message).

- **receiver.go**  
  - Implements **RPC methods** for:
    - Connecting clients
    - Sending messages
    - Creating and listing chat rooms
    - Joining and leaving chat rooms
    - Changing client names
    - Managing chat room communication

---

### **2. Client Directory**
The **Client** directory contains:
- **main.go**  
  - Establishes a connection to the RPC server.
  - Allows users to send commands and messages.
  - Displays messages received from the server.
  - Commands supported:
    - `/create name` → Create a chat room.
    - `/join name` → Join a chat room.
    - `/help` → Display available commands.

---

### **3. Shared Directory**
The **Shared** directory contains:
- **shared.go**  
  - Defines **constants and shared structures** used between the server and client.

---

## **Network Configuration**
- **Protocol**: TCP  
- **Port**: 3333  

---

## **How to Run the Chat Application**
### **1. Start the Server**

cd /Sol_chat_go
cd /go-chat-rpc
go run server/*.go

cd /Sol_chat_go
cd /go-chat-rpc
go run server/*.go

