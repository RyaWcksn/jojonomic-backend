CREATE TABLE IF NOT EXISTS tbl_harga (
    reff_id VARCHAR (15) PRIMARY KEY NOT NULL,
    admin_id VARCHAR (15) NOT NULL,
    harga_topup DECIMAL(12, 2) NOT NULL,
    harga_buyback DECIMAL(12, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
