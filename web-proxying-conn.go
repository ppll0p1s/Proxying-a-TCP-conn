package main

import (
    "fmt"
    "log"
    "net/http"
    "io"
    "net"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Pepito Flores\nTrabajador Ejemplar\nValencia\n27/08/1994 %s", r.URL.Path[1:])
}

func handleConns(src net.Conn) {
    // Realizamos peticion desde "pepito.es" a "comohackearamijefe.saw"
    dst, err := net.Dial("tcp","comohackearamijefe.saw")
    if err!= nil {
        log.Fatalln("No se ha podido realizar la conexion")
    }
    defer dst.Close()

    // Copiamos la conexion entrante a la que se acaba de realizar
    // Lo ejecutamos en una goroutina para prevenir bloqueos producidos por io.Copy
    go func() {
        if _, err := io.Copy(dst,src); err != nil{
            log.Fatalln(err)
        }
    }()

    // Copiamos el resultado de la peticion a "comohackearamijefe.saw" a peticion origen
    if _,err := io.Copy(src,dst); err != nil{
        log.Fatalln(err)
    }
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))

    // Escuchamos conexiones al puerto 8080, donde se aloja "pepito.es"
    listener, err := net.Listen("tcp",":8080")
    if err!= nil {
        log.Fatalln("Incapaz de enlazar con el puerto indicado.")
    }

    for {
    	// Para todas las conexiones al puerto 8080...
        conn, err := listener.Accept()
        if err!= nil {
            log.Fatalln("Incapaz de acceptar la comunicacion")
        }
        // Ejecutamos la funcion handleConns
        go handleConns(conn)
    }
}
