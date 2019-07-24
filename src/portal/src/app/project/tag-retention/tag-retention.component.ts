// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
import {Component, OnInit, ViewChild} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {AddRuleComponent} from "./add-rule/add-rule.component";
import {ClrDatagridStringFilterInterface} from "@clr/angular";
import {TagRetentionService} from "./tag-retention.service";
import {Retention} from "./retention";
import {Project} from "../project";
import {clone, ErrorHandler} from "@harbor/ui";
import {EditRuleComponent} from "./edit-rule/edit-rule.component";
const MIN = 60000;
const SEC = 1000;
const MIN_STR = "min";
const SEC_STR = "sec";
@Component({
    selector: 'tag-retention',
    templateUrl: './tag-retention.component.html',
    styleUrls: ['./tag-retention.component.scss']
})
export class TagRetentionComponent implements OnInit {
    serialFilter: ClrDatagridStringFilterInterface<any> = {
        accepts(item: any, search: string): boolean {
            return item.id.toString().indexOf(search) !== -1;
        }
    };
    statusFilter: ClrDatagridStringFilterInterface<any> = {
        accepts(item: any, search: string): boolean {
            return item.status.toLowerCase().indexOf(search.toLowerCase()) !== -1;
        }
    };
    dryRunFilter: ClrDatagridStringFilterInterface<any> = {
        accepts(item: any, search: string): boolean {
            let str = item.dry_run ? 'YES' : 'NO';
            return str.indexOf(search) !== -1;
        }
    };
    projectId: number;
    addRuleOpened: boolean = false;
    isRetentionRunOpened: boolean = false;
    isAbortedOpened: boolean = false;
    selectedItem;
    ruleIndex: number = -1;
    index: number = -1;
    retentionId: number;
    retention: Retention = new Retention();
    editIndex: number;
    executionList = [];
    historyList = [];
    loadingExecutions: boolean = false;
    loadingHistories: boolean = false;
    metadata: any;
    dryRun: boolean = false;
    @ViewChild('addRule') addRuleComponent: AddRuleComponent;
    @ViewChild('editRule') editRuleComponent: EditRuleComponent;

    constructor(
        private route: ActivatedRoute,
        private tagRetentionService: TagRetentionService,
        private errorHandler: ErrorHandler,
    ) {
    }

    ngOnInit() {
        this.projectId = +this.route.snapshot.parent.params['id'];
        this.retention.scope = {
            level: "project",
            ref: this.projectId
        };
        let resolverData = this.route.snapshot.parent.data;
        if (resolverData) {
            let project = <Project>resolverData["projectResolver"];
            if (project.metadata && project.metadata.retention_id) {
                this.retentionId = project.metadata.retention_id;
            }
        }
        this.getRetention();
        this.getMetadata();
        this.refreshList();
    }
    getMetadata() {
        this.tagRetentionService.getRetentionMetadata().subscribe(
            response => {
               this.metadata = response;
            }, error => {
                this.errorHandler.error(error);
            });
    }
    getRetention() {
        if (this.retentionId) {
            this.tagRetentionService.getRetention(this.retentionId).subscribe(
                response => {
                    this.retention = response;
                }, error => {
                    this.errorHandler.error(error);
                });
        }
    }

    repositoryToString(repositories): string {
        let arr = [];
        repositories.forEach(rep => {
            arr.push(rep.pattern);
        });
        return arr.join(",");
    }

    editRuleByIndex(index) {
        this.editIndex = index;
        this.editRuleComponent.rule = clone(this.retention.rules[index]);
        this.editRuleComponent.init();
        this.editRuleComponent.open();
    }

    clickSave(rule) {
        let retention: Retention = clone(this.retention);
        retention.rules[this.editIndex] = rule;
        this.tagRetentionService.updateRetention(this.retentionId, retention).subscribe(
            response => {
                this.retention = retention;
            }, error => {
                this.errorHandler.error(error);
            });
    }

    disableRule(index) {
        let retention: Retention = clone(this.retention);
        retention.rules[index].isDisabled = true;
        this.tagRetentionService.updateRetention(this.retentionId, retention).subscribe(
            response => {
                this.retention = retention;
            }, error => {
                this.errorHandler.error(error);
            });
    }

    enableRule(index) {
        let retention: Retention = clone(this.retention);
        retention.rules[index].isDisabled = false;
        this.tagRetentionService.updateRetention(this.retentionId, retention).subscribe(
            response => {
                this.retention = retention;
            }, error => {
                this.errorHandler.error(error);
            });
    }

    deleteRule(index) {
        let retention: Retention = clone(this.retention);
        retention.rules.splice(index, 1);
        this.tagRetentionService.updateRetention(this.retentionId, retention).subscribe(
            response => {
                this.retention = retention;
            }, error => {
                this.errorHandler.error(error);
            });
    }

    openAddRule() {
        this.addRuleComponent.open();
    }

    runRetention() {
        this.isRetentionRunOpened = false;
        this.tagRetentionService.runNowTrigger(this.retentionId).subscribe(
            response => {
                this.refreshList();
            }, error => {
                this.errorHandler.error(error);
            });
    }
    whatIfRun() {
        this.tagRetentionService.whatIfRunTrigger(this.retentionId).subscribe(
            response => {
                this.refreshList();
            }, error => {
                this.errorHandler.error(error);
            });
    }

    refreshList() {
        if (this.retentionId) {
            this.loadingExecutions = true;
            this.tagRetentionService.getRunNowList(this.retentionId).subscribe(
                res => {
                    this.loadingExecutions = false;
                    this.executionList = res;
                    TagRetentionComponent.calculateDuration(this.executionList);
                }, error => {
                    this.loadingExecutions = false;
                    this.errorHandler.error(error);
                });
        }
    }
    static calculateDuration(arr: Array<any>) {
        if (arr && arr.length > 0) {
            for (let i = 0; i < arr.length; i++) {
                let duration = new Date(arr[i].end_time).getTime() - new Date(arr[i].start_time).getTime();
                let min = Math.floor(duration / MIN);
                let sec = Math.floor((duration % MIN) / SEC);
                arr[i]['duration'] = "";
                if ((min || sec) && duration > 0) {
                    if (min) {
                        arr[i]['duration'] += '' + min + MIN_STR;
                    }
                    if (sec) {
                        arr[i]['duration'] += '' + sec + SEC_STR;
                    }
                } else {
                    arr[i]['duration'] = "NA";
                }
            }
        }
    }

    abortRun() {
        this.isAbortedOpened = true;
        this.tagRetentionService.AbortRun(this.retentionId, this.selectedItem.id).subscribe(
            res => {
                this.refreshList();
            }, error => {
                this.errorHandler.error(error);
            });
    }
    abortRetention() {
        this.isAbortedOpened = false;
    }
    openEditor(index) {
        if (this.ruleIndex !== index) {
            this.ruleIndex = index;
        } else {
            this.ruleIndex = -1;
        }
    }

    openDetail(index, executionId) {
        if (this.index !== index) {
            this.index = index;
            this.historyList = [];
            this.loadingHistories = true;
            this.tagRetentionService.getExecutionHistory(this.retentionId, executionId).subscribe(
                res => {
                    this.loadingHistories = false;
                    this.historyList = res;
                    TagRetentionComponent.calculateDuration(this.historyList);
                }, error => {
                    this.loadingHistories = false;
                    this.errorHandler.error(error);
                });
        } else {
            this.index = -1;
        }
    }

    refreshAfterCreatRetention() {
        this.tagRetentionService.getProjectInfo(this.projectId).subscribe(
            response => {
                this.retentionId = response.metadata.retention_id;
                this.getRetention();
            }, error => {
                this.errorHandler.error(error);
            });
    }

    clickAdd(rule) {
        let retention: Retention = clone(this.retention);
        retention.rules.push(rule);
        if (!this.retentionId) {
            this.tagRetentionService.createRetention(retention).subscribe(
                response => {
                    this.refreshAfterCreatRetention();
                }, error => {
                    this.errorHandler.error(error);
                });
        } else {
            this.tagRetentionService.updateRetention(this.retentionId, retention).subscribe(
                response => {
                    this.retention = retention;
                }, error => {
                    this.errorHandler.error(error);
                });
        }
    }
    seeLog(executionId, taskId) {
       this.tagRetentionService.seeLog(this.retentionId , executionId, taskId);
    }
}
