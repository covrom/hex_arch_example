package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/covrom/hex_arch_example/best_practice/internal/entities/user"
	"github.com/covrom/hex_arch_example/best_practice/internal/logic/app/repos/userrepo"

	"github.com/google/uuid"
)

type Router struct {
	*http.ServeMux
	us *userrepo.Users
}

func NewRouter(us *userrepo.Users) *Router {
	r := &Router{
		ServeMux: http.NewServeMux(),
		us:       us,
	}
	r.Handle("/create",
		// r.AuthMiddleware(
		r.AuthMiddleware(
			http.HandlerFunc(r.CreateUser),
		),
		// ),
	)
	r.Handle("/read", r.AuthMiddleware(http.HandlerFunc(r.ReadUser)))
	r.Handle("/delete", r.AuthMiddleware(http.HandlerFunc(r.DeleteUser)))
	r.Handle("/search", r.AuthMiddleware(http.HandlerFunc(r.SearchUser)))
	return r
}

type User struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Data       string    `json:"data"`
	Permission int       `json:"perms"`
}

func (rt *Router) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if u, p, ok := r.BasicAuth(); !ok || !(u == "admin" && p == "admin") {
				http.Error(w, "unautorized", http.StatusUnauthorized)
				return
			}
			// r = r.WithContext(context.WithValue(r.Context(), 1, 0))
			next.ServeHTTP(w, r)
		},
	)
}

func (rt *Router) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	u := User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	bu := user.User{
		Name: u.Name,
		Data: u.Data,
	}

	nbu, err := rt.us.Create(r.Context(), bu)
	if err != nil {
		http.Error(w, "error when creating", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(
		User{
			ID:         nbu.ID,
			Name:       nbu.Name,
			Data:       nbu.Data,
			Permission: nbu.Permissions,
		},
	)
}

// read?uid=...
func (rt *Router) ReadUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}

	suid := r.URL.Query().Get("uid")
	if suid == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	uid, err := uuid.Parse(suid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if (uid == uuid.UUID{}) {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	nbu, err := rt.us.Read(r.Context(), uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(
		User{
			ID:         nbu.ID,
			Name:       nbu.Name,
			Data:       nbu.Data,
			Permission: nbu.Permissions,
		},
	)
}

func (rt *Router) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}

	suid := r.URL.Query().Get("uid")
	if suid == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	uid, err := uuid.Parse(suid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if (uid == uuid.UUID{}) {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	nbu, err := rt.us.Delete(r.Context(), uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(
		User{
			ID:         nbu.ID,
			Name:       nbu.Name,
			Data:       nbu.Data,
			Permission: nbu.Permissions,
		},
	)
}

// /search?q=...
func (rt *Router) SearchUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}

	q := r.URL.Query().Get("q")
	if q == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	ch, err := rt.us.SearchUsers(r.Context(), q)
	if err != nil {
		http.Error(w, "error when reading", http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)

	first := true
	fmt.Fprintf(w, "[")
	defer fmt.Fprintln(w, "]")

	for {
		select {
		case <-r.Context().Done():
			return
		case u, ok := <-ch:
			if !ok {
				return
			}
			if first {
				first = false
			} else {
				fmt.Fprintf(w, ",")
			}
			_ = enc.Encode(
				User{
					ID:         u.ID,
					Name:       u.Name,
					Data:       u.Data,
					Permission: u.Permissions,
				},
			)
			w.(http.Flusher).Flush()
		}
	}
}
