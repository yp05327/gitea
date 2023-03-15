// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_20 //nolint

import (
	"code.gitea.io/gitea/modules/container"
	"code.gitea.io/gitea/modules/log"

	"xorm.io/builder"
	"xorm.io/xorm"
)

// FixIncorrectProjectType: set individual project's type from 3(TypeOrganization) to 1(TypeIndividual)
func FixIncorrectProjectType(x *xorm.Engine) error {
	type User struct {
		ID   int64 `xorm:"pk autoincr"`
		Type int
	}

	const (
		UserTypeIndividual int = 0

		TypeIndividual   uint8 = 1
		TypeOrganization uint8 = 3
	)

	type Project struct {
		ID      int64 `xorm:"pk autoincr"`
		OwnerID int64 `xorm:"INDEX"`
		Type    uint8
		Owner   *User `xorm:"extends"`
	}

	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return err
	}

	var ps []Project
	err := sess.Table("project").
		Join("INNER", "user", "user.id = project.owner_id").
		Where("project.type = ? AND user.type = ?", TypeOrganization, UserTypeIndividual).
		Find(&ps)
	if err != nil {
		return err
	}

	pids := make(container.Set[int64], len(ps))
	for _, p := range ps {
		if p.ID == 0 {
			continue
		}
		pids.Add(p.ID)
	}

	sql, args, err := builder.ToSQL(builder.Eq{"id": pids.Values()})
	if err != nil {
		return err
	}

	count, err := sess.Table("project").
		Where(sql, args...).
		Update(&Project{
			Type: TypeIndividual,
		})
	if err != nil {
		return err
	}
	log.Debug("Updated %d projects to belong to a user instead of an organization", count)

	return sess.Commit()
}
