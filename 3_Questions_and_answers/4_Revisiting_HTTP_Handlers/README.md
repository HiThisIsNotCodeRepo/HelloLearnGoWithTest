# Revisiting HTTP Handler

## Use interface to separate responsibilities from HTTP handler

1. Actual data persistent/mock persistent , this is achieved by interface.
2. Routing , this is achieved by `http.Handler`

```
type User struct {
	Name string
}
type UserService interface {
	Register(user User) (insertedID string, err error)
}

// UserServer struct contains services for data persistence or testing
type UserServer struct {
	service UserService
}

// NewUserServer constructs UserServer , to enable router you have to embed a http.Handler in struct
func NewUserServer(service UserService) *UserServer {
	return &UserServer{service: service}
}

// RegisterUser describe a HandlerFunc which we will use to call the service in the struct
func (u *UserServer) RegisterUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// request parsing and validation
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, fmt.Sprintf("could not decode user payload: %v", err), http.StatusBadRequest)
		return
	}

	// call a service thing to take care of the hard work
	insertedID, err := u.service.Register(newUser)

	// depending on what we get back, respond accordingly
	if err != nil {
		//todo: handle different kinds of errors differently
		http.Error(w, fmt.Sprintf("problem registering new user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, insertedID)
}

```