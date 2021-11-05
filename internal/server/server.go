package server

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/vSterlin/auth/internal/cache"
	"github.com/vSterlin/auth/internal/user"
)

type Server struct {
	addr  string
	db    *sql.DB
	cache cache.Cache
}

func NewServer(addr int, db *sql.DB, c cache.Cache) *Server {
	strAddr := strconv.Itoa(addr)
	return &Server{addr: strAddr, db: db, cache: c}
}

func (s *Server) Init() {

	s.db.Exec(user.CreateUserTableSQL)

	ur := user.NewUserRepo(s.db)
	cur := user.NewCachedUserRepo(ur, s.cache)
	us := user.NewUserService(cur)
	as := user.NewAuthService(us)
	am := user.NewAuthMiddleware(us)
	uc := user.NewUserController(us, as)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(am.CurrentUser)

	r.Route("/users", func(r chi.Router) {

		r.With(am.IsAuthenticated).Get("/", uc.GetUsers)
		r.Get("/current-user", uc.GetCurrentUser)
		r.Post("/signup", uc.SignUp)
		r.Post("/signin", uc.SignIn)

		r.Get("/refresh-token", uc.RefreshToken)

	})

	http.ListenAndServe(":"+s.addr, r)
}

func (s *Server) Shutdown() {
	s.db.Close()
}
