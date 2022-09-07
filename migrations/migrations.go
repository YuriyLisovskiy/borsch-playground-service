/*
 * Borsch Playground API
 *
 * Copyright (C) 2022 Yuriy Lisovskiy - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the MIT license.
 */

package migrations

import (
	"github.com/YuriyLisovskiy/borsch-playground-service/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.JobOutputRowDbModel{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.JobDbModel{}); err != nil {
		return err
	}

	return nil
}