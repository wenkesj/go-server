package httpserver;

import (
  "log"
  "sync";
);

type HTTPServerGroup []*HTTPServer

func (serverGroup HTTPServerGroup) Listen() {
  waitGroup := sync.WaitGroup{}
  log.Printf("Starting server group ... \n");
  waitGroup.Add(len(serverGroup));
  for _, server := range serverGroup {
    go func(server *HTTPServer) {
      if err := server.Listen(); err != nil {
        log.Printf("Error starting server failed on " + server.Address + server.Port);
      }
      log.Printf("Server HTTP(s) listening on " + server.Address + server.Port + "\n");
      waitGroup.Done();
    }(server);
  }
  waitGroup.Wait();
  log.Printf("All servers closed \n");
}
