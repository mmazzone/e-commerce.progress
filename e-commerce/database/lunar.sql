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
    product_name VARCHAR (255),
    product_desc VARCHAR(255),
    product_price DECIMAL (5,2),
    info_link VARCHAR (255),
    img_src VARCHAR(255),
    categories VARCHAR(255)
);


INSERT INTO products (product_name, product_desc, product_price, info_link, img_src, categories) 
VALUES 
    ('32oz. Bodum Chambord French Press','Brew coffee to perfection in just 4 minutes',19.99,'productinfo.html','images/brewers/french.press.jpg','brewers'),
    ('20oz. Bialetti Moka Pot','Enjoy a rich, authentic espresso in just minutes',22.99,'productinfo.html','images/featured.products/bialetti.moka.jpg','brewers'), 
    ('24oz. Classic Chemex','Enjoy the perfect cup of pour over coffee',34.99,'productinfo.html','images/featured.products/chemex.jpg','brewers'),
    ('Midnight Ceramic Mug','12oz. handmade glazed ceramic mug',10.99,'productinfo.html','images/ceramics/mug.2.jpg','ceramics'),
    ('Black Ceramic Mug Set','Set of 4, 12oz. handmade glazed ceramic mugs',34.99,'productinfo.html','images/ceramics/black.mugs.jpg','ceramics'),
    ('Handheld Coffee Grinder','Handheld coffee grinder holds 8oz. with two storage containers',12.99,'productinfo.html','images/coffee.grinder.jpg','tools'),
    ('Golden Scoop-Clip','1 Tablespoon scoop doubles as a clip to seal coffee bags',9.99,'productinfo.html','images/featured.products/coffee.scoop.jpg','accessories'),
    ('Re-usable Chemex Filters','Set of 5 reusable cotton Chemex coffee filters',12.99,'productinfo.html','images/chemex.filter.jpg','accessories');

