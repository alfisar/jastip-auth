Create Table logs (
  id INT AUTO_INCREMENT PRIMARY KEY,
  `event` varchar(255),
  `desc` varchar(255),
  user_id INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

Create Table status_order (
  id INT AUTO_INCREMENT PRIMARY KEY,
  `desc` varchar(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);

Create Table status_product (
  id INT AUTO_INCREMENT PRIMARY KEY,
  `desc` varchar(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);


Create Table payment_method (
  id INT AUTO_INCREMENT PRIMARY KEY,
  `desc` varchar(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);

Create Table address (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT,
  street text,
  city varchar(255),
  district varchar(255),
  subdistrict varchar(255),
  postcode varchar(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

Create Table traveler_schedule (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT,
  locations varchar(255),
  period_start datetime,
  period_end datetime,
  status INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

Create Table traveler_history (
  id INT AUTO_INCREMENT PRIMARY KEY,
  travel_schedule_id INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (travel_schedule_id) REFERENCES traveler_schedule(id) ON DELETE CASCADE ON UPDATE CASCADE
);

Create Table products_history (
  id INT AUTO_INCREMENT PRIMARY KEY,
  product_id INT,
  price DECIMAL(10, 2),
  traveler_history_id INT,
  images varchar(255),
  status INT,
  quantity INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (traveler_history_id) REFERENCES traveler_history(id) ON DELETE CASCADE ON UPDATE CASCADE

);


Create Table products (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name varchar(255),
  `desc` text,
  price DECIMAL(10, 2),
  traveler_schedule_id INT,
  images varchar(255),
  status INT,
  quantity INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (traveler_schedule_id) REFERENCES traveler_schedule(id) ON DELETE CASCADE ON UPDATE CASCADE
);

Create Table orders (
  id INT AUTO_INCREMENT PRIMARY KEY,
  buyer_id INT,
  traveler_schedule_id INT,
  traveler_history_id INT,
  price DECIMAL(10, 2),
  status INT,
  payment_status INT,
  payment_method INT,
  address_id INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (traveler_schedule_id) REFERENCES traveler_schedule(id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (traveler_history_id) REFERENCES traveler_history(id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (buyer_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

Create Table orders_detail (
  id INT AUTO_INCREMENT PRIMARY KEY,
  products_id INT,
  order_id INT ,
  quantity INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE ON UPDATE CASCADE
);