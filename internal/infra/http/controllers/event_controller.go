package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Cry-coder/smpl_srvr/internal/domain/event"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/upper/db/v4"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type EventController struct {
	service *event.Service
}

const (
	sessionKeyUserId   = "userId"
	sessionKeyUserName = "userName"
	sessionKeyUserRole = "userRole"
)

// init sessions
var SessionManager *scs.SessionManager

func init() {
	SessionManager = scs.New()
	SessionManager.Lifetime = 24 * time.Hour
	SessionManager.Cookie.Persist = true
	SessionManager.Cookie.Name = "HelpDesk"
	SessionManager.Cookie.SameSite = http.SameSiteStrictMode
	SessionManager.Store = postgresstore.New(event.Pool)
}

func NewEventController(s *event.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) AuthRequired(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := SessionManager.Get(r.Context(), sessionKeyUserId)
		if userId == nil || SessionManager.Get(r.Context(), sessionKeyUserRole).(string) != "user" {

			http.Redirect(w, r, "http://localhost:8007/v1/login", http.StatusSeeOther)
			return
		}

		w.Header().Add("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	}
}

func (c *EventController) AdminAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := SessionManager.Get(r.Context(), sessionKeyUserId)

		if userId == nil || SessionManager.Get(r.Context(), sessionKeyUserRole).(string) != "admin" {

			http.Redirect(w, r, "http://localhost:8007/v1/login", http.StatusSeeOther)
			return
		}

		w.Header().Add("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//m := make(map[event.St][]event.Questions)
		events, err := (*c.service).FindAll()

		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s", err)
			}
			return
		}

		err = success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}

		//SessionManager.Put(r.Context(), "r", "Hello from a session!")

	}
}

func (c *EventController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := SessionManager.Get(r.Context(), sessionKeyUserId)
		event, slice, err := (*c.service).FindOne(id.(int64))
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}

		if SessionManager.Get(r.Context(), sessionKeyUserRole).(string) == "admin" {
			err = success(w, event)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		err = success(w, event, slice)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}

func (c *EventController) FindOneQuestion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		qid, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOneQuestion(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOneQuestion(): %s", err)
			}
			return
		}

		y, err := (*c.service).FindOneQuestion(int(qid))
		if err != nil {
			id := SessionManager.Get(r.Context(), sessionKeyUserId)
			_, slice, err := (*c.service).FindOne(id.(int64))
			//fmt.Print(len(*slice))
			if len(*slice) == 0 {
				w.Write([]byte("Seems you have not created questions yet"))
				http.Redirect(w, r, "http://localhost:8007/v1/user/cr", http.StatusSeeOther)
			}
			w.Write([]byte("Seems you trying wrong question id try another."))

			fmt.Printf("EventController.FindOneQuestion(): %s", err)
			return
		}
		err = success(w, y)
		if err != nil {
			fmt.Printf("EventController.FindOneQuestion(): %s", err)
		}
	}
}

func (c *EventController) OneQuestion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		qid, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOneQuestion(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOneQuestion(): %s", err)
			}
			return
		}

		y, err := (*c.service).FindOneQuestion(int(qid))
		if err != nil {
			id := SessionManager.Get(r.Context(), sessionKeyUserId)
			_, slice, err := (*c.service).FindOne(id.(int64))
			//fmt.Print(len(*slice))
			if len(*slice) == 0 {
				w.Write([]byte("Seems you have not created questions yet"))
				http.Redirect(w, r, "http://localhost:8007/v1/user/cr", http.StatusSeeOther)
			}
			w.Write([]byte("Seems you trying wrong question id try another."))

			fmt.Printf("EventController.FindOneQuestion(): %s", err)
			return
		}
		err = success(w, y)
		if err != nil {
			fmt.Printf("EventController.FindOneQuestion(): %s", err)
		}
	}
}

func (c *EventController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Delete(): %s", err)
			}
			return
		}
		err = (*c.service).Delete(id)
		if err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
		}
		err = success(w, "Successful deleted.")
		if err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
		}
	}

}

func (c *EventController) DeleteQuestion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		qid, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.DeleteQuestion(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.DeleteQuestion(): %s", err)
			}
			return
		}
		err = (*c.service).DeleteQuestion(qid)
		if err != nil {
			fmt.Printf("EventController.DeleteQuestion(): %s", err)
		}
		err = success(w, "Successfully deleted.")
		if err != nil {
			fmt.Printf("EventController.DeleteQuestion(): %s", err)
		}
	}
}

func (c *EventController) UserSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var g event.St
		err := json.NewDecoder(r.Body).Decode(&g)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(g.Password), 8)
		g.Password = string(hashedPass)
		g.Id, err = strconv.ParseInt(strconv.Itoa(rand.Intn(100000000)), 10, 64)
		if err != nil {
			fmt.Printf("EventController.Create(): %s", err)
		}
		g.Role = "user"
		check1, err := (*c.service).UserCheck(g.Email)

		if err != nil {
			fmt.Printf("EventController.Create(): %s", err)
		}
		if check1 == true {
			w.Write([]byte("Email already in use."))
			return
		} else {
			res, er := (*c.service).Create(&g)
			if er != nil {
				fmt.Printf("EventController.Create(): %s", er)
			}
			err = created(w, res)
			if er != nil {
				fmt.Printf("EventController.Create(): %s", er)
			}
		}
	}
}

func (c *EventController) AdminSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var g event.St
		err := json.NewDecoder(r.Body).Decode(&g)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(g.Password), 8)
		g.Password = string(hashedPass)
		g.Id, err = strconv.ParseInt(strconv.Itoa(rand.Intn(100)), 10, 64)
		if err != nil {
			fmt.Printf("EventController.Create(): %s", err)
		}
		g.Role = "admin"
		check0, err := (*c.service).AdminCheck()
		if err != nil {
			fmt.Printf("EventController.Create(): %s", err)
		}
		check1, err := (*c.service).UserCheck(g.Email)
		if err != nil {
			fmt.Printf("EventController.Create(): %s", err)
		}

		if check0 == true {
			w.Write([]byte("Admin account already created. Login for further expereince."))
			http.Redirect(w, r, "http://localhost:8007/v1/login", http.StatusSeeOther)
		} else if check1 == true {
			w.Write([]byte("Email already in use. Login for further experience."))
			return
		} else {
			res, er := (*c.service).Create(&g)
			if er != nil {

				fmt.Printf("EventController.Create(): %s", er)
				w.Write([]byte(fmt.Sprintf("EventController.Create(): %s", er)))

			}
			err = created(w, res) // changed without testing
			if er != nil {
				fmt.Printf("EventController.Create(): %s", er)
				w.Write([]byte(fmt.Sprintf("EventController.Create(): %s", er)))
			}
		}

	}
}

func (c *EventController) CreateQuestion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var g event.Questions
		err := json.NewDecoder(r.Body).Decode(&g)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Print("Error while reading request body")
		}
		userId := SessionManager.Get(r.Context(), sessionKeyUserId)
		if userId == nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		g.StId = int(userId.(int64))
		g.CreatedAt = time.Now()
		g.Status = false
		res, err := (*c.service).CreateQuestion(&g)
		if err != nil {
			fmt.Printf("EventController.CreateQuestion(): %s", err)
		}
		err = created(w, res)
	}
}

func (c *EventController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var g event.St
		err := json.NewDecoder(r.Body).Decode(&g)
		//fmt.Println(g)
		if err != nil {
			fmt.Print("Error while reading request body")
		}
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(g.Password), 8)

		if err != nil {
			internalServerError(w, err)
		}
		g.Password = string(hashedPass)
		g.Id = SessionManager.Get(r.Context(), sessionKeyUserId).(int64)
		err = (*c.service).UpdatePass(&g)
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
			err = notFound(w, err)
			if err != nil {
				fmt.Printf("EventController.Delete(): %s", err)
			}
		} else {
			err = success(w, "Successful updated.")
		}
	}
}

func (c *EventController) UpdateQuestion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var g event.Questions
		err := json.NewDecoder(r.Body).Decode(&g)
		if err != nil {
			fmt.Printf("EventController.UpdateQuestion(): %s", err)
			err = notFound(w, err)
			if err != nil {
				fmt.Printf("EventController.UpdateQuestion(): %s", err)
			}
		}
		if g.Id == 0 {
			badRequest(w, errors.New("bad question id"))
			return
		}
		y, err := (*c.service).FindOneQuestion(g.Id)
		if err != nil {
			fmt.Printf("EventController.UpdateQuestion(): %s", err)
			badRequest(w, errors.New("you trying to change unexisting question"))
			return
		}
		userId := SessionManager.Get(r.Context(), sessionKeyUserId).(int64)
		if userId == int64(y.StId) {
			g.StId = int(userId)
			g.CreatedAt = y.CreatedAt
			g.Status = y.Status
			err = (*c.service).UpdateQuestion(&g)
			if err != nil {
				fmt.Printf("EventController.UpdateQuestion(): %s", err)
			}
			err = success(w, "Successful updated.")
			return
		}

	}
}

func (c *EventController) AdminUpdateQuestion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var g event.Questions
		err := json.NewDecoder(r.Body).Decode(&g)
		if err != nil {
			fmt.Printf("EventController.AdminUpdateQuestion(): %s", err)
			err = notFound(w, err)
			if err != nil {
				fmt.Printf("EventController.AdminUpdateQuestion(): %s", err)
			}
		}
		err = (*c.service).UpdateQuestion(&g)
		if err != nil {
			fmt.Printf("EventController.AdminUpdateQuestion(): %s", err)
		}
		err = success(w, "Successful updated.")

	}
}

func (c *EventController) LoginPutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var d event.St
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			er := badRequest(w, err)
			if er != nil {
				fmt.Printf("EventController.LoginPutHandler(): %s", err)
			}
			return
		}
		t, err := (*c.service).GetPass(&d)
		if err != nil {
			if err == db.ErrNoMoreRows {
				w.Write([]byte("Email or/and password are wrong try one more time"))
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(t.Password), []byte(d.Password)); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		}
		err = SessionManager.RenewToken(r.Context())
		if err != nil {
			fmt.Println(err)
		}
		SessionManager.Put(r.Context(), sessionKeyUserId, t.Id)
		SessionManager.Put(r.Context(), sessionKeyUserName, t.Fn)
		SessionManager.Put(r.Context(), sessionKeyUserRole, t.Role)
		http.Redirect(w, r, fmt.Sprintf("http://localhost:8007/v1/"+t.Role+"/profile"), http.StatusSeeOther)
	}
}

func (c *EventController) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("You have no permission to be on requested page."))
		if err != nil {
			fmt.Printf("EventController.LoginHandler(): %s", err)

		}
	}
}

func (c *EventController) LogOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		SessionManager.Remove(r.Context(), sessionKeyUserId)
		SessionManager.Remove(r.Context(), sessionKeyUserName)
		err := SessionManager.Destroy(r.Context())
		if err != nil {
			fmt.Printf("EventController.LogOut(): %s", err)
		}
		er := SessionManager.RenewToken(r.Context())
		if er != nil {
			fmt.Printf("EventController.LogOut(): %s", err)
		}

		http.Redirect(w, r, "http://localhost:8007/v1/login", http.StatusSeeOther)
		return
	}
}

//func (c *EventController) AdminLoginPutHandler() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var d event.St
//		err := json.NewDecoder(r.Body).Decode(&d)
//		if err != nil {
//			er := badRequest(w, err)
//			if er != nil {
//				fmt.Printf("EventController.LoginPutHandler(): %s", err)
//			}
//			return
//		}
//		t, err := (*c.service).GetPass(&d)
//		if err != nil {
//			if err == db.ErrNoMoreRows { //sql.ErrNoRows {
//				w.Write([]byte("Email or/and password are wrong try one more time"))
//				w.WriteHeader(http.StatusUnauthorized)
//				return
//			}
//			fmt.Println(err)
//			w.WriteHeader(http.StatusInternalServerError)
//			return
//		}
//		if err = bcrypt.CompareHashAndPassword([]byte(t.Password), []byte(d.Password)); err != nil {
//			w.WriteHeader(http.StatusUnauthorized)
//		}
//		err = SessionManager.RenewToken(r.Context())
//		if err != nil {
//			fmt.Println(err)
//		}
//		SessionManager.Put(r.Context(), sessionKeyUserId, t.Id)
//		SessionManager.Put(r.Context(), sessionKeyUserName, t.Fn)
//		//SessionManager.Put(r.Context(), sessionKeyUserRole, t.Role)
//		http.Redirect(w, r, "http://localhost:8007/v1/user/account", http.StatusSeeOther)
//	}
//}
//
//func (c *EventController) AdminLoginHandler() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		_, err := w.Write([]byte("Login for further experience."))
//		if err != nil {
//			fmt.Printf("EventController.LoginHandler(): %s", err)
//
//		}
//	}
//}

func (c *EventController) FindOneAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOneQuestion(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOneQuestion(): %s", err)
			}
			return
		}
		userSt, questionSlice, err := (*c.service).FindOne(id)
		if err != nil {
			fmt.Printf("EventController.FindOneQuestion(): %s", err)
		}
		err = success(w, userSt, questionSlice)
		if err != nil {
			internalServerError(w, err)
			fmt.Printf("EventController.FindOneQuestion(): %s", err)
		}
	}
}

func (c *EventController) FindAllQuestions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		i, err := (*c.service).FindAllQuestions()
		if err != nil {
			internalServerError(w, err)
			fmt.Printf("EventController.FindAllQuestions(): %s", err)
		}
		err = success(w, i)
		if err != nil {
			fmt.Printf("EventController.FindAllQuestions(): %s", err)
		}
	}
}

func (c *EventController) FindOneQuestionAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOneQuestion(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOneQuestion(): %s", err)
			}
			return
		}

		y, err := (*c.service).FindOneQuestion(int(id))
		err = success(w, y)
		if err != nil {
			internalServerError(w, err)
		}
	}
}
