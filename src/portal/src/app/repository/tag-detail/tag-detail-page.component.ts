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
import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import {AppConfigService} from "../../app-config.service";
import { SessionService } from '../../shared/session.service';
import { ArtifactService } from '../../../lib/services';

@Component({
  selector: 'repository',
  templateUrl: 'tag-detail-page.component.html',
  styleUrls: ["tag-detail-page.component.scss"]
})
export class TagDetailPageComponent implements OnInit, OnDestroy {
  tagId: string;
  artifactDigest: string;
  repositoryName: string;
  projectId: string | number;
  referArtifactNameArray: string[];
  constructor(
    private route: ActivatedRoute,
    private appConfigService: AppConfigService,
    private router: Router,
    private artifactService: ArtifactService,
    private session: SessionService
  ) {
  }

  ngOnInit(): void {
    this.repositoryName = this.route.snapshot.params["repo"];
    this.artifactDigest = this.route.snapshot.params["digest"];
    this.projectId = this.route.snapshot.params["id"];
    this.referArtifactNameArray = this.artifactService.referenceSummary;
  }

  get withAdmiral(): boolean {
    return this.appConfigService.getConfig().with_admiral;
  }

  goBack(repositoryName: string): void {
    this.router.navigate(["harbor", "projects", this.projectId, "repositories", repositoryName]);
  }
  goBackRep(): void {
    this.router.navigate(["harbor", "projects", this.projectId, "repositories"]);
  }
  goBackPro(): void {
    this.router.navigate(["harbor", "projects"]);
  }
  ngOnDestroy(): void {
    this.artifactService.referenceSummary = [];
  }
  jumpDigest(referArtifactNameArray: string[], index: number) {
    this.artifactService.referenceSummary = [];
    this.artifactService.reference = referArtifactNameArray.slice(index);

    this.router.navigate(["harbor", "projects", this.projectId, "repositories", this.repositoryName]);
  }
}
