INSERT INTO products
    (name, category_id, brand_id, rating)
VALUES
    ('Starbucks® Blonde Espresso Roast', (SELECT id FROM categories WHERE name = 'coffee'), (SELECT id FROM brands WHERE name = 'Starbucks'), 4.3),
    ('Lavazza Classico Ground Coffee Blend, Medium Roast, 12-Ounce Bags (Pack of 6) , Value Pack, Rich Flavor with Notes of Dried Fruit', (SELECT id FROM categories WHERE name = 'coffee'), (SELECT id FROM brands WHERE name = 'Lavazza'), 3.2),
    ('Green Mountain Coffee K-Cup Pods Costa Rica Paraiso, 24/box', (SELECT id FROM categories WHERE name = 'coffee'), (SELECT id FROM brands WHERE name = 'Costa Coffee'), 4.5),
    ('NESCAFE CLASICO Dark Roast Instant Coffee 10.5 oz. Jar', (SELECT id FROM categories WHERE name = 'coffee'), (SELECT id FROM brands WHERE name = 'Nescafé'), 4.7),
    ('Folgers 100% Colombian Medium Roast Coffee, 72 Keurig K-Cup Pods', (SELECT id FROM categories WHERE name = 'coffee'), (SELECT id FROM brands WHERE name = 'Folgers'), 4.7),

    ('Pepsi Flavors Variety Pack, Wild Cherry, Mango, Original, 12 Ounce Cans (18 Pack)', (SELECT id FROM categories WHERE name = 'food'), (SELECT id FROM brands WHERE name = 'Pepsi'), 4.3),
    ('TWIX Fun Size Caramel Cookie Halloween Chocolate Bars - 10.83 oz Candy Bag (Pack of 3)', (SELECT id FROM categories WHERE name = 'food'), (SELECT id FROM brands WHERE name = 'Twix'), 4.5),
    ('YOGOURMET Dried Yogurt Starter 0.17 Ounce (Pack of 1)', (SELECT id FROM categories WHERE name = 'food'), (SELECT id FROM brands WHERE name = 'Danone'), 4.6),
    ('Original Cream Cheese (plain cream cheese)', (SELECT id FROM categories WHERE name = 'food'), (SELECT id FROM brands WHERE name = 'Philadelphia'), 4.8),
    ('Butter Cookie Cone Coffee', (SELECT id FROM categories WHERE name = 'food'), (SELECT id FROM brands WHERE name = 'Haagen-Dazs'), 4.2),
    ('Candy Retailer JuJu Cinnamon Hearts 1 Lb', (SELECT id FROM categories WHERE name = 'food'), (SELECT id FROM brands WHERE name = 'See’s Candies'), 4.2),

    ('LEGO Harry Potter Expecto Patronum 76414 Collectible 2-in-1 Building Set; Birthday Gift Idea for Teens or Fans Aged 14 and Up; Build and Display Patronus Set for Fans of The Wizarding World', (SELECT id FROM categories WHERE name = 'toys'), (SELECT id FROM brands WHERE name = 'Lego'), 4.8),
    ('Barbie Extra Fly Doll with Snow-Themed Travel Clothes & Accessories, Sparkly Pink Jumpsuit & Faux Fur Coat For 3Y+', (SELECT id FROM categories WHERE name = 'toys'), (SELECT id FROM brands WHERE name = 'Barbie'), 4.8),
    ('Hot Wheels Set of 50 Toy Trucks & Cars in 1:64 Scale, Individually Packaged Vehicles (Styles May Vary) (Amazon Exclusive)', (SELECT id FROM categories WHERE name = 'toys'), (SELECT id FROM brands WHERE name = 'Hot Wheels'), 4.8),
    ('Fisher-Price Little People Toddler Toys Disney Princess Moana & Maui’s Canoe Sail Boat with 2 Figures for Ages 18+ Months', (SELECT id FROM categories WHERE name = 'toys'), (SELECT id FROM brands WHERE name = 'Fisher-Price'), 4.9),
    ('NERF DinoSquad Rex-Rampage Motorized Dart Blaster, 10-Dart Clip, 20 Official Darts, 10-Dart Storage- T-Rex Dinosaur Design', (SELECT id FROM categories WHERE name = 'toys'), (SELECT id FROM brands WHERE name = 'Nerf'), 4.5),

    ('Miele Complete C3 Canister Vacuum, Marine Blue', (SELECT id FROM categories WHERE name = 'appliances'), (SELECT id FROM brands WHERE name = 'Miele'), 4.4),
    ('Bosch 800 Series 36 Inch Wide 21 Cu. Ft. Energy Star Rated French Door Refrigerator with Home Connect', (SELECT id FROM categories WHERE name = 'appliances'), (SELECT id FROM brands WHERE name = 'Bosch'), 4.4),
    ('Samsung 1.7 Cu. Ft. Fingerprint Resistant Black Stainless Steel Over-The-Range Microwave', (SELECT id FROM categories WHERE name = 'appliances'), (SELECT id FROM brands WHERE name = 'Samsung'), 4.7),
    ('Sanitaire SC689A Commercial Dust Cup Upright Vacuum Cleaner with Dirt Cup and 5 Amp Motor, 12" Cleaning Path', (SELECT id FROM categories WHERE name = 'appliances'), (SELECT id FROM brands WHERE name = 'Electrolux'), 3.0),
    ('36-inch Wide Side-by-Side Refrigerator - 28 cu. ft.', (SELECT id FROM categories WHERE name = 'appliances'), (SELECT id FROM brands WHERE name = 'Whirlpool'), 4.6)

;