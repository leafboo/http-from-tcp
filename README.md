# Building an HTTP Server from TCP
In this readme, I will document my journey in building an http server from scratch.

## Sections
[OSI Model](#the-osi-model)

## The OSI Model
|   | Layer | Protocol Data Unit <br> (PDU) |
|---|-------|--------------------------|
| 7 | Application | Data |
| 6 | Presentation | Data |
| 5 | Session | Data |
| 4 | Transport | Segments |
| 3 | Network | Packets | 
| 2 | Data Link | Frames |
| 1 | Physical | Bits |

**Top to Bottom** - Encapsulation. <br>
**Bottom to Top** - Decapsulation. <br> 
**PDU** - This is the form of the data while traversing each layer of the OSI model. <br> <br> <br>

### Flow of data going through the OSI Model from top to bottom <br>

_Session -> Transport_ <br> <br>
The transport layer uses various protocols like the TCP and UDP. In this project, I'll be using TCP. TCP breaks up the session data into segments and attaches a TCP header to each segments. Part of the TCP Header is the source of the application's port address and the intended destination of the application's port address.  
<br><br>
<img width="637" height="552" alt="image" src="https://github.com/user-attachments/assets/060340ee-749b-4702-b3c2-3c513ff23fc3" />
<br><br><br>

_Trasport -> Network_ <br><br>
The IP is used as the protocol in this layer. IP takes the segment and attaches the IP header to it, forming something that's called an IP packet. The IP header contains the source IP `(where it came from)` and the destination IP `(where it will go)` <br><br>
<img width="1121" height="413" alt="image" src="https://github.com/user-attachments/assets/b7fb3cc0-3369-466b-bba5-3d5e2fbb6aa5" />
<br><br><br>

_Network -> Data Link_ <br><br>
The src and the destination MAC address is then attached to the packet, which forms the PDU of this layer. <br>
```
Note: 
    --------------------------------------------------------
    |                 | src MAC address | dest MAC address |
    |-----------------|-----------------|------------------|
    |  encapsulation  |   sender's PC   |      router      |
    |-----------------|-----------------|------------------|
    |  decapsulation  |     router      |   receiver's PC  |
    --------------------------------------------------------

```
<br><br>
<img width="600" height="206" alt="image" src="https://github.com/user-attachments/assets/1acd4774-adfa-45c1-bc77-9810652fa4bc" />





