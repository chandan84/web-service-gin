package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

type movie struct {
    ID            string  `json:"id"`
    Title         string  `json:"title"`
    Genre         string  `json:"genre"`
    ReleasedDate  string  `json:"releaseDate"`
    Director      string  `json:"director"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var movies = []movie{
    {ID: "1", Title: "Catch me if you can", Genre: "Comedy", ReleasedDate: "2002/02/14", Director: "Steven Spielberg"},
    {ID: "2", Title: "Terminator", Genre: "Action", ReleasedDate: "1985/06/01", Director: "James Cameron"},
    {ID: "3", Title: "Terminator: Judgement Day", Genre: "Action", ReleasedDate: "1996/10/01", Director: "James Cameron"},
    {ID: "4", Title: "Avataar", Genre: "Action/Drama", ReleasedDate: "2009/08/01", Director: "James Cameron"},
}

//getMovies responds with the list of all movies as JSON
func getMovies(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, movies)
}

// postMovies adds an album from JSON received in the request body.
func postMovies(c *gin.Context) {
    var newMovie movie

    // Call BindJSON to bind the received JSON to
    // newMovie.
    if err := c.BindJSON(&newMovie); err != nil {
        return
    }

    // Add the new album to the slice.
    movies = append(movies, newMovie)
    c.IndentedJSON(http.StatusCreated, newMovie)
}

// getMovieByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getMovieByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of movies, looking for
    // an movie whose ID value matches the parameter.
    for _, a := range movies {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "movie not found"})
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.GET("/movies", getMovies)
    router.GET("/movies/:id", getMovieByID)
    router.POST("/movies", postMovies)

    router.Run("localhost:8080")
}