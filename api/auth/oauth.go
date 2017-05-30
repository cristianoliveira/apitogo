package auth

import(
  "net/http"
  "github.com/RangelReale/osin"
)

var server *osin.Server = osin.NewServer(osin.NewServerConfig(), NewFakeStorage())

func HandleAuthorization(w http.ResponseWriter, r *http.Request) {
  resp := server.NewResponse()
	defer resp.Close()

	if ar := server.HandleAuthorizeRequest(resp, r); ar != nil {

		ar.Authorized = true
		server.FinishAuthorizeRequest(resp, r, ar)

	}
	osin.OutputJSON(resp, w, r)
}

func HandleToken(w http.ResponseWriter, r *http.Request) {
	resp := server.NewResponse()
	defer resp.Close()

	if ar := server.HandleAccessRequest(resp, r); ar != nil {
		ar.Authorized = true
		server.FinishAccessRequest(resp, r, ar)
	}
	osin.OutputJSON(resp, w, r)
}
