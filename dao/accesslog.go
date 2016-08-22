/*
   Copyright (c) 2016 VMware, Inc. All Rights Reserved.
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package dao

import (
	"strings"

	"github.com/vmware/harbor/models"
	"github.com/vmware/harbor/utils/log"
)

// AddAccessLog persists the access logs
func AddAccessLog(accessLog models.AccessLog) error {
	o := GetOrmer()
	p, err := o.Raw(`insert into access_log
		 (user_id, project_id, repo_name, repo_tag, guid, operation, op_time)
		 values (?, ?, ?, ?, ?, ?, now())`).Prepare()
	if err != nil {
		return err
	}
	defer p.Close()

	_, err = p.Exec(accessLog.UserID, accessLog.ProjectID, accessLog.RepoName, accessLog.RepoTag, accessLog.GUID, accessLog.Operation)

	return err
}

// GetTotalOfAccessLogs ...
func GetTotalOfAccessLogs(query models.AccessLog) (int64, error) {
	o := GetOrmer()

	queryParam := []interface{}{}

	sql := `select count(*) from access_log al 
		where al.project_id = ?`
	queryParam = append(queryParam, query.ProjectID)

	if query.Username != "" {
		sql = `select count(*) from access_log al 
			left join user u 
			on al.user_id = u.user_id 
			where al.project_id = ? and u.username like ? `
		queryParam = append(queryParam, "%"+query.Username+"%")
	}

	sql += genFilterClauses(query, &queryParam)

	var total int64
	if err := o.Raw(sql, queryParam).QueryRow(&total); err != nil {
		return 0, err
	}
	return total, nil
}

//GetAccessLogs gets access logs according to different conditions
func GetAccessLogs(query models.AccessLog, limit, offset int64) ([]models.AccessLog, error) {
	o := GetOrmer()

	queryParam := []interface{}{}
	sql := `select al.log_id, u.username, al.repo_name, 
			al.repo_tag, al.operation, al.op_time 
		from access_log al 
		left join user u 
		on al.user_id = u.user_id
		where al.project_id = ? `
	queryParam = append(queryParam, query.ProjectID)

	if query.Username != "" {
		sql += ` and u.username like ? `
		queryParam = append(queryParam, "%"+query.Username+"%")
	}

	sql += genFilterClauses(query, &queryParam)

	sql += ` order by al.op_time desc `

	sql = paginateForRawSQL(sql, limit, offset)

	logs := []models.AccessLog{}
	_, err := o.Raw(sql, queryParam).QueryRows(&logs)
	if err != nil {
		return logs, err
	}

	return logs, nil
}

func genFilterClauses(query models.AccessLog, queryParam *[]interface{}) string {
	sql := ""

	if query.Operation != "" {
		sql += ` and al.operation = ? `
		*queryParam = append(*queryParam, query.Operation)
	}
	if query.RepoName != "" {
		sql += ` and al.repo_name = ? `
		*queryParam = append(*queryParam, query.RepoName)
	}
	if query.RepoTag != "" {
		sql += ` and al.repo_tag = ? `
		*queryParam = append(*queryParam, query.RepoTag)
	}
	if query.Keywords != "" {
		sql += ` and al.operation in ( `
		keywordList := strings.Split(query.Keywords, "/")
		num := len(keywordList)
		for i := 0; i < num; i++ {
			if keywordList[i] != "" {
				if i == num-1 {
					sql += `?)`
				} else {
					sql += `?,`
				}
				*queryParam = append(*queryParam, keywordList[i])
			}
		}
	}
	if query.BeginTimestamp > 0 {
		sql += ` and al.op_time >= ? `
		*queryParam = append(*queryParam, query.BeginTime)
	}
	if query.EndTimestamp > 0 {
		sql += ` and al.op_time <= ? `
		*queryParam = append(*queryParam, query.EndTime)
	}

	return sql
}

// AccessLog ...
func AccessLog(username, projectName, repoName, repoTag, action string) error {
	o := GetOrmer()
	sql := "insert into  access_log (user_id, project_id, repo_name, repo_tag, operation, op_time) " +
		"select (select user_id as user_id from user where username=?), " +
		"(select project_id as project_id from project where name=?), ?, ?, ?, now() "
	_, err := o.Raw(sql, username, projectName, repoName, repoTag, action).Exec()

	if err != nil {
		log.Errorf("error in AccessLog: %v ", err)
	}
	return err
}

//GetRecentLogs returns recent logs according to parameters
func GetRecentLogs(userID, linesNum int, startTime, endTime string) ([]models.AccessLog, error) {
	var recentLogList []models.AccessLog
	queryParam := make([]interface{}, 1)

	sql := "select log_id, access_log.user_id, project_id, repo_name, repo_tag, GUID, operation, op_time, username from access_log left join  user on access_log.user_id=user.user_id where project_id in (select distinct project_id from project_member where user_id = ?)"
	queryParam = append(queryParam, userID)
	if startTime != "" {
		sql += " and op_time >= ?"
		queryParam = append(queryParam, startTime)
	}

	if endTime != "" {
		sql += " and op_time <= ?"
		queryParam = append(queryParam, endTime)
	}

	sql += " order by op_time desc"
	if linesNum != 0 {
		sql += " limit ?"
		queryParam = append(queryParam, linesNum)
	}
	o := GetOrmer()
	_, err := o.Raw(sql, queryParam).QueryRows(&recentLogList)
	if err != nil {
		return nil, err
	}
	return recentLogList, nil
}

//GetTopRepos return top  accessed public repos
func GetTopRepos(countNum int) ([]models.TopRepo, error) {

	o := GetOrmer()
	// hide the where condition: project.public = 1, Can add to the sql when necessary.
	sql := `select repo_name, COUNT(repo_name) as access_count 
				from access_log 
				where access_log.operation = 'pull' 
				group by repo_name 
				order by access_count desc 
				limit ? `
	queryParam := []interface{}{}
	queryParam = append(queryParam, countNum)

	list := []models.TopRepo{}
	_, err := o.Raw(sql, queryParam).QueryRows(&list)

	return list, err
}
