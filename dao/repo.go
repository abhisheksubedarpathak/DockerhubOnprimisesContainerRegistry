package dao

import (
	"github.com/vmware/harbor/models"
)

// func AddLabel(repoLabel models.RepoLabel) (int64, error) {
// 	o := GetOrmer()
//
// 	sql := `insert into repo_label(repoName, label) values(?,?)`
//
//
// 	p,_ := o.Raw(sql).Prepare()
//
// 	defer p.Close()
//
// 	r, err := p.Exec(repoLabel.RepoName, repoLabel.Label)
//
// 	insertId,_ := r.LastInsertId()
// 	return insertId, err
// }

func AddLabel(repoLabel models.RepoLabel) (bool, error) {
	o := GetOrmer()

	sql := `select * from repo_label where repoName = ? and label = ?`
	type dummy struct{}
	var d []dummy
	_, err := o.Raw(sql, repoLabel.RepoName, repoLabel.Label).QueryRows(&d)
	if len(d) != 0 {
		return false, err
	}
	sql = `insert into repo_label(repoName, label) values(?,?)`

	p,_ := o.Raw(sql).Prepare()

	defer p.Close()

	_, err = p.Exec(repoLabel.RepoName, repoLabel.Label)

	return true, err
}

func DeletelLabel(repoLabel models.RepoLabel) (int64, error) {
	o := GetOrmer()

	sql := `delete from repo_label where repoName=? and label=?`


	p,_ := o.Raw(sql).Prepare()

	defer p.Close()

	r, err := p.Exec(repoLabel.RepoName, repoLabel.Label)

	affectedRows, _ := r.RowsAffected()

	return affectedRows, err
}


func GetRepoLabels(repoName string) ([]string, error){
	o := GetOrmer()

	sql := `select label from repo_label where repoName=?`

	var labels []string

	if _, err := o.Raw(sql, repoName).QueryRows(&labels); err != nil {
		return nil, err
	}

	return labels, nil
}


func GetRepoNames(label string) ([]string, error){
	o := GetOrmer()

	sql := `select repoName from repo_label where label=?`

	var repoNames []string

	if _, err := o.Raw(sql, label).QueryRows(&repoNames); err != nil {
		return nil, err
	}

	return repoNames, nil
}
