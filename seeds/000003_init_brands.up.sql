INSERT INTO brands
    (name, image_url, location_id)
VALUES
    ('Starbucks', 'https://upload.wikimedia.org/wikipedia/ru/thumb/d/d3/Starbucks_Corporation_Logo_2011.svg/1200px-Starbucks_Corporation_Logo_2011.svg.png', (SELECT id FROM locations WHERE city = 'Seattle')),
    ('Lavazza', 'https://e7.pngegg.com/pngimages/1004/368/png-clipart-lavazza-espresso-point-coffee-lavazza-espresso-point-coffee-capsule-lavazza-coffee-blue-text.png', (SELECT id FROM locations WHERE city = 'New York')),
    ('Costa Coffee', 'https://upload.wikimedia.org/wikipedia/en/thumb/f/fa/Costa_Coffee_logo.svg/1200px-Costa_Coffee_logo.svg.png', (SELECT id FROM locations WHERE city = 'Seattle')),
    ('Nescafé', 'https://www.comunicaffe.com/wp-content/uploads/2019/03/Logo_NESCAFE.jpg', (SELECT id FROM locations WHERE city = 'New York')),
    ('Folgers', 'https://mir-s3-cdn-cf.behance.net/project_modules/1400/a0a0f870749703.5bbaaa55e574b.png', (SELECT id FROM locations WHERE city = 'Dallas')),

    ('Pepsi', 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/68/Pepsi_2023.svg/280px-Pepsi_2023.svg.png', (SELECT id FROM locations WHERE city = 'Dallas')),
    ('Twix', 'https://1000logos.net/wp-content/uploads/2020/07/Twix-Logo.png', (SELECT id FROM locations WHERE city = 'Dallas')),
    ('Danone', 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Danone_spain.png/1599px-Danone_spain.png', (SELECT id FROM locations WHERE city = 'Miami')),
    ('Philadelphia', 'https://upload.wikimedia.org/wikipedia/commons/6/6f/Philadelphia_cheese_logo.png', (SELECT id FROM locations WHERE city = 'New York')),
    ('Haagen-Dazs', 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3c/H%C3%A4agen-Dazs_Logo.svg/1200px-H%C3%A4agen-Dazs_Logo.svg.png', (SELECT id FROM locations WHERE city = 'New York')),
    ('See’s Candies', 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/41/See%27s_Candies_logo.svg/1200px-See%27s_Candies_logo.svg.png', (SELECT id FROM locations WHERE city = 'Dallas')),

    ('Lego', 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/24/LEGO_logo.svg/1200px-LEGO_logo.svg.png', (SELECT id FROM locations WHERE city = 'Dallas')),
    ('Barbie', 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/0b/Barbie_Logo.svg/1200px-Barbie_Logo.svg.png', (SELECT id FROM locations WHERE city = 'Los Angeles')),
    ('Hot Wheels', 'https://blog.logomyway.com/wp-content/uploads/2021/09/hot-wheels-logo.png', (SELECT id FROM locations WHERE city = 'Los Angeles')),
    ('Fisher-Price', 'https://brand-info.com.ua/wp-content/uploads/2023/03/fisher-price-logo.png', (SELECT id FROM locations WHERE city = 'Los Angeles')),
    ('Nerf', 'https://static.brandirectory.com/logos/nerf001_1920px_nerf_logosvg.png', (SELECT id FROM locations WHERE city = 'Los Angeles')),

    ('Miele', 'https://1000logos.net/wp-content/uploads/2017/12/Miele-logo.png', (SELECT id FROM locations WHERE city = 'Los Angeles')),
    ('Bosch', 'https://i0.wp.com/bctechnologies.co.in/wp-content/uploads/2020/07/bosch-logo-big.png?resize=870%2C324&ssl=1', (SELECT id FROM locations WHERE city = 'Seattle')),
    ('Samsung', 'https://file.liga.net/images/general/2020/09/08/20200908141029-3295.jpg?v=1599567195', (SELECT id FROM locations WHERE city = 'Seattle')),
    ('Electrolux', 'https://5ypbvxa39ihl3fage541b0i.blob.core.windows.net/media/ElectroluxMedia/Home/Logo-Electrolux.svg', (SELECT id FROM locations WHERE city = 'Los Angeles')),
    ('Whirlpool', 'https://upload.wikimedia.org/wikipedia/commons/9/95/Whirlpool_Corporation_Logo_%28as_of_2017%29.svg', (SELECT id FROM locations WHERE city = 'Los Angeles'))
;