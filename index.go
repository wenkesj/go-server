package main;

import (
  "io/ioutil";
  "net/http";
  "encoding/json";
  "github.com/go-zoo/bone";
  "github.com/wenkesj/go-server/config";
  "github.com/wenkesj/go-server/httpserver";
  "github.com/robertkrimen/otto";
);

type JSRequest struct {
  JS string `json:"javascript"`;
};

// Set up JS VM.
var vm = otto.New();

func main() {
  // Get the config for the server.
  group, _ := config.Server("production");

  // Set up routes.
  router := make(map[string]httpserver.HTTPHandler);
  router["GET†/"] = GetRoot;
  router["GET†/:file"] = GetPublic;
  router["POST†/js"] = PostJS;
  group.Server.Router = router;

  // Server listening ...
  group.Server.Listen();
}

func GetRoot(res http.ResponseWriter, req *http.Request) {
  http.ServeFile(res, req, "public/index.html");
};

func GetPublic(res http.ResponseWriter, req *http.Request) {
  file := bone.GetValue(req, "file");
  http.ServeFile(res, req, "public/" + file);
};

func PostJS(res http.ResponseWriter, req *http.Request) {
  body, _ := ioutil.ReadAll(req.Body);
  jsReq := &JSRequest{};
  json.Unmarshal([]byte(string(body)), jsReq);
  ret, err := vm.Run(string(jsReq.JS));
  _ret, _ := ret.ToString();
  if err != nil {
    res.Write([]byte(err.Error()));
    return;
  }
  res.Write([]byte(_ret));
};
