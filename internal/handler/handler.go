package handler

import (
	"forum/internal/service"
	"html/template"
	"net/http"
	"path/filepath"
)

type Handler struct {
	Mux     *http.ServeMux
	Tmp     *template.Template
	Service *service.Service
}

func InitHandler(services *service.Service) *Handler {
	return &Handler{
		Mux:     http.NewServeMux(),
		Tmp:     template.Must(template.ParseGlob("./ui/template/*.html")),
		Service: services,
	}
}

func (h *Handler) Routers() {
	h.Mux.HandleFunc("/", h.middleWare(h.homePage))
	h.Mux.HandleFunc("/signUp", h.signUp)
	h.Mux.HandleFunc("/signIn", h.signIn)

	h.Mux.HandleFunc("/logOut", h.logOut)

	h.Mux.HandleFunc("/post/create", h.middleWare(h.postCreate))
	h.Mux.HandleFunc("/post/", h.middleWare(h.PostPage))

	h.Mux.HandleFunc("/reaction/post/", h.middleWare(h.reactionPost))
	h.Mux.HandleFunc("/reaction/comment/", h.middleWare(h.reactionComment))

	h.Mux.HandleFunc("/myPosts", h.middleWare(h.myPosts))
	h.Mux.HandleFunc("/likedPosts", h.middleWare(h.GetMyLikedPost))

	h.Mux.HandleFunc("/post/edit/", h.middleWare(h.editPost))
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	h.Mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
