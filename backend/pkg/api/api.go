package api

import (
	"drizzy-backend/pkg/db/models"
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
)

func StartAPI(pgdb *pg.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))

	r.Route("/bars", func(r chi.Router) {
		r.Get("/", GetBars)
		r.Get("/random", GetTwoRandomBars)
		r.Get("/leaderboard", GetLeaderboard)
		r.Patch("/{id}/win", UpdateBar)
		r.Post("/", CreateBar)
		r.Delete("/{id}", DeleteBar)
	})

	r.Route("/report", func(r chi.Router) {
		r.Get("/best_year", GetYear)
		r.Get("/worst_year", GetYear)
		r.Get("/best_album", GetAlbum)
		r.Get("/worst_album", GetAlbum)
		r.Get("/best_song", GetSong)
		r.Get("/worst_song", GetSong)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server running"))
	})

	return r
}

type GetBarsResponse struct {
	Success bool          `json:"success"`
	Error   string        `json:"error"`
	Bars    []*models.Bar `json:"bars"`
}

func GetBars(w http.ResponseWriter, r *http.Request) {
	pgdb, ok := r.Context().Value("DB").(*pg.DB)

	if !ok {
		res := &GetBarsResponse{
			Success: false,
			Error:   "couldn't get DB from context",
			Bars:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bars, err := models.GetBars(pgdb)
	if err != nil {
		res := &GetBarsResponse{
			Success: false,
			Error:   err.Error(),
			Bars:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// SUCCESS
	res := &GetBarsResponse{
		Success: true,
		Error:   "",
		Bars:    bars,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding bars: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type GetLeaderboardResponse struct {
	Success bool          `json:"success"`
	Error   string        `json:"error"`
	Bars    []*models.Bar `json:"bars"`
}

func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	pgdb, ok := r.Context().Value("DB").(*pg.DB)

	if !ok {
		res := &GetBarsResponse{
			Success: false,
			Error:   "couldn't get DB from context",
			Bars:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bars, err := models.GetLeaderboard(pgdb)
	if err != nil {
		res := &GetBarsResponse{
			Success: false,
			Error:   err.Error(),
			Bars:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// SUCCESS
	res := &GetBarsResponse{
		Success: true,
		Error:   "",
		Bars:    bars,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding bars: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type GetRandomBarResponse struct {
	Success bool          `json:"success"`
	Error   string        `json:"error"`
	Bars    []*models.Bar `json:"bars"`
}

func GetTwoRandomBars(w http.ResponseWriter, r *http.Request) {
	pgdb, ok := r.Context().Value("DB").(*pg.DB)

	if !ok {
		res := &GetBarsResponse{
			Success: false,
			Error:   "couldn't get DB from context",
			Bars:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bars, err := models.GetTwoRandomBars(pgdb)
	if err != nil {
		res := &GetBarsResponse{
			Success: false,
			Error:   err.Error(),
			Bars:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// SUCCESS
	res := &GetBarsResponse{
		Success: true,
		Error:   "",
		Bars:    bars,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding bars: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type UpdateBarResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func UpdateBar(w http.ResponseWriter, r *http.Request) {
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &UpdateBarResponse{
			Success: false,
			Error:   "couldn't get DB from context",
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		res := &UpdateBarResponse{
			Success: false,
			Error:   err.Error(),
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = models.UpdateBar(pgdb, id)
	if err != nil {
		res := &UpdateBarResponse{
			Success: false,
			Error:   err.Error(),
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// SUCCESS
	res := &UpdateBarResponse{
		Success: true,
		Error:   "",
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding bars: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type CreateBarRequest struct {
	ID       int64  `json:"bar"`
	Lyrics   string `json:"lyrics"`
	SongName string `json:"song_name"`
}

type CreateBarResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func CreateBar(w http.ResponseWriter, r *http.Request) {
	req := &CreateBarRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res := &CreateBarResponse{
			Success: false,
			Error:   err.Error(),
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &UpdateBarResponse{
			Success: false,
			Error:   "couldn't get DB from context",
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = models.CreateBar(pgdb, &models.Bar{
		ID:       req.ID,
		Lyrics:   req.Lyrics,
		Wins:     0,
		SongName: req.SongName,
	})
	if err != nil {
		res := &CreateBarResponse{
			Success: false,
			Error:   err.Error(),
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// SUCCESS
	res := &CreateBarResponse{
		Success: true,
		Error:   "",
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding bars: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type DeleteBarResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func DeleteBar(w http.ResponseWriter, r *http.Request) {
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &DeleteBarResponse{
			Success: false,
			Error:   "couldn't get DB from context",
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		res := &DeleteBarResponse{
			Success: false,
			Error:   err.Error(),
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = models.DeleteBar(pgdb, id)
	if err != nil {
		res := &DeleteBarResponse{
			Success: false,
			Error:   err.Error(),
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// SUCCESS
	res := &DeleteBarResponse{
		Success: true,
		Error:   "",
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding bars: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// REPORT
type GetYearResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Year    int64  `json:"year"`
}

func GetYear(w http.ResponseWriter, r *http.Request) {
	pgdb, ok := r.Context().Value("DB").(*pg.DB)

	if !ok {
		res := &GetYearResponse{
			Success: false,
			Error:   "couldn't get DB from context",
			Year:    -1,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var direction string
	if path.Base(r.URL.Path)[0] == 'b' {
		direction = "DESC"
	} else {
		direction = "ASC"
	}
	year, err := models.GetYearReport(pgdb, direction)
	if err != nil {
		res := &GetYearResponse{
			Success: false,
			Error:   err.Error(),
			Year:    -1,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// SUCCESS
	res := &GetYearResponse{
		Success: true,
		Error:   "",
		Year:    year,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding bars: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type GetAlbumResponse struct {
	Success bool          `json:"success"`
	Error   string        `json:"error"`
	Album   *models.Album `json:"album"`
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	pgdb, ok := r.Context().Value("DB").(*pg.DB)

	if !ok {
		res := &GetAlbumResponse{
			Success: false,
			Error:   "couldn't get DB from context",
			Album:   nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var direction string
	if path.Base(r.URL.Path)[0] == 'b' {
		direction = "DESC"
	} else {
		direction = "ASC"
	}
	album, err := models.GetAlbumReport(pgdb, direction)
	if err != nil {
		res := &GetAlbumResponse{
			Success: false,
			Error:   err.Error(),
			Album:   nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// SUCCESS
	res := &GetAlbumResponse{
		Success: true,
		Error:   "",
		Album:   album,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding bars: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type GetSongResponse struct {
	Success bool         `json:"success"`
	Error   string       `json:"error"`
	Song    *models.Song `json:"song"`
}

func GetSong(w http.ResponseWriter, r *http.Request) {
	pgdb, ok := r.Context().Value("DB").(*pg.DB)

	if !ok {
		res := &GetSongResponse{
			Success: false,
			Error:   "couldn't get DB from context",
			Song:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var direction string
	if path.Base(r.URL.Path)[0] == 'b' {
		direction = "DESC"
	} else {
		direction = "ASC"
	}
	song, err := models.GetSongReport(pgdb, direction)
	if err != nil {
		res := &GetSongResponse{
			Success: false,
			Error:   err.Error(),
			Song:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// SUCCESS
	res := &GetSongResponse{
		Success: true,
		Error:   "",
		Song:    song,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding bars: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
