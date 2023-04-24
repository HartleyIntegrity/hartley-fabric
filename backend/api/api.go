package api

import (
	"encoding/json"
	"net/http"

	"github.com/HartleyIntegrity/hartley-fabric/backend/auth"
	"github.com/HartleyIntegrity/hartley-fabric/backend/blockchain"
	"github.com/gorilla/mux"
)

var Blockchain []blockchain.Block

var (
	Username string = "admin"
	Password string = "password123"
)

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

type API struct {
	router *mux.Router
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	var req SignInRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Username != Username || req.Password != Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	res := SignInResponse{Token: "dummy_token"}
	json.NewEncoder(w).Encode(res)
}

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.router.ServeHTTP(w, r)
}

func NewAPI() *API {
	api := &API{router: mux.NewRouter()}
	api.initRoutes()
	return api
}

func (api *API) initRoutes() {
	api.router.HandleFunc("/transactions", api.getTransactions).Methods(http.MethodGet)
	api.router.HandleFunc("/transactions", api.createTransaction).Methods(http.MethodPost)
	api.router.HandleFunc("/signin", api.signIn).Methods(http.MethodPost)
	api.router.HandleFunc("/validate-token", api.validateToken).Methods(http.MethodPost)
}

func (api *API) getTransactions(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to retrieve transactions
	jsonResponse(w, http.StatusOK, Blockchain)

}

func (api *API) createTransaction(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to create a new transaction
	var transaction blockchain.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, "Invalid transaction data")
		return
	}

	prevBlock := Blockchain[len(Blockchain)-1]
	newBlock := blockchain.GenerateNewBlock(prevBlock, []blockchain.Transaction{transaction})
	Blockchain = append(Blockchain, newBlock)

	jsonResponse(w, http.StatusCreated, newBlock)
}

func (api *API) signIn(w http.ResponseWriter, r *http.Request) {

	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Replace the following line with actual user authentication logic
	isAuthenticated := credentials.Username == "admin" && credentials.Password == "password"

	if !isAuthenticated {
		jsonResponse(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	token, err := auth.GenerateToken(credentials.Username)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	jsonResponse(w, http.StatusOK, map[string]string{"token": token})
}

func (api *API) validateToken(w http.ResponseWriter, r *http.Request) {
	var token struct {
		Token string `json:"token"`
	}

	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	username, err := auth.ValidateToken(token.Token)
	if err != nil {
		jsonResponse(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	jsonResponse(w, http.StatusOK, map[string]string{"username": username})
}

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
