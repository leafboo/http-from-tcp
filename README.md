# Building an HTTP Server from TCP
In this readme, I will document my journey on how I will build an http server from scratch.

## Sections
[OSI Model](#the-osi-model)

### The OSI Model
|   | Layer | Protocol Data Unit <br> (PDU) |
|---|-------|--------------------------|
| 7 | Application | Data |
| 6 | Presentation | Data |
| 5 | Session | Data |
| 4 | Transport | Segments |
| 3 | Network | Packets | 
| 2 | Data Link | Frames |
| 1 | Physical | Bits |

**PDU** - This is the form of the data while traversing each layer of the OSI model.
