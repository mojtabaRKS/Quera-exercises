CREATE TABLE capacities(
    level INT PRIMARY KEY,
    capacity BIGINT NOT NULL
);
CREATE TABLE markets(
    m_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    m_name VARCHAR(255) NOT NULL,
    m_address TEXT NOT NULL,
    m_score BIGINT NOT NULL,
    level INT NOT NULL,
    FOREIGN KEY (level) REFERENCES capacities(level) 
);
CREATE TABLE products (
    p_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    p_name VARCHAR(255) NOT NULL,
    p_weight BIGINT NOT NULL
);
CREATE TABLE prices(
    p_id BIGINT NOT NULL,
    m_id BIGINT NOT NULL,
    price BIGINT NOT NULL,
    FOREIGN KEY (p_id) REFERENCES products(p_id),
    FOREIGN KEY (m_id) REFERENCES markets(m_id),
    PRIMARY KEY (m_id,p_id)
);
