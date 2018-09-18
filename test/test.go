package main

func main(){
    s := CreateServer()
    log.Fatal(s.httpserver.ListenAndServe())
}

