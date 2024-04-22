package repository

import (
	"database/sql"
	"fmt"

	"github.com/balazskvancz/image-gallery-hw/internal/models"
)

type Repository interface {
	Insert(string) error
	GetByName(string) *models.Image
	DeleteByName(string) error
	GetBy(string, string) models.Images
}

type repository struct {
	*sql.DB
}

// New létrehoz egy új egyedet.
func New(db *sql.DB) Repository {
	return &repository{
		DB: db,
	}
}

// Insert beszúr egy új egyedet.
func (r *repository) Insert(name string) error {
	_, err := r.Exec(`
		INSERT INTO images SET 
			name = ?,
			created_at = NOW()
	`, name)

	return err
}

// GetByName lekérdez egy egyedet azonosító szerint.
func (r *repository) GetByName(name string) *models.Image {
	row := r.QueryRow(`
		SELECT	
			name,
			created_at
		FROM images
		WHERE name = ?
	`, name)
	if row == nil {
		return nil
	}

	return scanImage(row)
}

// DeleteByName töröl egy egyedet név szerint.
func (r *repository) DeleteByName(name string) error {
	_, err := r.Exec(`
		DELETE FROM images
		WHERE name = ?
	`, name)

	return err
}

// GetBy lekérdezi a felvett egyedeket.
func (r *repository) GetBy(orderBy string, direction string) models.Images {
	dir := func() string {
		if direction == "asc" {
			return "ASC"
		}
		return "DESC"
	}()

	var (
		order  = fmt.Sprintf("ORDER BY %s %s", orderBy, dir)
		images = make(models.Images, 0)
	)

	rows, err := r.Query(`
		SELECT
			name,
			created_at
		FROM images ` + order,
	)
	if err != nil {
		fmt.Printf("[images.Get]: %v\n", err)

		return images
	}
	defer rows.Close()

	for rows.Next() {
		if im := scanImage(rows); im != nil {
			images = append(images, im)
		}
	}

	return images
}

type scanner interface {
	Scan(...any) error
}

func scanImage(sc scanner) *models.Image {
	var (
		name      string
		createdAt string
	)

	if err := sc.Scan(&name, &createdAt); err != nil {
		return nil
	}

	return &models.Image{
		Name:      name,
		CreatedAt: createdAt,
	}
}
