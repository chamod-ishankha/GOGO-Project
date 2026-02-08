-- FUNCTION (ONE TIME)
CREATE OR REPLACE FUNCTION gogo.update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ language 'plpgsql';

-- HOW TO CALL
CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON gogo.users
FOR EACH ROW
EXECUTE FUNCTION gogo.update_updated_at_column();

CREATE TRIGGER update_drivers_updated_at
BEFORE UPDATE ON gogo.drivers
FOR EACH ROW
EXECUTE FUNCTION gogo.update_updated_at_column();

CREATE TRIGGER update_vehicles_updated_at
BEFORE UPDATE ON gogo.vehicles
FOR EACH ROW
EXECUTE FUNCTION gogo.update_updated_at_column();

CREATE TRIGGER update_rides_updated_at
BEFORE UPDATE ON gogo.rides
FOR EACH ROW
EXECUTE FUNCTION gogo.update_updated_at_column();