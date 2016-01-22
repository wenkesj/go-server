package httpserver;

import (
  "log";
  "strings";
  "net/http";
  "github.com/go-zoo/bone";
);

type HTTPHandler func(http.ResponseWriter, *http.Request);

type HTTPServer struct {
  Address string `json:"address"`;
  Port string `json:"port"`;
  SSL bool `json:"ssl"`;
  Router map[string]HTTPHandler;
  KeyFile string `json:"keyfile"`;
  CertFile string `json:"certfile"`;
}

func (server *HTTPServer) Listen() error {
  if server.SSL {
    log.Printf("Server HTTPS listening on " + server.Address + server.Port);
    return http.ListenAndServeTLS(server.Port, server.CertFile, server.KeyFile, nil);
  }
  log.Printf("Server HTTP listening on " + server.Address + server.Port);
  return http.ListenAndServe(server.Port, NewRouter(server.Router));
}

func NewRouter(handler map[string]HTTPHandler) *bone.Mux {
  mux := bone.New();
  for key := range handler {
    reqestHandler := handler[key];
    requesyKeys := strings.Split(key, "†");
    requestType, requestUrl := requesyKeys[0], requesyKeys[1];
    switch requestType {
    case "GET":
      mux.Get(requestUrl, http.HandlerFunc(reqestHandler));
      break;
    case "POST":
      mux.Post(requestUrl, http.HandlerFunc(reqestHandler));
      break;
    case "HANDLE":
      mux.Handle(requestUrl, http.HandlerFunc(reqestHandler));
      break;
    }
  }
  return mux;
};
