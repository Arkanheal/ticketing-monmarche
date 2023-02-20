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

func CreateOrder(c *fiber.Ctx) error {
	var err error

	tx, err := database.DB.BeginTx(c.Context(), nil)
	defer tx.Rollback()

	order := new(model.Order)

	// Get payload datas
	data := string(c.Body()[:])
	lines := strings.Split(strings.TrimSuffix(data, "\n"), "\n")

	// Get data into struct (kinda sad to do like this tbh)
	order.Id, err = strconv.Atoi(parseOrderData(lines[0], `\d+`))
	if err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	order.VAT, err = strconv.ParseFloat(parseOrderData(lines[1], `\d+\.\d*`), 64)
	if err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	order.Price, err = strconv.ParseFloat(parseOrderData(lines[2], `\d+\.\d*`), 64)
	if err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	// Parse product data in CSV
	productsCSV := strings.Join(lines[4:], "\n")
	reader := csv.NewReader(strings.NewReader(productsCSV))

	field, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	// Insert order now (in case product wasn't valid)
	_, err = database.DB.Query("INSERT INTO orders (id, VAT, total) VALUES ($1, $2, $3)", order.Id, order.VAT, order.Price)
	if err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	// Insert product(s)
	for _, product := range field[1:] {
		// Imaginons qu'on change le prix du même produits
		_, err := database.DB.Query(`INSERT INTO products (id, price, name)
            VALUES ($1, $2, $3)
            ON CONFLICT (id) DO UPDATE SET price = EXCLUDED.price, name = EXCLUDED.name`, product[1], product[2], product[0])
		if err != nil {
			log.Println(err)
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		}

		// Ajout pour le lien avec les commandes
		_, err = database.DB.Query("INSERT INTO product_order (product_id, order_id) VALUES ($1, $2)", product[1], order.Id)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		}
	}

	// Commit les changements en DB si tout se passe bien
	if err = tx.Commit(); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	return c.JSON(&fiber.Map{
		"success":  true,
		"message":  "Order and products successfully created",
		"order_id": order.Id,
	})

}
