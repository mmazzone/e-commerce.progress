-- CREATE DATABASE lunar;

USE lunar;

CREATE TABLE contact_info (
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email_address VARCHAR(255),
    phone_number INT

);

CREATE TABLE products (
    product_id SERIAL PRIMARY KEY NOT NULL,
    discontinued TINYINT (1),
    product_name VARCHAR (255),
    product_desc VARCHAR(255),
    product_price DECIMAL (5,2),
    info_link VARCHAR (255),
    img_src VARCHAR(255),
    categories VARCHAR(255)

);

CREATE TABLE product_info (
    product_name VARCHAR (255),
    product_desc VARCHAR (1000),
    product_price DECIMAL (5,2),
    img_src VARCHAR (255)
    
);
-- CREATE TABLE errors (

-- )

INSERT INTO products (product_name, product_desc, product_price, info_link, img_src, categories) 
VALUES 
    ("32oz. Bodum Chambord French Press","Brew coffee to perfection in just 4 minutes",19.99,"productinfo.html","../images/brewers/french.press.jpg","BREWERS"),
    ("20oz. Bialetti Moka Pot","Enjoy a rich, authentic espresso in just minutes",22.99,"productinfo.html","../images/featured.products/bialetti.moka.jpg","BREWERS"), 
    ("24oz. Classic Chemex","Enjoy the perfect cup of pour over coffee",34.99,"productinfo.html","../images/featured.products/chemex.jpg","BREWERS"),
    ("Midnight Ceramic Mug","12oz. handmade glazed ceramic mug",10.99,"productinfo.html","../images/ceramics/mug.2.jpg","CERAMICS"),
    ("Black Ceramic Mug Set","Set of 4, 12oz. handmade glazed ceramic mugs",34.99,"productinfo.html","../images/ceramics/black.mugs.jpg","CERAMICS"),
    ("Handheld Coffee Grinder","Handheld coffee grinder holds 8oz. with two storage containers",12.99,"productinfo.html","../images/coffee.grinder.jpg","TOOLS"),
    ("Golden Scoop-Clip","1 Tablespoon scoop doubles as a clip to seal coffee bags",9.99,"productinfo.html","../images/featured.products/coffee.scoop.jpg","ACCESSORIES"),
    ("Reusable Chemex Filters","Set of 5 reusable cotton Chemex coffee filters",12.99,"productinfo.html","../images/chemex.filter.jpg","ACCESSORIES");

INSERT INTO product_info 
VALUES
    ("Bodum Chambord French Press","Chambord French press brews a premium cup of coffee in just 4 minutes. Simply add course ground coffee, hot
            water and press
            stainless steel: 3-part stainless steel plunger has a mesh filter that helps extract your coffee's aromatic
            oils and subtle
            flavors instead of being absorbed by a paper filter. Made of plastic durable design: coffee press features
            Bodum patented
            safety lid to keep contents from spilling and is dishwasher safe for easy cleaning. Maximum flavor: pressed
            coffee extracts
            the perfect amount of essentials oils and acids from the coffee bean for maximum flavor; the preferred
            method for brewing
            for coffee enthusiasts everywhere! Servings: premium French Press Coffee maker makes 4 cups of Coffee, 8 oz
            each",19.99,"../images/brewers/french.press.jpg");