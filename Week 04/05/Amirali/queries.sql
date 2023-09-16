CREATE TABLE products(
    p_id bigint AUTO_INCREMENT PRIMARY KEY,
    p_name varchar(255) NOT NULL,
    p_weight bigint NOT NULL
);
CREATE TABLE capacities(
    level int PRIMARY KEY,
    capacity bigint NOT NULL
);
CREATE TABLE markets(
    m_id bigint AUTO_INCREMENT PRIMARY KEY,
    m_name varchar(255) NOT NULL,
    m_address text NOT NULL,
    m_score bigint NOT NULL,
    level int NOT NULL,
    FOREIGN KEY (level) REFERENCES capacities(level)
);
CREATE TABLE prices(
    m_id bigint NOT NULL,
    p_id bigint NOT NULL,
    price bigint NOT NULL,
    PRIMARY KEY (m_id,p_id),
    FOREIGN KEY (m_id) REFERENCES markets(m_id),
    FOREIGN KEY (p_id) REFERENCES products(p_id)
);