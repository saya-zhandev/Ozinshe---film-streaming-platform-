package handlers

import (
	"Ozinshe---film-streaming-platform-/models"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	// "github.com/saya-zhandev/Ozinshe---film-streaming-platform-/models"
)

type MovieHandlers struct {
	db map[int]models.Movie //creates the dictionary of movies: integer to the field of info in the struct
}

func NewMovieHandlers() *MovieHandlers {
	return &MovieHandlers{
		db: map[int]models.Movie{
			//initial ize the db - important before use

			1: {
				Id:          1,
				Title:       "1+1",
				Description: "Пострадав в результате несчастного случая, богатый аристократ Филипп нанимает в помощники человека, который менее всего подходит для этой работы, – молодого жителя предместья Дрисса, только что освободившегося из тюрьмы. Несмотря на то, что Филипп прикован к инвалидному креслу, Дриссу удается привнести в размеренную жизнь аристократа дух приключений.",
				ReleaseYear: 2011,
				Director:    "Оливье Накаш",
				Rating:      0,
				IsWatched:   false,
				TrailerUrl:  "https://www.youtube.com/watch?v=m95M-I7Ij0o&ab_channel=%D0%9A%D0%B8%D0%BD%D0%BE%D0%92%D0%B8%D1%85%D1%80%D1%8C",
				PosterUrl:   "",
				Genres:      make([]string, 0),
			},
			2: {
				Id:          2,
				Title:       "Интерстеллар",
				Description: "Когда засуха, пыльные бури и вымирание растений приводят человечество к продовольственному кризису, коллектив исследователей и учёных отправляется сквозь червоточину (которая предположительно соединяет области пространства-времени через большое расстояние) в путешествие, чтобы превзойти прежние ограничения для космических путешествий человека и найти планету с подходящими для человечества условиями.",
				ReleaseYear: 2014,
				Director:    "Кристофер Нолан",
				Rating:      0,
				IsWatched:   false,
				TrailerUrl:  "https://www.youtube.com/watch?v=6ybBuTETr3U",
				PosterUrl:   "",
				Genres:      make([]string, 0),
			},
			3: {
				Id:          3,
				Title:       "Побег из Шоушенка",
				Description: "Бухгалтер Энди Дюфрейн обвинён в убийстве собственной жены и её любовника. Оказавшись в тюрьме под названием Шоушенк, он сталкивается с жестокостью и беззаконием, царящими по обе стороны решётки. Каждый, кто попадает в эти стены, становится их рабом до конца жизни. Но Энди, обладающий живым умом и доброй душой, находит подход как к заключённым, так и к охранникам, добиваясь их особого к себе расположения.",
				ReleaseYear: 1994,
				Director:    "Фрэнк Дарабонт",
				Rating:      0,
				IsWatched:   false,
				TrailerUrl:  "https://www.youtube.com/watch?v=kgAeKpAPOYk&ab_channel=%D0%A2%D1%80%D0%B5%D0%B9%D0%BB%D0%B5%D1%80%D1%8B%D0%BA%D1%84%D0%B8%D0%BB%D1%8C%D0%BC%D0%B0%D0%BC",
				PosterUrl:   "",
				Genres:      make([]string, 0),
			},
			4: {
				Id:          4,
				Title:       "Бойцовский клуб",
				Description: "Сотрудник страховой компании страдает хронической бессонницей и отчаянно пытается вырваться из мучительно скучной жизни. Однажды в очередной командировке он встречает некоего Тайлера Дёрдена — харизматического торговца мылом с извращенной философией. Тайлер уверен, что самосовершенствование — удел слабых, а единственное, ради чего стоит жить, — саморазрушение. Проходит немного времени, и вот уже новые друзья лупят друг друга почем зря на стоянке перед баром, и очищающий мордобой доставляет им высшее блаженство. Приобщая других мужчин к простым радостям физической жестокости, они основывают тайный Бойцовский клуб, который начинает пользоваться невероятной популярностью.",
				ReleaseYear: 1999,
				Director:    "Дэвид Финчер",
				Rating:      0,
				IsWatched:   false,
				TrailerUrl:  "https://www.youtube.com/watch?v=C7-7qQ61QHU&ab_channel=%D0%A2%D1%80%D0%B5%D0%B9%D0%BB%D0%B5%D1%80%D1%8B%D0%BA%D1%84%D0%B8%D0%BB%D1%8C%D0%BC%D0%B0%D0%BC",
				PosterUrl:   "",
				Genres:      make([]string, 0),
			},
			5: {
				Id:          5,
				Title:       "Остров проклятых",
				Description: "Два американских судебных пристава отправляются на один из островов в штате Массачусетс, чтобы расследовать исчезновение пациентки клиники для умалишенных преступников. При проведении расследования им придется столкнуться с паутиной лжи, обрушившимся ураганом и смертельным бунтом обитателей клиники.",
				ReleaseYear: 2009,
				Director:    "Мартин Скорсезе",
				Rating:      0,
				IsWatched:   false,
				TrailerUrl:  "https://www.youtube.com/watch?v=_l7R9Rz5URw&ab_channel=%D0%A2%D1%80%D0%B5%D0%B9%D0%BB%D0%B5%D1%80%D1%8B%D0%BA%D1%84%D0%B8%D0%BB%D1%8C%D0%BC%D0%B0%D0%BC",
				PosterUrl:   "",
				Genres:      make([]string, 0),
			},
		},
	}
}

func (h *MovieHandlers) FindById(c *gin.Context) {
	idStr := c.Param("id") // parameter functions from the gin context
	id, err := strconv.Atoi(idStr) //converting the id string to the integer 
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid Movie Id"))
		return //if the id is not valid case
	}	

	originalMovie, ok := h.db[id]
	if !ok {
		c.JSON(http.StatusBadRequest, models.NewApiError("Movie not found"))
		return
	}//checking whether the movie id is in the database

	c.JSON(http.StatusOK, originalMovie)//returning the status ok with the movie name
}

func (h *MovieHandlers) FindAll(c *gin.Context) {
	movies := make([]models.Movie, 0, len(h.db)) //make slice with the length of the database
	for _, v := range h.db {                     //gets the value, ignoring the id, and loops through all elements in the database
		movies = append(movies, v)
	}

	c.JSON(http.StatusOK, movies)
}

func (h *MovieHandlers) Create(c *gin.Context) {
	var m models.Movie

	err := c.BindJSON(&m)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Could not bind json"))
		return
	}

	id := len(h.db) + 1 //generate the id for each of the new movie

	m.Id = id                    //assign the id to the m variable
	m.Genres = make([]string, 0) //initializes the slice with empty value to actually present and save it there as values

	h.db[id] = m //saving up the value to the movie

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *MovieHandlers) Update(c *gin.Context) { //creating the function for the updating the films
	idStr := c.Param("id")         //c.param is the gin method for getting the parameter from the url - built in function
	id, err := strconv.Atoi(idStr) //converting the id string to the integer
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid Movie Id")) //when the id is not valid and can't be converted to int
		return
	}

	originalMovie, ok := h.db[id] //checks if the movie with the id exists in the DATABASE, h.db is the map structure, original movie (the movie itself) and ok is the boolean
	if !ok {
		c.JSON(http.StatusBadRequest, models.NewApiError("Movie not found"))
		return
	}

	var updatedMovie models.Movie   //created the variable for the updated movie, as a NEW request from the CLIENT
	err = c.BindJSON(&updatedMovie) //getting the json data for the updated movie
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Could not bind json"))
		return
	}

	originalMovie.Title = updatedMovie.Title
	originalMovie.Description = updatedMovie.Description
	originalMovie.ReleaseYear = updatedMovie.ReleaseYear
	originalMovie.Director = updatedMovie.Director
	originalMovie.Rating = updatedMovie.Rating
	// originalMovie.IsWatched = updatedMovie.IsWatched
	// originalMovie.TrailerUrl = updatedMovie.TrailerUrl

	h.db[id] = originalMovie //saving the updated movie to the db with the new values
	//originalMovie apparently contains both of the new and old info so we have to save it back once more to the db
	c.Status(http.StatusOK)
}

func (h *MovieHandlers) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewApiError("Invalid Movie Id"))
		return
	}

	delete(h.db, id)

	c.Status(http.StatusOK)
}
