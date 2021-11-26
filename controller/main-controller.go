package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_frontend_keluaran/config"
)

type response struct {
	Status int         `json:"status"`
	Record interface{} `json:"record"`
}
type responsekeluaran struct {
	Status       int         `json:"status"`
	Record       interface{} `json:"record"`
	Paito_minggu interface{} `json:"paito_minggu"`
	Paito_senin  interface{} `json:"paito_senin"`
	Paito_selasa interface{} `json:"paito_selasa"`
	Paito_rabu   interface{} `json:"paito_rabu"`
	Paito_kamis  interface{} `json:"paito_kamis"`
	Paito_jumat  interface{} `json:"paito_jumat"`
	Paito_sabtu  interface{} `json:"paito_sabtu"`
}
type clientlistkeluaran struct {
	Pasaran string `json:"pasaran"`
}
type clientbukumimpi struct {
	Bukumimpi_tipe string `json:"bukumimpi_tipe"`
	Bukumimpi_nama string `json:"bukumimpi_nama"`
}
type clienttafsirmimpi struct {
	Tafsirmimpi_search string `json:"tafsirmimpi_search"`
}

const PATH string = config.PATH_API

func ListPasaran(c *fiber.Ctx) error {
	render_page := time.Now()

	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company": "",
		}).
		Post(PATH + "api/pasaran")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)
	return c.JSON(fiber.Map{
		"status": http.StatusOK,
		"record": result.Record,
		"time":   time.Since(render_page).String(),
	})
}
func ListKeluaran(c *fiber.Ctx) error {
	client := new(clientlistkeluaran)
	render_page := time.Now()
	log.Println(client.Pasaran)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsekeluaran{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"pasaran_id": client.Pasaran,
		}).
		Post(PATH + "api/keluaran")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsekeluaran)
	return c.JSON(fiber.Map{
		"status":       http.StatusOK,
		"record":       result.Record,
		"paito_minggu": result.Paito_minggu,
		"paito_senin":  result.Paito_senin,
		"paito_selasa": result.Paito_selasa,
		"paito_rabu":   result.Paito_rabu,
		"paito_kamis":  result.Paito_kamis,
		"paito_jumat":  result.Paito_jumat,
		"paito_sabtu":  result.Paito_sabtu,
		"time":         time.Since(render_page).String(),
	})
}
func ListNews(c *fiber.Ctx) error {
	render_page := time.Now()

	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company": "",
		}).
		Post(PATH + "api/news")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)
	return c.JSON(fiber.Map{
		"status": http.StatusOK,
		"record": result.Record,
		"time":   time.Since(render_page).String(),
	})
}
func ListNewsMovie(c *fiber.Ctx) error {
	render_page := time.Now()

	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company": "",
		}).
		Post(PATH + "api/newsmovie")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)
	return c.JSON(fiber.Map{
		"status": http.StatusOK,
		"record": result.Record,
		"time":   time.Since(render_page).String(),
	})
}
func ListBukumimpi(c *fiber.Ctx) error {
	client := new(clientbukumimpi)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"tipe": client.Bukumimpi_tipe,
			"nama": client.Bukumimpi_nama,
		}).
		Post(PATH + "api/bukumimpi")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)
	return c.JSON(fiber.Map{
		"status": http.StatusOK,
		"record": result.Record,
		"time":   time.Since(render_page).String(),
	})
}
func ListTafsirmimpi(c *fiber.Ctx) error {
	client := new(clienttafsirmimpi)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"search": client.Tafsirmimpi_search,
		}).
		Post(PATH + "api/tafsirmimpi")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)
	return c.JSON(fiber.Map{
		"status": http.StatusOK,
		"record": result.Record,
		"time":   time.Since(render_page).String(),
	})
}
