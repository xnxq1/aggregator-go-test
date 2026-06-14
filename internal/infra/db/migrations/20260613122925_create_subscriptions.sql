-- +goose Up
CREATE TABLE subscriptions (
      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
      service_name TEXT NOT NULL,
      price INT NOT NULL,
      user_id UUID NOT NULL,
      start_date TEXT NOT NULL,
      end_date TEXT
  );

-- +goose Down
DROP TABLE subscriptions;
