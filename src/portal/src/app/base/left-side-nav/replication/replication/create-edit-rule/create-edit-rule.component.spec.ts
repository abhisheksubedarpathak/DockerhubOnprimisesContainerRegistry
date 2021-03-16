import {ComponentFixture, fakeAsync, TestBed, tick, waitForAsync} from "@angular/core/testing";
import { NoopAnimationsModule } from "@angular/platform-browser/animations";
import { ConfirmationDialogComponent } from "../../../../../shared/components/confirmation-dialog";
import { CronTooltipComponent } from "../../../../../shared/components/cron-schedule";
import { CreateEditRuleComponent } from "./create-edit-rule.component";
import { DatePickerComponent } from "../../../../../shared/components/datetime-picker/datetime-picker.component";
import { FilterComponent } from "../../../../../shared/components/filter/filter.component";
import { InlineAlertComponent } from "../../../../../shared/components/inline-alert/inline-alert.component";
import {
  ReplicationRule,
  ReplicationJob,
  Endpoint,
  ReplicationJobItem
} from "../../../../../shared/services";
import { ErrorHandler } from "../../../../../shared/units/error-handler";
import {
  ReplicationService,
} from "../../../../../shared/services";
import {LabelPieceComponent} from "../../../../../shared/components/label/label-piece/label-piece.component";
import { RouterTestingModule } from '@angular/router/testing';
import { of } from "rxjs";
import {HttpHeaders, HttpResponse} from "@angular/common/http";
import {delay} from "rxjs/operators";
import { SharedTestingModule } from "../../../../../shared/shared.module";
import { EndpointService } from "../../../../../shared/services/endpoint.service";

describe("CreateEditRuleComponent (inline template)", () => {
  let mockRules: ReplicationRule[] = [
    {
      id: 1,
      name: "sync_01",
      description: "",
      src_registry: {id: 2},
      src_namespaces: ["name1", "name2"],
      trigger: {
        type: "Manual",
        trigger_settings: {}
      },
      filters: [],
      deletion: false,
      enabled: true,
      override: true
    }
  ];
  let mockJobs: ReplicationJobItem[] = [
    {
      id: 1,
      status: "stopped",
      policy_id: 1,
      trigger: "Manual",
      total: 0,
      failed: 0,
      succeed: 0,
      in_progress: 0,
      stopped: 0
    },
    {
      id: 2,
      status: "stopped",
      policy_id: 1,
      trigger: "Manual",
      total: 1,
      failed: 0,
      succeed: 1,
      in_progress: 0,
      stopped: 0
    },
    {
      id: 3,
      status: "stopped",
      policy_id: 2,
      trigger: "Manual",
      total: 1,
      failed: 1,
      succeed: 0,
      in_progress: 0,
      stopped: 0
    }
  ];

  let mockJob: ReplicationJob = {
    metadata: { xTotalCount: 3 },
    data: mockJobs
  };

  let mockEndpoints: Endpoint[] = [
    {
      id: 1,
      credential: {
        access_key: "admin",
        access_secret: "",
        type: "basic"
      },
      description: "test",
      insecure: false,
      name: "target_01",
      type: "Harbor",
      url: "https://10.117.4.151"
    },
    {
      id: 2,
      credential: {
        access_key: "AAA",
        access_secret: "",
        type: "basic"
      },
      description: "test",
      insecure: false,
      name: "target_02",
      type: "Harbor",
      url: "https://10.117.5.142"
    },
    {
      id: 3,
      credential: {
        access_key: "admin",
        access_secret: "",
        type: "basic"
      },
      description: "test",
      insecure: false,
      name: "target_03",
      type: "Harbor",
      url: "https://101.1.11.111"
    },
    {
      id: 4,
      credential: {
        access_key: "admin",
        access_secret: "",
        type: "basic"
      },
      description: "test",
      insecure: true,
      name: "target_04",
      type: "Harbor",
      url: "https://4.4.4.4"
    }
  ];

  let mockRule: ReplicationRule = {
    id: 1,
    name: "sync_01",
    description: "",
    src_namespaces: ["namespace1", "namespace2"],
    src_registry: {id: 10 },
    dest_registry: {id: 0 },
    trigger: {
      type: "Manual",
      trigger_settings: {}
    },
    filters: [],
    deletion: false,
    enabled: true,
    override: true
  };

  let mockRegistryInfo = {
    "type": "harbor",
    "description": "",
    "supported_resource_filters": [
      {
        "type": "Name",
        "style": "input"
      },
      {
        "type": "Version",
        "style": "input"
      },
      {
        "type": "Label",
        "style": "input"
      },
      {
        "type": "Resource",
        "style": "radio",
        "values": [
          "repository",
          "chart"
        ]
      }
    ],
    "supported_triggers": [
      "manual",
      "scheduled",
      "event_based"
    ]
  };
  let fixture: ComponentFixture<CreateEditRuleComponent>;
  let comp: CreateEditRuleComponent;
  const fakedErrorHandler = {
    error() {
    }
  };
  const fakedReplicationService = {
    getReplicationRule() {
      return of(mockRule).pipe(delay(0));
    },
    getReplicationRulesResponse() {
      return of(new HttpResponse({
        body: mockRules,
        headers:  new HttpHeaders({
          "x-total-count": "2"
        })
      })).pipe(delay(0));
    },
    getExecutions() {
      return of(mockJob).pipe(delay(0));
    },
    getEndpoints() {
      return of(mockEndpoints).pipe(delay(0));
    },
    getRegistryInfo() {
      return  of(mockRegistryInfo).pipe(delay(0));
    }
  };
  const fakedEndpointService = {
    getEndpoints() {
      return of(mockEndpoints).pipe(delay(0));
    }
  };
  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [SharedTestingModule, NoopAnimationsModule, RouterTestingModule],
      declarations: [
        CreateEditRuleComponent,
        CronTooltipComponent,
        ConfirmationDialogComponent,
        DatePickerComponent,
        FilterComponent,
        InlineAlertComponent,
        LabelPieceComponent
      ],
      providers: [
        { provide: ErrorHandler, useValue: fakedErrorHandler },
        { provide: ReplicationService, useValue: fakedReplicationService },
        { provide: EndpointService, useValue: fakedEndpointService },
      ]
    });
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateEditRuleComponent);
    comp = fixture.componentInstance;
    fixture.detectChanges();
  });

  it("Should open creation modal and load endpoints", async () => {
    fixture.detectChanges();
    await fixture.whenStable();
    comp.openCreateEditRule();
    fixture.detectChanges();
    await fixture.whenStable();
    const modal =  fixture.nativeElement.querySelector("clr-modal");
    expect(modal).toBeTruthy();
    const selectionOptions = fixture.nativeElement.querySelectorAll("#dest_registry>option");
    expect(selectionOptions).toBeTruthy();
    expect(selectionOptions.length).toEqual(5);
  });

  it("Should open modal to edit replication rule", fakeAsync( () => {
    fixture.detectChanges();
    comp.openCreateEditRule(mockRule.id);
    fixture.detectChanges();
    tick(5000);
    const ruleNameInput: HTMLInputElement = fixture.nativeElement.querySelector("#ruleName");
    expect(ruleNameInput).toBeTruthy();
    expect(ruleNameInput.value.trim()).toEqual("sync_01");
  }));
});
