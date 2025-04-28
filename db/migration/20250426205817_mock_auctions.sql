-- +goose Up
INSERT INTO auctions (
    id,
    product_name,
    product_desc,
    auc_mode,
    auc_status,
    starting_price,
    target_price,
    seller_id
) VALUES
-- Auction 1
(gen_random_uuid(), 'Vintage Camera', 'A classic 1960s Leica camera in great condition.', 'manual', 'ongoing', 150.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 2
(gen_random_uuid(), 'Electric Guitar', 'Fender Stratocaster, sunburst finish, barely used.', 'manual', 'ongoing', 400.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 3
(gen_random_uuid(), 'Antique Desk', 'Solid oak desk from the 19th century, restored.', 'manual', 'ongoing', 250.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 4
(gen_random_uuid(), 'Gaming Laptop', 'RTX 4070, 32GB RAM, barely 3 months old.', 'manual', 'ongoing', 1200.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 5
(gen_random_uuid(), 'Mountain Bike', 'Trek X-Caliber 8, lightly ridden, tuned up.', 'manual', 'ongoing', 500.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 6
(gen_random_uuid(), 'Noise Cancelling Headphones', 'Sony WH-1000XM5, like new.', 'manual', 'ongoing', 220.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 7
(gen_random_uuid(), '4K Monitor', 'Dell UltraSharp 32-inch, stunning visuals.', 'manual', 'ongoing', 350.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 8
(gen_random_uuid(), 'Mechanical Keyboard', 'Custom-built with Gateron switches.', 'manual', 'ongoing', 120.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 9
(gen_random_uuid(), 'Espresso Machine', 'Breville Barista Express, excellent condition.', 'manual', 'ongoing', 400.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 10
(gen_random_uuid(), 'Leather Jacket', 'Schott NYC, size M, barely worn.', 'manual', 'ongoing', 180.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 11
(gen_random_uuid(), 'Smartphone', 'Google Pixel 7, unlocked, pristine.', 'manual', 'ongoing', 500.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 12
(gen_random_uuid(), 'Camping Tent', 'REI Half Dome 2 Plus, perfect for two.', 'manual', 'ongoing', 160.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 13
(gen_random_uuid(), 'Vinyl Record Player', 'Audio-Technica AT-LP60XBT, Bluetooth.', 'manual', 'ongoing', 130.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 14
(gen_random_uuid(), 'Fitness Watch', 'Garmin Forerunner 255, black.', 'manual', 'ongoing', 200.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 15
(gen_random_uuid(), 'Tablet', 'iPad Air 5th Gen, sky blue, 64GB.', 'manual', 'ongoing', 550.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 16
(gen_random_uuid(), 'Road Bike', 'Specialized Allez Sprint, fast and lightweight.', 'manual', 'ongoing', 1000.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 17
(gen_random_uuid(), 'Smart Home Speaker', 'Sonos One, great condition.', 'manual', 'ongoing', 180.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 18
(gen_random_uuid(), 'Photography Lighting Kit', 'Neewer softbox set, hardly used.', 'manual', 'ongoing', 110.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 19
(gen_random_uuid(), 'DJ Turntable', 'Pioneer PLX-500, like new.', 'manual', 'ongoing', 300.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66'),

-- Auction 20
(gen_random_uuid(), 'Ski Equipment Set', 'Atomic Vantage 79Ti skis + bindings.', 'manual', 'ongoing', 700.00, NULL, '23980949-4093-4f31-8458-a7a893a7ef66');


-- +goose Down
DELETE FROM auctions
WHERE seller_id = '23980949-4093-4f31-8458-a7a893a7ef66';