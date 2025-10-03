# Building an HTTP Server from TCP
In this readme, I will document my journey on how I will build an http server from scratch.

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
TCP breaks up the session data into segments and attaches a TCP header to each segments. <br><br>
<img width="637" height="552" alt="image" src="https://github.com/user-attachments/assets/060340ee-749b-4702-b3c2-3c513ff23fc3" />
<br><br><br>
_Trasport -> Network_ <br><br>
IP takes the segment and attaches the IP header to it, forming something that's called an IP packet. <br><br>
<img width="1121" height="413" alt="image" src="https://github.com/user-attachments/assets/b7fb3cc0-3369-466b-bba5-3d5e2fbb6aa5" />





