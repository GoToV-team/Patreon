package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"patreon/internal/app"
	"patreon/internal/app/store"
	"patreon/internal/models"
)

type RegisterHandler struct {
	baseHandler app.HandlerJoiner
	router      *mux.Router
	Store       store.Store
	log         *logrus.Logger
}

func NewRegisterHandler() *RegisterHandler {
	return &RegisterHandler{
		baseHandler: *app.NewHandlerJoiner([]app.Joinable{}, "/register"),
		log:         logrus.New(),
	}
}

func (h *RegisterHandler) SetStore(store store.Store) {
	h.Store = store
}
func (h *RegisterHandler) SetLogger(logger *logrus.Logger) {
	h.log = logger
}
func (h *RegisterHandler) Join(router *mux.Router) {
	router.HandleFunc(h.baseHandler.GetUrl(), h.ServeHTTP).Methods("POST", "GET")
	h.baseHandler.Join(router)
}
func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			h.log.Error(err)
		}
	}(r.Body)
	req := &request{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(req); err != nil {
		h.Error(w, r, http.StatusUnprocessableEntity, err)
		return
	}
	u := &models.User{
		Login:    req.Login,
		Password: req.Password,
	}

	logUser, _ := json.Marshal(u)
	logrus.Info("get: ", string(logUser))

	checkUser, _ := h.Store.User().FindByLogin(u.Login)
	if checkUser != nil {
		h.Error(w, r, http.StatusConflict, store.UserAlreadyExist)
		return
	}
	if err := h.Store.User().Create(u); err != nil {
		h.Error(w, r, http.StatusBadRequest, err)
		return
	}
	u.MakePrivateDate()
	h.Respond(w, r, http.StatusOK, u)
}
func (h *RegisterHandler) Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	h.Respond(w, r, code, map[string]string{"error": err.Error()})
}
func (h *RegisterHandler) Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	encoder := json.NewEncoder(w)
	w.WriteHeader(code)
	if data != nil {
		err := encoder.Encode(data)
		if err != nil {
			h.log.Error(err)
		}
	}
	logUser, _ := json.Marshal(data)
	logrus.Info("Respond data: ", string(logUser))
}