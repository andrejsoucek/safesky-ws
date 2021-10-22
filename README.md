# safesky-ws
Distribute SafeSky data through WS.

## How it works
```
                                                           clients
                                                         +---+  +---+
                                                   WS    |   |  |   |
+------------+        +------------------------+ ------> |   |  |   |
|  SafeSky   | 4 sec  |        safesky-ws      |         +---+  +---+
|            + <----- +------------------------+ ------>
| REST API   |        | filtering for viewport |         +---+  +---+
+------------+        +------------------------+ ------> |   |  |   |
                                                         |   |  |   |
                                                         +---+  +---+       
```
## Setting up
TBA
