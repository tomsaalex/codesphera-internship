-- +goose Up
-- +goose StatementBegin

INSERT INTO categories
(id, category_name)
VALUES
-- Change the IDs to random later, they're for the mocks.
('0db70426-00d5-48da-a33e-8a1a2724a710', 'Furniture'),
('454dbba9-9db3-487a-b768-5c74ba6c945d', 'Fashion'),
('66dc2127-96a1-4844-909a-68c32dd72183', 'Jewelry & Watches'),
('750f63df-0d03-4340-b910-4f423a0768f4', 'Books'),
('80f8dc08-f1c6-407d-9b8e-10a5db75164f', 'Fine Art'),
('a7b5bb7d-bf6a-4900-87cb-38b698faafd7', 'Antiques'),
('d55357b4-d8bf-4dca-a88b-f76d26ff6a2b', 'Media'),
('e489d0e1-2db2-4239-afdb-f29c9fd7bfaa', 'Books & Manuscripts'),
('f000e982-8063-4b4c-9cc1-ba52a74d3f97', 'Real Estate'),
('f2e52b15-a9b8-4f12-ab65-856bee467075', 'Miscellaneous');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DELETE FROM categories;

-- +goose StatementEnd
