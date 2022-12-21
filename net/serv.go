package net

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
	"task_rest/DataBase"
	"task_rest/crypting"
	"task_rest/middleware"
	"task_rest/model"
)

type handler struct {
	db *gorm.DB
}

type result struct {
	Result any `json:"result"`
	Error  any `json:"error"`
}

var (
	h        handler
	wrapServ http.Handler
)

func init() {
	middleware.Logs.Debug().Msgf("[http] init started")
	URL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Europe/Moscow",
		model.ConfigFile.Sql.Host, model.ConfigFile.Sql.User, model.ConfigFile.Sql.Password,
		model.ConfigFile.Sql.Database, model.ConfigFile.Sql.Port, model.ConfigFile.Sql.SslMode)
	db, err := DataBase.Init(URL)
	if err != nil {
		middleware.Logs.Err(err).Msgf("error with Data Base opened")
	} else {
		middleware.Logs.Info().Msgf("Data Base is open success")
	}
	h = handler{db}
	middleware.Logs.Debug().Msgf("[http] init finished")
}

// outJson function forms output info in json
func outJson(w http.ResponseWriter, body any, err error) {
	var res result
	if err != nil {
		res = result{Result: nil, Error: err.Error()}
	} else {
		res = result{Result: body, Error: nil}
	}
	middleware.Logs.Debug().Interface("any", res).Msgf("output=")
	w.Header().Set("Content-Type", "application/json")
	jByte, _err := json.Marshal(res)
	if _err != nil {
		middleware.Logs.Err(err).Msgf("[http] json encode failed")
	}
	w.Write(jByte)
	middleware.Logs.Debug().Msgf("[http] json forms end")
}

// StartNet function: started with settings and binds functions http client
func StartNet() {
	middleware.Logs.Debug().Msgf("[http] StartNet started")
	netServ := http.NewServeMux()

	netServ.HandleFunc("/test", h.test)
	netServ.HandleFunc("/encrypt", h.sendEncrypt)
	netServ.HandleFunc("/decrypt", h.sendDecrypt)
	netServ.HandleFunc("/history", h.sendHistory)

	wrapServ = middleware.Logging(netServ)
	middleware.Logs.Debug().Msgf("[http] StartNet finished")
}

// RunNet function: launch listen http client
func RunNet() {
	middleware.Logs.Debug().Msgf("[http] RunNet started")
	adr := ":" + model.ConfigFile.Api.Port
	middleware.Logs.Info().Msgf("listen http client from: %s%s", model.ConfigFile.Api.Host, adr)
	err := http.ListenAndServe(adr, wrapServ) //wrappedMux
	if err != nil {
		middleware.Logs.Err(err).Msgf("error listen http client")
	}
	middleware.Logs.Debug().Msgf("[http] RunNet finished")
}

// Test function: test connection.
func (h handler) test(w http.ResponseWriter, r *http.Request) {
	middleware.Logs.Debug().Msgf("[http] test started")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Work!"))
	middleware.Logs.Debug().Msgf("[http] test finished")
}

// SendEncrypt function: crypt string in request and added value in Data Base
func (h handler) sendEncrypt(w http.ResponseWriter, r *http.Request) {
	middleware.Logs.Debug().Msgf("[http] sendEncrypt started")
	defer r.Body.Close()
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		middleware.Logs.Err(err).Msgf("error read [Request]")
	} else {
		middleware.Logs.Info().Msgf("read [Request] is success")
	}

	var body model.Body
	if err = json.Unmarshal(req, &body); err != nil {
		middleware.Logs.Err(err).Msgf("error Unmarshal json")
		outJson(w, nil, err)
		middleware.Logs.Debug().Msgf("[http] sendEncrypt exit")
		return
	} else {
		middleware.Logs.Info().Msgf("Unmarshal json is success")
	}

	if body.Encrypt == "" {
		err = errors.New("ERROR, no value")
		middleware.Logs.Err(err).Msgf("value is empty")
		outJson(w, body.Encrypt, err)
		return
	}
	body.Decrypt = crypting.Encrypt(body.Encrypt)
	if err = DataBase.AddRec(h.db, "encrypt", body.Encrypt, body.Decrypt); err != nil {
		middleware.Logs.Err(err).Msgf("can't add new record in Data Base")
	} else {
		middleware.Logs.Info().Msgf("add new record in Data Base is success")
	}
	outJson(w, body.Decrypt, err)
	middleware.Logs.Debug().Msgf("[http] sendEncrypt finished")
}

// SendDecrypt function: decrypt string in request and added value in Data Base
func (h handler) sendDecrypt(w http.ResponseWriter, r *http.Request) {
	middleware.Logs.Debug().Msgf("[http] sendDecrypt started")
	defer r.Body.Close()
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		middleware.Logs.Err(err).Msgf("error read [Request]")
	} else {
		middleware.Logs.Info().Msgf("read [Request] is success")
	}

	var body model.Body
	if err = json.Unmarshal(req, &body); err != nil {
		middleware.Logs.Err(err).Msgf("error Unmarshal json")
		outJson(w, nil, err)
		middleware.Logs.Debug().Msgf("[http] sendDecrypt exit")
		return
	} else {
		middleware.Logs.Info().Msgf("Unmarshal json is success")
	}

	if body.Decrypt == "" {
		err = errors.New("ERROR, no value")
		middleware.Logs.Err(err).Msgf("value is empty")
		outJson(w, body.Decrypt, err)
		return
	}
	body.Encrypt = crypting.Decrypt(body.Decrypt)
	if err = DataBase.AddRec(h.db, "decrypt", body.Decrypt, body.Encrypt); err != nil {
		middleware.Logs.Err(err).Msgf("can't add new record in Data Base")
	} else {
		middleware.Logs.Info().Msgf("add new record in Data Base is success")
	}
	outJson(w, body.Encrypt, err)
	middleware.Logs.Debug().Msgf("[http] sendDecrypt finished")
}

// SendHistory function: show History request's from Data Base
func (h handler) sendHistory(w http.ResponseWriter, r *http.Request) {
	middleware.Logs.Debug().Msgf("[http] sendHistory started")
	defer r.Body.Close()
	strLimit, strOffset := r.URL.Query().Get("limit"), r.URL.Query().Get("offset")
	var limit, offset int
	var err error
	limit, err = strconv.Atoi(strLimit)
	if err != nil {
		middleware.Logs.Err(err).Msgf("can't convert value of limit")
		outJson(w, nil, err)
		return
	}
	offset, err = strconv.Atoi(strOffset)
	if err != nil {
		middleware.Logs.Err(err).Msgf("can't convert value of offset")
		outJson(w, nil, err)
		return
	}

	res, err := DataBase.Show(h.db, limit, offset)
	if err != nil {
		middleware.Logs.Err(err).Msgf("can't show list from table")
	}
	outJson(w, res, err)
	middleware.Logs.Debug().Msgf("[http] sendHistory finished")
}
