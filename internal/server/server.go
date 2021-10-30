package server

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/vSterlin/auth/internal/user"
)

type Server struct {
	addr string
	// db *sql.DB
}

func NewServer(addr int) *Server {
	strAddr := strconv.Itoa(addr)
	return &Server{addr: strAddr}
}

func (s *Server) Init() {

	ur := user.NewUserRepo()
	us := user.NewUserService(ur)
	as := user.NewAuthService(us)
	am := user.NewAuthMiddleware(us)
	uc := user.NewUserController(us, as)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(am.CurrentUser)

	r.Route("/users", func(r chi.Router) {

		// r.Get("/", uc.GetUsers)
		r.With(am.IsAuthenticated).Get("/", uc.GetUsers)
		r.Get("/current-user", uc.GetCurrentUser)
		r.Post("/signup", uc.SignUp)
		r.Post("/signin", uc.SignIn)

	})

	http.ListenAndServe(":"+s.addr, r)
}

func (s *Server) Shutdown() {
	// close db
}
