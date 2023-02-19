package handler

import (
	"encoding/csv"
	"log"
	"regexp"
	"strconv"
	"strings"

	"tickets/database"
	"tickets/model"

	"github.com/gofiber/fiber/v2"
)

func parseOrderData(line string, regex string) string {
    re := regexp.MustCompile(regex)
    data := string(re.Find([]byte(line)))
    return data
}

func returnErr(c *fiber.Ctx, err error, s int) {
    log.Println(err)
    c.Status(s).JSON(&fiber.Map{
        "success": false,
        "message": err,
    })
}

func CreateOrder(c *fiber.Ctx) error {
    var err error

    order := new(model.Order)

    // Get payload datas
    data := string(c.Body()[:])
    lines := strings.Split(strings.TrimSuffix(data, "\n"), "\n")

    // Get data into struct (kinda sad to do like this tbh)
    order.Id, err = strconv.Atoi(parseOrderData(lines[0], `\d+`))
    if err != nil {
        returnErr(c, err, 400)
    }
    order.VAT, err = strconv. ParseFloat(parseOrderData(lines[1], `\d+\.\d*`), 64)
    if err != nil {
        returnErr(c, err, 400)
    }
    order.Price, err = strconv.ParseFloat(parseOrderData(lines[2], `\d+\.\d*`), 64)
    if err != nil {
        returnErr(c, err, 400)
    }

    // Parse product data in CSV
    productsCSV := strings.Join(lines[4:], "\n")
    reader := csv.NewReader(strings.NewReader(productsCSV))

    field, err := reader.ReadAll()
    if err != nil {
        returnErr(c, err, 400)
    }
    log.Println(field)

    // Insert order now (in case product wasn't valid)
    _, err = database.DB.Query("INSERT INTO orders (id, VAT, total) VALUES ($1, $2, $3)", order.Id, order.VAT, order.Price)
    if err != nil {
        returnErr(c, err, 500)
    }

    // Insert product(s)
    for _, product := range field[1:] {
        _, err := database.DB.Query("INSERT INTO products (id, price, name) VALUES ($1, $2, $3)", product[1], product[2], product[0])
        if err != nil {
            returnErr(c, err, 500)
        }
        _, err = database.DB.Query("INSERT INTO product_order (product_id, order_id) VALUES ($1, $2)", product[1], order.Id)
        if err != nil {
            returnErr(c, err, 500)
        }
    }

    // Return order infos (STATUS 200) or 500 with err
    if err := c.JSON(&fiber.Map{
        "success": true,
        "message": "Order and products successfully created",
        /* "order_id": order.id */
    }); err != nil {
        c.Status(500).JSON(&fiber.Map{
            "success": false,
            "message": "Error creating order",
        })
    }

    return c.SendString(data)

}
