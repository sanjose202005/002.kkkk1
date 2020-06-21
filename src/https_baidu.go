// https://www.example-code.com/golang/socket_socks_proxy.asp

// This example requires the Chilkat API to have been previously unlocked.
// See Global Unlock Sample for sample code.

socket := chilkat.NewSocket()

// To use a SOCKS4 or SOCKS5 proxy, simply set the following
// properties prior to calling Connect.  The connection may be SSL/TLS or 
// non-secure - both will work with a SOCKS proxy.
// The SOCKS hostname may be a domain name or 
// IP address:
socket.SetSocksHostname("www.mysocksproxyserver.com")
socket.SetSocksPort(1080)
socket.SetSocksUsername("myProxyLogin")
socket.SetSocksPassword("myProxyPassword")
// Set the SOCKS version to 4 or 5 based on the version
// of the SOCKS proxy server:
socket.SetSocksVersion(5)
// Note: SOCKS4 servers only support usernames without passwords.
// SOCKS5 servers support full login/password authentication.

// Connect to port 5555 of 192.168.1.108.
// hostname may be a domain name or IP address.
hostname := "192.168.1.108"
ssl := false
maxWaitMillisec := 20000
success := socket.Connect(hostname,5555,ssl,maxWaitMillisec)
if success != true {
    fmt.Println(socket.LastErrorText())
    socket.DisposeSocket()
    return
}

// Set maximum timeouts for reading an writing (in millisec)
socket.SetMaxReadIdleMs(10000)
socket.SetMaxSendIdleMs(10000)

// The server (in this example) is going to send a "Hello World!" 
// message.  Read it:
receivedMsg := socket.ReceiveString()
if socket.LastMethodSuccess() != true {
    fmt.Println(socket.LastErrorText())
    socket.DisposeSocket()
    return
}

// Close the connection with the server
// Wait a max of 20 seconds (20000 millsec)
success = socket.Close(20000)

fmt.Println(*receivedMsg)

socket.DisposeSocket()
