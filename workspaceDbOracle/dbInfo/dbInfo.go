package dbinfo

import (
	"database/sql"
	"fmt"
	"time"
)

type Claim struct {
	ID          int
	CreateDate  time.Time
	Description string
	Priority    string
	Progress    int
}

func (c Claim) IsEmpty() bool {
	return c.Description == ""
}

func GetClaimsByPriority(priority string, db *sql.DB) ([]Claim, error) {
	var claims []Claim
	rows, error := db.Query("SELECT ID, create_date, description, priority, progress from claim1_tmp_remove where priority = :1", priority)
	if error != nil {
		return nil, fmt.Errorf("GetClaimsByPriority can not get rows - %w", error)
	}
	defer rows.Close()
	for rows.Next() {
		var clm Claim
		if error := rows.Scan(&clm.ID, &clm.CreateDate, &clm.Description, &clm.Priority, &clm.Progress); error != nil {
			return nil, fmt.Errorf("GetClaimsByPriority can not scan rows - %w", error)
		}
		claims = append(claims, clm)
	}
	return claims, nil
}

func GetClaimById(id int, db *sql.DB) (Claim, error) {

	row := db.QueryRow("SELECT ID, create_date, description, priority, progress from claim1_tmp_remove where ID = :1", id)
	var clm Claim
	if error := row.Scan(&clm.ID, &clm.CreateDate, &clm.Description, &clm.Priority, &clm.Progress); error != nil {
		return Claim{}, fmt.Errorf("GetClaimsByPriority can not scan row - %w", error)
	}
	return clm, nil
}

func ModifyClaimById(id int, db *sql.DB) (int, error) {

	res, error := db.Exec("update claim1_tmp_remove set priority = 'Major' where ID = :1", id)
	_ = res
	if error != nil {
		return 0, fmt.Errorf("ModifyClaimById can not update rows - %w", error)
	}
	return id, nil
}

func AddNewClaim(newClaim Claim, db *sql.DB) (int, error) {

	res, error := db.Exec("insert into claim1_tmp_remove(id, create_date, description, priority, progress) values(:1, :2, :3, :4, :5)",
		newClaim.ID, newClaim.CreateDate, newClaim.Description, newClaim.Priority, newClaim.Progress)
	_ = res
	if error != nil {
		return 0, fmt.Errorf("AddNewClaim can not insert rows - %w", error)
	}
	return newClaim.ID, nil
}

func GetDesciptionFromPLSQL(claimId int, db *sql.DB) (string, error) {

	row := db.QueryRow("select claim_pack.get_description_by_claim_id(:1) from dual", claimId)
	var clmDescription string
	if error := row.Scan(&clmDescription); error != nil {
		return "", fmt.Errorf("GetClaimsByPriority can not scan row - %w", error)
	}
	return clmDescription, nil
}

func ModifyWithPLSQLFunctionById(id int, db *sql.DB) (int, error) {

	res, error := db.Query("begin claim_pack.modify_claim1_tmp(:1); end;", id)
	_ = res
	if error != nil {
		return 0, fmt.Errorf("ModifyWithPLSQLFunctionById can not update rows - %w", error)
	}
	return id, nil
}
