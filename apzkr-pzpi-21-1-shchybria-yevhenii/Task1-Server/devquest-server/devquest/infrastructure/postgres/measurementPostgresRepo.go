package postgres

import (
	"context"
	"database/sql"
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/repositories"
	"devquest-server/devquest/infrastructure"

	"github.com/google/uuid"
)

type MeasurementPostgresRepo struct {
	db infrastructure.Database
}

func NewMeasurementPostgresRepo(db infrastructure.Database) repositories.MeasurementRepo {
	return &MeasurementPostgresRepo{db: db}
}

func (m *MeasurementPostgresRepo) GetDeviceByID(deviceID uuid.UUID) (*entities.MeasurementDevice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id, type_id, owner_id
		FROM measurement_devices
		WHERE id = $1
	`

	row := m.db.GetDB().QueryRowContext(ctx, query, deviceID)

	var device entities.MeasurementDevice
	err := row.Scan(&device.ID, &device.TypeID, &device.OwnerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
			return nil, err
	}

	return &device, nil
}

func (m *MeasurementPostgresRepo) GetTypeByID(typeID uuid.UUID) (*entities.MeasurementType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id, name, minimum_norm, maximum_norm
		FROM measurement_types
		WHERE id = $1
	`

	row := m.db.GetDB().QueryRowContext(ctx, query, typeID)

	var deviceType entities.MeasurementType
	err := row.Scan(&deviceType.ID, &deviceType.Name, &deviceType.MinimumNorm, &deviceType.MaximumNorm)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
			return nil, err
	}

	return &deviceType, nil
}

func (m *MeasurementPostgresRepo) GetTypeByName(typeName string) (*entities.MeasurementType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id, name, minimum_norm, maximum_norm
		FROM measurement_types
		WHERE name = $1
	`

	row := m.db.GetDB().QueryRowContext(ctx, query, typeName)

	var deviceType entities.MeasurementType
	err := row.Scan(&deviceType.ID, &deviceType.Name, &deviceType.MinimumNorm, &deviceType.MaximumNorm)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
			return nil, err
	}

	return &deviceType, nil
}

func (m *MeasurementPostgresRepo) CheckForDevice(deviceID uuid.UUID, typeID uuid.UUID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id
		FROM measurement_devices
		WHERE id = $1 AND type_id = $2
	`

	row := m.db.GetDB().QueryRowContext(ctx, query, deviceID, typeID)

	var foundDeviceID uuid.UUID
	err := row.Scan(&foundDeviceID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
			return false, err
	}

	return true, nil
}

func (m *MeasurementPostgresRepo) AddOwnerToDevice(deviceID uuid.UUID, ownerID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.db.GetDBTimeout())
	defer cancel()

	execute := `
		UPDATE measurement_devices
		SET owner_id = $1
		WHERE id = $2
	`

	_, err := m.db.GetDB().ExecContext(ctx, execute, ownerID, deviceID)
	if err != nil {
		return err
	}

	return nil
}

func (m *MeasurementPostgresRepo) AddMeasurementResult(newMeasurement entities.Measurement) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.db.GetDBTimeout())
	defer cancel()

	execute := `
		INSERT INTO measurements
		(id, device_id, measured_at, value)
		VALUES ($1, $2, $3, $4)
	`

	_, err := m.db.GetDB().ExecContext(ctx, execute, newMeasurement.ID, newMeasurement.DeviceID, newMeasurement.MeasuredAt, newMeasurement.Value)
	if err != nil {
			return err
	}

	return nil
}

func (m *MeasurementPostgresRepo) GetLatestMeasurementsForDeveloper(developerID uuid.UUID) ([]*entities.Measurement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT m.id, m.device_id, m.measured_at, m.value
		FROM measurements m
		LEFT JOIN measurement_devices md ON m.device_id = md.id
		WHERE m.measured_at = (
			SELECT MAX(m2.measured_at)
			FROM measurements m2
			LEFT JOIN measurement_devices md2 ON m2.device_id = md2.id
			WHERE md2.type_id = md.type_id
			AND md2.owner_id = $1
		)
	`

	rows, err := m.db.GetDB().QueryContext(ctx, query, developerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var measurements []*entities.Measurement
	for rows.Next() {
		var measurement entities.Measurement

		err := rows.Scan(&measurement.ID, &measurement.DeviceID, &measurement.MeasuredAt, &measurement.Value)
		if err != nil {
			return nil, err
		}

		measurements = append(measurements, &measurement)
	}

	return measurements, nil
}