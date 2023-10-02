INSERT INTO brands
    (name, image_url, location_id)
VALUES
    ('Pepsi', 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/68/Pepsi_2023.svg/280px-Pepsi_2023.svg.png', (SELECT id FROM locations WHERE city = 'Seattle')),
    ('Mars', 'https://www.mars.com/sites/g/files/jydpyr316/files/2019-07/Mars%20Brand%20Logos%20Web%20Confectionery%20Mars%20Large.png', (SELECT id FROM locations WHERE city = 'Dallas')),
    ('Danone', 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Danone_spain.png/1599px-Danone_spain.png', (SELECT id FROM locations WHERE city = 'Miami')),
    ('Starbucks', 'https://upload.wikimedia.org/wikipedia/ru/thumb/d/d3/Starbucks_Corporation_Logo_2011.svg/1200px-Starbucks_Corporation_Logo_2011.svg.png', (SELECT id FROM locations WHERE city = 'Seattle')),
    ('Philadelphia', 'https://upload.wikimedia.org/wikipedia/commons/6/6f/Philadelphia_cheese_logo.png', (SELECT id FROM locations WHERE city = 'New York')),
    ('Haagen-Dazs', 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/3c/H%C3%A4agen-Dazs_Logo.svg/1200px-H%C3%A4agen-Dazs_Logo.svg.png', (SELECT id FROM locations WHERE city = 'New York')),
    ('See’s Candies', 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/41/See%27s_Candies_logo.svg/1200px-See%27s_Candies_logo.svg.png', (SELECT id FROM locations WHERE city = 'Los Angeles'))
;