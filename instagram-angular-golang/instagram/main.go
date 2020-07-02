package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "instagram"
)

var jwtKey = []byte("my_secret_key")

//NewUser structure
type NewUser struct {
	Firstname string `json"firstname"`
	LastName  string `json"lastname"`
	Email     string `json"email"`
	UserName  string `json"username"`
	Password  string `json"password"`
}

//JwtToken struct
type JwtToken struct {
	Token string `json:"token"`
}

//Exception struct
type Exception struct {
	Message string `json:"message"`
}

//Claims struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//Posts structure
type Posts struct {
	PostID     int    `json"postid"`
	PostURL    string `json"posturl"`
	PostLiked  bool   `json"postliked"`
	PostSaved  bool   `json"postsaved"`
	PosterName string `json"postername"`
}

//Stories structure
type Stories struct {
	StoryID    int    `json"storyid"`
	StoryURL   string `json"storyurl"`
	StoryOwner string `json"storyowner"`
	UniqueID   string `json"uniqueid"`
}

//SuggestionTable structure
type SuggestionTable struct {
	ID         int    `json"id"`
	UserName   string `json"username"`
	Followed   bool   `json"followed"`
	ProfileURL string `json"profileurl"`
}

var newdata []NewUser = []NewUser{}
var newpost []Posts = []Posts{}
var postcount int = 0

func main() {
	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
	})
	handler := c.Handler(router)
	router.Handle("/", middleware(http.HandlerFunc(checkcookie))).Methods("GET")
	router.HandleFunc("/removecookie", removecookie)
	router.HandleFunc("/newuser", newuser).Methods("POST")                                 //entering new user
	router.HandleFunc("/getuser/{username}/{password}", getuser).Methods("GET")            //retrieving user details
	router.Handle("/newpost", middleware(http.HandlerFunc(newpostupload))).Methods("POST") //upload new post
	router.Handle("/getpost", middleware(http.HandlerFunc(getpost))).Methods("GET")
	router.Handle("/suggestiontable/{username}", middleware(http.HandlerFunc(suggestionTable))).Methods("GET")
	router.Handle("/requesting/{username}/{friendname}", middleware(http.HandlerFunc(requesting))).Methods("GET")
	router.Handle("/accepting/{username}/{friendname}", middleware(http.HandlerFunc(accepting))).Methods("GET")
	router.Handle("/checklist/{username}", middleware(http.HandlerFunc(checklist))).Methods("GET")
	router.Handle("/friends/{username}", middleware(http.HandlerFunc(friends))).Methods("GET")
	router.Handle("/getprofile/{username}", middleware(http.HandlerFunc(getprofile))).Methods("GET")
	router.Handle("/getabout/{username}", middleware(http.HandlerFunc(getabout))).Methods("GET")
	//api calls
	router.Handle("/loadpost", middleware(http.HandlerFunc(loadpost))).Methods("GET")       //posts
	router.Handle("/loadstories", middleware(http.HandlerFunc(loadstories))).Methods("GET") //stories
	srv := &http.Server{
		Handler: handler,
		Addr:    ":8080",
	}

	log.Fatal(srv.ListenAndServe())
}

//checkcookie functio for not authenticate routers
func checkcookie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	json.NewEncoder(w).Encode(true)
}

//removecookie for logout
func removecookie(w http.ResponseWriter, r *http.Request) {
	db := dbconnect()
	_, dberr := db.Exec(`Delete from friends_suggestion`)
	if dberr != nil {
		log.Fatalf("Unable to execute the query. %v", dberr)
	}
	expire := time.Now().Add(-7 * 24 * time.Hour)
	c, err := r.Cookie("token")
	if err == nil {
		cookie := http.Cookie{
			Name:    "token",
			Value:   c.Value,
			Expires: expire,
		}
		http.SetCookie(w, &cookie)
	}
}

//onsignup new user
func newuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	db := dbconnect()
	var id int64
	w.Header().Set("Content-Type", "application/json")
	var user NewUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != io.EOF {
		sqlstatement := `INSERT INTO instagram_user (firstname,lastname,username,email,password) VALUES ($1,$2,$3,$4,$5) RETURNING user_id`
		dberr := db.QueryRow(sqlstatement, user.Firstname, user.LastName, user.UserName, user.Email, user.Password).Scan(&id)
		if dberr != nil {
			log.Fatalf("Unable to execute the query. %v", dberr)
		}
		expirationTime := time.Now().Add(24 * time.Minute)
		// Create the JWT claims, which includes the username and expiry time
		claims := &Claims{
			Username: user.UserName,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Create the JWT string
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			fmt.Println(err)
		}

		newdata = append(newdata, user)
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
			Path:    "/",
		})
		json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
		//json.NewEncoder(w).Encode(newdata)
	}

}

//on user login fetching user details for validation
func getuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	username := mux.Vars(r)["username"]
	password := mux.Vars(r)["password"]
	db := dbconnect()
	query := `SELECT * FROM instagram_user 	WHERE username = $1 and password = $2`
	dbres, dberr := db.Query(query, username, password)
	if dberr != nil {
		log.Fatalf("Unable to execute the query. %v", dberr)
	}
	if dbres.Next() {
		expirationTime := time.Now().Add(24 * time.Minute)
		// Create the JWT claims, which includes the username and expiry time
		claims := &Claims{
			Username: username,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}
		// Declare the token with the algorithm used for signing, and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Create the JWT string
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
		}

		// Finally, we set the client cookie for "token" as the JWT we just generated
		// we also set an expiry time which is the same as the token itself
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
			Path:    "/",
		})
		json.NewEncoder(w).Encode(username)
	}
}

//suggestiontable
func suggestionTable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	username := mux.Vars(r)["username"]
	var suggestions []SuggestionTable = []SuggestionTable{}
	db := dbconnect()
	query := `SELECT friends FROM user_friends where username = $1`
	query2, dberr := db.Query(query, username)
	if dberr != nil {
		log.Fatalf("Unable to execute the query. %v", dberr)
	}
	var userfriendsarray []string
	var friendoffriendsarray []string
	var suggestedarray []string
	row, erro := db.Query(`select username from friends_suggestion`)
	if erro != nil {
		log.Fatalf("Unable to execute the query. %v", r)
	}
	var usr string
	for row.Next() {
		err := row.Scan(&usr)
		if err != nil {
			panic(err)
		}
		suggestedarray = append(suggestedarray, usr)
	}
	for query2.Next() {
		err := query2.Scan(pq.Array(&userfriendsarray))
		if err != nil {
			panic(err)
		}
		for i := range userfriendsarray {
			q, e := db.Query(`SELECT friends FROM user_friends where username = $1`, userfriendsarray[i])
			if e != nil {
				log.Fatalf("Unable to execute the query. %v", e)
			}
			for q.Next() {
				er := q.Scan(pq.Array(&friendoffriendsarray))
				if er != nil {
					panic(er)
				}
				for j := range friendoffriendsarray {
					var flag = 0
					if friendoffriendsarray[j] != username {
						for k := range suggestedarray {
							if suggestedarray[k] == friendoffriendsarray[j] {
								flag = 1
							}
						}
						if flag == 0 {
							_, rr := db.Exec(`Insert into friends_suggestion (username,followed) values ($1,$2)`, friendoffriendsarray[j], false)
							if rr != nil {
								panic(rr)
							}
						}
					}
				}
			}
		}
	}

	rows, errr := db.Query(`select * from friends_suggestion`)
	if errr != nil {
		log.Fatalf("Unable to execute the query. %v", r)
	}
	for rows.Next() {
		suggestion := SuggestionTable{}
		err := rows.Scan(&suggestion.ID, &suggestion.UserName, &suggestion.Followed)
		row, dberr := db.Query(`SELECT profileurl from user_details where username = $1`, suggestion.UserName)
		for row.Next() {
			er := row.Scan(&suggestion.ProfileURL)
			if dberr != nil || er != nil {
				fmt.Println("er")
			}
		}
		if err != nil {
			panic(err)
		}
		suggestions = append(suggestions, suggestion)
	}
	err := rows.Err()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(suggestions)
}

//friends
func requesting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	username := mux.Vars(r)["username"]
	friendname := mux.Vars(r)["friendname"]
	db := dbconnect()
	query := `Select requested from user_friends where username = $1`
	rows, dberr := db.Query(query, friendname)
	if dberr != nil {
		log.Fatalf("Unable to execute the query. %v", dberr)
	}
	type request struct {
		RequestedArray []string `json"requestedarray`
	}
	var req = request{}
	for rows.Next() {
		er := rows.Scan(pq.Array(&req.RequestedArray))
		if er != nil {
			panic(er)
		}
	}
	req.RequestedArray = append(req.RequestedArray, username)
	_, er := db.Exec(`Update user_friends set requested = $1 where username = $2`, pq.Array(req.RequestedArray), friendname)
	if er != nil {
		panic(er)
	}
	json.NewEncoder(w).Encode(req)
}

//friends
func accepting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	username := mux.Vars(r)["username"]
	friendname := mux.Vars(r)["friendname"]
	db := dbconnect()
	type request struct {
		RequestedArray []string `json"requestedarray`
	}
	var req = request{}
	var AcceptedArray []string
	dbqu, dber := db.Query(`Select friends from user_friends where username = $1`, username)
	if dber != nil {
		panic(dber)
	}
	for dbqu.Next() {
		er := dbqu.Scan(pq.Array(&AcceptedArray))
		if er != nil {
			panic(er)
		}
	}
	var FriendArray []string
	dbquF, dberF := db.Query(`Select friends from user_friends where username = $1`, friendname)
	if dberF != nil {
		panic(dber)
	}
	for dbquF.Next() {
		er := dbquF.Scan(pq.Array(&FriendArray))
		if er != nil {
			panic(er)
		}
	}
	FriendArray = append(FriendArray, username)
	query := `Select requested from user_friends where username = $1`
	rows, dberr := db.Query(query, username)
	if dberr != nil {
		log.Fatalf("Unable to execute the query. %v", dberr)
	}
	for rows.Next() {
		er := rows.Scan(pq.Array(&req.RequestedArray))
		if er != nil {
			panic(er)
		}
	}
	for i := range req.RequestedArray {
		if req.RequestedArray[i] == friendname {
			AcceptedArray = append(AcceptedArray, friendname)
			req.RequestedArray[i] = req.RequestedArray[len(req.RequestedArray)-1]
			req.RequestedArray[len(req.RequestedArray)-1] = ""
		}
	}
	req.RequestedArray = req.RequestedArray[:len(req.RequestedArray)-1]
	_, er := db.Exec(`Update user_friends set requested = $1 where username = $2`, pq.Array(req.RequestedArray), username)
	if er != nil {
		panic(er)
	}
	_, e := db.Exec(`Update user_friends set friends = $1 where username = $2`, pq.Array(AcceptedArray), username)
	if e != nil {
		panic(e)
	}
	_, el := db.Exec(`Update user_friends set friends = $1 where username = $2`, pq.Array(FriendArray), friendname)
	if el != nil {
		panic(el)
	}
	json.NewEncoder(w).Encode(req)
}

//friends
func friends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	username := mux.Vars(r)["username"]
	db := dbconnect()
	var Friendsarray []string
	dbquery, dberror := db.Query(`Select friends from user_friends where username = $1`, username)
	if dberror != nil {
		panic(dberror)
	}
	for dbquery.Next() {
		er := dbquery.Scan(pq.Array(&Friendsarray))
		if er != nil {
			panic(er)
		}
	}
	json.NewEncoder(w).Encode(Friendsarray)
}

//about
func getabout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	username := mux.Vars(r)["username"]
	db := dbconnect()
	var about string
	dbquery, dberror := db.Query(`Select about from user_details where username = $1`, username)
	if dberror != nil {
		panic(dberror)
	}
	for dbquery.Next() {
		er := dbquery.Scan(&about)
		if er != nil {
			panic(er)
		}
	}
	json.NewEncoder(w).Encode(about)
}

//checking requesting list
func checklist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	username := mux.Vars(r)["username"]
	db := dbconnect()
	dbqu, dber := db.Query(`Select requested from user_friends where username = $1`, username)
	if dber != nil {
		panic(dber)
	}
	var RequestedList []string
	for dbqu.Next() {
		er := dbqu.Scan(pq.Array(&RequestedList))
		if er != nil {
			panic(er)
		}
	}
	json.NewEncoder(w).Encode(RequestedList)
}

//new post creation
func newpostupload(w http.ResponseWriter, r *http.Request) {
	postcount++
	w.Header().Set("Content-Type", "application/json")
	var post Posts
	err := json.NewDecoder(r.Body).Decode(&post)
	post.PostID = postcount
	if err != io.EOF {
		newpost = append(newpost, post)
	}
	json.NewEncoder(w).Encode(newpost)
}

func getprofile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	username := mux.Vars(r)["username"]
	db := dbconnect()
	query := `Select profileurl from user_details where username = $1`
	rows, dberr := db.Query(query, username)
	if dberr != nil {
		panic(dberr)
		log.Fatalf("Unable to execute the query. %v", dberr)
	}
	var ProfileURL string
	for rows.Next() {
		er := rows.Scan(&ProfileURL)
		if er != nil {
			panic(er)
		}

	}
	json.NewEncoder(w).Encode(ProfileURL)
}

//posts retrieving
func getpost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newpost)
}

//api call for getting post - predeclared
func loadpost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	var apipost = []Posts{
		Posts{PostID: 1, PostURL: "https://wallpaperplay.com/walls/full/c/8/5/128260.jpg", PostLiked: false, PostSaved: false, PosterName: "jeewa"},
		Posts{PostID: 2, PostURL: "https://www.chromethemer.com/google-chrome/backgrounds/download/random-hd-background-for-google-chrome-ct1404.jpg", PostLiked: false, PostSaved: false, PosterName: "dev"},
		Posts{PostID: 3, PostURL: "https://thewallpaper.co//wp-content/uploads/2016/10/adidas-logo-red-ferrari-nature-strike-random-wallpaper-hd-wallpapers-desktop-images-download-free-windows-wallpapers-amazing-colourful-4k-lovely-2560x1600.jpg", PostLiked: false, PostSaved: false, PosterName: "dev"},
		Posts{PostID: 4, PostURL: "https://c1.wallpaperflare.com/preview/800/288/19/cube-play-random-luck.jpg", PostLiked: false, PostSaved: false, PosterName: "hello"},
	}
	json.NewEncoder(w).Encode(apipost)
}

//api call for getting stories - predeclared
func loadstories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	var apistories = []Stories{
		Stories{StoryID: 1, StoryURL: "https://www.pixelstalk.net/wp-content/uploads/images1/HD-artistic-backgrounds.jpg", StoryOwner: "jeewa", UniqueID: "strory1"},
		Stories{StoryID: 2, StoryURL: "https://data.1freewallpapers.com/download/peek-out-hide-hat-eyes-art.png", StoryOwner: "hello", UniqueID: "strory2"},
		Stories{StoryID: 3, StoryURL: "https://images2.alphacoders.com/303/thumb-350-30312.png", StoryOwner: "dev", UniqueID: "strory3"},
		Stories{StoryID: 4, StoryURL: "https://lh3.googleusercontent.com/proxy/TJ9fWFM4jqvxGgBWGgbf69RFXye6s8DxBe7V8b-3iRz8hMZS0eqgJUC8HQOxWI7V5sG5tGl_TBrHbdWv-aav4Cj47omLQ-otex2SZPXiysexPqIliSQIJoGH6WiXYT4Kt62Cs301pzga=s0-d", StoryOwner: "jeewa", UniqueID: "strory4"},
		Stories{StoryID: 5, StoryURL: "https://cdn.wallpaperhi.com/1920x1200/20120227/artistic%20coffee%20funny%20owls%201920x1200%20wallpaper_www.wallpaperhi.com_76.jpg", StoryOwner: "hello", UniqueID: "strory5"},
	}
	json.NewEncoder(w).Encode(apistories)
}

//db
func dbconnect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

//Middleware
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tknStr := c.Value
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

/*create table instagram_user
(
 user_id serial primary key,
 firstname varchar,
 lastname varchar,
 username varchar,
 email varchar,
 password varchar);
*/
