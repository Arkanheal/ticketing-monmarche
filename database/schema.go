package database

func CreateTables() {
    // Use of numeric because it's monetary
    DB.Query(`CREATE TABLE IF NOT EXISTS orders (
        id int PRIMARY KEY NOT NULL,
        vat numeric NOT NULL,
        total numeric NOT NULL
    )`)
    DB.Query(`CREATE TABLE IF NOT EXISTS products (
        id VARCHAR(4) PRIMARY KEY NOT NULL,
        price numeric NOT NULL,
        name VARCHAR(255) NOT NULL
    )`)
    DB.Query(`CREATE TABLE IF NOT EXISTS product_order (
        product_id VARCHAR(4) NOT NULL,
        order_id int NOT NULL,
        CONSTRAINT fk_product
            FOREIGN KEY(product_id)
                REFERENCES products(id)
        ON DELETE CASCADE,
        CONSTRAINT fk_order
            FOREIGN KEY(order_id)
                REFERENCES orders(id)
        ON DELETE CASCADE)
    `)
}
