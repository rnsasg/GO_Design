# Adapter in Go

- [ ] Adapter is a structural design pattern, which allows incompatible objects to collaborate.
- [ ] The Adapter acts as a wrapper between two objects. It catches calls for one object and transforms them to format and interface recognizable by the second object.

## Conceptual Example

We have a client code that expects some features of an object (Lightning port), but we have another object called adaptee (Windows laptop) which offers the same functionality but through a different interface (USB port)

This is where the Adapter pattern comes into the picture. We create a struct type known as adapter that will:

+ Adhere to the same interface which the client expects (Lightning port).

+ Translate the request from the client to the adaptee in the form that the adaptee expects. The adapter accepts a Lightning connector and then translates its signals into a USB format and passes them to the USB port in windows laptop.

### client.go: Client code

```
package main

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
    fmt.Println("Client inserts Lightning connector into computer.")
    com.InsertIntoLightningPort()
}
```

### computer.go: Client interface

```
package main

type Computer interface {
    InsertIntoLightningPort()
}
```

### mac.go: Service
```
package main

import "fmt"

type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
    fmt.Println("Lightning connector is plugged into mac machine.")
}
```
###  windows.go: Unknown service
```
package main

import "fmt"

type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
    fmt.Println("USB connector is plugged into windows machine.")
}
```

### windowsAdapter.go: Adapter
```
package main

import "fmt"

type WindowsAdapter struct {
    windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
    fmt.Println("Adapter converts Lightning signal to USB.")
    w.windowMachine.insertIntoUSBPort()
}
```

### main.go
```
package main

func main() {

    client := &Client{}
    mac := &Mac{}

    client.InsertLightningConnectorIntoComputer(mac)

    windowsMachine := &Windows{}
    windowsMachineAdapter := &WindowsAdapter{
        windowMachine: windowsMachine,
    }

    client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
```
###  output.txt: Execution result

```
Client inserts Lightning connector into computer.
Lightning connector is plugged into mac machine.
Client inserts Lightning connector into computer.
Adapter converts Lightning signal to USB.
USB connector is plugged into windows machine.
```