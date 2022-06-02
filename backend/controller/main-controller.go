package controller

import (
	"log"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_movie2/config"
	"github.com/nikitamirzani323/isbpanel_movie2/helpers"
)

type response struct {
	Status int         `json:"status"`
	Record interface{} `json:"record"`
}
type responseerror struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type responseinit struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}

const PATH string = config.PATH_API
const Fieldmovie_home_redis = "LISTMOVIE_ISBPANEL"

type Model_moviecategory struct {
	Movie_idcategory string      `json:"movie_idcategory"`
	Movie_category   string      `json:"movie_category"`
	Movie_list       interface{} `json:"movie_list"`
}
type Model_movie struct {
	Movie_id        int         `json:"movie_id"`
	Movie_type      string      `json:"movie_type"`
	Movie_title     string      `json:"movie_title"`
	Movie_label     string      `json:"movie_label"`
	Movie_thumbnail string      `json:"movie_thumbnail"`
	Movie_video     interface{} `json:"movie_video"`
}
type Model_movievideo struct {
	Movie_src string `json:"movie_src"`
}

func Init(c *fiber.Ctx) error {
	type payload_init struct {
		Client_hostname string `json:"client_hostname"`
	}
	client := new(payload_init)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("HOSTNAME-CLIENT: ", client.Client_hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responseinit{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": client.Client_hostname,
		}).
		Post(PATH + "api/init")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responseinit)
	return c.JSON(fiber.Map{
		"status": result.Status,
		"token":  result.Token,
		"time":   time.Since(render_page).String(),
	})
}
func Listmovie(c *fiber.Ctx) error {
	type payload_listmovie struct {
		Client_hostname string `json:"client_hostname"`
	}
	client := new(payload_listmovie)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	render_page := time.Now()

	var obj Model_moviecategory
	var arraobj []Model_moviecategory
	resultredis, flag := helpers.GetRedis(Fieldmovie_home_redis)
	jsonredis := []byte(resultredis)
	// message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		movie_idcategory, _ := jsonparser.GetString(value, "movie_idcategory")
		movie_category, _ := jsonparser.GetString(value, "movie_category")
		movie_list, _, _, _ := jsonparser.Get(value, "movie_list")
		var objchild Model_movie
		var arraobjchild []Model_movie
		jsonparser.ArrayEach(movie_list, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			movie_id, _ := jsonparser.GetInt(value, "movie_id")
			movie_type, _ := jsonparser.GetString(value, "movie_type")
			movie_title, _ := jsonparser.GetString(value, "movie_title")
			movie_label, _ := jsonparser.GetString(value, "movie_label")
			movie_thumbnail, _ := jsonparser.GetString(value, "movie_thumbnail")
			movie_video, _, _, _ := jsonparser.Get(value, "movie_video")
			var objmoviesrc Model_movievideo
			var arraobjmoviesrc []Model_movievideo
			jsonparser.ArrayEach(movie_video, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				movie_src, _ := jsonparser.GetString(value, "movie_src")
				objmoviesrc.Movie_src = movie_src
				arraobjmoviesrc = append(arraobjmoviesrc, objmoviesrc)
			})
			objchild.Movie_id = int(movie_id)
			objchild.Movie_type = movie_type
			objchild.Movie_title = movie_title
			objchild.Movie_label = movie_label
			objchild.Movie_thumbnail = movie_thumbnail
			objchild.Movie_video = arraobjmoviesrc
			arraobjchild = append(arraobjchild, objchild)
		})

		obj.Movie_idcategory = movie_idcategory
		obj.Movie_category = movie_category
		obj.Movie_list = arraobjchild
		arraobj = append(arraobj, obj)
	})
	if !flag {
		axios := resty.New()
		resp, err := axios.R().
			SetAuthToken(token[1]).
			SetResult(response{}).
			SetError(responseerror{}).
			SetHeader("Content-Type", "application/json").
			SetBody(map[string]interface{}{
				"client_hostname": client.Client_hostname,
			}).
			Post(PATH + "api/movie")
		if err != nil {
			log.Println(err.Error())
		}
		result := resp.Result().(*response)

		if result.Status == 200 {
			helpers.SetRedis(Fieldmovie_home_redis, result, time.Minute*300)
			log.Println("MOVIE MYSQL")
			return c.JSON(fiber.Map{
				"status": result.Status,
				"record": result.Record,
				"time":   time.Since(render_page).String(),
			})
		} else {
			result_error := resp.Error().(*responseerror)
			return c.JSON(fiber.Map{
				"status":  result_error.Status,
				"message": result_error.Message,
				"time":    time.Since(render_page).String(),
			})
		}
	} else {
		log.Println("MOVIE CACHE")
		return c.JSON(fiber.Map{
			"status": fiber.StatusOK,
			"record": arraobj,
			"time":   time.Since(render_page).String(),
		})
	}

}

type clientSeason struct {
	Movie_id int `json:"movie_id"`
}
type clientEpisode struct {
	Season_id int `json:"season_id"`
}

func Listseason(c *fiber.Ctx) error {
	client := new(clientSeason)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"movie_id":        client.Movie_id,
		}).
		Post(PATH + "api/season")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": result.Status,
			"record": result.Record,
			"time":   time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Listepisode(c *fiber.Ctx) error {
	client := new(clientEpisode)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"season_id":       client.Season_id,
		}).
		Post(PATH + "api/episode")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": result.Status,
			"record": result.Record,
			"time":   time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
