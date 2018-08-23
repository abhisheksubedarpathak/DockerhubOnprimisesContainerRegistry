// Copyright 2018 The Harbor Authors
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
import { Injectable } from '@angular/core';
import {
  CanActivate, Router,
  ActivatedRouteSnapshot,
  RouterStateSnapshot,
  CanActivateChild
} from '@angular/router';
import { CommonRoutes } from '../../shared/shared.const';
import { AppConfigService } from '../../app-config.service';

@Injectable()
export class ModeGuard implements CanActivate, CanActivateChild {
  constructor(
    private router: Router,
    private appConfigService: AppConfigService) { }

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Promise<boolean> | boolean {
    // Show the right sign-in page for different modes
    return new Promise((resolve, reject) => {
      if (this.appConfigService.isIntegrationMode()) {
        if (state.url.startsWith(CommonRoutes.SIGN_IN)) {
          this.router.navigate([CommonRoutes.EMBEDDED_SIGN_IN], route.queryParams);
          resolve(false);
        } else {
          resolve(true);
        }
      } else {
        if (state.url.startsWith(CommonRoutes.EMBEDDED_SIGN_IN)) {
          this.router.navigate([CommonRoutes.SIGN_IN], route.queryParams);
          resolve(false);
        } else {
          resolve(true);
        }
      }
    });
  }

  canActivateChild(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Promise<boolean> | boolean {
    return this.canActivate(route, state);
  }
}
