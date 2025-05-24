DROP TYPE IF EXISTS user_role;
CREATE TYPE user_role AS ENUM ('common', 'shopkeeper'); 

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(255) NOT NULL,
    role user_role NOT NULL,
    cpf CHAR(11) UNIQUE,
    cnpj CHAR(14) UNIQUE,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL, 
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

ALTER TABLE IF EXISTS users
DROP CONSTRAINT IF EXISTS cpf_or_cnpj_required;

ALTER TABLE IF EXISTS users
ADD CONSTRAINT cpf_or_cnpj_required
CHECK (cpf IS NOT NULL OR cnpj IS NOT NULL);