INSERT INTO products
    (name, category_id, brand_id)
VALUES
    ('Pepsi', (SELECT id FROM categories WHERE name = 'food'), (SELECT id FROM brands WHERE name = 'Pepsi')),
    ('Mars', (SELECT id FROM categories WHERE name = 'food'), (SELECT id FROM brands WHERE name = 'Mars')),
    ('Danone', (SELECT id FROM categories WHERE name = 'food'), (SELECT id FROM brands WHERE name = 'Danone')),
    ('Starbucks', (SELECT id FROM categories WHERE name = 'coffee'), (SELECT id FROM brands WHERE name = 'Starbucks')),
    ('Philadelphia', (SELECT id FROM categories WHERE name = 'food'), (SELECT id FROM brands WHERE name = 'Philadelphia')),
    ('Haagen-Dazs', (SELECT id FROM categories WHERE name = 'food'), (SELECT id FROM brands WHERE name = 'Haagen-Dazs')),
    ('See’s Candies', (SELECT id FROM categories WHERE name = 'food'), (SELECT id FROM brands WHERE name = 'See’s Candies'))
;